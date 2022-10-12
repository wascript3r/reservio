package errcode

type ErrName string

type Error struct {
	name     ErrName
	original error
	data     any
}

func New(name ErrName, original error) *Error {
	return &Error{
		name:     name,
		original: original,
	}
}

func (e *Error) Error() string {
	return e.original.Error()
}

func (e *Error) Name() string {
	return string(e.name)
}

func (e *Error) Original() error {
	return e.original
}

func (e *Error) Data() any {
	return e.data
}

func (e *Error) SetData(data any) *Error {
	return &Error{
		name:     e.name,
		original: e.original,
		data:     data,
	}
}

func UnwrapErr(err error, defaultErr *Error) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return defaultErr
}
