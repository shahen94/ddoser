package internal

import (
	"github.com/schollz/progressbar/v3"
	"github.com/shahen94/ddoser/pkg/domain"
)

// Coordinator is an interface for coordinators
type Coordinator struct {
	actor   *domain.Actor
	params  *domain.CliParams
	bar     *progressbar.ProgressBar
	results []domain.RequestResult
}

// Start starts the coordinator and accepts channel for communication
func (c *Coordinator) Start(stream chan bool) error {
	logger := NewLogger()
	reqChan := make(chan *domain.RequestResult)

	(*c.actor).UseParams(c.params)

	logger.Log("Starting the coordinator")
	go (*c.actor).Start(reqChan)

	for {
		select {
		case reqResult := <-reqChan:
			if reqResult == nil {
				stream <- true
				return nil
			}
			c.bar.Add(1)
			c.results = append(c.results, *reqResult)
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
	network := NewDDoserNetwork()

	params, err := NewCli().Parse()

	if err != nil {
		panic(err)
	}

	actor.UseParams(params)
	actor.UseNetworkInterface(network)

	if err != nil {
		panic(err)
	}

	return &Coordinator{
		actor:  &actor,
		params: params,
		bar:    progressbar.Default(int64(params.RequestCount)),
	}
}
