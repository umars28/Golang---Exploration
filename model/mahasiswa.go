package model

import (
	"time"
)

type (
	Mahasiswa struct {
		ID        int
		NIM       int
		Name      string
		Semester  int
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)
