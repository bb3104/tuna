package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bb3104/tuna/internal/apps/lambda_app/pkg/models"
)

type ParseRss interface {
	ParseRss(url string)
}

func main() {
	lambda.Start(create_rss)
}

func create_rss() {

	var hatena_xml models.HatenaXML
	var itmedia_xml models.ItmediaXML
	var rss ParseRss

	rss = hatena_xml
	rss.ParseRss("http://b.hatena.ne.jp/hotentry/it.rss")

	rss = itmedia_xml
	rss.ParseRss("https://rss.itmedia.co.jp/rss/1.0/news_bursts.xml")
}
