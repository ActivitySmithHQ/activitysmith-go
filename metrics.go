package activitysmith

import (
	"fmt"
	"time"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

type MetricsService struct {
	client *Client
}

type MetricValueInput struct {
	Value     any
	Timestamp *time.Time
}

func (s *MetricsService) Update(key string, input any) (*generated.MetricValueUpdateResponse, error) {
	request, err := normalizeMetricValueUpdateRequest(input)
	if err != nil {
		return nil, err
	}

	response, _, err := s.client.apiClient.MetricsAPI.
		UpdateMetricValue(s.client.ctx, key).
		MetricValueUpdateRequest(request).
		Execute()

	return response, err
}

// MetricInput defines a widget metric. Empty fields are omitted: an existing metric keeps
// its current values, and a new metric gets a label made from the key and the number format.
type MetricInput struct {
	Label        string
	Format       string
	CurrencyCode string
	Unit         string
	UnitSpacing  string
}

// Create creates the metric, or updates the fields sent for an existing one.
// Omitted fields keep their current values, so it is safe to call on every start.
func (s *MetricsService) Create(key string, input MetricInput) (*generated.MetricResponse, error) {
	return s.UpsertMetric(key, input.toGenerated())
}

// Backward-compatible alias.
func (s *MetricsService) UpsertMetric(key string, request generated.MetricUpsertRequest) (*generated.MetricResponse, error) {
	response, _, err := s.client.apiClient.MetricsAPI.
		UpsertMetric(s.client.ctx, key).
		MetricUpsertRequest(request).
		Execute()

	return response, err
}

func (in MetricInput) toGenerated() generated.MetricUpsertRequest {
	request := generated.MetricUpsertRequest{}
	if in.Label != "" {
		request.SetLabel(in.Label)
	}
	if in.Format != "" {
		request.SetFormat(in.Format)
	}
	if in.CurrencyCode != "" {
		request.SetCurrencyCode(in.CurrencyCode)
	}
	if in.Unit != "" {
		request.SetUnit(in.Unit)
	}
	if in.UnitSpacing != "" {
		request.SetUnitSpacing(in.UnitSpacing)
	}
	return request
}

// Backward-compatible alias.
func (s *MetricsService) UpdateMetricValue(key string, request generated.MetricValueUpdateRequest) (*generated.MetricValueUpdateResponse, error) {
	return s.Update(key, request)
}

func normalizeMetricValueUpdateRequest(input any) (generated.MetricValueUpdateRequest, error) {
	switch v := input.(type) {
	case MetricValueInput:
		return v.toGenerated()
	case *MetricValueInput:
		if v == nil {
			return generated.MetricValueUpdateRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return v.toGenerated()
	case generated.MetricValueUpdateRequest:
		return v, nil
	case *generated.MetricValueUpdateRequest:
		if v == nil {
			return generated.MetricValueUpdateRequest{}, fmt.Errorf("activitysmith: input cannot be nil")
		}
		return *v, nil
	default:
		value, err := metricValueUpdateRequestValue(input)
		if err != nil {
			return generated.MetricValueUpdateRequest{}, err
		}

		return *generated.NewMetricValueUpdateRequest(value), nil
	}
}

func (in MetricValueInput) toGenerated() (generated.MetricValueUpdateRequest, error) {
	value, err := metricValueUpdateRequestValue(in.Value)
	if err != nil {
		return generated.MetricValueUpdateRequest{}, err
	}

	request := *generated.NewMetricValueUpdateRequest(value)
	if in.Timestamp != nil {
		request.SetTimestamp(*in.Timestamp)
	}

	return request, nil
}

func metricValueUpdateRequestValue(input any) (generated.MetricValueUpdateRequestValue, error) {
	switch v := input.(type) {
	case string:
		return generated.StringAsMetricValueUpdateRequestValue(&v), nil
	case float32:
		return generated.Float32AsMetricValueUpdateRequestValue(&v), nil
	case float64:
		value := float32(v)
		return generated.Float32AsMetricValueUpdateRequestValue(&value), nil
	case int:
		value := float32(v)
		return generated.Float32AsMetricValueUpdateRequestValue(&value), nil
	case int32:
		value := float32(v)
		return generated.Float32AsMetricValueUpdateRequestValue(&value), nil
	case int64:
		value := float32(v)
		return generated.Float32AsMetricValueUpdateRequestValue(&value), nil
	default:
		return generated.MetricValueUpdateRequestValue{}, fmt.Errorf(
			"activitysmith: unsupported metric value input type %T",
			input,
		)
	}
}
