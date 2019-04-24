package main

import (
	"github.com/bb3104/sample_dynamodb/internal/apps/lambda_app/pkg/models"
)

func main() {

	var hatena_xml models.HatenaXML
	var itmedia_xml models.ItmediaXML
	var rss ParseRss

	rss = hatena_xml
	rss.ParseRss("http://b.hatena.ne.jp/hotentry/it.rss")

	rss = itmedia_xml
	rss.ParseRss("https://rss.itmedia.co.jp/rss/1.0/news_bursts.xml")

}

type ParseRss interface {
	ParseRss(url string)
}
