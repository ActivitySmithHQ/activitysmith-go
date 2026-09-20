package activitysmith

import (
	"fmt"
	"math"
	"reflect"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

// liveActivityValue adapts the same plain string/number input used by other SDKs.
func liveActivityValue(input any) generated.LiveActivityValue {
	if v, ok := input.(generated.LiveActivityValue); ok {
		return v
	}
	if v, ok := input.(string); ok {
		return generated.StringAsLiveActivityValue(&v)
	}
	var n float64
	v := reflect.ValueOf(input)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n = float64(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n = float64(v.Uint())
	case reflect.Float32, reflect.Float64:
		n = v.Float()
	default:
		panic(fmt.Sprintf("activitysmith: value must be a string or finite number, got %T", input))
	}
	f := float32(n)
	if math.IsNaN(n) || math.IsInf(n, 0) || math.IsInf(float64(f), 0) {
		panic("activitysmith: value must be finite")
	}
	return generated.Float32AsLiveActivityValue(&f)
}
