package customValidator

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	// validation "github.com/go-ozzo/ozzo-validation"
)


type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.validator.Struct(i); err != nil {
    // Optionally, you could return the error to give each route more control over the status code
    return echo.NewHTTPError(http.StatusBadRequest, err.Error())
  }
  return nil
}