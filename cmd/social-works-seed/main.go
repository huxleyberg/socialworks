package main

import (
	"os"

	"github.com/huxleyberg/socialworks/internal/db"
	"github.com/huxleyberg/socialworks/internal/posts"
	"github.com/huxleyberg/socialworks/internal/users"
)

func main() {
	postgresDbUrl := os.Getenv("POSTGRES_DB_URL")
	dbConn := db.ProvidePostgres(postgresDbUrl)
	defer db.Close(dbConn)

	postsRepo := posts.NewPostRepository(dbConn)
	commentsRepo := posts.NewCommentRepository(dbConn)
	usersRepo := users.NewUserRepository(dbConn)

	db.Seed(commentsRepo, postsRepo, usersRepo)
}
