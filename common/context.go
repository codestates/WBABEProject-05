package common

import (
	"context"
	"time"
)

const TotalRequestTimeOut = 2_500 * time.Millisecond
const ControllerContextTimeOut = 0_500 * time.Millisecond
const ServiceContextTimeOut = 0_500 * time.Millisecond
const ModelContextTimeOut = 2_000 * time.Millisecond
const DatabaseClientTimeOut = 2 * time.Millisecond

func NewContext(t time.Duration) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), t)
	return ctx, cancel
}
