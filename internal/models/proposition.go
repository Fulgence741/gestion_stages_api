package models

import (
	"time"
)

type Proposition struct {
	IDEtablissement int64
	IDFiliere       int64
	DateProposition time.Time
}
