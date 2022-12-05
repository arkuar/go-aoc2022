package utils

import "sync"

type Stack[T any] struct {
	items []T
	mutex sync.Mutex
}

func (stack *Stack[T]) Push(item T) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

func (stack *Stack[T]) PushN(items []T) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, items...)
}

func (stack *Stack[T]) Pop() (T, bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var item T

	if len(stack.items) == 0 {
		return item, false
	}

	item = stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]
	return item, true
}

func (stack *Stack[T]) PopN(n int) ([]T, bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	var items []T
	if len(stack.items) == 0 {
		return items, false
	}
	items = make([]T, n)
	copy(items, stack.items[len(stack.items)-n:])
	stack.items = stack.items[:len(stack.items)-n]
	return items, true
}

func (stack *Stack[T]) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}

func (stack *Stack[T]) Peek() (T, bool) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	var item T
	if len(stack.items) == 0 {
		return item, false
	}

	return stack.items[len(stack.items)-1], true
}
