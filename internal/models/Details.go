package models

import "fmt"

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
}

func NewDetails(nm, ttl, xpOvw, msc string, exp []PrevWork) *Details {
	return &Details{
		Name:        nm,
		Title:       ttl,
		Exp:         exp,
		ExpOverview: xpOvw,
	}
}

func NewPrevWork(company, descript string, duration float32) *PrevWork {
	return &PrevWork{
		Company:     company,
		Description: descript,
		Duration:    duration,
	}
}

func TempPrevWorkFactory() []PrevWork {
	return []PrevWork{
		*NewPrevWork("Panera Bread", "Android Developer", 0.8),
		*NewPrevWork("1218 Global", "Full-stack Developer", 2.0),
	}
}

func (p PrevWork) FormattedDuration() string {
	return fmt.Sprintf("%.2f", p.Duration)
}

func TempDetailFactory() *Details {
	pWork := TempPrevWorkFactory()
	return &Details{
		Name:        "Zach Palmer",
		Title:       "Software Engineer",
		Exp:         pWork,
		ExpOverview: "3 Yrs of fullstack development experience",
	}
}

// <li>Go</li>
// <li>Python</li>
// <li>Java</li>
// <li>SQL</li>
// <li>Postgres</li>
// <li>MongoDB</li>
