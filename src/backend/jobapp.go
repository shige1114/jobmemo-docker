package backend

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Email string    `db:"email"`

	Future    string `db:"future"`
	Pr        string `db:"pr"`
	GoodPoint string `db:"good_point"`
	BadPoint  string `db:"bad_point"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
type Cores struct {
	Id     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	reason string    `db:"reason"`

	UsersId uuid.UUID `db:"users_id"`
}
type Companies struct {
	id       uuid.UUID `db:"id"`
	title    string    `db:"name"`
	industry string    `db:"industry"`
}
type Recruits struct {
	UsersId     uuid.UUID `db:"users_id"`
	CompaniesId uuid.UUID `db:"companies_id"`

	Reject     bool   `db:"reject"`
	Offer      bool   `db:"offer"`
	Motivation string `db:"motivation"`
	GoodPoint  string `db:"good_point"`
	Concerns   string `db:"concerns"`
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

	UsersId      string `db:"users_id"`
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
	Companies_id   uuid.UUID `db:"companies_id"`
	Companies_name string    `db:"companies_name"`

	Selections_level string    `db:"selections_level"`
	Selections_date  time.Time `db:"selections_date"`
	Adjusting        bool      `db:"adjusting"`

	Recruits_reject bool `db:"recruits_reject"`
	Recruits_offer  bool `db:"recruits_offer"`
}
type Main struct {
	Recruit Recruits
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
