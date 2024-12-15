package main

// import (
// 	"fmt"
// 	"net/http"
// 	"os"

// 	db "practice_todo/db/my.cnf"

// 	"github.com/labstack/echo/v4"
// )

// type Todo struct {
// 	ID     uint `gorm:"primaryKey";autoIncrement`
// 	Title string
// 	Description string
// 	Status int `gorm:"type:int"`
// }

// func main() {
// 	db := db.NewDB()
// 	// db.AutoMigrate(&Todo{})

// 	e := echo.New()
// 	e.GET("/", func(c echo.Context) error {
// 		var todo = Todo{}
// 		db.First(&todo)
// 		return c.JSON(http.StatusOK, todo)
// 	})

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080" // デフォルトのポート番号を指定
// 	}
// 	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
// }
import (
	"practice_todo/controller"
	"practice_todo/db"
	"practice_todo/repository"
	"practice_todo/router"
	"practice_todo/usecase"
)

func main(){
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	todoRepository := repository.NewTodoRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	userController := controller.NewUserController(userUsecase)
	todoController := controller.NewTodoController(todoUsecase)
	mailUsecase := usecase.NewMailUsecase()
	mailController := controller.NewMailController(mailUsecase)
	e := router.NewRouter(userController, todoController, mailController)
	e.Logger.Fatal(e.Start(":8080"))
}