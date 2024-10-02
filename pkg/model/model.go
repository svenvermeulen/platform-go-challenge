package model

import (
	"time"

	"github.com/google/uuid"
)

type Chart struct {
	Id         uuid.UUID
	Title      string
	XAxisTitle string
	YAxisTitle string
	DataPoints []DataPoint
}

type DataPoint struct {
	X float64
	Y float64
}

type Insight struct {
	Id uuid.UUID
	Description string
}

type Audience struct {
	Id uuid.UUID
	Gender rune
	BirthCountry string
	AgeFrom int
	AgeTo int
	HoursSpentFrom int
	HoursSpentTo int
	PurchasesFrom int
	PurchasesTo int
}

type User struct {
	Id uuid.UUID
	FirstName string
	LastName string
	Gender rune
	BirthDate time.Time
}

type UserFavourites struct {
	Charts    []Chart
	Insights  []Insight
	Audiences []Audience
}