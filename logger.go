package mail

type Logger interface {
	Errorln(args ...interface{})
}
