package main

import (
	"flag"
	"github.com/ByteBam/thirftbam/biz/repository"
	"github.com/ByteBam/thirftbam/pkg/utils/config"
	"github.com/ByteBam/thirftbam/pkg/utils/log"
	"gorm.io/gen"
)

func main() {
	var envConf = flag.String("conf", "biz/config/config.yaml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger := log.NewLog(conf)
	db := repository.NewDB(conf, logger)

	g := gen.NewGenerator(gen.Config{
		OutPath:      "biz/repository/query",
		ModelPkgPath: "biz/model",

		Mode: gen.WithQueryInterface,

		FieldNullable: false,

		FieldCoverable: false,

		FieldSignable: true,

		FieldWithIndexTag: true,

		FieldWithTypeTag: true,

		WithUnitTest: false,
	})

	g.UseDB(db)

	// dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
	// 	"int":     func(columnType gorm.ColumnType) (dataType string) { return "uint64" },
	// 	"tinyint": func(columnType gorm.ColumnType) (dataType string) { return "int8" },
	// }
	// g.WithDataTypeMap(dataMap)
	g.ApplyBasic(
		g.GenerateModel("branch"),
		g.GenerateModel("user"),
		g.GenerateModel("module_info"),
		g.GenerateModel("interface_info"),
	)

	g.Execute()
}
