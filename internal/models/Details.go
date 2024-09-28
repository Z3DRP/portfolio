package models

type XpItem struct {
	Item string
	Yrs  float32
}

type PrevWork struct {
	Company     string
	Duration    float32
	Description string
}

type Details struct {
	Name        string
	Title       string
	Exp         []PrevWork //
	ExpOverview string     // description of experience
	Misc        string
}

func NewDetails(nm, ttl, xpOvw, msc string, exp []PrevWork) *Details {
	return &Details{
		Name:        nm,
		Title:       ttl,
		Exp:         exp,
		ExpOverview: xpOvw,
	}
}

func NewXpItem(item string, yrs float32) *XpItem {
	return &XpItem{
		Item: item,
		Yrs:  yrs,
	}
}

func NewPrevWork(company, descript string, duration float32) *PrevWork {
	return &PrevWork{
		Company:     company,
		Description: descript,
		Duration:    duration,
	}
}

// <li>Go</li>
// <li>Python</li>
// <li>Java</li>
// <li>SQL</li>
// <li>Postgres</li>
// <li>MongoDB</li>
