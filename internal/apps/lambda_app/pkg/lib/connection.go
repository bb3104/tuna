package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Dynamo_connect() *dynamo.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessKey := os.Getenv("DYNAMO_DB_ACCESS_KEY")
	secretKey := os.Getenv("DYNAMO_DB_SECRET_KEY")

	cred := credentials.NewStaticCredentials(accessKey, secretKey, "") // 最後の引数は[セッショントークン]

	db := dynamo.New(session.New(), &aws.Config{
		Credentials: cred,
		Region:      aws.String("ap-northeast-1"),
	})

	return db
}
