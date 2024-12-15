package controller

import (
	"net/http"
	"practice_todo/mail"
	"practice_todo/usecase"

	"github.com/labstack/echo/v4"
)

type IMailController interface {
	SendMail(c echo.Context) error
}

type mailController struct {
	mu usecase.IMailUsecase
}

func NewMailController(mu usecase.IMailUsecase) mailController {
	return mailController {mu}
}


func (mc mailController) SendMail(c echo.Context) error {
	var mail = mail.MessageParam{}

	if err := c.Bind(&mail); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "送れたよ")
	}

	if err := mc.mu.Send(mail); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "送れたよ")
}
