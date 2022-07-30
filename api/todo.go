package api

import (
	"context"
	"log"

	todo "github.com/morning-night-dream/play-cicd/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(logger *log.Logger) todo.Service {
	return &todosrvc{logger}
}

// ReadTodoList implements ReadTodoList.
func (s *todosrvc) ReadTodoList(ctx context.Context) (*todo.ReadTodoListResponseBody, error) {
	res := &todo.ReadTodoListResponseBody{}

	s.logger.Print("todo.ReadTodoList")

	return res, nil
}

// CreateTodo implements CreateTodo.
func (s *todosrvc) CreateTodo(
	ctx context.Context,
	p *todo.CreateTodoRequestBody,
) (*todo.CreateTodoResponseBody, error) {
	res := &todo.CreateTodoResponseBody{}

	s.logger.Print("todo.CreateTodo")

	return res, nil
}

// UpdateTodo implements UpdateTodo.
func (s *todosrvc) UpdateTodo(
	ctx context.Context,
	p *todo.UpdateTodoPayload,
) (*todo.UpdateTodoResponseBody, error) {
	res := &todo.UpdateTodoResponseBody{}

	s.logger.Print("todo.UpdateTodo")

	return res, nil
}

// DeleteTodo implements DeleteTodo.
func (s *todosrvc) DeleteTodo(ctx context.Context, p *todo.DeleteTodoPayload) error {
	s.logger.Print("todo.DeleteTodo")

	return nil
}
