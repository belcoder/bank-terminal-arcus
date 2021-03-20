package types

type CommandResult struct {
	Code   string `json:"code"`
	Status string `json:"string"`
	Data   string `json:"data"`
	Cheque string `json:"cheque"`
}
