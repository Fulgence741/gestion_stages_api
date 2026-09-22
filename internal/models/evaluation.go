package models

import (
	"time"
)

type Evaluation struct {
	ID             int64
	Note           float64
	Commentaire    string
	DateEvaluation time.Time
	IDRapport      int64
	IDUser         int64
}
