package sentry

import (
	"fmt"
	"net/http"
	"time"
)

type alertRuleMatchPolicy string

var (
	AlertRuleMatchAll  = alertRuleMatchPolicy("all")
	AlertRuleMatchAny  = alertRuleMatchPolicy("any")
	AlertRuleMatchNone = alertRuleMatchPolicy("none")
)

// AlertRuleCondition represents alert rule condition.
// Refer to https://github.com/getsentry/sentry/tree/master/src/sentry/rules/conditions or GUI
// to get detailed information.
type AlertRuleCondition struct {
	ID        string      `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Interval  string      `json:"interval,omitempty"` // 1m, 1w, 30d etc
	Value     interface{} `json:"value,omitempty"`
	Attribute string      `json:"attribute,omitempty"`
	Key       string      `json:"key,omitempty"`
}

// AlertRuleAction represents alert rule action.
// Refer to https://github.com/getsentry/sentry/tree/master/src/sentry/rules/actions or GUI
// to get detailed information.
type AlertRuleAction struct {
	ID               string `json:"id,omitempty"`
	TargetType       string `json:"targetType,omitempty"`
	FallThroughType  string `json:"fallthroughType,omitempty"`
	TargetIdentifier uint   `json:"targetIdentifier,omitempty"`
	Name             string `json:"name,omitempty"`
	Service          string `json:"service,omitempty"`
	Workspace        string `json:"workspace,omitempty"`
	Channel          string `json:"channel,omitempty"`
	ChannelID        string `json:"channel_id,omitempty"`
	Tags             string `json:"tags,omitempty"`
	Account          string `json:"account,omitempty"`
}

// AlertFilter represents alert filter
// Refer to https://github.com/getsentry/sentry/tree/master/src/sentry/rules/filters or GUI
// to get detailed information.
type AlertFilters struct {
	ID               string      `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	Attribute        string      `json:"attribute,omitempty"`
	Match            string      `json:"match,omitempty"`
	Value            interface{} `json:"value,omitempty"`
	Level            string      `json:"level,omitempty"`
	ComparisonType   string      `json:"comparison_type,omitempty"`
	Time             string      `json:"time,omitempty"`
	TargetType       string      `json:"targetType,omitempty"`
	TargetIdentifier string      `json:"targetIdentifier,omitempty"`
	Key              string      `json:"key,omitempty"`
}

type AlertRule struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	DateCreated *time.Time           `json:"dateCreated,omitempty"`
	Environment *string              `json:"environment,omitempty"`
	ActionMatch alertRuleMatchPolicy `json:"actionMatch,omitempty"`
	FilterMatch alertRuleMatchPolicy `json:"filterMatch,omitempty"`
	Frequency   uint                 `json:"frequency,omitempty"` // run actions at most once every Frequency minutes
	Conditions  []AlertRuleCondition `json:"conditions,omitempty"`
	Filters     []AlertFilters       `json:"filters,omitempty"`
	Actions     []AlertRuleAction    `json:"actions,omitempty"`
	Owner       string               `json:"owner,omitempty"`
}

func (c *Client) GetAlertRules(o Organization, p Project) ([]AlertRule, *Link, error) {
	var rules []AlertRule
	link, err := c.doWithPagination(http.MethodGet, fmt.Sprintf("projects/%s/%s/rules/", *o.Slug, *p.Slug), &rules, nil)

	return rules, link, err
}

func (c *Client) AddAlertRule(o Organization, p Project, r AlertRule) (AlertRule, error) {
	err := c.do(http.MethodPost, fmt.Sprintf("projects/%s/%s/rules/", *o.Slug, *p.Slug), &r, r)

	return r, err
}

func (c *Client) UpdateAlertRule(o Organization, p Project, r AlertRule) (AlertRule, error) {
	err := c.do(http.MethodPut, fmt.Sprintf("projects/%s/%s/rules/%s/", *o.Slug, *p.Slug, r.ID), &r, r)

	return r, err
}

func (c *Client) DeleteAlertRule(o Organization, p Project, r AlertRule) error {
	return c.do(http.MethodDelete, fmt.Sprintf("projects/%s/%s/rules/%s/", *o.Slug, *p.Slug, r.ID), nil, nil)
}
