package schemas

import "time"

type VacancySchema struct {
	Name        string    `json:"name"`
	CompanyID   uint      `json:"company_id"`
	VacancyLink string    `json:"vacancy_link"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Experience  int       `json:"experience"`
}
