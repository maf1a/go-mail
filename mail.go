// copied from code.marcobeierer.com/utils/mail
package mail

import (
	"fmt"
	"log"
	"mime"
	"net/mail"
	"time"
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
	// https://tools.ietf.org/html/rfc5322#section-3.3 defines own date format, but is identical to RFC1123Z
	// CRLF (\r\n) according to https://tools.ietf.org/html/rfc5322#section-2.3
	return fmt.Sprintf("Date: %s\r\nFrom: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s", time.Now().Format(time.RFC1123Z), qm.from, qm.to, qm.subject, qm.body)
}
