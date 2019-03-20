package test

import "context"

type Bar struct{}

type Anything interface {
	Foo(hoge *string, fuga bool) error
	Bar(ctx context.Context, baz, qux string) (*Bar, error)
}
