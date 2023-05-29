package internal

import (
	"sync/atomic"

	"github.com/shahen94/ddoser/pkg/domain"
)

// DDoser is an interface for DDosers
type DDoser struct {
	params  *domain.CliParams
	network domain.NetworkInterface

	onFlyRequests    atomic.Int32
	finishedRequests atomic.Int32

	transportChan chan *domain.RequestResult
	parentChan    chan *domain.RequestResult
}

// Use sets params
func (d *DDoser) UseParams(params *domain.CliParams) {
	d.params = params
}

// UseNetworkInterface sets network interface
func (d *DDoser) UseNetworkInterface(network domain.NetworkInterface) {
	d.network = network
}

// Start starts the DDoser and accepts channel for communication
func (d *DDoser) Start(stream chan *domain.RequestResult) error {
	d.parentChan = stream

	go d.ddos()

	for {
		if d.params.RequestCount <= int(d.finishedRequests.Load()) {
			d.parentChan <- nil
			return nil
		}

		reqResult := <-d.transportChan

		d.parentChan <- reqResult
		d.finishedRequests.Add(1)
		d.onFlyRequests.Add(-1)
		go d.ddos()
	}
}

// Stop stops the DDoser
func (d *DDoser) Stop() error {
	close(d.transportChan)
	return nil
}

func (d *DDoser) ddos() {
	inProgress := d.onFlyRequests.Load()

	for i := 0; i < d.params.Concurrent-int(inProgress); i++ {
		d.onFlyRequests.Add(1)
		go d.network.Request(d.params.BaseUrl, d.transportChan)
	}
}

// NewDDoser creates a new DDoser
func NewDDoser() domain.Actor {
	return &DDoser{
		onFlyRequests:    atomic.Int32{},
		finishedRequests: atomic.Int32{},
		transportChan:    make(chan *domain.RequestResult),
	}
}
