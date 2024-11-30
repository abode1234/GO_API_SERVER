package migrations

import (
	"database/sql"
)

func CreateImagesTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS images (
		id INT AUTO_INCREMENT PRIMARY KEY,
		path VARCHAR(255) NOT NULL
	);
	`
	_, err := db.Exec(query)
	return err
}
