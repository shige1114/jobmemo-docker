package backend

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Future    string    `db:"future"`
	Pr        string    `db:"pr"`
	GoodPoint string    `db:"good_point"`
	BadPoint  string    `db:"bad_point"`
	//CreatedAt string    `db:"created_at"`
	//UpdatedAt string    `db:"updated_at"`
}
type Cores struct {
	Id     uuid.UUID `db:"id"`
	Title  string    `db:"title"`
	Reason string    `db:"reason"`

	UsersId uuid.UUID `db:"users_id"`
}
type Companies struct {
	Id       uuid.UUID `db:"id"`
	Name     string    `db:"name"`
	Industry string    `db:"industry"`
}
type Recruits struct {
	UsersId      uuid.UUID `db:"users_id"`
	Companies_id uuid.UUID `db:"companies_id"`

	Reject        bool      `db:"reject"`
	Offer         bool      `db:"offer"`
	Selections_id uuid.UUID `db:"selections_id"`
	Motivation    string    `db:"motivation"`
	Good_point    string    `db:"good_point"`
	Concerns      string    `db:"concerns"`
}
type Selections struct {
	Id uuid.UUID `db:"id"`

	Level string `db:"level"`
	Types string `db:"type"`

	Person    string    `db:"person"`
	Adjusting bool      `db:"adjusting"`
	Date      time.Time `db:"date"`

	Pass bool `db:"pass"`
	Fail bool `db:"fail"`

	Good_point bool   `db:"good_point"`
	Memo       string `db:"memo"`

	UsersId      string `db:"users_id"`
	Companies_id string `db:"companies_id"`
}

type Questions struct {
	Id uuid.UUID `db:"id"`

	Title  string `db:"title"`
	Answer string `db:"answer"`
}

type UsersStore interface {
	Users(ctx context.Context, id uuid.UUID) (Users, error)
	Insert(ctx context.Context, user Users) error
	Update(ctx context.Context, user Users) error
	Delete(ctx context.Context, id string) error
}
type CoresStore interface {
	Core(users_id string) (Cores, error)
	Cores(users_id string) ([]Cores, error)
	Insert(core *Cores) error
	Update(core *Cores) error
	Delete(id uuid.UUID) error
}
type CompaniesStore interface {
	Companies(id uuid.UUID) (Companies, error)
	Insert(company *Companies) error
	Update(company *Companies) error
	Delete(id uuid.UUID) error
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
	Companies_id    uuid.UUID `db:"companies_id"`
	Companies_title string    `db:"companies_title"`

	Selections_level    string    `db:"selections_level"`
	Selections_adusting bool      `db:"selections_adusting"`
	Selections_date     time.Time `db:"selections_date"`

	Recruits_reject bool `db:"recruits_reject"`
	Recruits_offer  bool `db:"recruits_offer"`
}
type Main struct {
	Companies_name   string
	recruits         Recruits
	latest_selection Selections
	Selections_ids   []uuid.UUID
}

type SideStore interface {
	Sides(users_id uuid.UUID) ([]Side, error)
}
type MainStore interface {
	Main(users_id, companis_id string) (Main, error)
}

type Store interface {
	SideStore
	MainStore
	UsersStore
}
