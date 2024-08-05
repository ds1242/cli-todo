package main

import (
	"time"
)

type Row struct {
	ID			int
	Description string
	CreatedAt	time.Time
	IsComplete	bool
}	