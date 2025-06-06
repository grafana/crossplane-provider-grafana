/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DefaultRouteInitParameters struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.EscalationChain
	// +crossplane:generate:reference:refFieldName=EscalationChainRef
	// +crossplane:generate:reference:selectorFieldName=EscalationChainSelector
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// Reference to a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainRef *v1.Reference `json:"escalationChainRef,omitempty" tf:"-"`

	// Selector for a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainSelector *v1.Selector `json:"escalationChainSelector,omitempty" tf:"-"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	Msteams []MsteamsInitParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	Slack []SlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	Telegram []TelegramInitParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type DefaultRouteObservation struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	Msteams []MsteamsObservation `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	Slack []SlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	Telegram []TelegramObservation `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type DefaultRouteParameters struct {

	// (String) The ID of the escalation chain.
	// The ID of the escalation chain.
	// +crossplane:generate:reference:type=github.com/grafana/crossplane-provider-grafana/apis/oncall/v1alpha1.EscalationChain
	// +crossplane:generate:reference:refFieldName=EscalationChainRef
	// +crossplane:generate:reference:selectorFieldName=EscalationChainSelector
	// +kubebuilder:validation:Optional
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// Reference to a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainRef *v1.Reference `json:"escalationChainRef,omitempty" tf:"-"`

	// Selector for a EscalationChain in oncall to populate escalationChainId.
	// +kubebuilder:validation:Optional
	EscalationChainSelector *v1.Selector `json:"escalationChainSelector,omitempty" tf:"-"`

	// specific settings for a route. (see below for nested schema)
	// MS teams-specific settings for a route.
	// +kubebuilder:validation:Optional
	Msteams []MsteamsParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Slack-specific settings for a route.
	// +kubebuilder:validation:Optional
	Slack []SlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Telegram-specific settings for a route.
	// +kubebuilder:validation:Optional
	Telegram []TelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type EmailInitParameters struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type EmailObservation struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type EmailParameters struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type IntegrationInitParameters struct {

	// (Block List, Min: 1, Max: 1) The Default route for all alerts from the given integration (see below for nested schema)
	// The Default route for all alerts from the given integration
	DefaultRoute []DefaultRouteInitParameters `json:"defaultRoute,omitempty" tf:"default_route,omitempty"`

	// to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	DynamicLabels []map[string]*string `json:"dynamicLabels,omitempty" tf:"dynamic_labels,omitempty"`

	// to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	Labels []map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the service integration.
	// The name of the service integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team (using the grafana_oncall_team datasource).
	// The ID of the OnCall team (using the `grafana_oncall_team` datasource).
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (Block List, Max: 1) Jinja2 templates for Alert payload. An empty templates block will be ignored. (see below for nested schema)
	// Jinja2 templates for Alert payload. An empty templates block will be ignored.
	Templates []TemplatesInitParameters `json:"templates,omitempty" tf:"templates,omitempty"`

	// (String) The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	// The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IntegrationObservation struct {

	// (Block List, Min: 1, Max: 1) The Default route for all alerts from the given integration (see below for nested schema)
	// The Default route for all alerts from the given integration
	DefaultRoute []DefaultRouteObservation `json:"defaultRoute,omitempty" tf:"default_route,omitempty"`

	// to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	DynamicLabels []map[string]*string `json:"dynamicLabels,omitempty" tf:"dynamic_labels,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	Labels []map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The link for using in an integrated tool.
	// The link for using in an integrated tool.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`

	// (String) The name of the service integration.
	// The name of the service integration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team (using the grafana_oncall_team datasource).
	// The ID of the OnCall team (using the `grafana_oncall_team` datasource).
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (Block List, Max: 1) Jinja2 templates for Alert payload. An empty templates block will be ignored. (see below for nested schema)
	// Jinja2 templates for Alert payload. An empty templates block will be ignored.
	Templates []TemplatesObservation `json:"templates,omitempty" tf:"templates,omitempty"`

	// (String) The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	// The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IntegrationParameters struct {

	// (Block List, Min: 1, Max: 1) The Default route for all alerts from the given integration (see below for nested schema)
	// The Default route for all alerts from the given integration
	// +kubebuilder:validation:Optional
	DefaultRoute []DefaultRouteParameters `json:"defaultRoute,omitempty" tf:"default_route,omitempty"`

	// to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for dynamic labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	// +kubebuilder:validation:Optional
	DynamicLabels []map[string]*string `json:"dynamicLabels,omitempty" tf:"dynamic_labels,omitempty"`

	// to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the grafana_oncall_label datasource).
	// A list of string-to-string mappings for static labels. Each map must include one key named "key" and one key named "value" (using the `grafana_oncall_label` datasource).
	// +kubebuilder:validation:Optional
	Labels []map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the service integration.
	// The name of the service integration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ID of the OnCall team (using the grafana_oncall_team datasource).
	// The ID of the OnCall team (using the `grafana_oncall_team` datasource).
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// (Block List, Max: 1) Jinja2 templates for Alert payload. An empty templates block will be ignored. (see below for nested schema)
	// Jinja2 templates for Alert payload. An empty templates block will be ignored.
	// +kubebuilder:validation:Optional
	Templates []TemplatesParameters `json:"templates,omitempty" tf:"templates,omitempty"`

	// (String) The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	// The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging, jira, zendesk.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type MicrosoftTeamsInitParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MicrosoftTeamsObservation struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MicrosoftTeamsParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MobileAppInitParameters struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MobileAppObservation struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MobileAppParameters struct {

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MsteamsInitParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MsteamsObservation struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MsteamsParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in MS teams. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PhoneCallInitParameters struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PhoneCallObservation struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type PhoneCallParameters struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SMSInitParameters struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SMSObservation struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SMSParameters struct {

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SlackInitParameters struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SlackObservation struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type SlackParameters struct {

	// (String) Slack channel id. Alerts will be directed to this channel in Slack.
	// Slack channel id. Alerts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Slack. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TelegramInitParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TelegramObservation struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TelegramParameters struct {

	// (Boolean) Enable notification in MS teams. Defaults to true.
	// Enable notification in Telegram. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The ID of this resource.
	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TemplatesInitParameters struct {

	// (String) Template for sending a signal to acknowledge the Incident.
	// Template for sending a signal to acknowledge the Incident.
	AcknowledgeSignal *string `json:"acknowledgeSignal,omitempty" tf:"acknowledge_signal,omitempty"`

	// (Block List, Max: 1) Templates for Email. (see below for nested schema)
	// Templates for Email.
	Email []EmailInitParameters `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Template for the key by which alerts are grouped.
	// Template for the key by which alerts are grouped.
	GroupingKey *string `json:"groupingKey,omitempty" tf:"grouping_key,omitempty"`

	// (Block List, Max: 1) Templates for Microsoft Teams. NOTE: Microsoft Teams templates are only available on Grafana Cloud. (see below for nested schema)
	// Templates for Microsoft Teams. **NOTE**: Microsoft Teams templates are only available on Grafana Cloud.
	MicrosoftTeams []MicrosoftTeamsInitParameters `json:"microsoftTeams,omitempty" tf:"microsoft_teams,omitempty"`

	// (Block List, Max: 1) Templates for Mobile app push notifications. (see below for nested schema)
	// Templates for Mobile app push notifications.
	MobileApp []MobileAppInitParameters `json:"mobileApp,omitempty" tf:"mobile_app,omitempty"`

	// (Block List, Max: 1) Templates for Phone Call. (see below for nested schema)
	// Templates for Phone Call.
	PhoneCall []PhoneCallInitParameters `json:"phoneCall,omitempty" tf:"phone_call,omitempty"`

	// (String) Template for sending a signal to resolve the Incident.
	// Template for sending a signal to resolve the Incident.
	ResolveSignal *string `json:"resolveSignal,omitempty" tf:"resolve_signal,omitempty"`

	// (Block List, Max: 1) Templates for SMS. (see below for nested schema)
	// Templates for SMS.
	SMS []SMSInitParameters `json:"sms,omitempty" tf:"sms,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Slack.
	Slack []TemplatesSlackInitParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) Template for a source link.
	// Template for a source link.
	SourceLink *string `json:"sourceLink,omitempty" tf:"source_link,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Telegram.
	Telegram []TemplatesTelegramInitParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`

	// (Block List, Max: 1) Templates for Web. (see below for nested schema)
	// Templates for Web.
	Web []WebInitParameters `json:"web,omitempty" tf:"web,omitempty"`
}

type TemplatesObservation struct {

	// (String) Template for sending a signal to acknowledge the Incident.
	// Template for sending a signal to acknowledge the Incident.
	AcknowledgeSignal *string `json:"acknowledgeSignal,omitempty" tf:"acknowledge_signal,omitempty"`

	// (Block List, Max: 1) Templates for Email. (see below for nested schema)
	// Templates for Email.
	Email []EmailObservation `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Template for the key by which alerts are grouped.
	// Template for the key by which alerts are grouped.
	GroupingKey *string `json:"groupingKey,omitempty" tf:"grouping_key,omitempty"`

	// (Block List, Max: 1) Templates for Microsoft Teams. NOTE: Microsoft Teams templates are only available on Grafana Cloud. (see below for nested schema)
	// Templates for Microsoft Teams. **NOTE**: Microsoft Teams templates are only available on Grafana Cloud.
	MicrosoftTeams []MicrosoftTeamsObservation `json:"microsoftTeams,omitempty" tf:"microsoft_teams,omitempty"`

	// (Block List, Max: 1) Templates for Mobile app push notifications. (see below for nested schema)
	// Templates for Mobile app push notifications.
	MobileApp []MobileAppObservation `json:"mobileApp,omitempty" tf:"mobile_app,omitempty"`

	// (Block List, Max: 1) Templates for Phone Call. (see below for nested schema)
	// Templates for Phone Call.
	PhoneCall []PhoneCallObservation `json:"phoneCall,omitempty" tf:"phone_call,omitempty"`

	// (String) Template for sending a signal to resolve the Incident.
	// Template for sending a signal to resolve the Incident.
	ResolveSignal *string `json:"resolveSignal,omitempty" tf:"resolve_signal,omitempty"`

	// (Block List, Max: 1) Templates for SMS. (see below for nested schema)
	// Templates for SMS.
	SMS []SMSObservation `json:"sms,omitempty" tf:"sms,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Slack.
	Slack []TemplatesSlackObservation `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) Template for a source link.
	// Template for a source link.
	SourceLink *string `json:"sourceLink,omitempty" tf:"source_link,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Telegram.
	Telegram []TemplatesTelegramObservation `json:"telegram,omitempty" tf:"telegram,omitempty"`

	// (Block List, Max: 1) Templates for Web. (see below for nested schema)
	// Templates for Web.
	Web []WebObservation `json:"web,omitempty" tf:"web,omitempty"`
}

type TemplatesParameters struct {

	// (String) Template for sending a signal to acknowledge the Incident.
	// Template for sending a signal to acknowledge the Incident.
	// +kubebuilder:validation:Optional
	AcknowledgeSignal *string `json:"acknowledgeSignal,omitempty" tf:"acknowledge_signal,omitempty"`

	// (Block List, Max: 1) Templates for Email. (see below for nested schema)
	// Templates for Email.
	// +kubebuilder:validation:Optional
	Email []EmailParameters `json:"email,omitempty" tf:"email,omitempty"`

	// (String) Template for the key by which alerts are grouped.
	// Template for the key by which alerts are grouped.
	// +kubebuilder:validation:Optional
	GroupingKey *string `json:"groupingKey,omitempty" tf:"grouping_key,omitempty"`

	// (Block List, Max: 1) Templates for Microsoft Teams. NOTE: Microsoft Teams templates are only available on Grafana Cloud. (see below for nested schema)
	// Templates for Microsoft Teams. **NOTE**: Microsoft Teams templates are only available on Grafana Cloud.
	// +kubebuilder:validation:Optional
	MicrosoftTeams []MicrosoftTeamsParameters `json:"microsoftTeams,omitempty" tf:"microsoft_teams,omitempty"`

	// (Block List, Max: 1) Templates for Mobile app push notifications. (see below for nested schema)
	// Templates for Mobile app push notifications.
	// +kubebuilder:validation:Optional
	MobileApp []MobileAppParameters `json:"mobileApp,omitempty" tf:"mobile_app,omitempty"`

	// (Block List, Max: 1) Templates for Phone Call. (see below for nested schema)
	// Templates for Phone Call.
	// +kubebuilder:validation:Optional
	PhoneCall []PhoneCallParameters `json:"phoneCall,omitempty" tf:"phone_call,omitempty"`

	// (String) Template for sending a signal to resolve the Incident.
	// Template for sending a signal to resolve the Incident.
	// +kubebuilder:validation:Optional
	ResolveSignal *string `json:"resolveSignal,omitempty" tf:"resolve_signal,omitempty"`

	// (Block List, Max: 1) Templates for SMS. (see below for nested schema)
	// Templates for SMS.
	// +kubebuilder:validation:Optional
	SMS []SMSParameters `json:"sms,omitempty" tf:"sms,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Slack.
	// +kubebuilder:validation:Optional
	Slack []TemplatesSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// (String) Template for a source link.
	// Template for a source link.
	// +kubebuilder:validation:Optional
	SourceLink *string `json:"sourceLink,omitempty" tf:"source_link,omitempty"`

	// specific settings for a route. (see below for nested schema)
	// Templates for Telegram.
	// +kubebuilder:validation:Optional
	Telegram []TemplatesTelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`

	// (Block List, Max: 1) Templates for Web. (see below for nested schema)
	// Templates for Web.
	// +kubebuilder:validation:Optional
	Web []WebParameters `json:"web,omitempty" tf:"web,omitempty"`
}

type TemplatesSlackInitParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesSlackObservation struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesSlackParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesTelegramInitParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesTelegramObservation struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesTelegramParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebInitParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebObservation struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebParameters struct {

	// (String) Template for Alert image url.
	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// (String) Template for Alert message.
	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// (String) Template for Alert title.
	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider IntegrationInitParameters `json:"initProvider,omitempty"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Integration is the Schema for the Integrations API. Official documentation https://grafana.com/docs/oncall/latest/configure/integrations/HTTP API https://grafana.com/docs/oncall/latest/oncall-api-reference/
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultRoute) || (has(self.initProvider) && has(self.initProvider.defaultRoute))",message="spec.forProvider.defaultRoute is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   IntegrationSpec   `json:"spec"`
	Status IntegrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IntegrationList contains a list of Integrations
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

// Repository type metadata.
var (
	Integration_Kind             = "Integration"
	Integration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Integration_Kind}.String()
	Integration_KindAPIVersion   = Integration_Kind + "." + CRDGroupVersion.String()
	Integration_GroupVersionKind = CRDGroupVersion.WithKind(Integration_Kind)
)

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
