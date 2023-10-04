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

type DefaultRouteObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DefaultRouteParameters struct {

	// The ID of the escalation chain.
	// +kubebuilder:validation:Optional
	EscalationChainID *string `json:"escalationChainId,omitempty" tf:"escalation_chain_id,omitempty"`

	// MS teams-specific settings for a route.
	// +kubebuilder:validation:Optional
	Msteams []MsteamsParameters `json:"msteams,omitempty" tf:"msteams,omitempty"`

	// Slack-specific settings for a route.
	// +kubebuilder:validation:Optional
	Slack []SlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// Telegram-specific settings for a route.
	// +kubebuilder:validation:Optional
	Telegram []TelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`
}

type EmailObservation struct {
}

type EmailParameters struct {

	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type IntegrationObservation struct {

	// The Default route for all alerts from the given integration
	// +kubebuilder:validation:Required
	DefaultRoute []DefaultRouteObservation `json:"defaultRoute,omitempty" tf:"default_route,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The link for using in an integrated tool.
	Link *string `json:"link,omitempty" tf:"link,omitempty"`
}

type IntegrationParameters struct {

	// The Default route for all alerts from the given integration
	// +kubebuilder:validation:Required
	DefaultRoute []DefaultRouteParameters `json:"defaultRoute" tf:"default_route,omitempty"`

	// The name of the service integration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the OnCall team. To get one, create a team in Grafana, and navigate to the OnCall plugin (to sync the team with OnCall). You can then get the ID using the `grafana_oncall_team` datasource.
	// +kubebuilder:validation:Optional
	TeamID *string `json:"teamId,omitempty" tf:"team_id,omitempty"`

	// Jinja2 templates for Alert payload.
	// +kubebuilder:validation:Optional
	Templates []TemplatesParameters `json:"templates,omitempty" tf:"templates,omitempty"`

	// The type of integration. Can be grafana, grafana_alerting, webhook, alertmanager, kapacitor, fabric, newrelic, datadog, pagerduty, pingdom, elastalert, amazon_sns, curler, sentry, formatted_webhook, heartbeat, demo, manual, stackdriver, uptimerobot, sentry_platform, zabbix, prtg, slack_channel, inbound_email, direct_paging.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type MicrosoftTeamsObservation struct {
}

type MicrosoftTeamsParameters struct {

	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type MsteamsObservation struct {
}

type MsteamsParameters struct {

	// Enable notification in MS teams. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// MS teams channel id. Alerts will be directed to this channel in Microsoft teams.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PhoneCallObservation struct {
}

type PhoneCallParameters struct {

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SMSObservation struct {
}

type SMSParameters struct {

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type SlackObservation struct {
}

type SlackParameters struct {

	// Slack channel id. Alerts will be directed to this channel in Slack.
	// +kubebuilder:validation:Optional
	ChannelID *string `json:"channelId,omitempty" tf:"channel_id,omitempty"`

	// Enable notification in Slack. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
}

type TelegramObservation struct {
}

type TelegramParameters struct {

	// Enable notification in Telegram. Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Telegram channel id. Alerts will be directed to this channel in Telegram.
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type TemplatesObservation struct {
}

type TemplatesParameters struct {

	// Template for sending a signal to acknowledge the Incident.
	// +kubebuilder:validation:Optional
	AcknowledgeSignal *string `json:"acknowledgeSignal,omitempty" tf:"acknowledge_signal,omitempty"`

	// Templates for Email.
	// +kubebuilder:validation:Optional
	Email []EmailParameters `json:"email,omitempty" tf:"email,omitempty"`

	// Template for the key by which alerts are grouped.
	// +kubebuilder:validation:Optional
	GroupingKey *string `json:"groupingKey,omitempty" tf:"grouping_key,omitempty"`

	// Templates for Microsoft Teams.
	// +kubebuilder:validation:Optional
	MicrosoftTeams []MicrosoftTeamsParameters `json:"microsoftTeams,omitempty" tf:"microsoft_teams,omitempty"`

	// Templates for Phone Call.
	// +kubebuilder:validation:Optional
	PhoneCall []PhoneCallParameters `json:"phoneCall,omitempty" tf:"phone_call,omitempty"`

	// Template for sending a signal to resolve the Incident.
	// +kubebuilder:validation:Optional
	ResolveSignal *string `json:"resolveSignal,omitempty" tf:"resolve_signal,omitempty"`

	// Templates for SMS.
	// +kubebuilder:validation:Optional
	SMS []SMSParameters `json:"sms,omitempty" tf:"sms,omitempty"`

	// Templates for Slack.
	// +kubebuilder:validation:Optional
	Slack []TemplatesSlackParameters `json:"slack,omitempty" tf:"slack,omitempty"`

	// Template for a source link.
	// +kubebuilder:validation:Optional
	SourceLink *string `json:"sourceLink,omitempty" tf:"source_link,omitempty"`

	// Templates for Telegram.
	// +kubebuilder:validation:Optional
	Telegram []TemplatesTelegramParameters `json:"telegram,omitempty" tf:"telegram,omitempty"`

	// Templates for Web.
	// +kubebuilder:validation:Optional
	Web []WebParameters `json:"web,omitempty" tf:"web,omitempty"`
}

type TemplatesSlackObservation struct {
}

type TemplatesSlackParameters struct {

	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type TemplatesTelegramObservation struct {
}

type TemplatesTelegramParameters struct {

	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type WebObservation struct {
}

type WebParameters struct {

	// Template for Alert image url.
	// +kubebuilder:validation:Optional
	ImageURL *string `json:"imageUrl,omitempty" tf:"image_url,omitempty"`

	// Template for Alert message.
	// +kubebuilder:validation:Optional
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	// Template for Alert title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// IntegrationSpec defines the desired state of Integration
type IntegrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IntegrationParameters `json:"forProvider"`
}

// IntegrationStatus defines the observed state of Integration.
type IntegrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IntegrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Integration is the Schema for the Integrations API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationSpec   `json:"spec"`
	Status            IntegrationStatus `json:"status,omitempty"`
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
