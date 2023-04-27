package backend

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	id    uuid.UUID `db:"id"`
	name  string    `db:"name"`
	email string    `db:"email"`

	future     string `db:"future"`
	pr         string `db:"pr"`
	good_point string `db:"good_point"`
	bad_point  string `db:"bad_point"`

	created_at time.Time `db:"created_at"`
	updated_at time.Time `db:"updated_at"`
}
type Cores struct {
	id     uuid.UUID `db:"id"`
	title  string    `db:"title"`
	reason string    `db:"reason"`

	users_id uuid.UUID `db:"users_id"`
}
type Companies struct {
	id       uuid.UUID `db:"id"`
	title    string    `db:"title"`
	industry string    `db:"industry"`
}
type Recruits struct {
	users_id     uuid.UUID `db:"users_id"`
	companies_id uuid.UUID `db:"companies_id"`

	reject bool `db:"reject"`
	offer  bool `db:"offer"`

	motivation string `db:"motivation"`
	good_point string `db:"good_point"`
	concerns   string `db:"concerns"`
}
type Selections struct {
	id uuid.UUID `db:"id"`

	level string `db:"level"`
	types string `db:"type"`

	person    string    `db:"person"`
	adjusting bool      `db:"adjusting"`
	date      time.Time `db:"date"`

	pass bool `db:"pass"`
	fail bool `db:"fail"`

	good_point bool   `db:"good_point"`
	memo       string `db:"memo"`

	users_id     string `db:"users_id"`
	companies_id string `db:"companies_id"`
}

type Questions struct {
	id uuid.UUID `db:"id"`

	title  string `db:"title"`
	answer string `db:"answer"`
}

type UsersStore interface {
	Users(id uuid.UUID) (Users, error)
	UsersByMain(email string) (Users, error)
	CreateUsers(user *Users) error
	UpdateUsers(user *Users) error
	DeleteUsers(id uuid.UUID) error
}
type CoresStore interface {
	Cores(users_id string) ([]Cores, error)
	CreateCores(core *Cores) error
	UpdateCores(core *Cores) error
	DeleteCores(id uuid.UUID) error
}
type CompaniesStore interface {
	Companies(id uuid.UUID) (Companies, error)
	CompaniesByTitle(title uuid.UUID) (Companies, error)
	CreateCompanies(company *Companies) error
	UpdateCompanies(company *Companies) error
	DeleteCompanies(id uuid.UUID) error
}
type RecruitsStore interface {
	RecruitsByUsers(users_id uuid.UUID) ([]Recruits, error)
	CreateRecruits(recruit *Recruits) error
	UpdateRecruits(recruit *Recruits) error
	DeleteRecruitsByUsersID(users_id uuid.UUID) error
}
type SelectionsStore interface {
	SelectionsByUser(users_id, company_id uuid.Domain) ([]Selections, error)
	CreateSelections(selection *Selections) error
	UpdateSelections(seleciton *Selections) error
	DeleteSelections(id uuid.UUID) error
}
type QuestionsStore interface {
	QuestionsBySelection(selections_id uuid.UUID) (Questions, error)
	CreateQuestions(question *Questions) error
	UpdateQuestions(question *Questions) error
	DeleteQuestions(id uuid.UUID) error
}

type Side struct {
	companies_id    uuid.UUID `db:"companies_id"`
	companies_title string    `db:"companies_title"`

	selections_level    string    `db:"selections_level"`
	selections_adusting bool      `db:"selections_adusting"`
	selections_date     time.Time `db:"selections_date"`

	recruits_reject bool `db:"recruits_reject"`
	recruits_offer  bool `db:"recruits_offer"`
}
type Main struct {
}

type SideStore interface {
	Side(users_id uuid.UUID) ([]Side, error)
}
type MainStore interface {
}

type Store interface {
	SideStore
	MainStore
}
