package backend

import "github.com/google/uuid"

type Company struct {
	ID uuid.UUID `db:id`
	Title string `db:title`

}