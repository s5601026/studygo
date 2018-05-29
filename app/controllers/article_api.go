package controllers

import (
	"github.com/masato-kataoka/studygo/app/models"
	"github.com/revel/revel"
)

type ArticleApi struct {
	*revel.Controller
}

func (c ArticleApi) GetArticles() revel.Result {

	// articleモデルを利用
	articles := []models.Article{}
	// Idが降順になるように取得
	DB.Order("id desc").Find(&articles)

	response := JsonResponse{}
	response.Response = articles // 結果を格納する

	return c.RenderJSON(response)
}

func (c ArticleApi) GetArticle() revel.Result {

	// ルーティングで設定したurlに含まれる :id とかの部分はc.Params.Route.Getで取得
	id := c.Params.Route.Get("id")

	article := []models.Article{}
	// DB.Firstの第二引数にそのidを渡してあげると第一引数のモデルからidが一致するデータを検索してくれる
	DB.First(&article, id)

	response := JsonResponse{}
	response.Response = article

	return c.RenderJSON(response)
}

func (c ArticleApi) PostArticle() revel.Result {

	// articleモデルに値を格納
	article := &models.Article{
		// x-www-form-urlencodeで飛んできたデータはc.Params.Form.Getで受け取れます
		Title: c.Params.Form.Get("title"),
		Text:  c.Params.Form.Get("text"),
	}

	// DBで保存
	DB.Create(article)

	response := JsonResponse{}
	// この時点でarticleにはidが振られているのでそのまま返してあげる
	response.Response = article

	return c.RenderJSON(response)
}

func (c ArticleApi) PutArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "put article"

	return c.RenderJSON(response)
}

func (c ArticleApi) DeleteArticle() revel.Result {

	response := JsonResponse{}
	response.Response = "delete article"

	return c.RenderJSON(response)
}
