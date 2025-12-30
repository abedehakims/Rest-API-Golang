package main

import (
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	e := echo.New()
	// Deretan Controller
	controller.NewGetAllTodosController(e, db)     // Controller untuk Mendapatkan data todos
	controller.NewDeleteTodosController(e, db)     // Controller untuk Menghapus data todos
	controller.NewCreatePostTodosController(e, db) // Controller untuk Membuat data todos
	controller.NewUpdateDataTodosCotroller(e, db)  // Controller untuk Mengupdate data todos
	controller.NewCheckTodosController(e, db)      // Controller untuk Check status data todos
	e.Start(":8080")
}
