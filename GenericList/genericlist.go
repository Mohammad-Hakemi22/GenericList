package genericlist

import (
	"errors"
)

type GenericList[T comparable] struct {
	data []T
}

func New[T comparable]() *GenericList[T] {
	return &GenericList[T]{
		data: []T{},
	}
}

func (lst *GenericList[T]) Insert(value T) {
	lst.data = append(lst.data, value)
}

func (lst *GenericList[T]) GetByIndex(idx int) (T, error) {
	for i := 0; i < len(lst.data); i++ {
		if i == idx {
			return lst.data[idx], nil
		}
	}
	return *new(T), errors.New("index does not exists")
}
