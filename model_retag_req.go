/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RetagReq struct {
	// new tag to be created
	Tag string `json:"tag,omitempty"`
	// Source image to be retagged, e.g. 'stage/app:v1.0'
	SrcImage string `json:"src_image,omitempty"`
	// If target tag already exists, whether to override it
	Override bool `json:"override,omitempty"`
}