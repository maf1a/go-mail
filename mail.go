// copied from code.marcobeierer.com/utils/mail
package mail

import (
	"fmt"
	"log"
	"mime"
	"net/mail"
)

type Mail struct {
	from    *mail.Address
	to      *mail.Address
	subject string
	body    string
}

func NewMail(fromStr, toStr, subject, body string) (*Mail, error) {
	// parsing is necessary to handle special chars like ä, ö, ü; they could cause errors with some mail server

	from, err := mail.ParseAddress(fromStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	to, err := mail.ParseAddress(toStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// encoding is necessary for special chars like ä, ö, ü
	// if base64 is used for the message/body in the future, BEncoding could be used instead of QEncoding
	subject = mime.QEncoding.Encode("UTF-8", subject)

	return &Mail{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}, nil
}

func (qm *Mail) To() *mail.Address {
	return qm.to
}

func (qm *Mail) From() *mail.Address {
	return qm.from
}

func (qm *Mail) Message() string {
	return fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nContent-Type: text/plain; charset=\"UTF-8\"\n\n%s", qm.from, qm.to, qm.subject, qm.body)
}
