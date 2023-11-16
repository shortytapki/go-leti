package servicedb

// Описание сущности рифмы
type Rhyme struct {
	Text string
}

// Описание сущности песни
type Song struct {
	Title  string
	Artist string
}

// Описание структуры базы данных
type DB struct {
	Rhymes []Rhyme
	Songs  []Song
}

// Создаёт и возвращает указатель
// на структуру базы данных
func New(rhymes []Rhyme, songs []Song) *DB {
	return &DB{Rhymes: rhymes, Songs: songs}
}
