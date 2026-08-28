package config

import "time"

type Runtime struct {
	Timeout  time.Duration
	ReadOnly bool
}

func DefaultRuntime() Runtime { return Runtime{Timeout: 5 * time.Second} }
func (r Runtime) Effective() Runtime {
	if r.Timeout <= 0 {
		r.Timeout = 5 * time.Second
	}
	return r
}
