package parser

var basicMap = map[string]bool{
	"i8":     true,
	"i16":    true,
	"i32":    true,
	"i64":    true,
	"double": true,
	"string": true,
	"byte":   true,
	"binary": true,
	"bool":   true,
	"set":    true,
	"list":   true,
	"map":    true,
}

func IsBasic(name string) bool {
	return basicMap[name]
}
