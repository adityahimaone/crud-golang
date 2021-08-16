package main

import (
	"context"
	"fmt"
	"onepage/entity"
	repo "onepage/repository"
	"testing"
)

func TestInsert(t *testing.T) {
	memberRepository := repo.NewMemberRepository(repo.OpenConnection())
	ctx := context.Background()
	member := entity.Fellow{
		Name:  "request.PostFormValue(name)",
		Job:   "request.PostFormValue(job)",
		Hobby: "request.PostFormValue(hobby)",
	}
	insert, err := memberRepository.Insert(ctx, member)
	if err != nil {
		panic(err)
	}
	fmt.Println(insert)
}
