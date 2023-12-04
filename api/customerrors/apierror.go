package customerrors

import "fmt"

// NotFoundError representa un error cuando un registro no se encuentra en la base de datos.
type NotFoundError struct {
	Resource string
	ID       string
}

// NewNotFoundError crea un nuevo error de "no encontrado" con el recurso y el ID especificados.
func NewNotFoundError(resource, id string) *NotFoundError {
	return &NotFoundError{
		Resource: resource,
		ID:       id,
	}
}

// Error implementa la interfaz error para NotFoundError.
func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s con ID %s no encontrado", e.Resource, e.ID)
}

// AlreadyExistsError representa un error cuando se intenta crear un registro que ya existe en la base de datos.
type AlreadyExistsError struct {
	Resource string
	ID       string
}

// NewAlreadyExistsError crea un nuevo error de "ya existe" con el recurso y el ID especificados.
func NewAlreadyExistsError(resource, id string) *AlreadyExistsError {
	return &AlreadyExistsError{
		Resource: resource,
		ID:       id,
	}
}

// Error implementa la interfaz error para AlreadyExistsError.
func (e *AlreadyExistsError) Error() string {
	return fmt.Sprintf("%s con ID %s ya existe", e.Resource, e.ID)
}

// ValidationError representa un error cuando los datos ingresados no pasan la validación.
type ValidationError struct {
	Message string
}

// NewValidationError crea un nuevo error de validación con el mensaje especificado.
func NewValidationError(message string) *ValidationError {
	return &ValidationError{Message: message}
}

// Error implementa la interfaz error para ValidationError.
func (e *ValidationError) Error() string {
	return e.Message
}
