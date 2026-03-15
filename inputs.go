package activitysmith

import "github.com/ActivitySmithHQ/activitysmith-go/generated"

// PushNotificationInput is a handwritten DX input with plain optional values.
type PushNotificationInput struct {
	Title       string
	Message     string
	Subtitle    string
	Media       string
	Redirection string
	Actions     []generated.PushNotificationAction
	Channels    []string
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
	if in.Media != "" {
		req.SetMedia(in.Media)
	}
	if in.Redirection != "" {
		req.SetRedirection(in.Redirection)
	}
	if len(in.Actions) > 0 {
		req.SetActions(append([]generated.PushNotificationAction{}, in.Actions...))
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
	Percentage    float32
	Value         float32
	UpperLimit    float32
	Type          string
	Subtitle      string
	Color         string
	StepColor     string
	Channels      []string

	numberOfStepsSet bool
	percentageSet    bool
	valueSet         bool
	upperLimitSet    bool
}

func (in LiveActivityStartInput) toGenerated() generated.LiveActivityStartRequest {
	req := generated.LiveActivityStartRequest{
		ContentState: *generated.NewContentStateStart(in.Title, in.Type),
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.CurrentStep != 0 {
		req.ContentState.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		req.ContentState.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		req.ContentState.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		req.ContentState.SetUpperLimit(in.UpperLimit)
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

// WithNumberOfSteps forces inclusion of number_of_steps, including explicit zero.
func (in LiveActivityStartInput) WithNumberOfSteps(v int32) LiveActivityStartInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

// WithPercentage forces inclusion of percentage, including explicit zero.
func (in LiveActivityStartInput) WithPercentage(v float32) LiveActivityStartInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

// WithValue forces inclusion of value, including explicit zero.
func (in LiveActivityStartInput) WithValue(v float32) LiveActivityStartInput {
	in.Value = v
	in.valueSet = true
	return in
}

// WithUpperLimit forces inclusion of upper_limit, including explicit zero.
func (in LiveActivityStartInput) WithUpperLimit(v float32) LiveActivityStartInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

// LiveActivityUpdateInput is a handwritten DX input with plain optional values.
type LiveActivityUpdateInput struct {
	ActivityID    string
	Title         string
	CurrentStep   int32
	Percentage    float32
	Value         float32
	UpperLimit    float32
	Type          string
	Subtitle      string
	Color         string
	StepColor     string
	NumberOfSteps int32

	numberOfStepsSet bool
	percentageSet    bool
	valueSet         bool
	upperLimitSet    bool
}

func (in LiveActivityUpdateInput) toGenerated() generated.LiveActivityUpdateRequest {
	req := generated.LiveActivityUpdateRequest{
		ActivityId:   in.ActivityID,
		ContentState: *generated.NewContentStateUpdate(in.Title),
	}
	if in.CurrentStep != 0 {
		req.ContentState.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		req.ContentState.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		req.ContentState.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		req.ContentState.SetUpperLimit(in.UpperLimit)
	}
	if in.Type != "" {
		req.ContentState.SetType(in.Type)
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

// WithPercentage forces inclusion of percentage, including explicit zero.
func (in LiveActivityUpdateInput) WithPercentage(v float32) LiveActivityUpdateInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

// WithValue forces inclusion of value, including explicit zero.
func (in LiveActivityUpdateInput) WithValue(v float32) LiveActivityUpdateInput {
	in.Value = v
	in.valueSet = true
	return in
}

// WithUpperLimit forces inclusion of upper_limit, including explicit zero.
func (in LiveActivityUpdateInput) WithUpperLimit(v float32) LiveActivityUpdateInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

// LiveActivityEndInput is a handwritten DX input with plain optional values.
type LiveActivityEndInput struct {
	ActivityID         string
	Title              string
	CurrentStep        int32
	Percentage         float32
	Value              float32
	UpperLimit         float32
	Type               string
	Subtitle           string
	Color              string
	StepColor          string
	NumberOfSteps      int32
	AutoDismissMinutes int32

	numberOfStepsSet      bool
	percentageSet         bool
	valueSet              bool
	upperLimitSet         bool
	autoDismissMinutesSet bool
}

func (in LiveActivityEndInput) toGenerated() generated.LiveActivityEndRequest {
	req := generated.LiveActivityEndRequest{
		ActivityId:   in.ActivityID,
		ContentState: *generated.NewContentStateEnd(in.Title),
	}
	if in.CurrentStep != 0 {
		req.ContentState.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		req.ContentState.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		req.ContentState.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		req.ContentState.SetUpperLimit(in.UpperLimit)
	}
	if in.Type != "" {
		req.ContentState.SetType(in.Type)
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

// WithPercentage forces inclusion of percentage, including explicit zero.
func (in LiveActivityEndInput) WithPercentage(v float32) LiveActivityEndInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

// WithValue forces inclusion of value, including explicit zero.
func (in LiveActivityEndInput) WithValue(v float32) LiveActivityEndInput {
	in.Value = v
	in.valueSet = true
	return in
}

// WithUpperLimit forces inclusion of upper_limit, including explicit zero.
func (in LiveActivityEndInput) WithUpperLimit(v float32) LiveActivityEndInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

// WithAutoDismissMinutes forces inclusion of auto_dismiss_minutes, including explicit zero.
func (in LiveActivityEndInput) WithAutoDismissMinutes(v int32) LiveActivityEndInput {
	in.AutoDismissMinutes = v
	in.autoDismissMinutesSet = true
	return in
}
