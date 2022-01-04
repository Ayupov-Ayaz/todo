package list

import "github.com/ayupov-ayaz/todo/pkg/repository"

type Repository struct {
	db repository.DbRepository
}

func NewRepository(db repository.DbRepository) Repository {
	return Repository{
		db: db,
	}
}
