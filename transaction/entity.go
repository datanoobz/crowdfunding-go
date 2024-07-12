package transaction

import "time"

type Transaction struct {
	ID        int
	Campaign  int
	UserID    int
	Amount    int
	Status    string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
