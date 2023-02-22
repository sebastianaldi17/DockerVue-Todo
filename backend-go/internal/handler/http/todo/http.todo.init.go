package todo

type Handler struct {
	todoUsecase todoUsecase
}

func New(todoUsecase todoUsecase) *Handler {
	handler := Handler{
		todoUsecase: todoUsecase,
	}
	return &handler
}
