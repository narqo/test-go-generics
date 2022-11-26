package account

import (
	"fmt"
	"time"

	"github.com/narqo/test-go-generics/es"
	"github.com/narqo/test-go-generics/es/interest"
)

type AccountState string

const (
	AccountStateOpened = "OPENED"
	AccountStateClosed = "CLOSED"
)

// Account is an aggregate, that holds an account.
type Account struct {
	ID        string
	Email     string
	State     AccountState
	UpdatedAt time.Time
}

func (a *Account) Apply(e es.DomainEvent) error {
	switch e := e.(type) {
	case EventOpen:
		a.ID = e.ID
		a.State = AccountStateOpened

	case EventUpdate:
		if a.State == AccountStateClosed {
			return fmt.Errorf("update closed account: id %v", a.ID)
		}
		if err := a.updateDetails(e.NewDetails); err != nil {
			return fmt.Errorf("update account: %w", err)
		}

	default:
		return fmt.Errorf("unexpected domain event %T", e)
	}
	return nil
}

func (a *Account) updateDetails(b AccountDetails) error {
	if b.Email != "" {
		a.Email = b.Email
	}
	if b.State != "" {
		a.State = b.State
	}
	return nil
}

func (a Account) validateState() error {
	if a.State == AccountStateClosed {
		return fmt.Errorf("account closed: id %v", a.ID)
	}
	return nil
}

func (a *Account) Close() ([]es.DomainEvent, error) {
	if err := a.validateState(); err != nil {
		return nil, err
	}
	e := EventUpdate{
		NewDetails: AccountDetails{
			State: AccountStateClosed,
		},
	}
	return []es.DomainEvent{e}, nil
}

func (a *Account) AttachInterestProfile() ([]es.DomainEvent, error) {
	if err := a.validateState(); err != nil {
		return nil, err
	}
	e := interest.EventAttachProfile{
		AccID: a.ID,
		Rate:  0,
	}
	return []es.DomainEvent{e}, nil
}
