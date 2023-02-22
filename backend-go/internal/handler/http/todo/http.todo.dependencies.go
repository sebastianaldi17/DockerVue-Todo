package todo

// go:generate mockgen -build_flags=-mod=mod -source=http.todo.dependencies.go -package=todo -destination=http.dependencies.mock_test.go
type todoUsecase interface {
}
