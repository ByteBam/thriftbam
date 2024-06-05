package interface_info

import (
	"errors"
	"github.com/ByteBam/thirftbam/pkg/parser"
	"github.com/bytedance/sonic"
	"sync"
)

func GetParameter(structure *sync.Map, function *parser.Function) (*[]byte, error) {
	var structures []*parser.StructLike
	for _, field := range function.GetArguments() {
		param, ok := structure.Load(field.GetType().GetName())
		if !ok {
			return nil, errors.New("no such struct")
		}
		structures = append(structures, param.(*parser.StructLike))
	}
	paramBytes, err := sonic.Marshal(structures)
	if err != nil {
		return nil, err
	}
	return &paramBytes, nil
}
