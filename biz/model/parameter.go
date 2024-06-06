package model

type Category int64

const (
	CategoryInteger Category = iota
	CategoryNumber
	CategoryString
	CategoryBoolean
	CategoryArray
	CategoryObject
)

func (c Category) String() string {
	switch c {
	case CategoryInteger:
		return "integer"
	case CategoryNumber:
		return "number"
	case CategoryString:
		return "string"
	case CategoryBoolean:
		return "boolean"
	case CategoryArray:
		return "array"
	case CategoryObject:
		return "object"
	default:
		return "unknown"
	}
}

type FieldType int64

const (
	FieldTypeOptional FieldType = iota
	FieldTypeRequired
)

func (t FieldType) String() string {
	switch t {
	case FieldTypeOptional:
		return "optional"
	case FieldTypeRequired:
		return "required"
	default:
		return "unknown"
	}
}

type Structure struct { // 请求参数
	Category     string       `json:"category"`     // 类型
	Name         string       `json:"name"`         // 名称
	Requiredness FieldType    `json:"requiredness"` // 是否必须
	Fields       []*Structure `json:"fields"`       // 字段
}

type Field struct {
	Name         string     `json:"name"`         // 字段名称
	Requiredness FieldType  `json:"requiredness"` // 是否必须
	Type         *Structure `json:"type"`         // 字段类型
}
