// copied from code.marcobeierer.com/utils/mail
package mail

import (
	"crypto/tls"
	"fmt"
	"log"
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
	To() string
	From() string
	Message() string
}

func (qm *Mailer) SendMail(mail Mailable) bool {
	address := fmt.Sprintf("%s:%d", qm.host, qm.port)

	connection, err := smtp.Dial(address)
	if err != nil {
		log.Println(err)
		return false
	}
	defer connection.Close()

	if err = connection.Hello(qm.host); err != nil {
		log.Println(err)
		return false
	}

	supportsStartTLS, _ := connection.Extension("STARTTLS")

	if supportsStartTLS {
		config := &tls.Config{InsecureSkipVerify: qm.InsecureSkipVerify}
		if err = connection.StartTLS(config); err != nil {
			log.Println(err)
			return false
		}
	} else {
		log.Println("STARTTLS is not supported")
	}

	if !(qm.username == "" && qm.password == "") {
		auth := smtp.PlainAuth("", qm.username, qm.password, qm.host)
		if ok, _ := connection.Extension("AUTH"); ok {
			if err = connection.Auth(auth); err != nil {
				log.Println(err)
				return false
			}
		}
	}

	if err := connection.Mail(mail.From()); err != nil {
		log.Println(err)
		return false
	}

	if err := connection.Rcpt(mail.To()); err != nil {
		log.Println(err)
		return false
	}

	data, err := connection.Data()
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = fmt.Fprint(data, mail.Message())
	if err != nil {
		log.Println(err)
		return false
	}

	err = data.Close()
	if err != nil {
		log.Println(err)
		return false
	}

	if err := connection.Quit(); err != nil {
		log.Println(err)
		return false
	}

	return true
}
