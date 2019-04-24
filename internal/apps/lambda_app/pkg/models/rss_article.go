package models

type RssArticle struct {
	RssTypeName string `dynamo:"rss_type_name"`
	Title       string `dynamo:"title"`
	Description string `dynamo:"description"`
	Image       string `dynamo:"image"`
	Url         string `dynamo:"url"`
	PublishedAt int64  `dynamo:"published_at"`
	CreatedAt   int64  `dynamo:"created_at"`
}
