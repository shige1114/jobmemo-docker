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
	Name   string    `db:"name"`
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

	Reject bool `db:"reject"`
	Offer  bool `db:"offer"`

	Motivation string `db:"motivation"`
	Good_point string `db:"good_point"`
	Concerns   string `db:"concerns"`
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
	Id   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name"  json:"name"`

	Level     int       `db:"level" json:"level"`
	Type      int       `db:"type" json:"type"`
	Adjusting bool      `db:"adjusting" json:"adjusting"`
	Date      time.Time `db:"date" json:"date"`

	Reject bool `db:"reject" json:"reject"`
	Offer  bool `db:"offer" json:"offer"`
}
type Main struct {
	Name string `db:"name"  json:"name"`

	Reject    bool      `db:"reject" json:"reject"`
	Offer     bool      `db:"offer" json:"offer"`
	Adjusting bool      `db:"adjusting" json:"adjusting"`
	Date      time.Time `db:"date" json:"date"`
	Level     int       `db:"level" json:"level"`
	Type      int       `db:"type" json:"type"`

	Motivation string `db:"motivation" json:"motivation"`
	Good_point string `db:"good_point" json:"good_point"`
	Concerns   string `db:"concerns" json:"concerns"`

	Selections map[int]int `json:"selections"`
}

type SideStore interface {
	Sides(users_id uuid.UUID) ([]Side, error)
}
type MainStore interface {
	Main(users_id, companis_id string) (Main, error)
}
type Error struct {
	Err string `json:"err"`
}

type Store interface {
	SideStore
	MainStore
	UsersStore
}
