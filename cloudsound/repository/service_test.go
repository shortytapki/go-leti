package repository_test

import (
	_ "fmt"
	"golangCourse/cloudsound/entities"
	"golangCourse/cloudsound/repository"
	"reflect"
	"testing"
)
func TestSongCrud(t *testing.T) {
	// Тест создания записи в БД
	db, err := repository.New("../..",)
	if err != nil {
		t.Fatal(err)
	}
	initialRecords, err := db.GetSongs()
	if err != nil {
		t.Fatal(err)
	}
	err = db.AddSong(entities.Song{Name: "Kissing The Shadows", Duration: 311, GenreId: 4, ArtistId: 5})
	if err != nil {
		t.Fatal(err)
	}
	newRecords, err := db.GetSongs()
	if err != nil {
		t.Fatal(err)
	}
	if len(newRecords) != len(initialRecords) + 1 {
		t.Error("Wrong amount of records after addition.")
	}

	// Тест получения песни по id
	mockSong := entities.Song{Id: 4, Name: "Are You Dead Yet?", Duration: 237, GenreId: 4, ArtistId: 5}
	song, err := db.GetSongById(4)
	if err != nil {
		t.Fatal(err)
	}
	songReflect := reflect.ValueOf(&song).Elem()
	songType:= songReflect.Type()
	mockReflect := reflect.ValueOf(&mockSong).Elem()
	mockType:= mockReflect.Type()
	for i := 0; i < songType.NumField(); i++ {
		field := songType.Field(i)
		mockField := mockType.Field(i)
		rv := reflect.ValueOf(&song)
		mockRv := reflect.ValueOf(&mockSong)
		songValue := reflect.Indirect(rv).FieldByName(field.Name)
		mockValue := reflect.Indirect(mockRv).FieldByName(mockField.Name)
		if !(songValue.Equal(mockValue))  {
			t.Fatal("Wrong song struct values.")
		}
	}

	// Тест обновления песни
	newSong := entities.Song{Id: 1, Name: "No More Years", Duration: 159}
	db.UpdateSong(&newSong)
	updatedSong, err := db.GetSongById(newSong.Id)
	if err != nil {
		t.Fatal(err.Error())
	}
	newSongReflect := reflect.ValueOf(&newSong).Elem()
	newSongType := newSongReflect.Type()
	updatedSongReflect := reflect.ValueOf(&updatedSong).Elem()
	updatedSongType:= updatedSongReflect.Type()
	for i := 0; i < songType.NumField(); i++ {
		field := newSongType.Field(i)
		mockField := updatedSongType.Field(i)
		rv := reflect.ValueOf(&song)
		mockRv := reflect.ValueOf(&mockSong)
		songValue := reflect.Indirect(rv).FieldByName(field.Name)
		mockValue := reflect.Indirect(mockRv).FieldByName(mockField.Name)
		if !(songValue.Equal(mockValue))  {
			t.Fatal("Wrong song struct values.")
		}
	}

	// Тест удаления песни из базы
	db.DeleteSong(25)
	currentRecords, err := db.GetSongs()
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(currentRecords) != len(initialRecords) {
		t.Fatal("Failed deleting a record.")
	}
}