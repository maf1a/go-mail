package mail

import (
	"log"
	"testing"
)

func TestMail(qt *testing.T) {
	mail, err := NewMail("Märco Béierer <email+from@marcobeierer.ch>", "Marco Beierer <email+to@marcobeierer.ch>", "test äöü", "test äöü")
	if err != nil {
		qt.Fatal(err)
	}

	if mail.subject != "=?UTF-8?q?test_=C3=A4=C3=B6=C3=BC?=" {
		qt.Fatalf("subject was %s, expected %s", mail.subject, "=?UTF-8?q?test_=C3=A4=C3=B6=C3=BC?=")
	}

	log.Println(mail.From().String())
	log.Println(mail.To().String())
	log.Println(mail.Message())
}
