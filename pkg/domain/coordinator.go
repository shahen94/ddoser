package domain

// Coordinator is an interface for coordinators
type Coordinator interface {
	Start(chan bool) error
	GetResults() []RequestResult
}
