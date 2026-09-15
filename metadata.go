package activitysmith

import (
	"encoding/json"
	"fmt"
	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

func metadataValues(values map[string]any) (map[string]generated.MetadataValue, error) {
	if values == nil {
		return nil, nil
	}
	result := make(map[string]generated.MetadataValue, len(values))
	for key, value := range values {
		raw, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("activitysmith: metadata %q: %w", key, err)
		}
		var scalar generated.MetadataValue
		if err := json.Unmarshal(raw, &scalar); err != nil {
			return nil, fmt.Errorf("activitysmith: metadata %q must be a string, number, or boolean: %w", key, err)
		}
		result[key] = scalar
	}
	return result, nil
}
