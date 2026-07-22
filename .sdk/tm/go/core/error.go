package core

type InfranodeOpenDataError struct {
	IsInfranodeOpenDataError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewInfranodeOpenDataError(code string, msg string, ctx *Context) *InfranodeOpenDataError {
	return &InfranodeOpenDataError{
		IsInfranodeOpenDataError: true,
		Sdk:              "InfranodeOpenData",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *InfranodeOpenDataError) Error() string {
	return e.Msg
}
