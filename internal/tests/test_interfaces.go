package test

import "context"

type Anything interface {
	Foo(hoge *string, fuga bool) error
	Bar(ctx context.Context, baz, qux string) (int, error)
}
