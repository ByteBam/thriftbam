package interface_info

import (
	"github.com/ByteBam/thirftbam/biz/model"
	"github.com/ByteBam/thirftbam/pkg/parser"
)

func GetCategory(c string) model.Category {
	switch c {
	case "bool":
		return model.CategoryBoolean
	case "i8", "i16", "i32", "i64":
		return model.CategoryInteger
	case "double":
		return model.CategoryNumber
	case "binary", "map", "list", "set":
		return model.CategoryArray
	case "string", "byte":
		return model.CategoryString
	default:
		return model.CategoryObject
	}
}

func GetRequierdness(requiredness parser.FieldType) model.FieldType {

	if requiredness == parser.FieldType_Default || requiredness == parser.FieldType_Required {
		return model.FieldTypeRequired
	}

	return model.FieldTypeOptional
}
