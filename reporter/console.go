package reporter

import (
	"log"

	"github.com/wavefronthq/wavefront-opentracing-sdk-go/tracer"
	"github.com/wavefronthq/wavefront-sdk-go/senders"
)

// ConsoleSpanReporter send spand to STDOUT.
type ConsoleSpanReporter struct {
	source string
}

// NewConsoleSpanReporter returns a ConsoleSpanReporter.
func NewConsoleSpanReporter(source string) tracer.SpanReporter {
	return &ConsoleSpanReporter{source}
}

// ReportSpan complies with the SpanReporter interface.
func (r *ConsoleSpanReporter) ReportSpan(span tracer.RawSpan) {
	sampled := ""
	if !span.Context.Sampled {
		sampled = " [not sampled]"
	}

	tags := prepareTags(span)
	parents, followsFrom := prepareReferences(span)

	line, err := senders.SpanLine(span.Operation, span.Start.UnixNano()/1000000, span.Duration.Nanoseconds()/1000000, r.source,
		span.Context.TraceID, span.Context.SpanID, parents, followsFrom, tags, nil, "")

	if err != nil {
		log.Printf("SpanLine Error: %v", err)
	} else {
		log.Printf("SpanLine%s: %v", sampled, line)
	}
}
