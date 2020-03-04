package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/asaskevich/govalidator"
)

var StructNamesExcludedFromGenerator = []string{
	"secureBytes", "secureString", "bytes", "vector", "key",
}

var SkipMethodNames = []string{
	"sync", "query.estimateFees",
}

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

	structsContent += "type tonCommon struct {\n" +
		"Type string `json:\"@type\"`\n" +
		"Extra string `json:\"@extra\"`\n" +
		"}\n\n"

	structsContent += `
	type SecureBytes   []byte
	type SecureString  string
	type Bytes         []byte
	type TvmStackEntry interface {}
	type SmcMethodId   interface {} 
	type TvmNumber     interface {} 
	type GenericAccountState string
	`

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
		// TonMessage is the interface for all messages send and received to/from tonlib
		type TonMessage interface{
			MessageType() string
		}
`

	for _, enum := range *enums {

		structsContent += fmt.Sprintf(`
				// %s Alias for abstract %s 'Sub-Classes', used as constant-enum here
				type %s string
				`,
			enum.EnumType,
			enum.EnumType[:len(enum.EnumType)-len("Enum")],
			enum.EnumType)

		consts := ""
		for _, item := range enum.Items {
			consts += item + "Type " + enum.EnumType + " = \"" +
				strings.ToLower(item[:1]) + item[1:] + "\"\n"

		}
		structsContent += fmt.Sprintf(`
				// %s enums
				const (
					%s
				)`, enum.EnumType[:len(enum.EnumType)-len("Enum")], consts)
	}

	for _, interfaceInfo := range *interfaces {
		interfaceInfo.Name = interfaceInfo.Name
		typesCases := ""

		structsContent += fmt.Sprintf("// %s %s \ntype %s interface {\nGet%sEnum() %sEnum\n}\n\n",
			interfaceInfo.Name, interfaceInfo.Description, interfaceInfo.Name, interfaceInfo.Name, interfaceInfo.Name)

		for _, enum := range *enums {
			if enum.EnumType == interfaceInfo.Name+"Enum" {
				for _, enumItem := range enum.Items {
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

	// gen entity`s structs
	for _, itemInfo := range *entities {
		// skip generation
		skip := false
		for _, name := range StructNamesExcludedFromGenerator {
			if itemInfo.Name == name {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		if !itemInfo.IsFunction {
			structName := getStructName(itemInfo.Name)
			structNameCamel := strings.ToLower(structName[0:1]) + structName[1:]

			hasInterfaceProps := false
			propsStr := ""
			propsStrWithoutInterfaceOnes := ""
			assignStr := fmt.Sprintf("%s.tonCommon = tempObj.tonCommon\n", structNameCamel)
			assignInterfacePropsStr := ""

			// sort params to enshure the same params order in each generation
			sort.Slice(itemInfo.Properties, func(i, j int) bool {
				if (itemInfo.Properties[i].Name > itemInfo.Properties[j].Name){
					return false
				}
				return true
			})

			for i, prop := range itemInfo.Properties {
				propName := govalidator.UnderscoreToCamelCase(prop.Name)
				propName = propName

				dataType, isPrimitive := convertDataType(prop.Type, true)
				propsStrItem := ""
				if isPrimitive || checkIsInterface(dataType, interfaces) {
					propsStrItem += fmt.Sprintf("%s %s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
				} else {
					propsStrItem += fmt.Sprintf("%s *%s `json:\"%s\"` // %s", propName, dataType, prop.Name, prop.Description)
				}
				if i < len(itemInfo.Properties)-1 {
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
				"tonCommon\n"+
				"%s\n"+
				"}\n\n", structName, itemInfo.Description, structName, propsStr)

			structsContent += fmt.Sprintf("// MessageType return the string telegram-type of %s \nfunc (%s *%s) MessageType() string {\n return \"%s\" }\n\n",
				structName, structNameCamel, structName, itemInfo.Name)

			// empty parms thats uses for multiply lines
			paramsStr := ""
			paramsDesc := ""
			assingsStr := ""

			for i, param := range itemInfo.Properties {
				propName := govalidator.UnderscoreToCamelCase(param.Name)
				propName = propName
				dataType, isPrimitive := convertDataType(param.Type, true)
				paramName := convertToArgumentName(param.Name)

				if isPrimitive || checkIsInterface(dataType, interfaces) {
					paramsStr += paramName + " " + dataType

				} else { // if is not a primitive, use pointers
					paramsStr += paramName + " *" + dataType
				}

				if i < len(itemInfo.Properties)-1 {
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
						tonCommon: tonCommon {Type: "%s"},
						%s
					}

					return &%sTemp
				}
				`, structName, structName, paramsDesc,
				structName, paramsStr, structName, structNameCamel,
				structName, itemInfo.Name, assingsStr, structNameCamel)

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
							tonCommon
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
			if checkIsInterface(itemInfo.RootName, interfaces) {
				rootName := itemInfo.RootName
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
			skip := false
			for _, name := range SkipMethodNames {
				if name == itemInfo.Name {
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			methodName := convertToExternalMethodName(itemInfo.Name)
			returnType := getStructName(itemInfo.RootName)
			returnTypeCamel := strings.ToLower(returnType[:1]) + returnType[1:]

			returnIsInterface := checkIsInterface(returnType, interfaces)

			asterike := "*"
			ampersign := "&"
			if returnIsInterface {
				asterike = ""
				ampersign = ""
			}

			paramsStr := ""
			clientCallStructAttrs := ""
			paramsDesc := ""

			// sort params to enshure the same params order in each generation
			sort.Slice(itemInfo.Properties, func(i, j int) bool {
				if (itemInfo.Properties[i].Name > itemInfo.Properties[j].Name){
					return false
				}
				return true
			})

			for i, param := range itemInfo.Properties {
				paramName := convertToArgumentName(param.Name)
				dataType, isPrimitive := convertDataType(param.Type, false)
				if isPrimitive || checkIsInterface(dataType, interfaces) {
					paramsStr += paramName + " " + dataType
					clientCallStructAttrs += fmt.Sprintf("%s %s `json:\"%s\"`\n", convertToExternalArgumentName(param.Name), dataType, param.Name)

				} else {
					paramsStr += paramName + " " + dataType
					clientCallStructAttrs += fmt.Sprintf("%s %s `json:\"%s\"`\n", convertToExternalArgumentName(param.Name), dataType, param.Name)
				}

				if i < len(itemInfo.Properties)-1 {
					paramsStr += ", "
				}
				paramsDesc += "\n// @param " + paramName + " " + param.Description
			}

			methodsContent += fmt.Sprintf(`
				// %s %s %s
				func (client *Client) %s(%s) (%s%s, error)`, methodName, itemInfo.Description, paramsDesc, methodName,
				paramsStr, asterike, returnType)

			paramsStr = ""
			for i, param := range itemInfo.Properties {
				paramName := convertToArgumentName(param.Name)

				paramsStr += fmt.Sprintf(`%s:   %s,`, convertToExternalArgumentName(param.Name), paramName)
				if i < len(itemInfo.Properties)-1 {
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

				for _, enum := range *enums {
					if enum.EnumType == enumType {
						for _, item := range enum.Items {
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
					result, err := client.executeAsynchronously(
						struct {
							Type string `+"`json:\"@type\"`"+`
							%s	
						}{
							Type: "%s",
							%s
						},
					)
	
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
					
					`, clientCallStructAttrs, itemInfo.Name, paramsStr, illStr,
					enumType, casesStr)

			} else {
				methodsContent += fmt.Sprintf(` {
					result, err := client.executeAsynchronously(
						struct {
							Type string `+"`json:\"@type\"`"+`
							%s	
						}{
							Type: "%s",
							%s
						},
					)
	
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
					
					`, clientCallStructAttrs, itemInfo.Name, paramsStr, illStr, returnTypeCamel,
					returnType, returnTypeCamel, ampersign, returnTypeCamel)
			}

		}
	}

	structsContent += "\n\n" + structUnmarshals
	return &structsContent, &methodsContent
}
