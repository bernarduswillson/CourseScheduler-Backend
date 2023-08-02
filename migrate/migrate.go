package migrate

import (
	"singleservice/initializers"
	model "singleservice/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Migrate(){
	fmt.Println("Migrating........................")
	if !initializers.DB.Migrator().HasTable(&model.Matakuliah{}) {
		initializers.DB.Migrator().CreateTable(&model.Matakuliah{})
	}
}
