package interface_info

import (
	"github.com/ByteBam/thirftbam/biz/model"
	"github.com/ByteBam/thirftbam/pkg/parser"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/thriftgo/utils"
	"sync"
)

func GetParameter(structures *sync.Map, fields []*parser.Field) (*[]byte, error) {
	var parameters []*model.Structure

	for _, field := range fields {
		parameter := getParameter(structures, field)
		parameters = append(parameters, parameter)
	}
	bytes, err := sonic.Marshal(parameters)
	if err != nil {
		return nil, err
	}
	return &bytes, nil
}

func getParameter(structures *sync.Map, field *parser.Field) *model.Structure {
	var parameter model.Structure

	parameter.Name = field.GetName()
	parameter.Requiredness = GetRequierdness(field.GetRequiredness())
	parameter.Category = GetCategory(field.GetType().GetName()).String()
	ok := utils.IsBasic(field.GetType().GetName())
	if !ok {
		parameterField := getParameterFields(structures, field.GetType().GetName())
		parameter.Fields = append(parameter.Fields, parameterField...)
	}
	return &parameter
}

func getParameterFields(structures *sync.Map, field string) []*model.Structure {
	var fields []*model.Structure
	loads, ok := structures.Load(field)
	if !ok {
		return nil
	}
	for _, load := range loads.(*parser.StructLike).GetFields() {
		var param model.Structure
		param.Name = load.GetName()
		param.Requiredness = GetRequierdness(load.GetRequiredness())
		param.Category = GetCategory(load.GetType().GetName()).String()
		if !utils.IsBasic(load.GetType().GetName()) {
			param.Fields = append(param.Fields, getParameter(structures, load))
		}
		fields = append(fields, &param)
	}
	return fields
}
