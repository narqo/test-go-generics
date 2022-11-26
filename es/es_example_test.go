package es_test

import (
	"context"
	"time"

	"github.com/narqo/test-go-generics/es"
	"github.com/narqo/test-go-generics/es/account"
)

func Example() {
	ctx := context.Background()

	data := make(map[string][]es.DomainEvent)
	store := make(chan map[string][]es.DomainEvent, 1)
	store <- data

	aggregateID := "120-4123"

	events := make([]es.DomainEvent, 0)
	events = append(events, account.EventOpen{
		ID: "e62777ac118bc627e8fe27b6c729b7bf886e2590",
	})
	events = append(events, account.EventUpdate{
		NewDetails: account.AccountDetails{
			Email: "user@example.com",
		},
	})
	events = append(events, account.EventUpdate{
		NewDetails: account.AccountDetails{
			Email: "user+test1@example.com",
		},
	})

	data[aggregateID] = events

	ee := newEventsEmitter()

	go func(aggregateID string) {
		h := &account.CloseCmdHandler{
			Now: time.Now,
			EE:  ee,
		}
		handle := makeEventsHandler[account.Account](h)

		for e := range ee.events {
			data := <-store

			events := data[aggregateID]
			data[aggregateID] = append(events, e)

			store <- data
		}

		err := handle(ctx, aggregateID, events)
		if err != nil {
			panic(err)
		}
	}(aggregateID)

	// Output:
	//
}

type eventsHandlerFunc func(ctx context.Context, aggregateID string, events []es.DomainEvent) error

func makeEventsHandler[T any, PT es.Applier[T]](h es.Handler[T]) eventsHandlerFunc {
	return func(ctx context.Context, aggregateID string, events []es.DomainEvent) error {
		var aggregate T
		for _, e := range events {
			err := PT(&aggregate).Apply(e)
			if err != nil {
				return err
			}
		}
		return h.Handle(ctx, aggregateID, aggregate)
	}
}

type eventsEmitter struct {
	events chan es.DomainEvent
}

func newEventsEmitter() *eventsEmitter {
	return &eventsEmitter{
		events: make(chan es.DomainEvent, 1),
	}
}

func (ee *eventsEmitter) Emit(ctx context.Context, events ...es.DomainEvent) error {
	for i := range events {
		select {
		case ee.events <- events[i]:
		case <-ctx.Done():
			return ctx.Err()
		}
	}
	return nil
}
