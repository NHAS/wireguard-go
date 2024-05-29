package device

type EventType int

const (
	EventEndpointChange = 1 << iota
)

type Event struct {
	Type EventType
	Pk   NoisePublicKey
}
