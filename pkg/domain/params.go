package domain

// CliParams is a struct for command line parameters
type CliParams struct {
	BaseUrl      string
	Concurrent   int
	RequestCount int
	Timeout      int
	Verbose      bool
}
