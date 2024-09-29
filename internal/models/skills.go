package models

import "fmt"

type XpItem struct {
	Name string
	Yrs  float32
}

func NewXpItem(name string, yrs float32) *XpItem {
	return &XpItem{
		Name: name,
		Yrs:  yrs,
	}
}

func (xp XpItem) FormattedYrs() string {
	return fmt.Sprintf("%.2f", xp.Yrs)
}

type Skills []XpItem

func TempSkillFactory() Skills {
	return []XpItem{
		{Name: "Golang", Yrs: 1.0},
		{Name: "Javascript", Yrs: 2.5},
		{Name: "React", Yrs: 2.0},
		{Name: "Android", Yrs: 1.0},
		{Name: "Java", Yrs: 2.2},
		{Name: "Kotlin", Yrs: 0.8},
		{Name: "SQL", Yrs: 3.0},
		{Name: "Postgres", Yrs: 1.5},
		{Name: "Docker", Yrs: 1.5},
		{Name: "Git", Yrs: 4.0},
		{Name: "Unit Testing", Yrs: 2.0},
	}
}
