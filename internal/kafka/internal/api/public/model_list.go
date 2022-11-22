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

// List struct for List
type List struct {
	Kind  string            `json:"kind"`
	Page  int32             `json:"page"`
	Size  int32             `json:"size"`
	Total int32             `json:"total"`
	Items []ObjectReference `json:"items"`
}
