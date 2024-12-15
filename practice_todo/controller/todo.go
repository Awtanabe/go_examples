package controller

import (
	"log"
	"net/http"
	"practice_todo/models"
	"practice_todo/usecase"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITodoController interface {
	GetAllTodos(c echo.Context) error
	GetTodoById(c echo.Context) error
	CreateTodo(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type todoController struct {
	tu usecase.ITodoUsecase
}

func NewTodoController(tu usecase.ITodoUsecase) ITodoController {
	return &todoController{tu}
}

// func getUserIDFromJWT(c echo.Context) (uint, error) {
// 	user := c.Get("user")

// 	if user == nil {
// 		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User not found in context")
// 	}

// 	log.Print("user: Type", reflect.TypeOf(user))
// 	token, ok := user.(*jwt.Token)
// 	if !ok {
// 		log.Print("aaa", ok)
// 		return 0, echo.NewHTTPError(http.StatusInternalServerError, "Invalid token type")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		log.Print("bbb")

// 		return 0, echo.NewHTTPError(http.StatusInternalServerError, "Invalid claims type")
// 	}

// 	log.Printf("JWT Claims: %+v", claims)

// 	// user_id の型を確認
// 	rawUserID := claims["user_id"]
// 	userIDFloat, ok := claims["user_id"].(float64)
// 	log.Printf("user_id raw value: %v, type: %T", rawUserID, rawUserID)

// 	if !ok {
// 		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in claims")
// 	}

// 	userID := uint(userIDFloat)
// 	log.Printf("user_id: %v", userID)
// 	return userID, nil
// }

func getUserIDFromJWT(c echo.Context) (uint, error) {
	// JWT ミドルウェアからユーザー情報を取得
	user := c.Get("user")
	if user == nil {
		log.Print("User not found in context")
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User not found in context")
	}
	// トークンの型をチェック
	token, ok := user.(*jwt.Token)
	if !ok {
		log.Print("Invalid token type")
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "Invalid token type")
	}

	// クレームを取得
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Print("Invalid claims type")
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "Invalid claims type")
	}

	log.Printf("JWT Claims: %+v", claims)

	// クレームから user_id を取得
	rawUserID, ok := claims["user_id"]
	if !ok {
		log.Print("user_id not found in claims")
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in claims")
	}

	// 型に応じて user_id を変換
	var userID uint
	switch v := rawUserID.(type) {
	case float64:
		userID = uint(v)
	case int:
		userID = uint(v)
	case string:
		// 数字が文字列の場合
		parsedID, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			log.Printf("Failed to parse user_id string: %v", err)
			return 0, echo.NewHTTPError(http.StatusInternalServerError, "Invalid user_id format")
		}
		userID = uint(parsedID)
	default:
		log.Printf("Unexpected user_id type: %T", v)
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "Invalid user_id type")
	}

	log.Printf("user_id: %v", userID)
	return userID, nil
}

func (tc *todoController) GetAllTodos(c echo.Context) error {
    // JWTミドルウェアから取得

		userID, _ := getUserIDFromJWT(c)

			
	todos, err := tc.tu.GetAllTodos(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, todos)
}

func (tc *todoController) GetTodoById(c echo.Context) error {
	user, _ := c.Get("user").(map[string]interface{}) // JWTミドルウェアから取得
	userIdFloat, ok := user["user_id"].(float64) // JSONでは数値がfloat64として扱われる
if !ok {
    return c.JSON(http.StatusUnauthorized, echo.Map{"message": "Invalid user ID"})
}

// userId を uint に変換
userId := uint(userIdFloat)
	todoId, err := strconv.ParseUint(c.Param("todoId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid todo ID"})
	}
	todo, err := tc.tu.GetTodoById(userId, uint(todoId))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, todo)
}

func (tc *todoController) CreateTodo(c echo.Context) error {
	userID, _ := getUserIDFromJWT(c)

	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	if err := c.Validate(todo); err != nil {
		return err
	}

	todo.UserId = userID
	createdTodo, err := tc.tu.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, createdTodo)
}

func (tc *todoController) UpdateTodo(c echo.Context) error {
	userID, _ := getUserIDFromJWT(c)
	todoId, err := strconv.ParseUint(c.Param("todoId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid todo ID"})
	}
	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}
	updatedTodo, err := tc.tu.UpdateTodo(todo, userID, uint(todoId))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedTodo)
}

func (tc *todoController) DeleteTodo(c echo.Context) error {
	userID, _ := getUserIDFromJWT(c)
	todoId, err := strconv.ParseUint(c.Param("todoId"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid todo ID"})
	}
	if err := tc.tu.DeleteTodo(userID, uint(todoId)); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"message": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
