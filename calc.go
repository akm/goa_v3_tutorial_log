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
	M string
	A int
	B int
	R int
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	s.logger.Print("calc.add")
	return s.saveCalc(ctx, &Calc{"add", p.A, p.B, p.A + p.B})
}

// Multiply implements multiply.
func (s *calcsrvc) Multiply(ctx context.Context, p *calc.MultiplyPayload) (res int, err error) {
	s.logger.Print("calc.multiply")
	return
}

// Devide implements devide.
func (s *calcsrvc) Devide(ctx context.Context, p *calc.DevidePayload) (res int, err error) {
	s.logger.Print("calc.devide")
	return
}

func (s *calcsrvc) saveCalc(ctx context.Context, c *Calc) (int, error) {
	dsClient, err := datastore.NewClient(ctx, "")
	if err != nil {
		s.logger.Printf("Failed to get datastore client because of [%T] %+v", err, err)
		return 0, err
	}

	k := datastore.IncompleteKey("calcsvc", nil)
	if key, err := dsClient.Put(ctx, k, c); err != nil {
		s.logger.Printf("Failed to put calculation %v into datastore client because of [%T] %+v", *c, err, err)
		return 0, err
	} else {
		s.logger.Printf("Calculation %v was saved successfully with key: %v", *c, *key)
		return c.R, nil
	}
}
