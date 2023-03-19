package errors

type InvalidFields struct {
	Name        string
	Description string
}

func (instance InvalidFields) GetName() string {
	return instance.Name
}

func (instance InvalidFields) GetDescription() string {
	return instance.Description
}
