package interface_info

import (
	"github.com/ByteBam/thirftbam/pkg/parser"
)

func ParserUrl(annotation *parser.Annotation) string {
	var url string
	for _, v := range annotation.GetValues() {
		url += v
	}
	return url
}

func ParserMethodAndUrl(function *parser.Function) (string, string, error) {
	var url string
	for _, annotation := range function.GetAnnotations() {
		switch annotation.GetKey() {
		case "api.get":
			return "GET", ParserUrl(annotation), nil
		case "api.post":
			return "POST", ParserUrl(annotation), nil
		case "api.put":
			return "PUT", ParserUrl(annotation), nil
		case "api.delete":
			return "DELETE", ParserUrl(annotation), nil
		case "api.patch":
			return "PATCH", ParserUrl(annotation), nil
		case "api.options":
			return "OPTIONS", ParserUrl(annotation), nil
		case "api.head":
			return "HEAD", ParserUrl(annotation), nil
		case "api.any":
			return "ANY", ParserUrl(annotation), nil
		}
	}
	return "", url, nil
}
