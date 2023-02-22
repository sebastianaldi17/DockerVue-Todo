package todo

// go:generate mockgen -build_flags=-mod=mod -source=uc.todo.dependencies.go -package=todo -destination=uc.dependencies.mock_test.go
type todoResource interface {
}
