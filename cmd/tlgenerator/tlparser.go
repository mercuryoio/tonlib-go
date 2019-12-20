package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ClassInfo holds info of a Class in .tl file
type ClassInfo struct {
	Name        string          `json:"name"`
	Properties  []ClassProperty `json:"properties"`
	Description string          `json:"description"`
	RootName    string          `json:"rootName"`
	IsFunction  bool            `json:"isFunction"`
}

// ClassProperty holds info about properties of a class (or function)
type ClassProperty struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

// InterfaceInfo equals to abstract base classes in .tl file
type InterfaceInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// EnumInfo ...
type EnumInfo struct {
	EnumType string   `json:"enumType"`
	Items    []string `json:"description"`
}

func parseTlFile(tlReader *bufio.Reader) (error, *[]ClassInfo, *[]InterfaceInfo, *[]EnumInfo) {
	isFunctions := false
	var entityDesc string
	var paramDescs map[string]string
	var params map[string]string
	var interfaces []InterfaceInfo
	var enums []EnumInfo
	var entities []ClassInfo

	for {
		line, err := tlReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return fmt.Errorf("read file line error: %v", err), nil, nil, nil
		}
		if strings.Contains(line, "---functions---") {
			isFunctions = true
			continue
		}

		if strings.HasPrefix(line, "//@class ") {
			line = line[len("//@class "):]
			interfaceName := line[:strings.Index(line, " ")]
			line = line[len(interfaceName):]
			line = line[len(" @description "):]
			entityDesc = line[:len(line)-1]
			interfaceInfo := InterfaceInfo{
				Name:        interfaceName,
				Description: entityDesc,
			}
			interfaces = append(interfaces, interfaceInfo)
			enums = append(enums, EnumInfo{EnumType: interfaceName + "Enum"})

		} else if strings.HasPrefix(line, "//@description ") { // Entity description
			line = line[len("//@description "):]
			indexOfFirstSign := strings.Index(line, "@")

			entityDesc = line[:len(line)-1]
			if indexOfFirstSign != -1 {
				entityDesc = line[:indexOfFirstSign]
			}

			if indexOfFirstSign != -1 { // there is some parameter description inline, parse them
				line = line[indexOfFirstSign+1:]
				rd2 := bufio.NewReader(strings.NewReader(line))
				for {
					paramName, _ := rd2.ReadString(' ')
					if paramName == "" {
						break
					}
					paramName = paramName[:len(paramName)-1]
					paramDesc, _ := rd2.ReadString('@')
					if paramDesc == "" {
						paramDesc, _ = rd2.ReadString('\n')

						paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
						break
					}

					paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
				}
			}
		} else if entityDesc != "" && strings.HasPrefix(line, "//@") {
			line = line[len("//@"):]
			rd2 := bufio.NewReader(strings.NewReader(line))
			for {
				paramName, _ := rd2.ReadString(' ')
				if paramName == "" {
					break
				}
				paramName = paramName[:len(paramName)-1]
				paramDesc, _ := rd2.ReadString('@')
				if paramDesc == "" {
					paramDesc, _ = rd2.ReadString('\n')

					paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
					break
				}

				paramDescs[paramName] = paramDesc[:len(paramDesc)-1]
			}

		} else if !strings.HasPrefix(line, "//") && len(line) > 2 {
			entityName := line[:strings.Index(line, " ")]

			line = line[len(entityName)+1:]
			for {
				if strings.Index(line, ":") == -1 {
					break
				}
				paramName := line[:strings.Index(line, ":")]
				line = line[len(paramName)+1:]
				paramType := line[:strings.Index(line, " ")]
				params[paramName] = paramType
				line = line[len(paramType)+1:]
			}

			rootName := line[len("= ") : len(line)-2]

			var classProps []ClassProperty
			classProps = make([]ClassProperty, 0, 0)

			for paramName, paramType := range params {
				classProp := ClassProperty{
					Name:        paramName,
					Type:        paramType,
					Description: paramDescs[paramName],
				}
				classProps = append(classProps, classProp)
			}

			itemInfo := ClassInfo{
				Name:        entityName,
				Description: entityDesc,
				RootName:    rootName,
				Properties:  classProps,
				IsFunction:  isFunctions,
			}

			entities = append(entities, itemInfo)
			entityDesc = ""
			paramDescs = make(map[string]string)
			params = make(map[string]string)

			// update enum`s items list
			if !itemInfo.IsFunction {
				for i, enumInfo := range enums {
					if enumInfo.EnumType == itemInfo.RootName+"Enum" {
						enumInfo.Items = append(enumInfo.Items,
							strings.ToUpper(itemInfo.Name[0:1])+itemInfo.Name[1:])
						enums[i] = enumInfo
						break
					}
				}
			}
		}
	}
	return nil, &entities, &interfaces, &enums
}
