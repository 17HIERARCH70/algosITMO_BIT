package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color" // Импорт пакета для работы с цветом вывода
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

// Node представляет собой узел в связном списке.
type Node struct {
	value int   // Значение узла
	next  *Node // Следующий узел в списке
}

// CircularQueue представляет собой структуру кольцевой очереди.
type CircularQueue struct {
	head *Node // Начало очереди
	tail *Node // Конец очереди
	size int   // Размер очереди
}

// Enqueue добавляет новый элемент в конец очереди.
func (q *CircularQueue) Enqueue(value int) {
	newNode := &Node{value: value} // Создание нового узла
	if q.head == nil {             // Если очередь пуста
		q.head = newNode
	} else { // Если в очереди уже есть элементы
		q.tail.next = newNode
	}
	q.tail = newNode // Обновление указателя на конец очереди
	// Связываем последний и первый элементы для обеспечения кольцевой структуры
	q.tail.next = q.head
	q.size++ // Увеличиваем размер очереди
}

// Dequeue удаляет элемент из начала очереди и возвращает его значение.
func (q *CircularQueue) Dequeue() (int, error) {
	if q.head == nil { // Если очередь пуста
		return 0, errors.New("очередь пуста")
	}
	value := q.head.value // Значение удаляемого элемента
	if q.head == q.tail { // Если в очереди остался один элемент
		q.head = nil
		q.tail = nil
	} else {
		// Обновление указателя на начало очереди
		q.head = q.head.next
		// Обновление указателя последнего элемента на новый первый элемент
		q.tail.next = q.head
	}
	q.size-- // Уменьшаем размер очереди
	return value, nil
}

// Display выводит элементы очереди.
func (q *CircularQueue) Display() {
	if q.head == nil { // Если очередь пуста
		fmt.Println("Очередь пуста")
		return
	}
	current := q.head
	for {
		fmt.Printf("%d ", current.value)
		current = current.next // Переход к следующему узлу
		if current == q.head { // Проверка на завершение обхода очереди
			break
		}
	}
	fmt.Println()
}

// ClearConsole очищает консоль в зависимости от операционной системы.
func ClearConsole() {
	switch runtime.GOOS {
	case "linux", "darwin": // Для Linux и macOS
		cmd := exec.Command("clear") // Команда для очистки консоли
		cmd.Stdout = os.Stdout
		err := cmd.Run() // Выполнение команды
		if err != nil {
			// Обработка ошибки очистки консоли
			_ = errors.New("Не удалось очистить консоль: %s\n")
			return
		}
	case "windows": // Для Windows
		// Команда для очистки консоли
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		// Выполнение команды
		err := cmd.Run()
		if err != nil {
			// Обработка ошибки очистки консоли
			_ = errors.New("Не удалось очистить консоль: %s\n")
			return
		}
	}
}

// BetonicSort начинает процесс сортировки элементов в очереди.
func (q *CircularQueue) BetonicSort() {
	var size int // Переменная для хранения размера очереди

	// Определение максимального количества элементов, которое является степенью двойки
	if q.size == 0 || (q.size&(q.size-1)) != 0 {
		fmt.Println("Количество элементов в очереди не является степенью двойки")
		fmt.Println("Будет взято максимальное количество элементов, которое является степенью двойки")
		size = q.size - 1
	} else {
		size = q.size
	}

	// Извлечение элементов из очереди в массив
	elements := make([]int, size)

	for i := 0; i < size; i++ {
		element, _ := q.Dequeue()
		elements[i] = element
	}

	// Битонная сортировка
	for k := 2; k <= size; k = k * 2 {
		for j := k / 2; j > 0; j = j / 2 {
			for i := 0; i < size; i++ {
				l := i ^ j
				if l > i {
					if ((i&k) == 0 && elements[i] > elements[l]) || ((i&k) != 0 && elements[i] < elements[l]) {
						// Обмен элементами
						elements[i], elements[l] = elements[l], elements[i]
					}
				}
			}
		}
	}
	// Возвращение отсортированных элементов в очередь
	for _, elem := range elements {
		q.Enqueue(elem)
	}
}

// main является точкой входа программы.
func main() {
	reader := bufio.NewReader(os.Stdin)
	queue := &CircularQueue{}

	// Функция для вывода сообщений об ошибке красным цветом
	errMsg := color.New(color.FgHiRed).PrintfFunc()
	// Функция для вывода успешных сообщений зеленым цветом
	success := color.New(color.FgHiGreen).PrintfFunc()

	for {
		ClearConsole() // Очистить консоль
		fmt.Printf("\nВведите команду. Доступные команды:\n1. enqueue [число]\n2. dequeue\n3. sort\n4. display\n5. exit\n")
		cmdString, _ := reader.ReadString('\n')
		cmdString = strings.TrimSpace(cmdString)
		cmdParts := strings.Split(cmdString, " ")

		switch cmdParts[0] {
		case "1":
			if len(cmdParts) != 2 {
				// Вывод сообщения об ошибке
				errMsg("Необходимо указать число для добавления в очередь\n")
				continue
			}
			value, err := strconv.Atoi(cmdParts[1])
			if err != nil {
				// Вывод сообщения об ошибке
				errMsg("Введите корректное число\n")
				continue
			}
			queue.Enqueue(value) // Добавление элемента в очередь
			// Вывод успешного сообщения
			success("Элемент добавлен в очередь\n")
		case "2":
			value, err := queue.Dequeue()
			if err != nil {
				// Вывод сообщения об ошибке
				errMsg("%s\n", err)
			} else {
				// Вывод успешного сообщения
				success("Из очереди удален элемент: %d\n", value)
			}
		case "3":
			queue.BetonicSort() // Битонная сортировка очереди
			// Вывод успешного сообщения
			success("Очередь отсортирована\n")
		case "4":
			queue.Display() // Отображение элементов очереди
		case "5":
			fmt.Print("Программа завершена\n")
			return
		default:
			// Вывод сообщения об ошибке
			errMsg("Неизвестная команда\n")
		}
		fmt.Printf("\nНажмите Enter для продолжения...")
		_, err := reader.ReadString('\n')
		if err != nil {
			// Вывод сообщения об ошибке
			errMsg("Не удалось прочитать строку: %s\n", err)
			return
		}
	}
}
