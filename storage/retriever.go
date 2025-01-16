package storage

import (
	"database/sql"
	"fmt"
)

type Retriever struct {
	DB *Database
}

func NewRetriever(db *Database) *Retriever {
	return &Retriever{DB: db}
}

func (r *Retriever) FetchByID(table string, id int) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table)
	row := r.DB.Connection.QueryRow(query, id)

	columns, err := row.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))
	for i := range values {
		valuePointers[i] = &values[i]
	}

	if err := row.Scan(valuePointers...); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	result := make(map[string]interface{})
	for i, colName := range columns {
		result[colName] = values[i]
	}

	return result, nil
}

func (r *Retriever) FetchAll(table string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := r.DB.Connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rows: %w", err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	var results []map[string]interface{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePointers := make([]interface{}, len(columns))
		for i := range values {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		rowResult := make(map[string]interface{})
		for i, colName := range columns {
			rowResult[colName] = values[i]
		}
		results = append(results, rowResult)
	}

	return results, nil
}