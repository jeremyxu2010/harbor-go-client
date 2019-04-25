/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Role struct {
	// ID in table.
	RoleId int32 `json:"role_id,omitempty"`
	// Description of permissions for the role.
	RoleCode string `json:"role_code,omitempty"`
	// Name the the role.
	RoleName string `json:"role_name,omitempty"`
	RoleMask string `json:"role_mask,omitempty"`
}
