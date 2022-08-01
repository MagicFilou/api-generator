package translator

import (
	"strings"
)

func ToTypescript(datatype string) string {

	switch datatype {
	case "integer", "unix", "serial":
		return "number"
	case "uid", "":
		return "string"
	default:
		return datatype
	}
}

func ToGo(datatype string) string {

	switch datatype {
	case "integer", "serial":
		return "int"
	case "unix":
		return "int64"
	case "uid", "":
		return "string"
	case "boolean":
		return "bool"
	default:
		return datatype
	}
}

func ToSQL(datatype string) string {

	switch datatype {
	case "string":
		return "TEXT"
	case "uid":
		return "VARCHAR(36)"
	case "unix":
		return "BIGINT"
	case "":
		return "TEXT"
	default:
		return strings.ToUpper(datatype)
	}
}

func ToOpenAPI(datatype string) string {

	switch datatype {
	case "integer", "serial":
		return "number"
	case "unix":
		return "\"#/components/schemas/unix\""
	case "uid":
		return "\"#/components/schemas/id\""
	case "":
		return "string"
	default:
		return datatype
	}
}
