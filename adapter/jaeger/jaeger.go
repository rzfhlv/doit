package jaeger

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type Jaeger struct {
	Tracer opentracing.Tracer
	Closer io.Closer
}

func NewJaeger() (*Jaeger, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		return nil, err
	}

	opentracing.SetGlobalTracer(tracer)

	return &Jaeger{
		Tracer: tracer,
		Closer: closer,
	}, nil
}
