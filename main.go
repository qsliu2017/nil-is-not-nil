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

func main() {
	println("(*A)(nil)")
	f((*A)(nil))
	println("nil")
	f(nil)
}
