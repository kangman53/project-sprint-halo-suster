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
			ON users (created_at DESC);
		CREATE INDEX IF NOT EXISTS users_created_at_asc
			ON users (created_at ASC);
        `,
		`
		CREATE TABLE IF NOT EXISTS patients (
			id VARCHAR(100) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			identity_number VARCHAR(20) UNIQUE,
			phone_number VARCHAR (20) NOT NULL,
            name VARCHAR(40) NOT NULL,
			gender VARCHAR(10) NOT NULL,
			birth_date TIMESTAMP NOT NULL,
			identity_card_scan_img TEXT,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS patients_identity_number
			ON patients (identity_number);
		CREATE INDEX IF NOT EXISTS patients_name
			ON patients USING HASH(lower(name));
		CREATE INDEX IF NOT EXISTS patients_phone_number
			ON patients (phone_number);
		CREATE INDEX IF NOT EXISTS patients_created_at_desc
			ON patients (created_at DESC);
		CREATE INDEX IF NOT EXISTS patients_created_at_asc
			ON patients (created_at ASC);
		`,
		`
		CREATE TABLE IF NOT EXISTS medical_records (
			id VARCHAR(100) PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
			patient_identity_number VARCHAR(20) NOT NULL,
			symptoms TEXT NOT NULL,
            medications TEXT NOT NULL,
			created_by VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (patient_identity_number) REFERENCES patients(identity_number) ON DELETE NO ACTION,
			FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE NO ACTION
		);
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
