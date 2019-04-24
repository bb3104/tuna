package models

import (
	"encoding/xml"
	"fmt"
	"github.com/bb3104/sample_dynamodb/internal/apps/lambda_app/pkg/lib"
	"time"
)

type HatenaXML struct {
	Bookmarks []struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Date        string `xml:"date"`
		Description string `xml:"description"`
		Count       int    `xml:"bookmarkcount"`
		ImageUrl    string `xml:"imageurl"`
	} `xml:"item"`
}

func (h HatenaXML) ParseRss(url string) {

	db := lib.Dynamo_connect()
	table := db.Table("engineee")

	data := lib.HttpGet(url)

	result := HatenaXML{}
	err := xml.Unmarshal([]byte(data), &result)

	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, bookmark := range result.Bookmarks {
		datetime, _ := time.Parse(time.RFC3339, bookmark.Date)
		h := RssArticle{RssTypeName: "hatena_bookmark", Title: bookmark.Title, Description: bookmark.Description, Url: bookmark.Link, Image: bookmark.ImageUrl, PublishedAt: datetime.Unix(), CreatedAt: time.Now().Unix()}

		if err := table.Put(h).Run(); err != nil {
			fmt.Println("err")
			panic(err.Error())
		}
	}
	fmt.Println("hatena ok")
}
