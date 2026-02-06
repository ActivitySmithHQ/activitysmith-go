package activitysmith

import "testing"

func TestClientConstructs(t *testing.T) {
	client, err := New("test-api-key", nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if client.Notifications == nil {
		t.Fatal("expected Notifications service")
	}

	if client.LiveActivities == nil {
		t.Fatal("expected LiveActivities service")
	}
}

func TestClientRequiresAPIKey(t *testing.T) {
	_, err := New("", nil)
	if !errorsIs(err, ErrAPIKeyRequired) {
		t.Fatalf("expected ErrAPIKeyRequired, got %v", err)
	}
}

func errorsIs(err error, target error) bool {
	if err == nil || target == nil {
		return err == target
	}
	return err.Error() == target.Error()
}
