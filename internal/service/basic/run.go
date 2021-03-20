package basic

import (
	"../../../internal"
	"../../../internal/types"
	"../../../pkg/logger"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Run(command string) types.CommandResult {
	//
	defer func() {
		internal.TerminalProcess = nil
	}()
	//
	result := types.CommandResult{}

	if internal.TerminalProcess != nil {
		result.Status = internal.ErrorTerminalBusy
		logger.New("Error:", result.Status)
		return result
	}

	// удаляем файлы с результатами последней операции
	_ = os.Remove(internal.PathArcus + "cheq.out")
	_ = os.Remove(internal.PathArcus + "output.out")

	// сформируем команду на оплату
	if os.Getenv("MODE") == "prod" {
		command = "./commandLineTool " + command
	} else {
		command = "./terminal-emulation " + command
	}

	logger.New("Run: ", command)

	internal.TerminalProcess = exec.Command("bash", "-c", command)
	internal.TerminalProcess.Dir = internal.PathArcus

	stdout, err := internal.TerminalProcess.Output()
	logger.New("Result:", string(stdout))

	if err != nil {
		result.Status = internal.ErrorSystem
		logger.New("Error:", result.Status, err.Error())
		return result
	}

	// достанем код и статус операции
	data := GetLastData()
	st := strings.Split(data, "\n")
	if len(st) < 6 {
		result.Status = internal.ErrorSystem
		logger.New("Error!!!", result.Status)
		return result
	}

	result.Code = strings.TrimSpace(st[0])
	result.Status = strings.TrimSpace(st[5])
	result.Cheque = GetLastCheque()
	result.Data = data

	fmt.Println(result.Cheque)

	return result
}
