package controllers

type Context interface {
	Param(string) string /*c.Param()でパスパラメータを名前で返す*/
	JSON(int, interface{}) error /*c.JSON()ステータスコード付きのJSONレスポンスを送信*/
	Bind(interface{}) error /*c.Bind()でリクエストボディからは更新データを取得*/
}