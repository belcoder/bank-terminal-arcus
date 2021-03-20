package basic

import (
	"../../../internal"
	"../../../pkg/logger"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetLastCheque() string {
	cheq, err := ioutil.ReadFile(fmt.Sprintf("%scheq.out", internal.PathArcus))
	if err != nil {
		logger.New("GetLastCheque", err.Error())
		return ""
	}

	lines := strings.Split(string(cheq), "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}

	return strings.Join(lines, "\n")
}