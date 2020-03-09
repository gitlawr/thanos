// Copyright (c) The Thanos Authors.
// Licensed under the Apache License 2.0.

package testpromcompatibility

import (
	"time"

	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/prometheus/prometheus/rules"
)

type RuleDiscovery struct {
	RuleGroups []*RuleGroup `json:"groups"`
}

// Same as https://github.com/prometheus/prometheus/blob/c530b4b456cc5f9ec249f771dff187eb7715dc9b/web/api/v1/api.go#L955
// but with Partial Response.
type RuleGroup struct {
	Name           string    `json:"name"`
	File           string    `json:"file"`
	Rules          []Rule    `json:"rules"`
	Interval       float64   `json:"interval"`
	EvaluationTime float64   `json:"evaluationTime"`
	LastEvaluation time.Time `json:"lastEvaluation"`

	DeprecatedPartialResponseStrategy string `json:"partial_response_strategy"`
	PartialResponseStrategy           string `json:"partialResponseStrategy"`
}

type Rule interface{}

type AlertAPI struct {
	Alerts []*Alert `json:"alerts"`
}

type Alert struct {
	Labels      labels.Labels `json:"labels"`
	Annotations labels.Labels `json:"annotations"`
	State       string        `json:"state"`
	ActiveAt    *time.Time    `json:"activeAt,omitempty"`
	Value       string        `json:"value"`

	PartialResponseStrategy string `json:"partialResponseStrategy"`
}

type AlertingRule struct {
	// State can be "pending", "firing", "inactive".
	State          string           `json:"state"`
	Name           string           `json:"name"`
	Query          string           `json:"query"`
	Duration       float64          `json:"duration"`
	Labels         labels.Labels    `json:"labels"`
	Annotations    labels.Labels    `json:"annotations"`
	Alerts         []*Alert         `json:"alerts"`
	Health         rules.RuleHealth `json:"health"`
	LastError      string           `json:"lastError,omitempty"`
	EvaluationTime float64          `json:"evaluationTime"`
	LastEvaluation time.Time        `json:"lastEvaluation"`
	// Type of an AlertingRule is always "alerting".
	Type string `json:"type"`
}

type RecordingRule struct {
	Name           string           `json:"name"`
	Query          string           `json:"query"`
	Labels         labels.Labels    `json:"labels,omitempty"`
	Health         rules.RuleHealth `json:"health"`
	LastError      string           `json:"lastError,omitempty"`
	EvaluationTime float64          `json:"evaluationTime"`
	LastEvaluation time.Time        `json:"lastEvaluation"`
	// Type of a prometheusRecordingRule is always "recording".
	Type string `json:"type"`
}
