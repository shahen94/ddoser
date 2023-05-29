package internal

import (
	"net/http"
	"time"

	"github.com/shahen94/ddoser/pkg/domain"
)

type DDoserNetwork struct{}

func (d *DDoserNetwork) Request(url string, transport chan *domain.RequestResult) {
	now := time.Now().UnixNano()

	data, err := http.Get(url)

	diff := time.Now().UnixNano() - now

	if err != nil {
		transport <- &domain.RequestResult{
			StatusCode:   0,
			ResponseTime: diff,
			Error:        true,
		}
		return
	}

	transport <- &domain.RequestResult{
		StatusCode:   data.StatusCode,
		ResponseTime: diff,
		Error:        err != nil,
	}
}

func NewDDoserNetwork() domain.NetworkInterface {
	return &DDoserNetwork{}
}
