package errors

// ServiceError used to return business error messages
type ServiceError struct {
	Message string `json:"message"`
}
