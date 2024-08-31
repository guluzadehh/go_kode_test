package auth

import (
	"math/rand"
)

type UserStorage interface {
	GetAll() []*User
	GetByUsername(username string) (*User, bool)
	Create(username, password string) (*User, error)
}

type MemoryStorage struct {
	memory map[string]*User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		memory: make(map[string]*User),
	}
}

func (s *MemoryStorage) GetByUsername(username string) (*User, bool) {
	user, ok := s.memory[username]
	return user, ok
}

func (s *MemoryStorage) Create(username, password string) (*User, error) {
	user := &User{
		Id:       rand.Intn(10000),
		Username: username,
		Password: password,
	}
	s.memory[user.Username] = user
	return user, nil
}

func (s *MemoryStorage) GetAll() []*User {
	users := make([]*User, 0, len(s.memory))
	for _, u := range s.memory {
		users = append(users, u)
	}
	return users
}
