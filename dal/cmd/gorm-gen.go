package main

import (
	"gorm.io/gen"
)

import (
	systemModel "github.com/MuserQuantity/gin-project-example-without-users/dal/model/system"
	"github.com/MuserQuantity/gin-project-example-without-users/global"
)

func genBasic() {
	db := global.SysDb
	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/dao/basic",
		Mode:    gen.WithoutContext,
	})
	g.UseDB(db)
	g.ApplyBasic(
		systemModel.TestStruct{},
	)
	g.Execute()
}

func main() {
	genBasic()
}
