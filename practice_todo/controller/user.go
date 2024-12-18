package controller

import (
	"fmt"
	"net/http"
	"os"
	"practice_todo/models"
	"practice_todo/usecase"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	GetUser(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
	
}

//usecaseの依存関係をcontrollerに注入
func NewUserController(uu usecase.IUserUsecase) IUserController{
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := models.User{}
	//c.Bind() メソッドはHTTPリクエストのボディデータを受け取り、指定した構造体にデータを関連付けます。
	//HTTPリクエストから送信されたデータをプログラム内のデータ構造にコピーする作業
	if err := c.Bind(&user); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusCreated,userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// tokenを作成して 下記でcookieを設定する
	tokenString, err := uc.uu.Login(user)
	fmt.Println("Generated JWT token:", tokenString)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	//JWTトークンをcookieに設定する
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}



func (uc *userController) GetUser(c echo.Context) error {
	idParam := c.QueryParam("id")
	userID, err := strconv.ParseUint(idParam, 10, 32)
	user, err := uc.uu.GetUser(uint(userID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}


func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}