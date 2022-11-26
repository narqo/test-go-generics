package account

import (
	"context"
	"fmt"
	"time"

	"github.com/narqo/test-go-generics/es"
)

type CloseCmdHandler struct {
	Now func() time.Time
	EE  es.Emitter
}

func (h *CloseCmdHandler) Handle(ctx context.Context, aggregateID string, acc Account) error {
	fmt.Printf("> CMD close: aggregate_id %v, account %+v\n", aggregateID, acc)

	events, err := acc.Close()
	if err != nil {
		return fmt.Errorf("close account %s: %w", acc.ID, err)
	}

	return h.EE.Emit(ctx, events...)
}

type AttachInterestProfileCmdHandler struct {
	EE es.Emitter
}

func (h *AttachInterestProfileCmdHandler) Handle(ctx context.Context, aggregateID string, acc Account) error {
	fmt.Printf("> CMD attach interest profile: aggregate_id %v, account %+v\n", aggregateID, acc)

	// TODO: cmd's payload: pass rate, etc
	events, err := acc.AttachInterestProfile()
	if err != nil {
		return fmt.Errorf("attach interest profile to account %s: %w", acc.ID, err)
	}

	return h.EE.Emit(ctx, events...)
}
