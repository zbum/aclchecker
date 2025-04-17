package report

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"time"
)

type Report struct {
	SrcAddress  string
	DestAddress string
	Success     bool
	Sent        string
	Receive     string
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
	time.Sleep(1 * time.Second)

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Source Address", "Destination Address", "SUCCESS", "SENT", "RECEIVE"})
	for i, report := range reports {
		t.AppendRow([]interface{}{i, report.SrcAddress, report.DestAddress, report.Success, report.Sent, report.Receive})
		t.AppendSeparator()
	}

	t.Render()

}

func Success(srcAddress string, destAddress string, sent string, recv string) Report {
	return Report{
		SrcAddress:  srcAddress,
		DestAddress: destAddress,
		Success:     true,
		Sent:        sent,
		Receive:     recv,
	}
}

func Fail(srcAddress string, destAddress string, sent string, recv string) Report {
	return Report{
		SrcAddress:  srcAddress,
		DestAddress: destAddress,
		Sent:        sent,
		Receive:     recv,
	}
}
