package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

// echo.Contextという型の引数cを受け取り、error型を返す
// echo.ContextはEchoフレームワークにおいて、HTTPリクエストやレスポンスを表すために使われる型
func handleWebSocket(c echo.Context) error {
	// wsは*websocket.Conn型のポインタ変数で、WebSocket接続のための構造体
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// 初回のメッセージを送信
		// websocket.Message.Send 関数はエラーが発生するかどうかを返すため、その戻り値をerr変数に代入する
		// エラーがなければ、errにはnilが代入
		err := websocket.Message.Send(ws, "Server: Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			// Client からのメッセージを読み込む
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			// Client に返すメッセージを作成し送信
			err := websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	// Logger()は、Echoフレームワークが提供するミドルウェアの一つで、HTTPリクエストをログに出力するために利用される
	// curl localhost:8080 などすればログが出力されていることを確認できる
	e.Use(middleware.Logger())
	// Webサーバーのルートディレクトリに"client"ディレクトリをマウントし、"/"にアクセスされたときは、そのディレクトリ内のファイルを返す
	e.Static("/", "client")
	// "/ws"へのリクエストがWebSocketリクエストとして扱われ、handleWebSocket関数が呼び出される
	e.GET("/ws", handleWebSocket)
	e.Logger.Fatal(e.Start(":8080"))
}
