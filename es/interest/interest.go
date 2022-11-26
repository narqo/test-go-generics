package interest

import (
	"fmt"
	"time"

	"github.com/narqo/test-go-generics/es"
)

// Profile is an aggregate, that holds an interest profile.
type Profile struct {
	AccID     string
	Rate      float32
	UpdatedAt time.Time
}

func (p *Profile) Apply(e es.DomainEvent) error {
	switch e := e.(type) {
	case EventAttachProfile:
		p.AccID = e.AccID
		p.Rate = e.Rate

	default:
		return fmt.Errorf("unexpected domain event %T", e)
	}
	return nil
}
