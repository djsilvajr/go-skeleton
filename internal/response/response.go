package response

import "github.com/gin-gonic/gin"

// errorBody follows the mandatory error contract from api-patterns.md.
type errorBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

// Success wraps any payload under the "data" key and sends the given HTTP status.
func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data})
}

// Error sends a structured error response following the mandatory format:
//
//	{ "error": { "code": <status>, "message": "<msg>", "details": <details> } }
//
// Pass nil for details to get an empty object.
func Error(c *gin.Context, status int, message string, details interface{}) {
	if details == nil {
		details = gin.H{}
	}
	c.JSON(status, gin.H{
		"error": errorBody{
			Code:    status,
			Message: message,
			Details: details,
		},
	})
}

// Deleted returns the standard soft-delete confirmation response.
func Deleted(c *gin.Context) {
	Success(c, 200, gin.H{"deleted": true})
}
