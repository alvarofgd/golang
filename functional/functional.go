package functional

type Predicate[T any] func(t T) bool

func Filter[T any](ts []T, f Predicate[T]) []T {
	return folding(ts, []T{}, func(ts []T, t T) []T {
		if f(t) {
			return append(ts, t)
		}

		return ts
	})
}

type Mapper[T any, R any] func(t T) R

func Map[T, R any](ts []T, f Mapper[T, R]) []R {
	return folding(ts, []R{}, func(rs []R, t T) []R {
		return append(rs, f(t))
	})
}

type Accumulator[T, R any] func(r R, t T) R

func Fold[T, R any](ts []T, acc R, f Accumulator[T, R]) R {
	return folding(ts, acc, f)
}

func folding[T any, R any](ts []T, acc R, f Accumulator[T, R]) R {
	res := acc
	for _, v := range ts {
		res = f(res, v)
	}

	return res
}
