package generic

type Struct struct {
	Foo string
	Baz int
}

var Object = Struct{
	Foo: "bar",
	Baz: 1,
}

var RawObject = []byte(`{"Foo":"bar","Baz":1}`)
