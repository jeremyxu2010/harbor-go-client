/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Resource struct {
	// The replication policy list.
	ReplicationPolicies []RepPolicy `json:"replication_policies,omitempty"`
}
