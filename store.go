package main

import (
	"errors"
	"fmt"
	"sync"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type Store struct {
	mu    sync.Mutex
	todos map[string]*Todo
	next  int
}

func NewStore() *Store {
	return &Store{
		todos: make(map[string]*Todo),
		next:  1,
	}
}

func (s *Store) Create(t *Todo) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := s.next
	s.next++
	t.ID = fmt.Sprintf("%d", id)
	s.todos[t.ID] = t
	return t
}

func (s *Store) Get(id string) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.todos[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return t, nil
}

func (s *Store) Update(id string, done bool) (*Todo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	t, ok := s.todos[id]
	if !ok {
		return nil, errors.New("not found")
	}
	t.Done = done
	return t, nil
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.todos[id]; !ok {
		return errors.New("not found")
	}
	delete(s.todos, id)
	return nil
}
