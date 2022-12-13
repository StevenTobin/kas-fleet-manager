/*
 * Kafka Management API
 *
 * Kafka Management API is a REST API to manage Kafka instances
 *
 * API version: 1.14.0
 * Contact: rhosak-support@redhat.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package public

// EnterpriseCluster Enterprise cluster registration endpoint response
type EnterpriseCluster struct {
	// OCM cluster id of the registered Enterprise cluster
	ClusterId string `json:"cluster_id,omitempty"`
	// status of registered Enterprise cluster
	Status               string                `json:"status,omitempty"`
	FleetshardParameters []FleetshardParameter `json:"fleetshard_parameters,omitempty"`
}
