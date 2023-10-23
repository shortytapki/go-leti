package linkedList

import (
	"errors"
	"fmt"
	"strconv"
)

// Описание ноды листа
type Node struct {
	value  int
	next *Node
}

// Описание листа
type LinkedList struct {
	head *Node
}

// Форматированный вывод
func (l *LinkedList) PrintList() {
	res := ""
	pointer := l.head

	for pointer.next != nil {
		res += strconv.Itoa(pointer.value) + " -> "
		pointer = pointer.next
	}
	fmt.Println(res + strconv.Itoa(pointer.value))
}

// Создаёт новый связанный список
// с количеством нод q
func New(q int) *LinkedList {
	// Создаём экземпляр списка
	// и первый элемент:
	head := &Node{}
	list := LinkedList{head: head}

	// Запомним первый
	// элемент:
	pointer := head

	// Создаём новые ноды в
	// количестве q с учётом первого:
	for nodeCreations := 1; nodeCreations < q; nodeCreations++ {
		newNode := &Node{}
		pointer.next = newNode
		pointer = newNode
	}

	return &list
}

// Добавляет ноду с переданным значением
// в конец списка
func (l *LinkedList) Add(value int) {
	
	// Создаём новую ноду c
	// терминирующим указателем
	newNode := Node{value: value}
	pointer := l.head

	if (l.Size() == 1) {
		pointer.next = &newNode
		return
	}

	// Ищем предпоследний элемент
	for pointer.next != nil {
		pointer = pointer.next
	}
	// Добавляем ноду
	pointer.next = &newNode
}

// Удаляет последний элемент списка
func (l *LinkedList) Pop() {
	// Выход из функции
	// если остался один элемент
	if (l.Size() == 1) {
		return
	}

	pointer := l.head
	// Ищем предпоследний элемент
	for pointer.next.next != nil {
		pointer = pointer.next
	}

	// Делаем предпоследний элемент
	// терминирующим
	pointer.next = nil
}

// Возвращает размер списка
func (l *LinkedList) Size() int {
	// Начинаем с головы
	counter := 1
	pointer := l.head

	for pointer.next != nil {
		counter++
		pointer = pointer.next
	}

	return counter
}

// Возвращает значение на позиции pos
// если 0 <= pos < размер списка
// иначе вернёт ошибку
func (l *LinkedList) At(pos int) (int, error) {
	// Проверка на дурика
	if pos < 0 || pos > l.Size() {
		return 0, errors.New("position is out of range")
	}

	pointer := l.head
	currentNodeIndex := 0

	// Ищем нужный элемент
	for currentNodeIndex < pos {
		pointer = pointer.next
		currentNodeIndex++
	}

	return pointer.value, nil
}

// Удаляет элемент на позиции pos
// (если 0 <= pos < размер списка)
// иначе вернёт ошибку
func (l *LinkedList) DeleteFrom(pos int) error {
	// Проверка на дурика
	if pos < 0 || pos > l.Size() {
		return errors.New("position is out of range")
	}
	// Проверяем удаление начальной
	// ноды
	pointer := l.head
	if pos == 0 {
		l.head = pointer.next
	} else {
		currentNodeIndex := 0
		isLastNode := pos == l.Size()

		// Идём до элемента перед
		// нужным
		for currentNodeIndex != pos-1 {
			// Проверяем удаление последней ноды
			if isLastNode && currentNodeIndex == pos-2 {
				pointer.next = nil
				return nil
			}
			pointer = pointer.next
			currentNodeIndex++
		}
		// Указываем на элемент после
		// нужного
		pointer.next = pointer.next.next
	}
	return nil
}

// Устанавливает переданное значение
// (если 0 <= pos < размер списка)
// на переданной позиции.
// Иначе вёрнёт ошибку.
func (l *LinkedList) UpdateAt(pos, value int) error {
	// Проверка на дурика
	if pos < 0 || pos > l.Size() {
		return errors.New("position is out of range")
	}
	// Проверяем обновление начальной
	// ноды
	pointer := l.head
	if pos == 0 {
		l.head.value = value
	} else {
		currentNodeIndex := 0
		// Идём до элемента
		for currentNodeIndex != pos {
			pointer = pointer.next
			currentNodeIndex++
		}
		pointer.value = value
	}
	return nil
}

func NewFromSlice(s []int) *LinkedList {
	//Сразу создаём список
	size := len(s)
	list := New(size)
	pointer := list.head

	// Проходимся с добавлением значений
	for idx := 0; idx < len(s); idx++ {
		pointer.value = s[idx]
		pointer = pointer.next
	}
	return list
}
