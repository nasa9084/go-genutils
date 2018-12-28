package test

// Something is a test struct.
// second line of doc comment.
type Something struct {
	Foo string
	Bar string `json:"barbar"`
	Baz bool
}

func DoSomething() {
	print(1)
}
