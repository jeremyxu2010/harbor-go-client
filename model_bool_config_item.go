/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BoolConfigItem struct {
	// The boolean value of current config item
	Value bool `json:"value,omitempty"`
	// The configure item can be updated or not
	Editable bool `json:"editable,omitempty"`
}
