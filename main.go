package main

import (
	"studying-golang-echo/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echoのインスタンス作る
	e := echo.New()

	// ミドルウェア
	e.Use(middleware.Logger()) //HTTPのリクエストのログを出力
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())
	e.Get("/hello/:username", handle.MainPage()) //セミコロンの場所がプレースホルダ

	// サーバー起動
	e.Start(":8000") //ポート番号8000指定
}
