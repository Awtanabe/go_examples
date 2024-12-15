package usecase

import (
	"practice_todo/mail"
)

type IMailUsecase interface {
	Send(mp mail.MessageParam) error
}

type mailUsecase struct {
}

func NewMailUsecase() IMailUsecase {
	return mailUsecase{}
}

func (mu mailUsecase) Send(mp mail.MessageParam) error {
	mailer := &mail.MailInstance{
		Message: mp.Message,
		To:      mp.To,
		From:    mp.From,
	}
	if err := mailer.Send(); err != nil {
		return err
	}
	return nil
}