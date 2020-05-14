package contextawarereader

import (
	"context"
	"io"
)

type CancellableReader struct {
	delegate io.Reader
	ctx      context.Context
}

func (r *CancellableReader) Read(p []byte) (n int, err error) {
	if err := r.ctx.Err(); err != nil {
		return 0, err
	}
	return r.delegate.Read(p)
}

func NewCancellableReader(ctx context.Context, r io.Reader) io.Reader {
	return &CancellableReader{delegate: r, ctx: ctx}
}
