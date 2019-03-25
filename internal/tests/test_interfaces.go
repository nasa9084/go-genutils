package test

import "context"

type Foo struct{}

type Bar struct{}

type Anything interface {
	Foo(hoge *string, fuga bool) ([]*Foo, error)
	Bar(ctx context.Context, baz, qux string) (*Bar, error)
}
