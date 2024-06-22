package metricfx

import (
	"errors"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type IMetrics interface {
	IncrementCounter(name string, opt prometheus.CounterOpts)
	GetCounter(name string) (prometheus.Counter, error)

	AddGauge(name string, opt prometheus.GaugeOpts)
	GetGauge(name string) (prometheus.Gauge, error)

	ObserveHistogram(name string, opt prometheus.HistogramOpts)
	GetHistogram(name string) (prometheus.Histogram, error)

	ObserveSummary(name string, opt prometheus.SummaryOpts)
	GetSummary(name string) (prometheus.Summary, error)

	GetMetrics() (string, error)
}

type prometheusMetrics struct {
	counters   map[string]prometheus.Counter
	gauges     map[string]prometheus.Gauge
	histograms map[string]prometheus.Histogram
	summaries  map[string]prometheus.Summary
}

func NewPrometheusMetrics() IMetrics {
	return &prometheusMetrics{
		counters:   make(map[string]prometheus.Counter),
		gauges:     make(map[string]prometheus.Gauge),
		histograms: make(map[string]prometheus.Histogram),
		summaries:  make(map[string]prometheus.Summary),
	}
}

func (m *prometheusMetrics) IncrementCounter(name string, opt prometheus.CounterOpts) {

	counter := prometheus.NewCounter(opt)
	prometheus.MustRegister(counter)
	m.counters[name] = counter

}

func (m *prometheusMetrics) GetCounter(name string) (prometheus.Counter, error) {
	counter, ok := m.counters[name]
	if !ok {
		return nil, errors.New("counter not found")
	}
	return counter, nil
}

func (m *prometheusMetrics) AddGauge(name string, opt prometheus.GaugeOpts) {

	gauge := prometheus.NewGauge(opt)
	prometheus.MustRegister(gauge)
	m.gauges[name] = gauge

	// Since Set() is not part of GaugeOpts, we cannot set a value here.
	// The value should be set via Gauge.Set() method elsewhere.
}

func (m *prometheusMetrics) GetGauge(name string) (prometheus.Gauge, error) {
	gauge, ok := m.gauges[name]
	if !ok {
		return nil, errors.New("gauge not found")
	}
	return gauge, nil
}

func (m *prometheusMetrics) ObserveHistogram(name string, opt prometheus.HistogramOpts) {

	histogram := prometheus.NewHistogram(opt)
	prometheus.MustRegister(histogram)
	m.histograms[name] = histogram

	// Since Observe() is not part of HistogramOpts, we cannot observe a value here.
	// The value should be observed via Histogram.Observe() method elsewhere.
}

func (m *prometheusMetrics) GetHistogram(name string) (prometheus.Histogram, error) {
	histogram, ok := m.histograms[name]
	if !ok {
		return nil, errors.New("histogram not found")
	}
	return histogram, nil
}

func (m *prometheusMetrics) ObserveSummary(name string, opt prometheus.SummaryOpts) {

	summary := prometheus.NewSummary(opt)
	prometheus.MustRegister(summary)
	m.summaries[name] = summary

	// Since Observe() is not part of SummaryOpts, we cannot observe a value here.
	// The value should be observed via Summary.Observe() method elsewhere.
}

func (m *prometheusMetrics) GetSummary(name string) (prometheus.Summary, error) {
	summary, ok := m.summaries[name]
	if !ok {
		return nil, errors.New("summary not found")
	}
	return summary, nil
}

func (m *prometheusMetrics) GetMetrics() (string, error) {
	var sb strings.Builder

	// Gather metrics from the default gatherer
	mfs, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		return "", err
	}

	// Append each metric family to the string builder
	for _, mf := range mfs {
		sb.WriteString(mf.String())
		sb.WriteString("\n")
	}

	return sb.String(), nil
}
