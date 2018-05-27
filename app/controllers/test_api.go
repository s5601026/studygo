package controllers

// revel本体を読み込む
import (
	"github.com/revel/revel"
)

// 構造体の初期化
type TestApi struct {
	*revel.Controller
}

// Index関数
// conf/routesに TestApi.Index() と書いたのでこれが実装される
func (c TestApi) Index() revel.Result {

	// json.goで指定したJsonResponcseを使う
	response := JsonResponse{}

	// JsonResponseで指定したResponseの中にデータを入れてみる
	response.Response = "Hello"

	// jsonとしてレスポンスを返す
	return c.RenderJSON(response)
}
