package tracing

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
)

func InitTracer(serviceName string, zipkinURL string) func(context.Context) error {
	exporter, err := zipkin.New(zipkinURL)
	if err != nil {
		log.Fatalf("Error creating Zipkin exporter: %v", err)
	}

	tp := traceSDK.NewTracerProvider(
		traceSDK.WithBatcher(exporter),
		traceSDK.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)

	return tp.Shutdown
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer("serviceB-tracer")
	return tracer.Start(ctx, name)
}

func InjectTraceIntoRequest(ctx context.Context, req *http.Request) {
	propagator := propagation.TraceContext{}
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))
}

func StartSpanFromRequest(r *http.Request, spanName string) (context.Context, trace.Span) {
	propagator := propagation.TraceContext{}
	ctx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	return StartSpan(ctx, spanName)
}
