/*
 * Connector Service Fleet Manager Admin APIs
 *
 * Connector Service Fleet Manager Admin is a Rest API to manage connector clusters.
 *
 * API version: 0.0.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package private

// ConnectorClusterPlatform information about the kubernetes platform
type ConnectorClusterPlatform struct {
	// the kubernetes cluster type
	Type string `json:"type,omitempty"`
	// uniquely identifies the kubernetes cluster
	Id string `json:"id,omitempty"`
	// optional version of the kubernetes cluster
	Version string `json:"version,omitempty"`
}