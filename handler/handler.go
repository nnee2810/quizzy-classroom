package handler

type Handler interface {
}

type handlerImpl struct {
}

type HandlerInject struct {
}

func New(inject *HandlerInject) Handler {
	return &handlerImpl{}
}
