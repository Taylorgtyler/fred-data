package database

import (
	"context"
	"fmt"
	"log"
	"sync"
)

// QueryOptions allows customization of query execution
type QueryOptions struct {
	MaxRows       int
	LogQuery      bool
	ContextCancel bool
}

// DefaultQueryOptions provides standard configuration
func DefaultQueryOptions() QueryOptions {
	return QueryOptions{
		MaxRows:       0, // 0 means no limit
		LogQuery:      false,
		ContextCancel: false,
	}
}

// ExecuteQuery executes a database query with improved performance and flexibility
func (ctx *DBContext) ExecuteQuery(
	query string,
	args []interface{},
	opts ...QueryOptions,
) ([]map[string]interface{}, error) {
	// Use default options if not provided
	queryOpts := DefaultQueryOptions()
	if len(opts) > 0 {
		queryOpts = opts[0]
	}

	// Optional query logging
	if queryOpts.LogQuery {
		log.Printf("Executing query: %s, args: %v", query, args)
	}

	// Use a context with optional cancellation
	var cancel context.CancelFunc
	var dbContext context.Context
	if queryOpts.ContextCancel {
		dbContext, cancel = context.WithCancel(context.Background())
		defer cancel()
	} else {
		dbContext = context.Background()
	}

	// Prepare the query with context
	rows, err := ctx.DB.QueryContext(dbContext, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query %q: %w", query, err)
	}
	defer rows.Close()

	// Get column information
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get column names for query %q: %w", query, err)
	}

	// Use a sync.Pool to reduce memory allocations
	var valuePool = sync.Pool{
		New: func() interface{} {
			return make([]interface{}, len(columns))
		},
	}

	var result []map[string]interface{}
	rowCount := 0

	for rows.Next() {
		// Check row limit if specified
		if queryOpts.MaxRows > 0 && rowCount >= queryOpts.MaxRows {
			break
		}

		// Get pooled values slice
		values := valuePool.Get().([]interface{})
		pointers := make([]interface{}, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			valuePool.Put(values)
			return nil, fmt.Errorf("failed to scan row in query %q: %w", query, err)
		}

		row := make(map[string]interface{}, len(columns))
		for i, column := range columns {
			row[column] = values[i]
		}
		result = append(result, row)

		// Return values to the pool
		valuePool.Put(values)
		rowCount++
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows for query %q: %w", query, err)
	}

	return result, nil
}
