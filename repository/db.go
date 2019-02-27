package repository

import (
	"fmt"
	"sync"
	"time"
)

type Storage struct {
	s	sync.RWMutex
	cells map[string]*Cell
}

type Cell struct {
	Value    string
	TimeLife int64
	SetTime int64
}

var storage *Storage
var once sync.Once

func GetInstance() *Storage {
	once.Do(func() {
		storage = &Storage{cells: make(map[string]*Cell)}
	})
	return storage
}

func (db *Storage) Set(key string, value string, timeLife int64) error {

	db.s.Lock()
	defer db.s.Unlock()
	db.cells[key] = &Cell{Value:value, TimeLife:timeLife, SetTime: time.Now().Unix()}
	return nil
}

func (db *Storage) Get(key string) (*Cell, error) {

	db.s.RLock()
	defer db.s.RUnlock()
	if cell, ok := db.cells[key]; ok {
		return cell, nil
	}
	return nil, fmt.Errorf("not found")
}

func (db *Storage) Delete(key string) error {

	db.s.Lock()
	defer db.s.Unlock()
	if _, ok := db.cells[key]; ok {
		delete(db.cells, key)
		return nil
	}
	return fmt.Errorf("not found")

}

func (db *Storage) Clean() {
	db.s.Lock()
	defer db.s.Unlock()
	now := time.Now().Unix()
	for key, cell := range db.cells {
		if cell.SetTime < (now - cell.TimeLife) {
			delete(db.cells, key)
		}
	}
}

func (db *Storage) StartClean(timeout int) {
	for {
		db.Clean()
		time.Sleep(time.Duration(timeout) * time.Second)
	}
}