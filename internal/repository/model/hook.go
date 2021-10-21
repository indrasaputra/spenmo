package model

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"

	"github.com/indrasaputra/spenmo/internal/repository/model/ent"
)

const (
	dbType = "postgres"
)

// HookTracing wraps the operation with tracing.
func HookTracing(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		span, ctx := opentracing.StartSpanFromContext(ctx, fmt.Sprintf("%s-%s", m.Type(), m.Op()))
		defer span.Finish()

		ext.DBType.Set(span, dbType)
		ext.DBStatement.Set(span, m.Op().String())

		return next.Mutate(ctx, m)
	})
}
