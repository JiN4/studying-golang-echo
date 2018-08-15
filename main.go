package main

import (
	"html/template"
	"io"
	"studying-golang-echo/handler"
	"studying-golang-echo/interceptor"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// レイアウト適用済のテンプレートを保存するmap
var templates map[string]*template.Template

// Template はHTMLテンプレートを利用するためのRenderer Interfaceです。
type Template struct {
}

// Render はHTMLテンプレートにデータを埋め込んだ結果をWriterに書き込みます。
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return templates[name].ExecuteTemplate(w, "layout.html", data) //(ライター/テンプレート名/埋め込むデータの構造体)
}

// 初期化を行います。
func init() {
	loadTemplates()
}

// 各HTMLテンプレートに共通レイアウトを適用した結果を保存（初期化時に実行）。
func loadTemplates() {
	var baseTemplate = "templates/layout.html"
	templates = make(map[string]*template.Template)
	templates["hello"] = template.Must( //正当性のチェック
		template.ParseFiles(baseTemplate, "./templates/hello.html")) //layout.htmlにhello.htmlを埋め込んだものをtemplates["index"]に格納
	templates["hello_form"] = template.Must(
		template.ParseFiles(baseTemplate, "templates/hello_form.html"))
}

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// テンプレートを利用するためのRendererの設定
	t := &Template{}
	e.Renderer = t

	// ミドルウェアを設定
	e.Use(middleware.Logger()) //HTTPのリクエストのログを出力
	e.Use(middleware.Recover())

	// 静的ファイルのパスを設定
	/*e.Static("/public/css/", "./public/css")
	e.Static("/public/js/", "./public/js/")
	e.Static("/public/img/", "./public/img/")*/

	// ルーティング
	e.GET("/hello", handler.HelloPage())                          //文字出力のみ
	e.GET("/basic", handler.HelloPage(), interceptor.BasicAuth()) //Basic認証
	e.GET("/json", handler.JsonPage())                            //json
	//e.GET("/websocket", handler.WebSocket())

	e.GET("/hello2", handler.HandleHelloGet())
	e.POST("/hello2", handler.HandleHelloPost())
	e.GET("/hello_form", handler.HandleHelloFormGet())

	// サーバー起動
	e.Start(":8000") //ポート番号8000指定
}
