package test

// Something is a test struct.
// second line of doc comment.
type Something struct {
	// Foo is something foo
	Foo string
	// Bar is not a bar
	// but bar
	Bar string `json:"barbar"`
	Baz bool
	/*
		qux is qux
	*/
	Qux int
}

func DoSomething() {
	print(1)
}

type Nested struct {
	Foo struct {
		Bar string
	}
}

type Parent struct {
	Child Child
}

type Child struct {
	Value int64
}
