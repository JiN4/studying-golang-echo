package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")                         //プレースホルダusernameの値取り出し
		return c.String(http.StatusOK, "Hello World "+username) //HTTPステータスコードが200と共に文字列”Hello World”とusernameを送信する
	}
}

func JsonPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		jsonMap := map[string]string{
			"foo":  "bar",
			"hoge": "fuga",
		}
		return c.JSON(http.StatusOK, jsonMap)
	}
}
