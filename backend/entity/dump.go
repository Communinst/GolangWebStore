package entities

import "time"

type Dump struct {
	ID        int       `db:"id" json:"id"`                 // Unique identifier for the dump record
	Filename  string    `db:"filename" json:"filename"`     // The name of the dump file
	CreatedAt time.Time `db:"created_at" json:"created_at"` // Timestamp when the dump was created
	Size      int64     `db:"size" json:"size"`             // Size of the dump file in bytes (optional)
}
