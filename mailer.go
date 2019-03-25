// copied from code.marcobeierer.com/utils/mail
package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

type Mailer struct {
	host               string
	port               int
	username           string
	password           string
	InsecureSkipVerify bool
}

func NewMailer(host string, port int, username, password string) *Mailer {
	return &Mailer{
		host:               host,
		port:               port,
		username:           username,
		password:           password,
		InsecureSkipVerify: false,
	}
}

type Mailable interface {
	To() *mail.Address
	From() *mail.Address
	Bcc() (*mail.Address, bool)
	Message() string
}

func (qm *Mailer) SendMail(mail Mailable) error {
	address := fmt.Sprintf("%s:%d", qm.host, qm.port)

	connection, err := smtp.Dial(address)
	if err != nil {
		log.Println(err)
		return err
	}
	defer connection.Close()

	if err = connection.Hello(qm.host); err != nil {
		log.Println(err)
		return err
	}

	supportsStartTLS, _ := connection.Extension("STARTTLS")

	if supportsStartTLS {
		config := &tls.Config{InsecureSkipVerify: qm.InsecureSkipVerify}
		if err = connection.StartTLS(config); err != nil {
			log.Println(err)
			return err
		}
	} else {
		log.Println("STARTTLS is not supported")
	}

	if !(qm.username == "" && qm.password == "") {
		auth := smtp.PlainAuth("", qm.username, qm.password, qm.host)
		if ok, _ := connection.Extension("AUTH"); ok {
			if err = connection.Auth(auth); err != nil {
				log.Println(err)
				return err
			}
		}
	}

	// the visible `from` with name and address is set in mail.go:Message()
	if err := connection.Mail(mail.From().Address); err != nil {
		log.Println(err)
		return err
	}

	// the visible `to` with name and address is set in mail.go:Message()
	if err := connection.Rcpt(mail.To().Address); err != nil {
		log.Println(err)
		return err
	}

	// bcc is handled as normal rcpt but not set in the message and thus invisible
	if bcc, set := mail.Bcc(); set {
		if err := connection.Rcpt(bcc.Address); err != nil {
			log.Println(err)
			return err
		}
	}

	data, err := connection.Data()
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = fmt.Fprint(data, mail.Message())
	if err != nil {
		log.Println(err)
		return err
	}

	err = data.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	if err := connection.Quit(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
