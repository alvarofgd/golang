package functional

import (
	"iter"
	"slices"
)

type Predicate[T any] func(t T) bool

func FilterSeq[T any](seq iter.Seq[T], p Predicate[T]) []T {
	return folding(seq, []T{}, func(s []T, t T) []T {
		if p(t) {
			return append(s, t)
		}

		return s
	})
}

func Filter[Slice ~[]T, T any](s Slice, p Predicate[T]) Slice {
	return FilterSeq(slices.Values(s), p)
}

func Find[Slice ~[]T, T any](s Slice, p Predicate[T]) (T, bool) {
	for _, v := range s {
		if p(v) {
			return v, true
		}
	}

	return *new(T), false
}

type Mapper[T any, R any] func(t T) R

func Map[Slice ~[]T, T, R any](s Slice, mapper Mapper[T, R]) []R {
	return folding(slices.Values(s), []R{}, func(rs []R, t T) []R {
		return append(rs, mapper(t))
	})
}

type Accumulator[T, R any] func(r R, t T) R

func Fold[Slice ~[]T, T, R any](s Slice, zero R, acc Accumulator[T, R]) R {
	return folding(slices.Values(s), zero, acc)
}

func folding[T any, R any](seq iter.Seq[T], zero R, acc Accumulator[T, R]) R {
	res := zero
	for v := range seq {
		res = acc(res, v)
	}

	return res
}
