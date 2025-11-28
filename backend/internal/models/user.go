package user

import "time"

type User struct {
	ID        int
	email     string
	password  string
	name      string
	createdAt time.Time
}
