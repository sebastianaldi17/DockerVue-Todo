package todo

type Usecase struct {
	todoRes todoResource
}

func New(todoRes todoResource) *Usecase {
	todoUsecase := Usecase{
		todoRes: todoRes,
	}
	return &todoUsecase
}
