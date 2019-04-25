/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Bad request
type BadRequestFormatedError struct {
	// The error message returned by the chart API
	Error_ string `json:"error"`
}
