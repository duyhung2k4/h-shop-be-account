package controller

type accessController struct{}

type AccessController interface {
}

func NewAccess() AccessController {
	return &accessController{}
}
