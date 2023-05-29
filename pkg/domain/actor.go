package domain

// Actor is an interface for actors
// Start starts the actor and accepts channel for communication
type Actor interface {
	UseParams(*CliParams)
	UseNetworkInterface(NetworkInterface)
	Start(stream chan *RequestResult) error
	Stop() error
}
