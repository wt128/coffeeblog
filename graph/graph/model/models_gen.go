// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Coffee struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	ProducingArea *string  `json:"producing_area"`
	HighPrice     *int     `json:"high_price"`
	LowPrice      *int     `json:"low_price"`
	StandardPrice []*Phase `json:"standard_price"`
	Richness      []*Phase `json:"richness"`
	Bitterness    []*Phase `json:"bitterness"`
	Acidity       []*Phase `json:"acidity"`
	Sweetness     []*Phase `json:"sweetness"`
}

type Comment struct {
	ID       string  `json:"id"`
	CoffeeID string  `json:"coffee_id"`
	UserID   string  `json:"user_id"`
	Title    string  `json:"title"`
	Comment  *string `json:"comment"`
}

type NewCoffee struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ProducingArea *string `json:"producing_area"`
	HighPrice     *int    `json:"high_price"`
	LowPrice      *int    `json:"low_price"`
	StandardPrice Phase   `json:"standard_price"`
	Richness      Phase   `json:"richness"`
	Bitterness    Phase   `json:"bitterness"`
	Acidity       Phase   `json:"acidity"`
	Sweetness     Phase   `json:"sweetness"`
}

type NewUser struct {
	Name string  `json:"name"`
	Mail *string `json:"mail"`
}

type User struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	LikeCoffee string  `json:"like_coffee"`
	Mail       *string `json:"mail"`
}

type Phase string

const (
	PhaseOne   Phase = "ONE"
	PhaseTwo   Phase = "TWO"
	PhaseThree Phase = "THREE"
	PhaseFour  Phase = "FOUR"
	PhaseFive  Phase = "FIVE"
)

var AllPhase = []Phase{
	PhaseOne,
	PhaseTwo,
	PhaseThree,
	PhaseFour,
	PhaseFive,
}

func (e Phase) IsValid() bool {
	switch e {
	case PhaseOne, PhaseTwo, PhaseThree, PhaseFour, PhaseFive:
		return true
	}
	return false
}

func (e Phase) String() string {
	return string(e)
}

func (e *Phase) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Phase(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PHASE", str)
	}
	return nil
}

func (e Phase) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
