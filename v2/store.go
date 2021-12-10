package main

import "io/fs"

func NewStore() *FileStore {
	p := Person{
		Name:   "Andy",
		Age:    20,
		Weight: 145,
	}
	fs.ReadFile(fs.FS, "data.json")
	return &FileStore{
		store: []Person{p},
	}
}

type FileStore struct {
	store []Person
}

func (f *FileStore) GetAll() []Person {
	return f.store
}

// func (p *PersonStore) Add(person Person) error {}
