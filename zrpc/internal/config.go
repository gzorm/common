package internal

import "github.com/gzorm/common/zrpc/internal/serverinterceptors"

type (
	// StatConf defines the stat config.
	StatConf = serverinterceptors.StatConf

	// ClientMiddlewaresConf defines whether to use client middlewares.
	ClientMiddlewaresConf struct {
		Trace      bool `json:",default=true"`
		Duration   bool `json:",default=true"`
		Prometheus bool `json:",default=true"`
		Breaker    bool `json:",default=true"`
		Timeout    bool `json:",default=true"`
	}

	// ServerMiddlewaresConf defines whether to use server middlewares.
	ServerMiddlewaresConf struct {
		Trace      bool     `json:",default=true"`
		Recover    bool     `json:",default=true"`
		Stat       bool     `json:",default=true"`
		StatConf   StatConf `json:",optional"`
		Prometheus bool     `json:",default=true"`
		Breaker    bool     `json:",default=true"`
	}

	// MethodTimeoutConf defines specified timeout for gRPC methods.
	MethodTimeoutConf = serverinterceptors.MethodTimeoutConf
)
