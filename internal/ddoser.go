package internal

import "github.com/shahen94/ddoser/pkg/domain"

// DDoser is an interface for DDosers
type DDoser struct {
	params *domain.CliParams
}

// Use sets params
func (d *DDoser) Use(params *domain.CliParams) {
	d.params = params
}

// Start starts the DDoser and accepts channel for communication
func (d *DDoser) Start(stream chan domain.RequestResult) error {
	return nil
}

// Stop stops the DDoser
func (d *DDoser) Stop() error {
	return nil
}

// NewDDoser creates a new DDoser
func NewDDoser() domain.Actor {
	return &DDoser{}
}
