package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitiateTables(dbPool *pgxpool.Pool) error {
	// Define table creation queries
	queries := []string{
		`
        CREATE TABLE IF NOT EXISTS users (
            id VARCHAR(100) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
            name VARCHAR(100) NOT NULL,
			nip VARCHAR(50) UNIQUE,
			password VARCHAR(255) NOT NULL,
            role VARCHAR(20) NOT NULL,
			identity_card_scan_img TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        );
		CREATE INDEX IF NOT EXISTS users_id
			ON users (id);
		CREATE INDEX IF NOT EXISTS users_nip
			ON users (nip);
		CREATE INDEX IF NOT EXISTS users_name
			ON users USING HASH(lower(name));
		CREATE INDEX IF NOT EXISTS users_role
				ON users (role);
		CREATE INDEX IF NOT EXISTS users_created_at_desc
			ON users(created_at DESC);
		CREATE INDEX IF NOT EXISTS users_created_at_asc
			ON users(created_at ASC);
        `,
		// Add more table creation queries here if needed
	}

	// Execute table creation queries
	for _, query := range queries {
		_, err := dbPool.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}

	return nil
}
