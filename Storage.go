package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	name   string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee whith such id doesn't exist")
	}
	return e, nil
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e

	return nil

}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func main() {

	var s storage

	fmt.Println("s ==nil", s == nil)
	fmt.Printf("type of s: %T\n:", s)

	s = newMemoryStorage()

	fmt.Println("s ==nil", s == nil)
	fmt.Printf("type of s: %T\n:", s)

}
