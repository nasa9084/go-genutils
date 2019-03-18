package test

type Anything interface {
	Foo(hoge string, fuga bool) error
	Bar(baz, qux string) (int, error)
}
