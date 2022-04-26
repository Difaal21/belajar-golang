package repository

import (
	"context"
	"fmt"
	"testing"

	golang_database "github.com/Difaal21/belajar-golang/golang-database"
	"github.com/Difaal21/belajar-golang/golang-database/entity"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "alfansha21@gmail.com",
		Comment: "Unit Testing Untuk Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 1)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {

	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
