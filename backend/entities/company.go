package entities

type Company struct {
	CompanyId int    `json:"company_id" db:"company_id"`
	Name      string `json:"name" db:"name"`
}
