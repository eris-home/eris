package storage

import (
	"fmt"
)

type Storage struct {
	Database *Database
	Indexer  *Indexer
}

func NewStorage(db *Database) *Storage {
	return &Storage{
		Database: db,
		Indexer:  NewIndexer(),
	}
}

func (s *Storage) SaveData(key string, documentID string, data string) error {
	query := "INSERT INTO documents (key, document_id, data) VALUES ($1, $2, $3)"
	_, err := s.Database.ExecuteQuery(query, key, documentID, data)
	if err != nil {
		return fmt.Errorf("failed to save data: %w", err)
	}

	s.Indexer.AddToIndex(key, documentID)
	fmt.Printf("Data saved and indexed: %s\n", documentID)
	return nil
}

func (s *Storage) RetrieveData(key string) ([]string, error) {
	documents := s.Indexer.SearchIndex(key)
	if len(documents) == 0 {
		return nil, fmt.Errorf("no documents found for key: %s", key)
	}

	fmt.Printf("Documents retrieved for key %s: %v\n", key, documents)
	return documents, nil
}