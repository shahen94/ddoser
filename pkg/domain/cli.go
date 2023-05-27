package domain

// CliParams is a interface for command line parameters
type Cli interface {
	Parse() (*CliParams, error)
}
