package main

import "fmt"

type Module struct {
	options Options
}

func (m *Module) PrintOption() {
	fmt.Printf("%+v\n", m.options)
}

type Options struct {
	tls           bool
	port          int
	maxConnection int
}

type OptFunc func(*Options)

func NewModule(opts ...OptFunc) Module {
	opt := defaultOptions()
	for _, f := range opts {
		f(&opt)
	}
	return Module{options: opt}
}

func defaultOptions() Options {
	return Options{
		tls:           false,
		port:          80,
		maxConnection: 20,
	}
}

func WithTLS(tls bool) OptFunc {
	return func(o *Options) {
		o.tls = tls
	}
}

func WithPort(port int) OptFunc {
	return func(o *Options) {
		o.port = port
	}
}

func WithMaxConnection(maxCon int) OptFunc {
	return func(o *Options) {
		o.maxConnection = maxCon
	}
}
