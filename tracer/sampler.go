package tracer

import "time"

// AllwaysSample basic sampler to sample all Spans
type AllwaysSample struct{}

// ShouldSample allways true
func (t AllwaysSample) ShouldSample(span RawSpan) bool {
	return true
}

// NeverSample basic sampler to not sample any Spans
type NeverSample struct{}

// ShouldSample allways false
func (t NeverSample) ShouldSample(span RawSpan) bool {
	return false
}

// DurationSampler allows spans above a given duration in milliseconds to be reported.
type DurationSampler struct {
	Duration time.Duration
}

// ShouldSample is span duration is bigger than Duration
func (t DurationSampler) ShouldSample(span RawSpan) bool {
	return span.Duration > t.Duration
}

// TODO: RateSampler
