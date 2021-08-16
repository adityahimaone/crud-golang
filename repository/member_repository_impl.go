package repository

import (
	"context"
	"database/sql"
	"fmt"
	"onepage/entity"
)

type memberRepositoryImpl struct {
	DB *sql.DB
}

func NewMemberRepository(db *sql.DB) MemberRepository {
	return &memberRepositoryImpl{DB: db}
}
func (repository *memberRepositoryImpl) FindAll(ctx context.Context) ([]entity.Fellow, error) {
	query := "SELECT id,name,job,hobby FROM member"
	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fellow := entity.Fellow{}
	var fellows []entity.Fellow
	for rows.Next() {
		var id int32
		var name, job, hobby string
		err := rows.Scan(&id, &name, &job, &hobby)
		if err != nil {
			panic(err)
		}
		fellow.Id = id
		fellow.Name = name
		fellow.Job = job
		fellow.Hobby = hobby
		fellows = append(fellows, fellow)
	}
	defer rows.Close()
	return fellows, nil
}
func (repository *memberRepositoryImpl) Insert(ctx context.Context, member entity.Fellow) (entity.Fellow, error) {
	query := "INSERT INTO member(name,job,hobby) VALUES(?,?,?)"
	result, err := repository.DB.ExecContext(ctx, query, member.Name, member.Job, member.Hobby)
	if result != nil{
		fmt.Println("Success Insert")
	}
	id, err := result.LastInsertId()
	if err != nil {
		return member, err
	}
	member.Id = int32(id)
	return member, nil
}
