package repositories

import db "forum/database"

type BaseRepo struct {
	DB        *db.Database
	TableName string
}

var (
	UserRepo  UserRepository
	PostRepo  PostRepository
	SessRepo  SessionRepository
	ReactRepo ReactionRepository
	CommRepo  CommentRepository
	CategRepo CatRepo
)

func init() {
	UserRepo.init()
	PostRepo.init()
	SessRepo.init()
	ReactRepo.init()
	CommRepo.init()
	CategRepo.init()
}
