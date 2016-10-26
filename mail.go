// copied from code.marcobeierer.com/utils/mail
package mail

import "fmt"

type Mail struct {
	from    string
	to      string
	subject string
	body    string
}

func NewMail(from, to, subject, body string) *Mail {
	return &Mail{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (qm *Mail) To() string {
	return qm.to
}

func (qm *Mail) From() string {
	return qm.from
}

func (qm *Mail) Message() string {
	return fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nContent-Type: text/plain; charset=UTF-8\n\n%s", qm.from, qm.to, qm.subject, qm.body)
}
