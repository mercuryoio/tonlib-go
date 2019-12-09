package main

import (
	"github.com/asaskevich/govalidator"
	"strings"
)

func checkIsInterface(input string, interfaces *[]InterfaceInfo) bool {
	for _, interfaceInfo := range *interfaces {
		if interfaceInfo.Name == input || replaceKeyWords(interfaceInfo.Name) == input {
			return true
		}
	}

	return false
}

func convertToArgumentName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = replaceKeyWords(paramName)
	paramName = strings.ToLower(paramName[0:1]) + paramName[1:]
	paramName = strings.Replace(paramName, "type", "typeParam", 1)

	return paramName
}

func convertDataType(input string) (string, bool) {
	propType := ""
	isPrimitiveType := true

	if strings.HasPrefix(input, "vector") {
		input = "[]" + input[len("vector<"):len(input)-1]
		isPrimitiveType = true
	}
	if strings.HasPrefix(input, "[]vector") {
		input = "[][]" + input[len("[]vector<"):len(input)-1]

	}
	if strings.Contains(input, "string") || strings.Contains(input, "int32") ||
		strings.Contains(input, "int64") {
		propType = strings.Replace(input, "int64", "JSONInt64", 1)

	} else if strings.Contains(input, "Bool") {
		propType = strings.Replace(input, "Bool", "bool", 1)

	} else if strings.Contains(input, "double") {
		propType = strings.Replace(input, "double", "float64", 1)

	} else if strings.Contains(input, "int53") {
		propType = strings.Replace(input, "int53", "int64", 1)

	} else if strings.Contains(input, "bytes") {
		propType = strings.Replace(input, "bytes", "[]byte", 1)

	} else {
		if strings.HasPrefix(input, "[][]") {
			propType = "[][]" + strings.ToUpper(input[len("[][]"):len("[][]")+1]) + input[len("[][]")+1:]
		} else if strings.HasPrefix(input, "[]") {
			propType = "[]" + strings.ToUpper(input[len("[]"):len("[]")+1]) + input[len("[]")+1:]
		} else {
			propType = strings.ToUpper(input[:1]) + input[1:]
			isPrimitiveType = false
		}
	}

	propType = replaceKeyWords(propType)

	return propType, isPrimitiveType
}
