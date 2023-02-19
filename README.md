# WebSocket 通信

## WebSocket 概要
Webにおいて双方向通信を低コストで行うためのプロトコル。
SNSなどではリアルタイムの画面更新が必須だが、HTTPでは難しい。ここでWebSocketが活躍する。

### 特徴
- Server Push
  - コネクションを確立すれば、サーバからも通信を行うことができる
- 通信量削減
  - 一度確立したコネクション上で通信を行うためHTTPのように、通信のたびにコネクションを作らない
  - ヘッダのサイズが最小2byte、最大14byteと小さい
  
WebSocketプロトコルではまずハンドシェイクを行う。

### ハンドシェイク


リクエスト: HTTPのUpgradeヘッダを使用し、プロトコルの変更を行う。

レスポンス: ステータスコード101「Switching Protocols」が返る。
コネクションが確立し、これ以降はHTTPではなくWebSocketのプロトコルで通信が行われる。

送信データはフレームという単位

# 選定技術
言語: Go

Web Framework: Echo

# install
1. `brew install go`
1. `go mod init myapp`
myapp以外でもok
1. `go get github.com/labstack/echo/v4`
1. `go get github.com/labstack/echo/v4/middleware@v4.10.0`


# build & run
`go run main.go`
