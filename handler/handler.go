package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
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

func JPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "http://example.com")
	}
}

func WebSocket() echo.HandlerFunc {
	var (
		upgrader = websocket.Upgrader{}
	)
	return func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		for {
			// Write
			err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			_, msg, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}
}

// HandleHelloGet は hello2 のGet時のHTMLデータ生成処理を行います。
func HandleHelloGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		greetingto := c.QueryParam("greetingto")
		return c.Render(http.StatusOK, "hello", greetingto)
	}
}

func HandleHelloPost() echo.HandlerFunc {
	return func(c echo.Context) error {
		greetingto := c.FormValue("greetingto")
		return c.Render(http.StatusOK, "hello", greetingto)
	}
}

// HandleHelloFormGet は /hello_form のGet時のHTMLデータ生成処理を行います。
func HandleHelloFormGet() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello_form", nil)
	}
}
