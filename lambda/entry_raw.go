package lambda

import (
	"context"
	"io"
)

func StartRawHandler(f func(ctx context.Context, payload []byte) (io.Reader, error), onSigterm ...func()) {
	start(&handlerOptions{
		baseContext:      context.Background(),
		handlerFunc:      f,
		enableSIGTERM:    true,
		sigtermCallbacks: onSigterm,
	})
}
