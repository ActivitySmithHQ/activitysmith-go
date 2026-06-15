package activitysmith

import (
	"fmt"
	"maps"

	"github.com/ActivitySmithHQ/activitysmith-go/generated"
)

const (
	LiveActivityTypeSegmentedProgress = "segmented_progress"
	LiveActivityTypeProgress          = "progress"
	LiveActivityTypeMetrics           = "metrics"
	LiveActivityTypeStats             = "stats"
	LiveActivityTypeAlert             = "alert"
	LiveActivityTypeTimer             = "timer"
)

type ActivityMetricOption func(*generated.ActivityMetric)

type ActivityMetric = generated.ActivityMetric
type PushNotificationAction = generated.PushNotificationAction

type LiveActivityAlertIconInput struct {
	Symbol string
	Color  string
}

type LiveActivityAlertBadgeInput struct {
	Title string
	Color string
}

func AlertIcon(symbol string, color ...string) LiveActivityAlertIconInput {
	icon := LiveActivityAlertIconInput{Symbol: symbol}
	if len(color) > 0 {
		icon.Color = color[0]
	}
	return icon
}

func AlertBadge(title string, color ...string) LiveActivityAlertBadgeInput {
	badge := LiveActivityAlertBadgeInput{Title: title}
	if len(color) > 0 {
		badge.Color = color[0]
	}
	return badge
}

func alertIconMap(icon LiveActivityAlertIconInput) map[string]interface{} {
	if icon.Symbol == "" {
		return nil
	}
	value := map[string]interface{}{"symbol": icon.Symbol}
	if icon.Color != "" {
		value["color"] = icon.Color
	}
	return value
}

func alertBadgeMap(badge LiveActivityAlertBadgeInput) map[string]interface{} {
	if badge.Title == "" {
		return nil
	}
	value := map[string]interface{}{"title": badge.Title}
	if badge.Color != "" {
		value["color"] = badge.Color
	}
	return value
}

func setAdditionalProperty(properties *map[string]interface{}, key string, value interface{}) {
	if value == nil {
		return
	}
	if *properties == nil {
		*properties = map[string]interface{}{}
	}
	(*properties)[key] = value
}

func Metric(label string, value any, options ...ActivityMetricOption) ActivityMetric {
	metric, err := NewActivityMetric(label, value, options...)
	if err != nil {
		panic(err)
	}
	return metric
}

func NewActivityMetric(label string, value any, options ...ActivityMetricOption) (generated.ActivityMetric, error) {
	metricValue, err := activityMetricValue(value)
	if err != nil {
		return generated.ActivityMetric{}, err
	}

	metric := generated.ActivityMetric{
		Label: label,
		Value: metricValue,
	}
	for _, option := range options {
		option(&metric)
	}
	return metric, nil
}

func MetricUnit(unit string) ActivityMetricOption {
	return WithActivityMetricUnit(unit)
}

func WithActivityMetricUnit(unit string) ActivityMetricOption {
	return func(metric *generated.ActivityMetric) {
		metric.SetUnit(unit)
	}
}

func MetricColor(color string) ActivityMetricOption {
	return WithActivityMetricColor(color)
}

func WithActivityMetricColor(color string) ActivityMetricOption {
	return func(metric *generated.ActivityMetric) {
		metric.SetColor(color)
	}
}

func activityMetricValue(input any) (generated.ActivityMetricValue, error) {
	switch v := input.(type) {
	case string:
		return generated.StringAsActivityMetricValue(&v), nil
	case float32:
		return generated.Float32AsActivityMetricValue(&v), nil
	case float64:
		value := float32(v)
		return generated.Float32AsActivityMetricValue(&value), nil
	case int:
		value := float32(v)
		return generated.Float32AsActivityMetricValue(&value), nil
	case int32:
		value := float32(v)
		return generated.Float32AsActivityMetricValue(&value), nil
	case int64:
		value := float32(v)
		return generated.Float32AsActivityMetricValue(&value), nil
	default:
		return generated.ActivityMetricValue{}, fmt.Errorf(
			"activitysmith: unsupported activity metric value input type %T",
			input,
		)
	}
}

type PushNotificationActionOption func(*generated.PushNotificationAction)

func PushAction(title string, type_ string, url string, options ...PushNotificationActionOption) PushNotificationAction {
	action := generated.PushNotificationAction{
		Title: title,
		Type:  generated.PushNotificationActionType(type_),
		Url:   url,
	}
	for _, option := range options {
		option(&action)
	}
	return action
}

func PushActionMethod(method string) PushNotificationActionOption {
	return func(action *generated.PushNotificationAction) {
		action.SetMethod(generated.PushNotificationWebhookMethod(method))
	}
}

func PushActionBody(body map[string]interface{}) PushNotificationActionOption {
	return func(action *generated.PushNotificationAction) {
		action.SetBody(maps.Clone(body))
	}
}

// PushNotificationInput is a handwritten DX input with plain optional values.
type PushNotificationInput struct {
	Title       string
	Message     string
	Subtitle    string
	Media       string
	Redirection string
	Actions     []PushNotificationAction
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

// LiveActivityContentStateInput keeps the visible Live Activity state separate
// from top-level routing fields such as action and channels.
type LiveActivityContentStateInput struct {
	Title              string
	NumberOfSteps      int32
	CurrentStep        int32
	Percentage         float32
	Value              float32
	UpperLimit         float32
	DurationSeconds    float32
	CountsDown         bool
	Type               string
	Subtitle           string
	Message            string
	Icon               LiveActivityAlertIconInput
	Badge              LiveActivityAlertBadgeInput
	Color              string
	StepColor          string
	AutoDismissMinutes int32
	Metrics            []ActivityMetric

	numberOfStepsSet      bool
	currentStepSet        bool
	percentageSet         bool
	valueSet              bool
	upperLimitSet         bool
	durationSecondsSet    bool
	countsDownSet         bool
	autoDismissMinutesSet bool
}

func (in LiveActivityContentStateInput) isSet() bool {
	return in.Title != "" ||
		in.Type != "" ||
		in.Subtitle != "" ||
		in.Message != "" ||
		in.Icon.Symbol != "" ||
		in.Badge.Title != "" ||
		in.Color != "" ||
		in.StepColor != "" ||
		in.NumberOfSteps != 0 ||
		in.CurrentStep != 0 ||
		in.Percentage != 0 ||
		in.Value != 0 ||
		in.UpperLimit != 0 ||
		in.DurationSeconds != 0 ||
		in.AutoDismissMinutes != 0 ||
		len(in.Metrics) > 0 ||
		in.numberOfStepsSet ||
		in.currentStepSet ||
		in.percentageSet ||
		in.valueSet ||
		in.upperLimitSet ||
		in.durationSecondsSet ||
		in.countsDownSet ||
		in.autoDismissMinutesSet
}

func (in LiveActivityContentStateInput) applyTimerFields(properties *map[string]interface{}) {
	if in.DurationSeconds != 0 || in.durationSecondsSet {
		setAdditionalProperty(properties, "duration_seconds", in.DurationSeconds)
	}
	if in.countsDownSet {
		setAdditionalProperty(properties, "counts_down", in.CountsDown)
	}
}

func (in LiveActivityContentStateInput) applyAlertFields(properties *map[string]interface{}) {
	if in.Message != "" {
		setAdditionalProperty(properties, "message", in.Message)
	}
	setAdditionalProperty(properties, "icon", alertIconMap(in.Icon))
	setAdditionalProperty(properties, "badge", alertBadgeMap(in.Badge))
}

func (in LiveActivityContentStateInput) applyStart(state *generated.ContentStateStart) {
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		state.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.CurrentStep != 0 || in.currentStepSet {
		state.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		state.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		state.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		state.SetUpperLimit(in.UpperLimit)
	}
	if in.Subtitle != "" {
		state.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		state.SetColor(in.Color)
	}
	if in.StepColor != "" {
		state.SetStepColor(in.StepColor)
	}
	if len(in.Metrics) > 0 {
		state.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	in.applyTimerFields(&state.AdditionalProperties)
	in.applyAlertFields(&state.AdditionalProperties)
}

func (in LiveActivityContentStateInput) applyUpdate(state *generated.ContentStateUpdate) {
	if in.CurrentStep != 0 || in.currentStepSet {
		state.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		state.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		state.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		state.SetUpperLimit(in.UpperLimit)
	}
	if in.Type != "" {
		state.SetType(in.Type)
	}
	if in.Subtitle != "" {
		state.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		state.SetColor(in.Color)
	}
	if in.StepColor != "" {
		state.SetStepColor(in.StepColor)
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		state.SetNumberOfSteps(in.NumberOfSteps)
	}
	if len(in.Metrics) > 0 {
		state.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	in.applyTimerFields(&state.AdditionalProperties)
	in.applyAlertFields(&state.AdditionalProperties)
}

func (in LiveActivityContentStateInput) applyEnd(state *generated.ContentStateEnd) {
	in.applyEndBase(state)
	if in.AutoDismissMinutes != 0 || in.autoDismissMinutesSet {
		state.SetAutoDismissMinutes(in.AutoDismissMinutes)
	}
}

func (in LiveActivityContentStateInput) applyEndBase(state *generated.ContentStateEnd) {
	if in.CurrentStep != 0 || in.currentStepSet {
		state.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		state.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		state.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		state.SetUpperLimit(in.UpperLimit)
	}
	if in.Type != "" {
		state.SetType(in.Type)
	}
	if in.Subtitle != "" {
		state.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		state.SetColor(in.Color)
	}
	if in.StepColor != "" {
		state.SetStepColor(in.StepColor)
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		state.SetNumberOfSteps(in.NumberOfSteps)
	}
	if len(in.Metrics) > 0 {
		state.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	in.applyTimerFields(&state.AdditionalProperties)
	in.applyAlertFields(&state.AdditionalProperties)
}

func (in LiveActivityContentStateInput) applyStream(state *generated.StreamContentState) {
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		state.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.CurrentStep != 0 || in.currentStepSet {
		state.SetCurrentStep(in.CurrentStep)
	}
	if in.Percentage != 0 || in.percentageSet {
		state.SetPercentage(in.Percentage)
	}
	if in.Value != 0 || in.valueSet {
		state.SetValue(in.Value)
	}
	if in.UpperLimit != 0 || in.upperLimitSet {
		state.SetUpperLimit(in.UpperLimit)
	}
	if in.Type != "" {
		state.SetType(in.Type)
	}
	if in.Subtitle != "" {
		state.SetSubtitle(in.Subtitle)
	}
	if in.Color != "" {
		state.SetColor(in.Color)
	}
	if in.StepColor != "" {
		state.SetStepColor(in.StepColor)
	}
	if len(in.Metrics) > 0 {
		state.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	if in.AutoDismissMinutes != 0 || in.autoDismissMinutesSet {
		state.SetAutoDismissMinutes(in.AutoDismissMinutes)
	}
	in.applyTimerFields(&state.AdditionalProperties)
	in.applyAlertFields(&state.AdditionalProperties)
}

func (in LiveActivityContentStateInput) WithNumberOfSteps(v int32) LiveActivityContentStateInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

func (in LiveActivityContentStateInput) WithCurrentStep(v int32) LiveActivityContentStateInput {
	in.CurrentStep = v
	in.currentStepSet = true
	return in
}

func (in LiveActivityContentStateInput) WithPercentage(v float32) LiveActivityContentStateInput {
	in.Percentage = v
	in.percentageSet = true
	return in
}

func (in LiveActivityContentStateInput) WithValue(v float32) LiveActivityContentStateInput {
	in.Value = v
	in.valueSet = true
	return in
}

func (in LiveActivityContentStateInput) WithUpperLimit(v float32) LiveActivityContentStateInput {
	in.UpperLimit = v
	in.upperLimitSet = true
	return in
}

func (in LiveActivityContentStateInput) WithDurationSeconds(v float32) LiveActivityContentStateInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityContentStateInput) WithCountsDown(v bool) LiveActivityContentStateInput {
	in.CountsDown = v
	in.countsDownSet = true
	return in
}

func (in LiveActivityContentStateInput) WithAutoDismissMinutes(v int32) LiveActivityContentStateInput {
	in.AutoDismissMinutes = v
	in.autoDismissMinutesSet = true
	return in
}

// LiveActivityStartInput is a handwritten DX input with plain optional values.
type LiveActivityStartInput struct {
	ContentState    LiveActivityContentStateInput
	Title           string
	NumberOfSteps   int32
	CurrentStep     int32
	Percentage      float32
	Value           float32
	UpperLimit      float32
	DurationSeconds float32
	CountsDown      bool
	Type            string
	Subtitle        string
	Message         string
	Icon            LiveActivityAlertIconInput
	Badge           LiveActivityAlertBadgeInput
	Color           string
	StepColor       string
	Metrics         []generated.ActivityMetric
	Action          *LiveActivityActionInput
	SecondaryAction *LiveActivityActionInput
	Channels        []string

	numberOfStepsSet   bool
	currentStepSet     bool
	percentageSet      bool
	valueSet           bool
	upperLimitSet      bool
	durationSecondsSet bool
	countsDownSet      bool
}

func (in LiveActivityStartInput) toGenerated() generated.LiveActivityStartRequest {
	contentState := in.ContentState
	if !contentState.isSet() {
		contentState = LiveActivityContentStateInput{
			Title:   in.Title,
			Type:    in.Type,
			Message: in.Message,
			Icon:    in.Icon,
			Badge:   in.Badge,
		}
	}
	req := generated.LiveActivityStartRequest{
		ContentState: *generated.NewContentStateStart(contentState.Title, contentState.Type),
	}
	contentState.applyStart(&req.ContentState)
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.CurrentStep != 0 || in.currentStepSet {
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
	if in.DurationSeconds != 0 || in.durationSecondsSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "duration_seconds", in.DurationSeconds)
	}
	if in.countsDownSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "counts_down", in.CountsDown)
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Message != "" {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "message", in.Message)
	}
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "icon", alertIconMap(in.Icon))
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "badge", alertBadgeMap(in.Badge))
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
	if in.SecondaryAction != nil {
		req.SetSecondaryAction(in.SecondaryAction.toGenerated())
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

func (in LiveActivityStartInput) WithCurrentStep(v int32) LiveActivityStartInput {
	in.CurrentStep = v
	in.currentStepSet = true
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

func (in LiveActivityStartInput) WithDurationSeconds(v float32) LiveActivityStartInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityStartInput) WithCountsDown(v bool) LiveActivityStartInput {
	in.CountsDown = v
	in.countsDownSet = true
	return in
}

func (in LiveActivityStartInput) WithAction(v LiveActivityActionInput) LiveActivityStartInput {
	in.Action = &v
	return in
}

func (in LiveActivityStartInput) WithSecondaryAction(v LiveActivityActionInput) LiveActivityStartInput {
	in.SecondaryAction = &v
	return in
}

// LiveActivityUpdateInput is a handwritten DX input with plain optional values.
type LiveActivityUpdateInput struct {
	ActivityID      string
	ContentState    LiveActivityContentStateInput
	Title           string
	CurrentStep     int32
	Percentage      float32
	Value           float32
	UpperLimit      float32
	DurationSeconds float32
	CountsDown      bool
	Type            string
	Subtitle        string
	Message         string
	Icon            LiveActivityAlertIconInput
	Badge           LiveActivityAlertBadgeInput
	Color           string
	StepColor       string
	NumberOfSteps   int32
	Metrics         []generated.ActivityMetric
	Action          *LiveActivityActionInput
	SecondaryAction *LiveActivityActionInput

	numberOfStepsSet   bool
	currentStepSet     bool
	percentageSet      bool
	valueSet           bool
	upperLimitSet      bool
	durationSecondsSet bool
	countsDownSet      bool
}

func (in LiveActivityUpdateInput) toGenerated() generated.LiveActivityUpdateRequest {
	contentState := in.ContentState
	if !contentState.isSet() {
		contentState = LiveActivityContentStateInput{
			Title:   in.Title,
			Type:    in.Type,
			Message: in.Message,
			Icon:    in.Icon,
			Badge:   in.Badge,
		}
	}
	req := generated.LiveActivityUpdateRequest{
		ActivityId:   in.ActivityID,
		ContentState: *generated.NewContentStateUpdate(contentState.Title),
	}
	contentState.applyUpdate(&req.ContentState)
	if in.CurrentStep != 0 || in.currentStepSet {
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
	if in.DurationSeconds != 0 || in.durationSecondsSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "duration_seconds", in.DurationSeconds)
	}
	if in.countsDownSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "counts_down", in.CountsDown)
	}
	if in.Type != "" {
		req.ContentState.SetType(in.Type)
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Message != "" {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "message", in.Message)
	}
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "icon", alertIconMap(in.Icon))
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "badge", alertBadgeMap(in.Badge))
	if in.Color != "" {
		req.ContentState.SetColor(in.Color)
	}
	if in.StepColor != "" {
		req.ContentState.SetStepColor(in.StepColor)
	}
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	if len(in.Metrics) > 0 {
		req.ContentState.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
	}
	if in.SecondaryAction != nil {
		req.SetSecondaryAction(in.SecondaryAction.toGenerated())
	}
	return req
}

// WithNumberOfSteps forces inclusion of number_of_steps, including explicit zero.
func (in LiveActivityUpdateInput) WithNumberOfSteps(v int32) LiveActivityUpdateInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

func (in LiveActivityUpdateInput) WithCurrentStep(v int32) LiveActivityUpdateInput {
	in.CurrentStep = v
	in.currentStepSet = true
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

func (in LiveActivityUpdateInput) WithDurationSeconds(v float32) LiveActivityUpdateInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityUpdateInput) WithCountsDown(v bool) LiveActivityUpdateInput {
	in.CountsDown = v
	in.countsDownSet = true
	return in
}

func (in LiveActivityUpdateInput) WithAction(v LiveActivityActionInput) LiveActivityUpdateInput {
	in.Action = &v
	return in
}

func (in LiveActivityUpdateInput) WithSecondaryAction(v LiveActivityActionInput) LiveActivityUpdateInput {
	in.SecondaryAction = &v
	return in
}

// LiveActivityEndInput is a handwritten DX input with plain optional values.
type LiveActivityEndInput struct {
	ActivityID         string
	ContentState       LiveActivityContentStateInput
	Title              string
	CurrentStep        int32
	Percentage         float32
	Value              float32
	UpperLimit         float32
	DurationSeconds    float32
	CountsDown         bool
	Type               string
	Subtitle           string
	Message            string
	Icon               LiveActivityAlertIconInput
	Badge              LiveActivityAlertBadgeInput
	Color              string
	StepColor          string
	NumberOfSteps      int32
	AutoDismissMinutes int32
	Metrics            []generated.ActivityMetric
	Action             *LiveActivityActionInput
	SecondaryAction    *LiveActivityActionInput

	numberOfStepsSet      bool
	currentStepSet        bool
	percentageSet         bool
	valueSet              bool
	upperLimitSet         bool
	durationSecondsSet    bool
	countsDownSet         bool
	autoDismissMinutesSet bool
}

func (in LiveActivityEndInput) toGenerated() generated.LiveActivityEndRequest {
	contentState := in.ContentState
	if !contentState.isSet() {
		contentState = LiveActivityContentStateInput{
			Title:   in.Title,
			Type:    in.Type,
			Message: in.Message,
			Icon:    in.Icon,
			Badge:   in.Badge,
		}
	}
	req := generated.LiveActivityEndRequest{
		ActivityId:   in.ActivityID,
		ContentState: *generated.NewContentStateEnd(contentState.Title),
	}
	contentState.applyEnd(&req.ContentState)
	if in.CurrentStep != 0 || in.currentStepSet {
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
	if in.DurationSeconds != 0 || in.durationSecondsSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "duration_seconds", in.DurationSeconds)
	}
	if in.countsDownSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "counts_down", in.CountsDown)
	}
	if in.Type != "" {
		req.ContentState.SetType(in.Type)
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Message != "" {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "message", in.Message)
	}
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "icon", alertIconMap(in.Icon))
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "badge", alertBadgeMap(in.Badge))
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
	if len(in.Metrics) > 0 {
		req.ContentState.SetMetrics(append([]generated.ActivityMetric{}, in.Metrics...))
	}
	if in.Action != nil {
		req.SetAction(in.Action.toGenerated())
	}
	if in.SecondaryAction != nil {
		req.SetSecondaryAction(in.SecondaryAction.toGenerated())
	}
	return req
}

// WithNumberOfSteps forces inclusion of number_of_steps, including explicit zero.
func (in LiveActivityEndInput) WithNumberOfSteps(v int32) LiveActivityEndInput {
	in.NumberOfSteps = v
	in.numberOfStepsSet = true
	return in
}

func (in LiveActivityEndInput) WithCurrentStep(v int32) LiveActivityEndInput {
	in.CurrentStep = v
	in.currentStepSet = true
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

func (in LiveActivityEndInput) WithDurationSeconds(v float32) LiveActivityEndInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityEndInput) WithCountsDown(v bool) LiveActivityEndInput {
	in.CountsDown = v
	in.countsDownSet = true
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

func (in LiveActivityEndInput) WithSecondaryAction(v LiveActivityActionInput) LiveActivityEndInput {
	in.SecondaryAction = &v
	return in
}

// LiveActivityStreamInput is a handwritten DX input with plain optional values.
type LiveActivityStreamInput struct {
	ContentState    LiveActivityContentStateInput
	Title           string
	NumberOfSteps   int32
	CurrentStep     int32
	Percentage      float32
	Value           float32
	UpperLimit      float32
	DurationSeconds float32
	CountsDown      bool
	Type            string
	Subtitle        string
	Message         string
	Icon            LiveActivityAlertIconInput
	Badge           LiveActivityAlertBadgeInput
	Color           string
	StepColor       string
	Metrics         []generated.ActivityMetric
	Action          *LiveActivityActionInput
	SecondaryAction *LiveActivityActionInput
	Alert           *generated.AlertPayload
	Channels        []string

	numberOfStepsSet   bool
	currentStepSet     bool
	percentageSet      bool
	valueSet           bool
	upperLimitSet      bool
	durationSecondsSet bool
	countsDownSet      bool
}

func (in LiveActivityStreamInput) toGenerated() generated.LiveActivityStreamRequest {
	contentState := in.ContentState
	if !contentState.isSet() {
		contentState = LiveActivityContentStateInput{
			Title:   in.Title,
			Type:    in.Type,
			Message: in.Message,
			Icon:    in.Icon,
			Badge:   in.Badge,
		}
	}
	req := generated.LiveActivityStreamRequest{
		ContentState: *generated.NewStreamContentState(contentState.Title),
	}
	contentState.applyStream(&req.ContentState)
	if in.NumberOfSteps != 0 || in.numberOfStepsSet {
		req.ContentState.SetNumberOfSteps(in.NumberOfSteps)
	}
	if in.CurrentStep != 0 || in.currentStepSet {
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
	if in.DurationSeconds != 0 || in.durationSecondsSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "duration_seconds", in.DurationSeconds)
	}
	if in.countsDownSet {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "counts_down", in.CountsDown)
	}
	if in.Type != "" {
		req.ContentState.SetType(in.Type)
	}
	if in.Subtitle != "" {
		req.ContentState.SetSubtitle(in.Subtitle)
	}
	if in.Message != "" {
		setAdditionalProperty(&req.ContentState.AdditionalProperties, "message", in.Message)
	}
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "icon", alertIconMap(in.Icon))
	setAdditionalProperty(&req.ContentState.AdditionalProperties, "badge", alertBadgeMap(in.Badge))
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
	if in.SecondaryAction != nil {
		req.SetSecondaryAction(in.SecondaryAction.toGenerated())
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

func (in LiveActivityStreamInput) WithCurrentStep(v int32) LiveActivityStreamInput {
	in.CurrentStep = v
	in.currentStepSet = true
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

func (in LiveActivityStreamInput) WithDurationSeconds(v float32) LiveActivityStreamInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityStreamInput) WithCountsDown(v bool) LiveActivityStreamInput {
	in.CountsDown = v
	in.countsDownSet = true
	return in
}

func (in LiveActivityStreamInput) WithAction(v LiveActivityActionInput) LiveActivityStreamInput {
	in.Action = &v
	return in
}

func (in LiveActivityStreamInput) WithSecondaryAction(v LiveActivityActionInput) LiveActivityStreamInput {
	in.SecondaryAction = &v
	return in
}

// LiveActivityStreamEndInput is an optional payload for ending a managed stream.
type LiveActivityStreamEndInput struct {
	ContentState    LiveActivityContentStateInput
	Title           string
	NumberOfSteps   int32
	CurrentStep     int32
	Percentage      float32
	Value           float32
	UpperLimit      float32
	DurationSeconds float32
	CountsDown      bool
	Type            string
	Subtitle        string
	Message         string
	Icon            LiveActivityAlertIconInput
	Badge           LiveActivityAlertBadgeInput
	Color           string
	StepColor       string
	Metrics         []generated.ActivityMetric
	Action          *LiveActivityActionInput
	SecondaryAction *LiveActivityActionInput
	Alert           *generated.AlertPayload

	numberOfStepsSet   bool
	currentStepSet     bool
	percentageSet      bool
	valueSet           bool
	upperLimitSet      bool
	durationSecondsSet bool
	countsDownSet      bool
}

func (in LiveActivityStreamEndInput) toGenerated() generated.LiveActivityStreamDeleteRequest {
	req := generated.NewLiveActivityStreamDeleteRequest()
	contentStateInput := in.ContentState
	if !contentStateInput.isSet() && in.Title != "" {
		contentStateInput = LiveActivityContentStateInput{
			Title:   in.Title,
			Type:    in.Type,
			Message: in.Message,
			Icon:    in.Icon,
			Badge:   in.Badge,
		}
	}
	if contentStateInput.isSet() {
		contentState := *generated.NewStreamContentState(contentStateInput.Title)
		contentStateInput.applyStream(&contentState)
		if in.NumberOfSteps != 0 || in.numberOfStepsSet {
			contentState.SetNumberOfSteps(in.NumberOfSteps)
		}
		if in.CurrentStep != 0 || in.currentStepSet {
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
		if in.DurationSeconds != 0 || in.durationSecondsSet {
			setAdditionalProperty(&contentState.AdditionalProperties, "duration_seconds", in.DurationSeconds)
		}
		if in.countsDownSet {
			setAdditionalProperty(&contentState.AdditionalProperties, "counts_down", in.CountsDown)
		}
		if in.Type != "" {
			contentState.SetType(in.Type)
		}
		if in.Subtitle != "" {
			contentState.SetSubtitle(in.Subtitle)
		}
		if in.Message != "" {
			setAdditionalProperty(&contentState.AdditionalProperties, "message", in.Message)
		}
		setAdditionalProperty(&contentState.AdditionalProperties, "icon", alertIconMap(in.Icon))
		setAdditionalProperty(&contentState.AdditionalProperties, "badge", alertBadgeMap(in.Badge))
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
	if in.SecondaryAction != nil {
		req.SetSecondaryAction(in.SecondaryAction.toGenerated())
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

func (in LiveActivityStreamEndInput) WithCurrentStep(v int32) LiveActivityStreamEndInput {
	in.CurrentStep = v
	in.currentStepSet = true
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

func (in LiveActivityStreamEndInput) WithDurationSeconds(v float32) LiveActivityStreamEndInput {
	in.DurationSeconds = v
	in.durationSecondsSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithCountsDown(v bool) LiveActivityStreamEndInput {
	in.CountsDown = v
	in.countsDownSet = true
	return in
}

func (in LiveActivityStreamEndInput) WithAction(v LiveActivityActionInput) LiveActivityStreamEndInput {
	in.Action = &v
	return in
}

func (in LiveActivityStreamEndInput) WithSecondaryAction(v LiveActivityActionInput) LiveActivityStreamEndInput {
	in.SecondaryAction = &v
	return in
}
