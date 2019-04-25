/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Manifest struct {
	// The detail of manifest.
	Manifest *interface{} `json:"manifest,omitempty"`
	// The config of the repository.
	Config string `json:"config,omitempty"`
}
