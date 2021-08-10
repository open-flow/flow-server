package registry

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"time"
)

type RestyOut struct {
	fx.Out
	Resty *resty.Client
}

func NewResty() (RestyOut, error) {
	out := RestyOut{}

	out.Resty = resty.New().
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	return out, nil
}
