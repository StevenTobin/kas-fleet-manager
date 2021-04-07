/*
 * Kafka Service Fleet Manager
 *
 * Kafka Service Fleet Manager is a Rest API to manage kafka instances and connectors.
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ServiceStatusKafkas The kafka resource api status
type ServiceStatusKafkas struct {
	// Indicates whether we have reached kafka maximum capacity
	MaxCapacityReached bool `json:"max_capacity_reached"`
}