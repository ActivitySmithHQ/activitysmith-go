package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

// PushNotificationInput is a handwritten DX input with plain optional values.
type PushNotificationInput struct {
	Title    string
	Message  string
	Subtitle string
	Channels []string
}

func (in PushNotificationInput) toGenerated() generated.PushNotificationRequest {
	req := generated.PushNotificationRequest{
		Title: in.Title,
	}
	if in.Message != "" {
		req.SetMessage(in.Message)
	}
	if in.Subtitle != "" {
		req.SetSubtitle(in.Subtitle)
	}
	if len(in.Channels) > 0 {
		req.SetTarget(generated.ChannelTarget{Channels: append([]string{}, in.Channels...)})
	}
	return req
}

// LiveActivityStartInput is a handwritten DX input with plain optional values.
type LiveActivityStartInput struct {
	Title         string
	NumberOfSteps int32
	CurrentStep   int32
	Type          string
	Subtitle      string
	Color         string
	StepColor     string
	Channels      []string
}

func (in LiveActivityStartInput) toGenerated() generated.LiveActivityStartRequest {
	req := generated.LiveActivityStartRequest{
		ContentState: generated.ContentStateStart{
			Title:         in.Title,
			NumberOfSteps: in.NumberOfSteps,
			CurrentStep:   in.CurrentStep,
			Type:          in.Type,
		},
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		req.ContentState.SetColor(in.Color)
	}
	if in.StepColor != "" {
		req.ContentState.SetStepColor(in.StepColor)
	}
	if len(in.Channels) > 0 {
		req.SetTarget(generated.ChannelTarget{Channels: append([]string{}, in.Channels...)})
	}
	return req
}

// LiveActivityUpdateInput is a handwritten DX input with plain optional values.
type LiveActivityUpdateInput struct {
	ActivityID    string
	Title         string
	CurrentStep   int32
	Subtitle      string
	Color         string
	StepColor     string
	NumberOfSteps int32

	numberOfStepsSet bool
}

func (in LiveActivityUpdateInput) toGenerated() generated.LiveActivityUpdateRequest {
	req := generated.LiveActivityUpdateRequest{
		ActivityId: in.ActivityID,
		ContentState: generated.ContentStateUpdate{
			Title:       in.Title,
			CurrentStep: in.CurrentStep,
		},
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		req.ContentState.SetColor(in.Color)
	}
	if in.StepColor != "" {
		req.ContentState.SetStepColor(in.StepColor)
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	return req
}

// WithNumberOfSteps forces inclusion of number_of_steps, including explicit zero.
func (in LiveActivityUpdateInput) WithNumberOfSteps(v int32) LiveActivityUpdateInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

// LiveActivityEndInput is a handwritten DX input with plain optional values.
type LiveActivityEndInput struct {
	ActivityID          string
	Title               string
	CurrentStep         int32
	Subtitle            string
	Color               string
	StepColor           string
	NumberOfSteps       int32
	AutoDismissMinutes  int32

	numberOfStepsSet      bool
	autoDismissMinutesSet bool
}

func (in LiveActivityEndInput) toGenerated() generated.LiveActivityEndRequest {
	req := generated.LiveActivityEndRequest{
		ActivityId: in.ActivityID,
		ContentState: generated.ContentStateEnd{
			Title:       in.Title,
			CurrentStep: in.CurrentStep,
		},
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		req.ContentState.SetColor(in.Color)
	}
	if in.StepColor != "" {
		req.ContentState.SetStepColor(in.StepColor)
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.AutoDismissMinutes != 0 || in.autoDismissMinutesSet {
		req.ContentState.SetAutoDismissMinutes(in.AutoDismissMinutes)
	}
	return req
}

// WithNumberOfSteps forces inclusion of number_of_steps, including explicit zero.
func (in LiveActivityEndInput) WithNumberOfSteps(v int32) LiveActivityEndInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

// WithAutoDismissMinutes forces inclusion of auto_dismiss_minutes, including explicit zero.
func (in LiveActivityEndInput) WithAutoDismissMinutes(v int32) LiveActivityEndInput {
	in.AutoDismissMinutes = v
	in.autoDismissMinutesSet = true
	return in
}
