/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The overview of the scan result.  This is an optional property.
type DetailedTagScanOverview struct {
	// The digest of the image.
	Digest string `json:"digest,omitempty"`
	// The status of the scan job, it can be \"pendnig\", \"running\", \"finished\", \"error\".
	ScanStatus string `json:"scan_status,omitempty"`
	// The ID of the job on jobservice to scan the image.
	JobId int32 `json:"job_id,omitempty"`
	// 0-Not scanned, 1-Negligible, 2-Unknown, 3-Low, 4-Medium, 5-High
	Severity int32 `json:"severity,omitempty"`
	// The top layer name of this image in Clair, this is for calling Clair API to get the vulnerability list of this image.
	DetailsKey string `json:"details_key,omitempty"`
	Components *DetailedTagScanOverviewComponents `json:"components,omitempty"`
}
