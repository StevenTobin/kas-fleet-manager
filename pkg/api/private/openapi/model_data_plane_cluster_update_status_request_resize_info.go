/*
 * Managed Service API
 *
 * Managed Service API
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DataPlaneClusterUpdateStatusRequestResizeInfo struct for DataPlaneClusterUpdateStatusRequestResizeInfo
type DataPlaneClusterUpdateStatusRequestResizeInfo struct {
	NodeDelta int32                                              `json:"nodeDelta,omitempty"`
	Delta     DataPlaneClusterUpdateStatusRequestResizeInfoDelta `json:"delta,omitempty"`
}