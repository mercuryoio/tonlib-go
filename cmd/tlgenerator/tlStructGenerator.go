package main

import (
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

func generateStructsFromTnEntities(
	packageName string, entities *[]ClassInfo, interfaces *[]InterfaceInfo, enums *[]EnumInfo) (*string, *string) {
	structsContent := fmt.Sprintf("package %s\n\n", packageName)
	structUnmarshals := ""
	structsContent += `
	
	import (
		"encoding/json"
		"fmt"
		"strconv"
		"strings"
	)
	
	`
	methodsContent := fmt.Sprintf("package %s\n\n", packageName)
	methodsContent += `
	
	import (
		"encoding/json"
		"fmt"
	)
	
	`

	structsContent += "type tdCommon struct {\n" +
		"Type string `json:\"@type\"`\n" +
		"Extra string `json:\"@extra\"`\n" +
		"}\n\n"

	structsContent += `
	// JSONInt64 alias for int64, in order to deal with json big number problem
	type JSONInt64 int64
	
	// MarshalJSON marshals to json
	func (jsonInt *JSONInt64) MarshalJSON() ([]byte, error) {
		intStr := strconv.FormatInt(int64(*jsonInt), 10)
		return []byte(intStr), nil
	}
	
	// UnmarshalJSON unmarshals from json
	func (jsonInt *JSONInt64) UnmarshalJSON(b []byte) error {
		intStr := string(b)
		intStr = strings.Replace(intStr, "\"", "", 2)
		jsonBigInt, err := strconv.ParseInt(intStr, 10, 64)
		if err != nil {
			return err
		}
		*jsonInt = JSONInt64(jsonBigInt)
		return nil
	}
`

	structsContent += `
		// TdMessage is the interface for all messages send and received to/from tdlib
		type TdMessage interface{
			MessageType() string
		}
`

	for _, enumInfoe := range *enums {

		structsContent += fmt.Sprintf(`
				// %s Alias for abstract %s 'Sub-Classes', used as constant-enum here
				type %s string
				`,
			enumInfoe.EnumType,
			enumInfoe.EnumType[:len(enumInfoe.EnumType)-len("Enum")],
			enumInfoe.EnumType)

		consts := ""
		for _, item := range enumInfoe.Items {
			consts += item + "Type " + enumInfoe.EnumType + " = \"" +
				strings.ToLower(item[:1]) + item[1:] + "\"\n"

		}
		structsContent += fmt.Sprintf(`
				// %s enums
				const (
					%s
				)`, enumInfoe.EnumType[:len(enumInfoe.EnumType)-len("Enum")], consts)
	}

	for _, interfaceInfo := range *interfaces {
		interfaceInfo.Name = replaceKeyWords(interfaceInfo.Name)
		typesCases := ""

		structsContent += fmt.Sprintf("// %s %s \ntype %s interface {\nGet%sEnum() %sEnum\n}\n\n",
			interfaceInfo.Name, interfaceInfo.Description, interfaceInfo.Name, interfaceInfo.Name, interfaceInfo.Name)

		for _, enumInfoe := range *enums {
			if enumInfoe.EnumType == interfaceInfo.Name+"Enum" {
				for _, enumItem := range enumInfoe.Items {
					typeName := enumItem
					typeNameCamel := strings.ToLower(typeName[:1]) + typeName[1:]
					typesCases += fmt.Sprintf(`case %s:
						var %s %s
						err := json.Unmarshal(*rawMsg, &%s)
						return &%s, err
						
						`,
						enumItem+"Type", typeNameCamel, typeName,
						typeNameCamel, typeNameCamel)
				}
				break
			}
		}

		structUnmarshals += fmt.Sprintf(`
				func unmarshal%s(rawMsg *json.RawMessage) (%s, error){

					if rawMsg == nil {
						return nil, nil
					}
					var objMap map[string]interface{}
					err := json.Unmarshal(*rawMsg, &objMap)
					if err != nil {
						return nil, err
					}

					switch %sEnum(objMap["@type"].(string)) {
						%s
					default:
						return nil, fmt.Errorf("Error unmarshaling, unknown type:" +  objMap["@type"].(string))
					}
				}
				`, interfaceInfo.Name, interfaceInfo.Name, interfaceInfo.Name,
			typesCases)
	}

	for _, classInfoe := range *entities {
		if !classInfoe.IsFunction {
			structName := strings.ToUpper(classInfoe.Name[:1]) + classInfoe.Name[1:]
			structName = replaceKeyWords(structName)
			structNameCamel := strings.ToLower(structName[0:1]) + structName[1:]

			hasInterfaceProps := false
			propsStr := ""
			propsStrWithoutInterfaceOnes := ""
			assignStr := fmt.Sprintf("%s.tdCommon = tempObj.tdCommon\n", structNameCamel)
			assignInterfacePropsStr := ""

			// sort.Sort(classInfoe.Properties)
			for i, prop := range classInfoe.Properties {
				propName := govalidator.UnderscoreToCamelCase(prop.Name)
				propName = replaceKeyWords(propName)

				dataType, isPrimitive := convertDataType(prop.Type)
				propsStrItem := ""
				if isPrimitive || checkIsInterface(dataType, interfaces) {
					propsStrItem += fmt.Sprintf("%s %s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
				} else {
					propsStrItem += fmt.Sprintf("%s *%s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
				}
				if i < len(classInfoe.Properties)-1 {
					propsStrItem += "\n"
				}

				propsStr += propsStrItem
				if !checkIsInterface(prop.Type, interfaces) {
					propsStrWithoutInterfaceOnes += propsStrItem
					assignStr += fmt.Sprintf("%s.%s = tempObj.%s\n", structNameCamel, propName, propName)
				} else {
					hasInterfaceProps = true
					assignInterfacePropsStr += fmt.Sprintf(`
						field%s, _  := 	unmarshal%s(objMap["%s"])
						%s.%s = field%s
						`,
						propName, dataType, prop.Name,
						structNameCamel, propName, propName)
				}
			}
			structsContent += fmt.Sprintf("// %s %s \ntype %s struct {\n"+
				"tdCommon\n"+
				"%s\n"+
				"}\n\n", structName, classInfoe.Description, structName, propsStr)

			structsContent += fmt.Sprintf("// MessageType return the string telegram-type of %s \nfunc (%s *%s) MessageType() string {\n return \"%s\" }\n\n",
				structName, structNameCamel, structName, classInfoe.Name)

			paramsStr := ""
			paramsDesc := ""
			assingsStr := ""
			for i, param := range classInfoe.Properties {
				propName := govalidator.UnderscoreToCamelCase(param.Name)
				propName = replaceKeyWords(propName)
				dataType, isPrimitive := convertDataType(param.Type)
				paramName := convertToArgumentName(param.Name)

				if isPrimitive || checkIsInterface(dataType, interfaces) {
					paramsStr += paramName + " " + dataType

				} else { // if is not a primitive, use pointers
					paramsStr += paramName + " *" + dataType
				}

				if i < len(classInfoe.Properties)-1 {
					paramsStr += ", "
				}
				paramsDesc += "\n// @param " + paramName + " " + param.Description

				if isPrimitive || checkIsInterface(dataType, interfaces) {
					assingsStr += fmt.Sprintf("%s : %s,\n", propName, paramName)
				} else {
					assingsStr += fmt.Sprintf("%s : %s,\n", propName, paramName)
				}
			}

			// Create New... constructors
			structsContent += fmt.Sprintf(`
				// New%s creates a new %s
				// %s
				func New%s(%s) *%s {
					%sTemp := %s {
						tdCommon: tdCommon {Type: "%s"},
						%s
					}

					return &%sTemp
				}
				`, structName, structName, paramsDesc,
				structName, paramsStr, structName, structNameCamel,
				structName, classInfoe.Name, assingsStr, structNameCamel)

			if hasInterfaceProps {
				structsContent += fmt.Sprintf(`
					// UnmarshalJSON unmarshal to json
					func (%s *%s) UnmarshalJSON(b []byte) error {
						var objMap map[string]*json.RawMessage
						err := json.Unmarshal(b, &objMap)
						if err != nil {
							return err
						}
						tempObj := struct {
							tdCommon
							%s
						}{}
						err = json.Unmarshal(b, &tempObj)
						if err != nil {
							return err
						}

						%s

						%s	
						
						return nil
					}
					`, structNameCamel, structName, propsStrWithoutInterfaceOnes,
					assignStr, assignInterfacePropsStr)
			}
			if checkIsInterface(classInfoe.RootName, interfaces) {
				rootName := replaceKeyWords(classInfoe.RootName)
				structsContent += fmt.Sprintf(`
					// Get%sEnum return the enum type of this object 
					func (%s *%s) Get%sEnum() %sEnum {
						 return %s 
					}

					`,
					rootName,
					strings.ToLower(structName[0:1])+structName[1:],
					structName, rootName, rootName,
					structName+"Type")
			}

		} else {
			methodName := strings.ToUpper(classInfoe.Name[:1]) + classInfoe.Name[1:]
			methodName = replaceKeyWords(methodName)
			returnType := strings.ToUpper(classInfoe.RootName[:1]) + classInfoe.RootName[1:]
			returnType = replaceKeyWords(returnType)
			returnTypeCamel := strings.ToLower(returnType[:1]) + returnType[1:]
			returnIsInterface := checkIsInterface(returnType, interfaces)

			asterike := "*"
			ampersign := "&"
			if returnIsInterface {
				asterike = ""
				ampersign = ""
			}

			paramsStr := ""
			paramsDesc := ""
			for i, param := range classInfoe.Properties {
				paramName := convertToArgumentName(param.Name)
				dataType, isPrimitive := convertDataType(param.Type)
				if isPrimitive || checkIsInterface(dataType, interfaces) {
					paramsStr += paramName + " " + dataType

				} else {
					paramsStr += paramName + " *" + dataType
				}

				if i < len(classInfoe.Properties)-1 {
					paramsStr += ", "
				}
				paramsDesc += "\n// @param " + paramName + " " + param.Description
			}

			methodsContent += fmt.Sprintf(`
				// %s %s %s
				func (client *Client) %s(%s) (%s%s, error)`, methodName, classInfoe.Description, paramsDesc, methodName,
				paramsStr, asterike, returnType)

			paramsStr = ""
			for i, param := range classInfoe.Properties {
				paramName := convertToArgumentName(param.Name)

				paramsStr += fmt.Sprintf("\"%s\":   %s,", param.Name, paramName)
				if i < len(classInfoe.Properties)-1 {
					paramsStr += "\n"
				}
			}

			illStr := `fmt.Errorf("error! code: %d msg: %s", result.Data["code"], result.Data["message"])`
			if strings.Contains(paramsStr, returnTypeCamel) {
				returnTypeCamel = returnTypeCamel + "Dummy"
			}
			if returnIsInterface {
				enumType := returnType + "Enum"
				casesStr := ""

				for _, enumInfoe := range *enums {
					if enumInfoe.EnumType == enumType {
						for _, item := range enumInfoe.Items {
							casesStr += fmt.Sprintf(`
								case %s:
									var %s %s
									err = json.Unmarshal(result.Raw, &%s)
									return &%s, err
									`, item+"Type", returnTypeCamel, item, returnTypeCamel,
								returnTypeCamel)
						}
						break
					}
				}

				methodsContent += fmt.Sprintf(` {
					result, err := client.SendAndCatch(UpdateData{
						"@type":       "%s",
						%s
					})
	
					if err != nil {
						return nil, err
					}
	
					if result.Data["@type"].(string) == "error" {
						return nil, %s
					}

					switch %s(result.Data["@type"].(string)) {
						%s
					default:
						return nil, fmt.Errorf("Invalid type")
					}
					}
					
					`, classInfoe.Name, paramsStr, illStr,
					enumType, casesStr)

			} else {
				methodsContent += fmt.Sprintf(` {
					result, err := client.SendAndCatch(UpdateData{
						"@type":       "%s",
						%s
					})
	
					if err != nil {
						return nil, err
					}
	
					if result.Data["@type"].(string) == "error" {
						return nil, %s
					}
	
					var %s %s
					err = json.Unmarshal(result.Raw, &%s)
					return %s%s, err
	
					}
					
					`, classInfoe.Name, paramsStr, illStr, returnTypeCamel,
					returnType, returnTypeCamel, ampersign, returnTypeCamel)
			}

		}
	}

	structsContent += "\n\n" + structUnmarshals
	return &structsContent, &methodsContent
}
