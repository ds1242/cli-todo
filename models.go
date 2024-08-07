package main

import (
	"time"
)

type Row struct {
	ID 			int
	Task		string
	CreatedAt	time.Time
	Completed	bool
}