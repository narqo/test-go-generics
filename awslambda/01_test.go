package awslambda

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type dummyHandler struct{}

func (h *dummyHandler) Handle(context.Context, events.SQSMessage) error

func ExampleLambdaHandler() {
	h := new(dummyHandler)

	handler := createLambdaHandler[events.SQSMessage](h)
	_ = handler

	// Output:
	//
}

/*
func ExampleLambdaBatchHandler() {
	h := new(dummyHandler)

	handler := createLambdaBatchHandler[events.SQSMessage](h)
	_ = handler

	// Output:
	//
}
*/
