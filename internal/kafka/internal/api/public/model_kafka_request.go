/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage Kafka instances.
 *
 * API version: 1.9.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

import (
	"time"
)

// KafkaRequest struct for KafkaRequest
type KafkaRequest struct {
	Id   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Href string `json:"href,omitempty"`
	// Values: [accepted, preparing, provisioning, ready, failed, deprovision, deleting]
	Status string `json:"status,omitempty"`
	// Name of Cloud used to deploy. For example AWS
	CloudProvider string `json:"cloud_provider,omitempty"`
	MultiAz       bool   `json:"multi_az"`
	// Values will be regions of specific cloud provider. For example: us-east-1 for AWS
	Region              string `json:"region,omitempty"`
	Owner               string `json:"owner,omitempty"`
	Name                string `json:"name,omitempty"`
	BootstrapServerHost string `json:"bootstrap_server_host,omitempty"`
	// The kafka admin server url to perform kafka admin operations e.g acl management etc. The value will be available when the Kafka has been fully provisioned i.e it reaches a 'ready' state
	AdminApiServerUrl           string     `json:"admin_api_server_url,omitempty"`
	CreatedAt                   time.Time  `json:"created_at,omitempty"`
	ExpiresAt                   *time.Time `json:"expires_at,omitempty"`
	UpdatedAt                   time.Time  `json:"updated_at,omitempty"`
	FailedReason                string     `json:"failed_reason,omitempty"`
	Version                     string     `json:"version,omitempty"`
	InstanceType                string     `json:"instance_type,omitempty"`
	InstanceTypeName            string     `json:"instance_type_name,omitempty"`
	ReauthenticationEnabled     bool       `json:"reauthentication_enabled"`
	KafkaStorageSize            string     `json:"kafka_storage_size,omitempty"`
	BrowserUrl                  string     `json:"browser_url,omitempty"`
	SizeId                      string     `json:"size_id,omitempty"`
	IngressThroughputPerSec     string     `json:"ingress_throughput_per_sec,omitempty"`
	EgressThroughputPerSec      string     `json:"egress_throughput_per_sec,omitempty"`
	TotalMaxConnections         int32      `json:"total_max_connections,omitempty"`
	MaxPartitions               int32      `json:"max_partitions,omitempty"`
	MaxDataRetentionPeriod      string     `json:"max_data_retention_period,omitempty"`
	MaxConnectionAttemptsPerSec int32      `json:"max_connection_attempts_per_sec,omitempty"`
	BillingCloudAccountId       string     `json:"billing_cloud_account_id,omitempty"`
	Marketplace                 string     `json:"marketplace,omitempty"`
}
