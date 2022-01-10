package solid

import (
	"fmt"
	"testing"
)

type Report interface {
	Report() string
}

type FinanceReport struct {
	report string
}

func (r *FinanceReport) Report() string {
	return r.report
}

type EmailReport struct {
	Email string
}

func (e *EmailReport) Report() string {
	return e.Email
}

type ReportSender struct {
	Sender string
}

func (s *ReportSender) SendReport(report Report) {
	fmt.Println(report.Report())
}

func TestSolid1(t *testing.T) {
	Fi := new(FinanceReport)
	Fi.report = "ooo@naver.com"
	Si := new(ReportSender)
	Si.SendReport(Fi)
	Ei := new(EmailReport)
	Ei.Email = "ooo@goolgle.com"
	Si.SendReport(Ei)
}
