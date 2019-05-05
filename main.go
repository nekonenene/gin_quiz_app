package main

import (
	"github.com/nekonenene/gin_quiz_app/model"
	"github.com/nekonenene/gin_quiz_app/registry"
	"github.com/nekonenene/gin_quiz_app/router"
)

func main() {
	registry.Init()
	defer registry.DB.Close()

	model.AutoMigrate()
	router.InitRouter()
}
