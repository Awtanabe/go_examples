package mail

import "log"

type IMailer interface {
	Send() error
}

type MessageParam struct {
	Message string
	To string
	From string
}

type MailInstance struct {
	Message string
	To string
	From string
}

func (mi MailInstance) Send() error {

	log.Print("メール送信", mi.Message, mi.To, mi.From)
	// メール送信処理
	return nil
}