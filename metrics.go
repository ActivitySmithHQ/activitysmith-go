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
