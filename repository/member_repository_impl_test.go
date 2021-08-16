package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func Test(t *testing.T) {
	memberRepository := NewMemberRepository(OpenConnection())
	ctx := context.Background()
	result, err := memberRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, value := range result{
		fmt.Println(value)
	}
}

