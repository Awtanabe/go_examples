package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendMail() {
	// SES の設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),       // SES リージョン
		Endpoint:    aws.String("http://localhost:4566"), // LocalStack のエンドポイント
		Credentials: credentials.NewStaticCredentials("ACCESS_KEY", "SECRET_KEY", ""),
	})

	if err != nil {
		log.Fatalf("failed to create session: %v", err)
	}

	// SES クライアントを作成
	svc := ses.New(sess)

	// メールの送信内容を設定
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String("panser098@gmail.com"), // 送信先メールアドレス
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String("This is a test email sent via LocalStack SES."),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Test Email"),
			},
		},
		Source: aws.String("sender@example.com"), // 送信元メールアドレス
	}

	// メールを送信
	result, err := svc.SendEmail(input)
	if err != nil {
		log.Fatalf("failed to send email: %v", err)
	}

	fmt.Printf("Email sent successfully: %v\n", result)
}


func main() {

	SendMail()
}