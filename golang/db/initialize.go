package database

import (
	"context"
	"fmt"

	passes "biblioteka-backend/packages"

	"github.com/jackc/pgx/v5"
)

func Initialize(query string) ([]interface{}, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, passes.Certs("nazwa"))
	if err != nil {
		return nil, fmt.Errorf("błąd połączenia z bazą: %v", err)
	}
	defer conn.Close(ctx)

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("błąd zapytania: %v", err)
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, fmt.Errorf("błąd pobierania wartości: %v", err)
		}
		results = append(results, values...)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("błąd podczas iteracji po wierszach: %v", rows.Err())
	}

	return results, nil
}
