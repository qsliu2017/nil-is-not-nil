package main

type A struct{}

func f(v any) {
	switch v {
	case nil:
		println("	is nil")
	default:
		println("	is not nil")
	}
	a, ok := v.(*A)
	switch {
	case !ok:
		println("	is not *A")
	case a == nil:
		println("	is (*A)(nil)")
	default:
		println("	is *A, but not nil")
	}
}

type I interface {
	i()
}

type T struct{}

func (t *T) i() {}

func g(v any) {
	switch v {
	case nil:
		println("	is nil")
	default:
		println("	is not nil")
	}
	i, ok := v.(I)
	switch {
	case !ok:
		println("	is not I")
	case i == nil:
		println("	is (I)(nil)")
	default:
		println("	is I, but not nil")
	}
	t, ok := v.(*T)
	switch {
	case !ok:
		println("	is not *T")
	case t == nil:
		println("	is (*T)(nil)")
	default:
		println("	is *T, but not nil")
	}
}

func main() {
	println("(*A)(nil)")
	f((*A)(nil))
	println("nil")
	f(nil)

	println("(*T)(nil)")
	g((*T)(nil))
	println("(I)(nil)")
	g((I)(nil))
	println("nil")
	g(nil)
}
