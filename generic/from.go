package generic

type Iterator[T any] func() (item T, ok bool)

type Query[T any] struct {
	Iterate func() Iterator[T]
}
type QueryTo[T, R any] struct {
	Query[T]
}

func From[T any](source []T) Query[T] {
	len := len(source)
	return Query[T]{
		Iterate: func() Iterator[T] {
			i := 0
			return func() (item T, ok bool) {
				ok = i < len
				if ok {
					item = source[i]
					i++
				}
				return
			}
		},
	}
}

func (q Query[T]) Where(predicate func(T) bool) Query[T] {
	return Query[T]{
		Iterate: func() Iterator[T] {
			next := q.Iterate()
			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}
				return
			}
		},
	}
}

func (q Query[T]) ToSlice() (r []T) {
	next := q.Iterate()
	for item, ok := next(); ok; item, ok = next() {
		r = append(r, item)
	}
	return
}
