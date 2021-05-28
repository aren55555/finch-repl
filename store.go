package main

import (
	"errors"
)

type Store struct {
	State   map[string]int
	Changes []map[string]int
}

func NewStore() *Store {
	return &Store{
		State:   map[string]int{},
		Changes: []map[string]int{},
	}
}

func (s *Store) currentChange() map[string]int {
	return s.Changes[len(s.Changes)-1]
}

func (s *Store) copyState() map[string]int {
	d := map[string]int{}
	for k, v := range s.State {
		d[k] = v
	}
	return d
}

func (s *Store) copyChanges() []map[string]int {
	d := []map[string]int{}
	for _, elem := range s.Changes {
		d = append(d, elem)
	}
	return d
}

func (s *Store) Begin() {
	if len(s.Changes) == 0 {
		s.Changes = append(s.copyChanges(), s.copyState())
		return
	}

	// append last
	s.Changes = append(s.copyChanges(), s.Changes[len(s.Changes)-1])
}

func (s *Store) Commit() {
	// apply all change, mutate the state
	s.State = s.Changes[len(s.Changes)-1]
	s.Changes = []map[string]int{}
}

func (s *Store) Rollback() {
	s.Changes = s.Changes[:len(s.Changes)-1]
}

func (s *Store) Get(key string) int {
	if len(s.Changes) == 0 {
		// operation on the state map directly
		return s.State[key]
	}

	change := s.currentChange()
	return change[key]
}

func (s *Store) Set(key string, value int) error {
	if len(s.Changes) == 0 {
		return errors.New("NO ACTIVE TRANSACTIONS")
	}

	// apply the change to the most recent change
	change := s.currentChange()
	change[key] = value
	return nil
}

func (s *Store) Del(key string) error {
	if len(s.Changes) == 0 {
		return errors.New("NO ACTIVE TRANSACTIONS")
	}

	// apply the change to the most recent change
	change := s.currentChange()
	delete(change, key)
	return nil
}

// func main() {
// 	s := NewStore()

// 	s.Begin()
// 	s.Set("x", 5)
// 	s.Set("y", 23)
// 	fmt.Println(s.Get("x"))

// 	s.Begin()
// 	fmt.Println(s.Get("y"))
// 	s.Set("y", 42)

// 	s.Commit()

// 	fmt.Println(s.Get("y"))
// 	fmt.Println(s.Set("z", 1)) // should error

// 	s.Begin()
// 	s.Set("x", 99)
// 	s.Rollback()
// 	fmt.Println(s.Get("x"))

// 	fmt.Println(s.Get("y"))

// 	s.Begin()
// 	fmt.Println(s.Del("y"))
// 	s.Rollback()

// 	fmt.Println(s.Get("y"))

// 	s.Begin()
// 	fmt.Println(s.Del("y"))
// 	s.Commit()

// 	fmt.Println(s.Get("y"))

// 	s.Begin()
// 	fmt.Println(s.Get("x"))
// 	s.Rollback()
// }
