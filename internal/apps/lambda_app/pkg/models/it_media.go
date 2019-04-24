package models

import (
	"encoding/xml"
	"fmt"
	"github.com/bb3104/sample_dynamodb/internal/apps/lambda_app/pkg/lib"
	"time"
)

type ItmediaXML struct {
	Itmedias []struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Date        string `xml:"date"`
		Description string `xml:"discription"`
	} `xml:"item"`
}

func (i ItmediaXML) ParseRss(url string) {

	data := lib.HttpGet(url)

	db := lib.Dynamo_connect()
	table := db.Table("engineee")

	result := ItmediaXML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	for _, itmedia := range result.Itmedias {
		datetime, _ := time.Parse(time.RFC3339, itmedia.Date)
		i := RssArticle{RssTypeName: "itmedia", Title: itmedia.Title, Description: itmedia.Description, Url: itmedia.Link, PublishedAt: datetime.Unix(), CreatedAt: time.Now().Unix()}

		if err := table.Put(i).Run(); err != nil {
			fmt.Println("err")
			panic(err.Error())
		}
		// fmt.Printf("%v\n", datetime.Format("2006/01/02 15:04:05"))
		// fmt.Printf("%v\n", itmedia.Title)
		// fmt.Printf("%v\n", itmedia.Link)
		// fmt.Printf("%v\n", itmedia.Description)
	}
	fmt.Println("itmedia ok")

}
