package report

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

type Report struct {
	Address string
	Success bool
	Sent    string
	Receive string
}

var reports []Report

func BuildChannel() chan Report {
	channel := make(chan Report, 5)
	reports = make([]Report, 0, 50)
	go func() {
		for {
			report := <-channel
			reports = append(reports, report)
		}
	}()
	return channel
}

func PrettyPrintReport() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Address", "SUCCESS", "SENT", "RECEIVE"})
	for i, report := range reports {
		t.AppendRow([]interface{}{i, report.Address, report.Success, report.Sent, report.Receive})
		t.AppendSeparator()
	}

	t.Render()
}

func Success(address string, sent string, recv string) Report {
	return Report{
		Address: address,
		Success: true,
		Sent:    sent,
		Receive: recv,
	}
}

func Fail(address string, sent string, recv string) Report {
	return Report{
		Address: address,
		Success: false,
		Sent:    sent,
		Receive: recv,
	}
}
