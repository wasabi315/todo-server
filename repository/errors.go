package repository

type ErrorCode uint8

const (
	ErrNotFound ErrorCode = iota
)

// TODO: better name
type GeneralError struct {
	Code ErrorCode
	Msg  string
}

func (g *GeneralError) Error() string {
	return g.Msg
}

func NotFound(msg string) *GeneralError {
	return &GeneralError{ErrNotFound, msg}
}
