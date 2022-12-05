package logger

import (
	"github.com/gin-gonic/gin"
	"log"
)

const (
	Error   = "Error"
	Info    = "Info"
	Success = "Success"
	Fatal   = "Fatal"
)

type Log struct {
	Context *gin.Context
	Message string
}

func (l Log) Error() {
	log.Printf("%s\n", Error)
}

func (l Log) Info() {
	log.Printf("%s\n", Info)
}

func (l Log) Success() {
	log.Printf("%s\n", Success)
}
func (l Log) Fatal() {
	log.Printf("%s\n", Fatal)
}
