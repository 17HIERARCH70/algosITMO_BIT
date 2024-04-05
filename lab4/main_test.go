package main

import (
	"testing"
)

// Функция для создания очереди с n элементами, где n - степень двойки.
func createQueueWithNElements(n int) *CircularQueue {
	q := &CircularQueue{}
	for i := 0; i < n; i++ {
		q.Enqueue(i)
	}
	return q
}

// TestEnqueueDequeue проверяет корректность добавления и удаления элементов из очереди.
func TestEnqueueDequeue(t *testing.T) {
	q := createQueueWithNElements(8) // Создаем очередь с 8 элементами

	// Проверяем размер очереди после добавления элементов
	if q.size != 8 {
		t.Errorf("Expected queue size of 8, got %d", q.size)
	}

	// Удаляем элементы и проверяем их значения
	for i := 0; i < 8; i++ {
		val, err := q.Dequeue()
		if err != nil {
			t.Fatalf("Dequeue error: %v", err)
		}
		if val != i {
			t.Errorf("Expected value %d, got %d", i, val)
		}
	}

	// Проверяем, что очередь пуста после удаления всех элементов
	if q.size != 0 {
		t.Errorf("Expected empty queue, got size %d", q.size)
	}
}

// TestBetonicSort проверяет корректность битонной сортировки.
func TestBetonicSort(t *testing.T) {
	q := createQueueWithNElements(8) // Создаем очередь с 8 элементами
	q.BetonicSort()                  // Сортировка

	// Проверяем, что элементы отсортированы
	prevVal, _ := q.Dequeue()
	for q.size > 0 {
		val, _ := q.Dequeue()
		if prevVal > val {
			t.Errorf("Queue is not sorted: %d > %d", prevVal, val)
		}
		prevVal = val
	}
}

// BenchmarkBetonicSort бенчмарк для битонной сортировки.
func BenchmarkBetonicSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := createQueueWithNElements(8) // Создаем очередь с 8 элементами
		q.BetonicSort()                  // Сортировка
	}
}
