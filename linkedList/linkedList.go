package linkedList

import (
	"errors"
	"fmt"
	"strconv"
)

// Описание ноды листа
type Node struct {
	Val  int
	Next *Node
}

// Описание листа
type LinkedList struct {
	Head *Node
}

// Форматированный вывод
func (l *LinkedList) PrintList() {
	res := ""
	pointer := l.Head

	for pointer.Next != nil {
		res += strconv.Itoa(pointer.Val) + " -> "
		pointer = pointer.Next
	}
	fmt.Println(res + strconv.Itoa(pointer.Val) + " -> nil")
}

// Создаёт новый связанный список
// с количеством нод q
func New(q int) *LinkedList {
	// Создаём экземпляр списка
	// и первый элемент:
	head := &Node{}
	list := LinkedList{Head: head}

	// Запомним первый
	// элемент:
	pointer := head

	// Создаём новые ноды в
	// количестве q с учётом первого:
	for nodeCreations := 1; nodeCreations < q; nodeCreations++ {
		newNode := &Node{}
		pointer.Next = newNode
		pointer = newNode
	}

	return &list
}

// Добавляет ноду с переданным значением
// в конец списка
func (l *LinkedList) Add(val int) {
	
	// Создаём новую ноду c
	// терминирующим указателем
	newNode := Node{Val: val}
	pointer := l.Head

	if (l.Size() == 1) {
		pointer.Next = &newNode
		fmt.Println(pointer.Next)
		return
	}

	// Ищем предпоследний элемент
	for pointer.Next != nil {
		pointer = pointer.Next
	}
	// Добавляем ноду
	pointer.Next = &newNode
}

// Удаляет последний элемент списка
func (l *LinkedList) Pop() {
	// Выход из функции
	// если остался один элемент
	if (l.Size() == 1) {
		return
	}

	pointer := l.Head
	// Ищем предпоследний элемент
	for pointer.Next.Next != nil {
		pointer = pointer.Next
	}

	// Делаем предпоследний элемент
	// терминирующим
	pointer.Next = nil
}

// Возвращает размер списка
func (l *LinkedList) Size() int {
	// Начинаем с головы
	counter := 1
	pointer := l.Head

	for pointer.Next != nil {
		counter++
		pointer = pointer.Next
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

	pointer := l.Head
	currentNodeIndex := 0

	// Ищем нужный элемент
	for currentNodeIndex < pos {
		pointer = pointer.Next
		currentNodeIndex++
	}

	return pointer.Val, nil
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
	pointer := l.Head
	if pos == 0 {
		l.Head = pointer.Next
	} else {
		currentNodeIndex := 0
		isLastNode := pos == l.Size()

		// Идём до элемента перед
		// нужным
		for currentNodeIndex != pos-1 {
			// Проверяем удаление последней ноды
			if isLastNode && currentNodeIndex == pos-2 {
				pointer.Next = nil
				return nil
			}
			pointer = pointer.Next
			currentNodeIndex++
		}
		// Указываем на элемент после
		// нужного
		pointer.Next = pointer.Next.Next
	}
	return nil
}

// Устанавливает переданное значение
// (если 0 <= pos < размер списка)
// на переданной позиции.
// Иначе вёрнёт ошибку.
func (l *LinkedList) UpdateAt(pos, val int) error {
	// Проверка на дурика
	if pos < 0 || pos > l.Size() {
		return errors.New("position is out of range")
	}
	// Проверяем обновление начальной
	// ноды
	pointer := l.Head
	if pos == 0 {
		l.Head.Val = val
	} else {
		currentNodeIndex := 0
		// Идём до элемента
		for currentNodeIndex != pos {
			pointer = pointer.Next
			currentNodeIndex++
		}
		pointer.Val = val
	}
	return nil
}

func NewFromSlice(s []int) *LinkedList {
	//Сразу создаём список
	size := len(s)
	list := New(size)
	pointer := list.Head

	// Проходимся с добавлением значений
	for idx := 0; idx < len(s); idx++ {
		pointer.Val = s[idx]
		pointer = pointer.Next
	}
	return list
}
