package v

type Queue[T any] struct {
	v []T
}

func (q *Queue[T]) Offer(e T) {
	q.v = append(q.v, e)
}

func (q *Queue[T]) Peek() T {
	return q.v[0]
}

func (q *Queue[T]) Poll() T {
	r := q.v[0]
	q.v = q.v[1:]
	return r
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.v) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.v)
}
