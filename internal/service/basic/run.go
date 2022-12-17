package basic

import (
	"../../../internal"
	"../../../internal/service/rabbit"
	"../../../internal/types"
	"../../../pkg/contains"
	"../../../pkg/logger"
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
	"time"
)

func Run(command string) types.CommandResult {
	//
	internal.RunProcess = true

	go ScanLog(time.Now())

	defer func() {
		internal.TerminalProcess = nil
		internal.RunProcess = false
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

// обработанные файлы
var fl = []string{}

func ScanLog(momentStart time.Time) {
	messages := []string{}
	fileName := ""

	for {
		if !internal.RunProcess {
			fl = append(fl, fileName)
			return
		}

		time.Sleep(time.Duration(1) * time.Second)

		// поиск текущего лога
		files, err := ioutil.ReadDir(internal.PathArcus + "logs")
		if err != nil {
			logger.New("ScanLog", err.Error())
			return
		}

		if len(files) == 0 {
			logger.New("ScanLog", "логи отсутствуют")
			return
		}

		sort.Slice(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})

		// самый новый файл
		fileLog := files[0]
		// logger.New("ScanLog", "лог", fileLog.Name())

		if (momentStart.Unix() - 5) > fileLog.ModTime().Unix() {
			logger.New("ScanLog", "лог старше чем momentStart - 5сек", fileName, momentStart)
			return
		}

		fileName = fileLog.Name()

		if contains.String(fl, fileName) {
			logger.New("ScanLog", "лог уже обработан", fileName)
			return
		}

		//
		file, err := os.Open(internal.PathArcus + "logs/" + fileLog.Name())
		if err != nil {
			logger.New(err.Error())
			return
		}

		decodingReader := transform.NewReader(file, charmap.Windows1251.NewDecoder())
		scanner := bufio.NewScanner(decodingReader)

		for scanner.Scan() {
			// <- STATUS
			r, _ := regexp.Compile(`^.*\<\- STATUS\:(.*)$`)
			v := r.FindStringSubmatch(scanner.Text())
			if len(v) == 2 {
				if !contains.String(messages, v[0]) {
					messages = append(messages, v[0])
					logger.New(v[1])
					rabbit.PublishWithContext("TERMINAL_EVENTS", v[1])
				}
			}
		}

		file.Close()
	}
}
