package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"strings"
)

func checkIsInterface(input string, interfaces *[]InterfaceInfo) bool {
	for _, interfaceInfo := range *interfaces {
		if interfaceInfo.Name == input {
			return true
		}
	}

	return false
}

func convertFromDots(paramName string) string {
	splited := strings.Split(paramName, ".")
	if len(splited) < 2 {
		return paramName
	}

	for i, sp := range splited {
		if i == 0 {
			paramName = sp
		} else if len(sp) > 0 {
			if len(sp) == 1 {
				paramName += strings.ToUpper(sp[0:1])
			} else {
				paramName += strings.ToUpper(sp[0:1]) + sp[1:]
			}
		}
	}
	return paramName

}

func convertToArgumentName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = strings.ToLower(paramName[0:1]) + paramName[1:]
	paramName = strings.Replace(paramName, "type", "typeParam", 1)
	paramName = convertFromDots(paramName)

	return paramName
}

func convertToExternalArgumentName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = strings.ToUpper(paramName[0:1]) + paramName[1:]
	paramName = convertFromDots(paramName)

	return paramName
}

func convertToExternalMethodName(input string) string {
	paramName := strings.ToUpper(input[0:1]) + input[1:]
	paramName = convertFromDots(paramName)

	return paramName
}

func convertToInternalAtributeName(input string) string {
	paramName := govalidator.UnderscoreToCamelCase(input)
	paramName = strings.ToLower(paramName[0:1]) + paramName[1:]
	paramName = convertFromDots(paramName)

	return paramName
}

func getStructName(input string) string {
	paramName := strings.ToUpper(input[0:1]) + input[1:]
	paramName = convertFromDots(paramName)

	return paramName
}

func convertDataType(input string) (string, bool) {
	propType := ""
	isPrimitiveType := true
	// cleanup from "="
	input = strings.Trim(input, "=")

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
			propType = "[][]" + getStructName(input[len("[][]"):])
		} else if strings.HasPrefix(input, "[]") {
			propType = "[]" + getStructName(input[len("[]"):])
		} else {
			propType = getStructName(input)
			isPrimitiveType = false
		}
	}

	return propType, isPrimitiveType
}

func ChangeType(paramName string, fromType string, toType string) string {
	if fromType == "SecureBytes" && toType == "string" {
		return fmt.Sprintf("base64.StdEncoding.EncodeToString(%s)", paramName)
	}
	return paramName

}
