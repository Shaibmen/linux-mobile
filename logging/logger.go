package logging 


type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
}

