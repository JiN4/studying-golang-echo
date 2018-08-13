package main

import (
	"studying-golang-echo/handler"
	"studying-golang-echo/interceptor"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger()) //HTTPのリクエストのログを出力
	e.Use(middleware.Recover())
	//e.Use(interceptor.BasicAuth())

	// ルーティング
	e.GET("/hello", handler.MainPage())                        //文字出力のみ
	e.GET("/app", handler.MainPage(), interceptor.BasicAuth()) //Basic認証
	// サーバー起動
	e.Start(":8000") //ポート番号8000指定
}
