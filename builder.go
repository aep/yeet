package yeet

import (
	"context"
	"net"
	"net/url"
	"time"
)

type Logger interface {
	Print(v ...interface{})
}

type Builder struct {
	Url            string
	Context        context.Context
	Dialer         *net.Dialer
	ConnectTimeout time.Duration
	Log            Logger
	KeepAlive      time.Duration
}

func New() *Builder {
	return &Builder{
		Context:        context.Background(),
		ConnectTimeout: 20 * time.Second,
		KeepAlive:      3 * time.Second,
	}
}

func (b *Builder) WithContext(ctx context.Context) *Builder {
	b.Context = ctx
	return b
}

func (b *Builder) WithDialer(dialer *net.Dialer) *Builder {
	b.Dialer = dialer
	return b
}

func (b *Builder) WithLogger(logger Logger) *Builder {
	b.Log = logger
	return b
}

func (b *Builder) WithKeepAlive(keepAlive time.Duration) *Builder {
	b.KeepAlive = keepAlive
	return b
}

func (b *Builder) Connect(uri string) (net.Conn, error) {

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	return fromBuilder(b, u, nil)
}

func (b *Builder) Accept(sock net.Conn) (net.Conn, error) {
	return fromBuilder(b, nil, sock)
}
