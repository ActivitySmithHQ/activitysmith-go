package activitysmith

import (
	"maps"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

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

// LiveActivityActionInput is a handwritten DX input for the optional Live Activity button.
type LiveActivityActionInput struct {
	Title  string
	Type   string
	URL    string
	Method string
	Body   map[string]interface{}
}

func (in LiveActivityActionInput) toGenerated() generated.LiveActivityAction {
	action := generated.LiveActivityAction{
		Title: in.Title,
		Type:  generated.LiveActivityActionType(in.Type),
		Url:   in.URL,
	}
	if in.Method != "" {
		action.SetMethod(generated.LiveActivityWebhookMethod(in.Method))
	}
	if in.Body != nil {
		action.SetBody(maps.Clone(in.Body))
	}
	return action
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
	Action        *LiveActivityActionInput
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
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
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

func (in LiveActivityStartInput) WithAction(v LiveActivityActionInput) LiveActivityStartInput {
	in.Action = &v
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
	Action        *LiveActivityActionInput

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
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
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

func (in LiveActivityUpdateInput) WithAction(v LiveActivityActionInput) LiveActivityUpdateInput {
	in.Action = &v
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
	Action             *LiveActivityActionInput

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
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
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

func (in LiveActivityEndInput) WithAction(v LiveActivityActionInput) LiveActivityEndInput {
	in.Action = &v
	return in
}

// LiveActivityStreamInput is a handwritten DX input with plain optional values.
type LiveActivityStreamInput struct {
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
	Metrics       []generated.ActivityMetric
	Action        *LiveActivityActionInput
	Alert         *generated.AlertPayload
	Channels      []string

	numberOfStepsSet bool
	percentageSet    bool
	valueSet         bool
	upperLimitSet    bool
}

func (in LiveActivityStreamInput) toGenerated() generated.LiveActivityStreamRequest {
	req := generated.LiveActivityStreamRequest{
		ContentState: *generated.NewStreamContentState(in.Title),
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
	if len(in.Metrics) > 0 {
		req.ContentState.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
	}
	if in.Alert != nil {
		req.SetAlert(*in.Alert)
	}
	if len(in.Channels) > 0 {
		req.SetTarget(generated.ChannelTarget{Channels: append([]string{}, in.Channels...)})
	}
	return req
}

func (in LiveActivityStreamInput) WithNumberOfSteps(v int32) LiveActivityStreamInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

func (in LiveActivityStreamInput) WithPercentage(v float32) LiveActivityStreamInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

func (in LiveActivityStreamInput) WithValue(v float32) LiveActivityStreamInput {
	in.Value = v
	in.valueSet = true
	return in
}

func (in LiveActivityStreamInput) WithUpperLimit(v float32) LiveActivityStreamInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

func (in LiveActivityStreamInput) WithAction(v LiveActivityActionInput) LiveActivityStreamInput {
	in.Action = &v
	return in
}

// LiveActivityStreamEndInput is an optional payload for ending a managed stream.
type LiveActivityStreamEndInput struct {
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
	Metrics       []generated.ActivityMetric
	Action        *LiveActivityActionInput
	Alert         *generated.AlertPayload

	numberOfStepsSet bool
	percentageSet    bool
	valueSet         bool
	upperLimitSet    bool
}

func (in LiveActivityStreamEndInput) toGenerated() generated.LiveActivityStreamDeleteRequest {
	req := generated.NewLiveActivityStreamDeleteRequest()
	if in.Title != "" {
		contentState := *generated.NewStreamContentState(in.Title)
		if in.NumberOfSteps != 0 || in.numberOfStepsSet {
			contentState.SetNumberOfSteps(in.NumberOfSteps)
		}
		if in.CurrentStep != 0 {
			contentState.SetCurrentStep(in.CurrentStep)
		}
		if in.Percentage != 0 || in.percentageSet {
			contentState.SetPercentage(in.Percentage)
		}
		if in.Value != 0 || in.valueSet {
			contentState.SetValue(in.Value)
		}
		if in.UpperLimit != 0 || in.upperLimitSet {
			contentState.SetUpperLimit(in.UpperLimit)
		}
		if in.Type != "" {
			contentState.SetType(in.Type)
		}
		if in.Subtitle != "" {
			contentState.SetSubtitle(in.Subtitle)
		}
		if in.Color != "" {
			contentState.SetColor(in.Color)
		}
		if in.StepColor != "" {
			contentState.SetStepColor(in.StepColor)
		}
		if len(in.Metrics) > 0 {
			contentState.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
		}
		req.SetContentState(contentState)
	}
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
	}
	if in.Alert != nil {
		req.SetAlert(*in.Alert)
	}
	return *req
}

func (in LiveActivityStreamEndInput) WithNumberOfSteps(v int32) LiveActivityStreamEndInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithPercentage(v float32) LiveActivityStreamEndInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithValue(v float32) LiveActivityStreamEndInput {
	in.Value = v
	in.valueSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithUpperLimit(v float32) LiveActivityStreamEndInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithAction(v LiveActivityActionInput) LiveActivityStreamEndInput {
	in.Action = &v
	return in
}
