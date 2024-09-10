package note

import (
	"math/rand"
	"sync"
	"time"
)

type NoteStorage interface {
	GetAllOfUser(uid int) []*Note
	Create(text string, uid int) (*Note, error)
}

type NoteList []*Note

type MemoryStorage struct {
	memory map[int]NoteList
	mu     sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{memory: make(map[int]NoteList)}
}

func (s *MemoryStorage) GetAllOfUser(uid int) []*Note {
	s.mu.RLock()
	defer s.mu.RUnlock()

	notes, ok := s.memory[uid]
	if !ok {
		return make([]*Note, 0)
	}
	return notes
}

func (s *MemoryStorage) Create(text string, uid int) (*Note, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

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
