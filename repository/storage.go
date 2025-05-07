package repository

import (
	"chi-books/entity"
	"sync"
)

type storage struct {
	books map[string]*entity.Book
}

var storageInstance *storage = nil
var lock = &sync.Mutex{}

// apply singleton to storage
func getStorage() *storage {
	if storageInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		storageInstance = &storage{books: map[string]*entity.Book{}}
	}

	return storageInstance
}
