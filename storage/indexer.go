package storage

import (
	"fmt"
	"strings"
)

type Indexer struct {
	Index map[string][]string
}

func NewIndexer() *Indexer {
	return &Indexer{Index: make(map[string][]string)}
}

func (i *Indexer) AddToIndex(key string, documentID string) {
	key = strings.ToLower(key)
	i.Index[key] = append(i.Index[key], documentID)
	fmt.Printf("Document %s added to index with key: %s\n", documentID, key)
}

func (i *Indexer) SearchIndex(key string) []string {
	key = strings.ToLower(key)
	return i.Index[key]
}

func (i *Indexer) ListIndex() {
	for key, documents := range i.Index {
		fmt.Printf("Key: %s, Documents: %v\n", key, documents)
	}
}