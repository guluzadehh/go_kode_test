package note

import (
	"math/rand"
	"time"
)

type NoteStorage interface {
	GetAllOfUser(uid int) []*Note
	Create(text string, uid int) (*Note, error)
}

type NoteList []*Note

type MemoryStorage struct {
	memory map[int]NoteList
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{memory: make(map[int]NoteList)}
}

func (s *MemoryStorage) GetAllOfUser(uid int) []*Note {
	notes, ok := s.memory[uid]
	if !ok {
		return make([]*Note, 0)
	}
	return notes
}

func (s *MemoryStorage) Create(text string, uid int) (*Note, error) {
	note := &Note{
		Id:        rand.Intn(10000),
		Text:      text,
		CreatedAt: time.Now().UTC(),
		UserId:    uid,
	}

	notes, ok := s.memory[uid]
	if !ok {
		notes = make([]*Note, 0, 1)
	}

	notes = append(notes, note)
	s.memory[uid] = notes

	return note, nil
}
