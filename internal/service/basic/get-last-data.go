package basic

import (
	"../../../internal"
	"../../../pkg/logger"
	"fmt"
	"io/ioutil")

func GetLastData() string {
	data, err := ioutil.ReadFile(fmt.Sprintf("%soutput.out", internal.PathArcus))
	if err != nil {
		logger.New("GetLastData", err.Error())
		return ""
	}

	return string(data)
}