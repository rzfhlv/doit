package jaeger

import (
	"io"
	"sync"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

var (
	tracer      opentracing.Tracer
	closer      io.Closer
	jaegerOnce  sync.Once
	jaegerError error
)

type Jaeger struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func NewJaeger() (*Jaeger, error) {
	jaegerOnce.Do(func() {
		cfg, err := config.FromEnv()
		if err != nil {
			jaegerError = err
		}

		cfg.Sampler.Type = jaeger.SamplerTypeConst
		cfg.Sampler.Param = 1
		cfg.Reporter.LogSpans = true

		tracer, closer, err = cfg.NewTracer(config.Logger(jaeger.StdLogger))
		if err != nil {
			jaegerError = err
		}

		opentracing.SetGlobalTracer(tracer)
	})

	if jaegerError != nil {
		return nil, jaegerError
	}

	return &Jaeger{
		Tracer: tracer,
		Closer: closer,
	}, nil
}
