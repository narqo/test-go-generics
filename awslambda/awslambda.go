package awslambda

import (
	"context"
)

type Handler[R any] interface {
	Handle(ctx context.Context, record R) error
}

type Event[R any] struct {
	Records []R `json:"Records"`
}

type lambdaHandler[R any] func(ctx context.Context, event Event[R]) error

func createLambdaHandler[R any](h Handler[R]) lambdaHandler[R] {
	return func(ctx context.Context, event Event[R]) error {
		for _, rec := range event.Records {
			err := h.Handle(ctx, rec)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

// FIXME: the code below doesn't compile
/*
type BatchEventRecord interface {
	SQSMessageRecord | StreamEventRecord
}

type SQSMessageRecord struct {
	MessageId string `json:"messageId"`
}

type StreamEventRecord struct {
	EventID string `json:"eventId"`
}

type BatchEventResponse struct {
	BatchItemFailures []BatchItemFailure `json:"batchItemFailures"`
}

type BatchItemFailure struct {
	ItemIdentifier string `json:"itemIdentifier"`
}

type lambdaBatchHandler[R BatchEventRecord] func(ctx context.Context, event Event[R]) (BatchEventResponse, error)

func createLambdaBatchHandler[R BatchEventRecord](h Handler[R]) lambdaBatchHandler[R] {
	return func(ctx context.Context, event Event[R]) (BatchEventResponse, error) {
		failedItemIDs := make(chan string, 1)

		var wg sync.WaitGroup
		wg.Add(len(event.Records))

		for _, rec := range event.Records {
			go func(rec R) {
				defer wg.Done()

				err := h.Handle(ctx, rec)
				if err != nil {
					switch rec := any(rec).(type) {
					case SQSMessageRecord:
						failedItemIDs <- rec.MessageId
					case StreamEventRecord:
						failedItemIDs <- rec.EventID
					}
				}
			}(rec)
		}

		go func() {
			wg.Wait()
			close(failedItemIDs)
		}()

		var itemFailures []BatchItemFailure
		for id := range failedItemIDs {
			item := BatchItemFailure{
				ItemIdentifier: id,
			}
			itemFailures = append(itemFailures, item)
		}

		resp := BatchEventResponse{
			BatchItemFailures: itemFailures,
		}
		return resp, nil
	}
}
*/
