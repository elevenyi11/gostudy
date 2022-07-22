package main

import (
	"context"
	"fmt"
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// InitJaeger ...
func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := jaegercfg.FromEnv()
	cfg.ServiceName = service

	cfg.Sampler.Type = jaeger.SamplerTypeConst
	cfg.Sampler.Param = 1
	//cfg.Reporter.LocalAgentHostPort = "127.0.0.1:6831"
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.NewTracer(jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {

	tracer, closer := InitJaeger("hello-world")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer)

	helloStr := "hello jaeger"
	span := tracer.StartSpan("say-hello")
	time.Sleep(time.Duration(2) * time.Millisecond)

	println(helloStr)
	ctx := context.Background()
	ctx = opentracing.ContextWithSpan(ctx, span)
	GetName(ctx, "aa")
	GetLife(ctx, "bb")
	span.Finish()
}

func GetName(ctx context.Context, userName string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetName")
	defer span.Finish()

	span.SetTag("userName", userName)
	span.LogFields(
		log.String("event", "sayhello"),
		log.String("value", userName),
	)
	time.Sleep(2 * time.Millisecond)
	println(userName)

	GetLife(ctx, userName+"11")
}

func GetLife(ctx context.Context, userName string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetLife")
	defer span.Finish()

	span.SetTag("userName", userName)
	span.LogFields(
		log.String("event", "get life"),
		log.String("value", userName),
	)
	time.Sleep(4 * time.Millisecond)
	println(userName)
}

/*
// TraceSpan is a middleware that initialize a tracing span and injects span
// context to r.Context(). In one word, this middleware kept an eye on the
// whole HTTP request that the server receives.
func TraceSpan(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			// Tracer not found, just skip.
			next.ServeHTTP(w, r)
		}

		buf := buffer.GlobalBytesPool().Get()
		buf.AppendString("HTTP ")
		buf.AppendString(r.Method)

		// Start span.
		span := opentracing.StartSpan(buf.String())
		rc := opentracing.ContextWithSpan(r.Context(), span)

		// Set request ID for context.
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			rc = context.WithValue(rc, constants.RequestID, sc.TraceID().String())
		}

		next.ServeHTTP(w, r.WithContext(rc))

		// Finish span.
		wrapper, ok := w.(WrapResponseWriter)
		if ok {
			ext.HTTPStatusCode.Set(span, uint16(wrapper.Status()))
		}
		span.Finish()
	}
	return http.HandlerFunc(fn)
}*/
