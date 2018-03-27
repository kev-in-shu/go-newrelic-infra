package api

import "errors"

var (
	// ErrNotFound is returned when the resource was not found in New Relic.
	ErrNotFound = errors.New("newrelic-infra: Resource not found")
)

// AlertInfraThreshold represents an Infra alerting condition
type AlertInfraThreshold struct {
	Value    int    `json:"value,omitempty"`
	Duration int    `json:"duration_minutes,omitempty"`
	Function string `json:"time_function,omitempty"`
}

// AlertInfraCondition represents a New Relic Infra Alert condition.
type AlertInfraCondition struct {
	PolicyID     int                  `json:"policy_id,omitempty"`
	ID           int                  `json:"id,omitempty"`
	Name         string               `json:"name,omitempty"`
	Type         string               `json:"type,omitempty"`
	Comparison   string               `json:"comparison,omitempty"`
	CreatedAt    int                  `json:"created_at_epoch_millis,omitempty"`
	UpdatedAt    int                  `json:"updated_at_epoch_millis,omitempty"`
	Enabled      bool                 `json:"enabled,omitempty"`
	Event        string               `json:"event_type,omitempty"`
	Select       string               `json:"select_value,omitempty"`
	Where        string               `json:"where_clause,omitempty"`
	ProcessWhere string               `json:"process_where_clause,omitempty"`
	Warning      *AlertInfraThreshold `json:"warning_threshold,omitempty"`
	Critical     *AlertInfraThreshold `json:"critical_threshold,omitempty"`
}
