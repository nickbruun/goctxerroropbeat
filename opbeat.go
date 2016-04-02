package ctxerroropbeat

import (
	"github.com/roncohen/opbeat-go"
	"golang.org/x/net/context"

	"github.com/nickbruun/goctxerror"
	"github.com/nickbruun/goctxreq"
)

// Opbeat error handler.
func opbeatErrorHandler(ctx context.Context, err error, msg string) {
	opts := &opbeat.Options{}

	if req, ok := ctxreq.FromContext(ctx); ok {
		opbeat.CaptureErrorWithRequest(err, req, opts)
	} else {
		opbeat.CaptureError(err, opts)
	}
}

// New context with Opbeat error handling assigned.
func NewContext(ctx context.Context) context.Context {
	return ctxerror.NewContext(ctx, opbeatErrorHandler)
}
