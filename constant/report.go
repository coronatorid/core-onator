package constant

// ReportedCasesStatus ...
type ReportedCasesStatus int

const (
	// ReportedCasesConfirmed ...
	ReportedCasesConfirmed ReportedCasesStatus = iota
	// ReportedCasesRejected ...
	ReportedCasesRejected
	// ReportedCasesPending ...
	ReportedCasesPending
)

// Int convert reported status to integer
func (r ReportedCasesStatus) Int() int {
	return int(r)
}

// Humanized reported cases status
func (r ReportedCasesStatus) Humanized() string {
	return [...]string{"confirmed", "rejected", "pending"}[r]
}
