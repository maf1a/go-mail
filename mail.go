// copied from code.marcobeierer.com/utils/mail
package mail

import (
	"fmt"
	"log"
	"mime"
	"net/mail"
)

type Mail struct {
	from      *mail.Address
	to        *mail.Address
	unsafeBcc *mail.Address
	subject   string
	body      string
}

func NewMail(fromStr, toStr, subject, body string) (*Mail, error) {
	return NewMailWithBcc(fromStr, toStr, "", subject, body)
}

func NewMailWithBcc(fromStr, toStr, bccStr, subject, body string) (*Mail, error) {
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

	mailx := &Mail{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}

	if bccStr != "" {
		bcc, err := mail.ParseAddress(bccStr)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		mailx.unsafeBcc = bcc
	}

	return mailx, nil
}

func (qm *Mail) To() *mail.Address {
	return qm.to
}

func (qm *Mail) From() *mail.Address {
	return qm.from
}

func (qm *Mail) Bcc() (*mail.Address, bool) {
	return qm.unsafeBcc, qm.unsafeBcc != nil
}

func (qm *Mail) Message() string {
	return fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nContent-Type: text/plain; charset=\"UTF-8\"\n\n%s", qm.from, qm.to, qm.subject, qm.body)
}
