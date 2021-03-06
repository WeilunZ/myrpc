package transport

import (
	"context"
	"time"
)

type ServerTransportOptions struct {
	Address         string
	Network         string
	Timeout         time.Duration
	Protocol        string // proto, json
	Handler         Handler
	Serialization   string        // serialization type
	KeepAlivePeriod time.Duration // keepalive period
}

type ServerTransportOption func(*ServerTransportOptions)

type Handler interface {
	Handle(context.Context, []byte) ([]byte, error)
}

// WithServerAddress returns a ServerTransportOption which sets the value for address
func WithServerAddress(address string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Address = address
	}
}

// WithServerNetwork returns a ServerTransportOption which sets the value for network
func WithServerNetwork(network string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Network = network
	}
}

// WithServerTimeout returns a ServerTransportOption which sets the value for timeout
func WithServerTimeout(timeout time.Duration) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Timeout = timeout
	}
}

func WithProtocol(protocol string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Protocol = protocol
	}
}

// WithHandler returns a ServerTransportOption which sets the value for handler
func WithHandler(handler Handler) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Handler = handler
	}
}

// WithSerialization returns a ServerTransportOption which sets the value for serialization
func WithSerialization(serialization string) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.Serialization = serialization
	}
}

// WithKeepAlivePeriod returns a ServerTransportOption which sets the value for keepAlivePeriod
func WithKeepAlivePeriod(keepAlivePeriod time.Duration) ServerTransportOption {
	return func(o *ServerTransportOptions) {
		o.KeepAlivePeriod = keepAlivePeriod
	}
}
