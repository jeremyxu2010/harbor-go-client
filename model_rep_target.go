/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type RepTarget struct {
	// The target ID.
	Id int64 `json:"id,omitempty"`
	// The target address URL string.
	Endpoint string `json:"endpoint,omitempty"`
	// The target name.
	Name string `json:"name,omitempty"`
	// The target server username.
	Username string `json:"username,omitempty"`
	// The target server password.
	Password string `json:"password,omitempty"`
	// Reserved field.
	Type_ int32 `json:"type,omitempty"`
	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`
	// The create time of the policy.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of the policy.
	UpdateTime string `json:"update_time,omitempty"`
}
