package logger

import (
	"fmt"
	"time"
)

func New(data ...interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), data)
}


func Continue(data interface{}) {
	fmt.Print(data)
}
