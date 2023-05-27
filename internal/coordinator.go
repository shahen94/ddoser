package internal

import "github.com/shahen94/ddoser/pkg/domain"

// Coordinator is an interface for coordinators
type Coordinator struct {
	actor   *domain.Actor
	params  *domain.CliParams
	results []domain.RequestResult
}

// Start starts the coordinator and accepts channel for communication
func (c *Coordinator) Start(stream chan bool) error {
	reqChan := make(chan domain.RequestResult)

	(*c.actor).Use(c.params)
	go (*c.actor).Start(reqChan)

	for {
		select {
		case reqResult := <-reqChan:
			c.results = append(c.results, reqResult)
		case <-stream:
			(*c.actor).Stop()
		}
	}
}

// GetResults returns results
func (c *Coordinator) GetResults() []domain.RequestResult {
	return c.results
}

// NewCoordinator creates a new coordinator
func NewCoordinator() domain.Coordinator {
	actor := NewDDoser()
	params, err := NewCli().Parse()

	if err != nil {
		panic(err)
	}

	return &Coordinator{
		actor:  &actor,
		params: params,
	}
}
