package pkg
import "github.com/google/uuid"
import "time"
type Company struct {
	ID uuid.UUID `db:"id"`
	Name string `db:"name"`
	Industry string `db:"industry"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type CompanyStore interface {
	Company(id uuid.UUID) (Company,error)
	CompaniesByUser(userID uuid.UUID)([]Company, error)
	CreateCompany(c *Company) error
	UpdateCompnay(c *Company) error
	DeleteCompany(id uuid.UUID) error
}
