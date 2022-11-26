package es

import "context"

type DomainEvent interface {
}

type Applier[P any] interface {
	*P
	Apply(DomainEvent) error
}

type Handler[T any] interface {
	Handle(context.Context, string, T) error
}

type HandlerFunc[T any] func(context.Context, string, T) error

func (f HandlerFunc[T]) Handle(ctx context.Context, aggregateID string, aggregate T) error {
	return f(ctx, aggregateID, aggregate)
}

type Emitter interface {
	Emit(ctx context.Context, events ...DomainEvent) error
}
