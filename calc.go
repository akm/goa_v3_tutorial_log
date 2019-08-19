package calcapi

import (
	calc "calcsvc/gen/calc"
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

type Calc struct {
	A int
	B int
	R int
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	res = p.A + p.B

	dsClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		s.logger.Printf("Failed to get datastore client because of [%T] %+v", err, err)
		return res, err
	}

	k := datastore.IncompleteKey("calcsvc", nil)
	calc := &Calc{
		A: p.A,
		B: p.B,
		R: res,
	}
	if key, err := dsClient.Put(ctx, k, calc); err != nil {
		s.logger.Printf("Failed to put calculation %v into datastore client because of [%T] %+v", *calc, err, err)
		return res, err
	} else {
		s.logger.Printf("Calculation %v was saved successfully with key: %v", *calc, *key)
	}

	return
}
