package validators

type Validator struct {
	Error *ValidationError
}

func GetValidator() *Validator {
	return &Validator{
		Error: &ValidationError{},
	}
}