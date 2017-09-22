package error

type StreamNotSupportedError struct {
	message string
}

func  NewStreamNotSupportedError(message string) * StreamNotSupportedError{
	err := StreamNotSupportedError{message:message}
	return &err
}

func ( e *StreamNotSupportedError) Error() string{
	return e.message
}