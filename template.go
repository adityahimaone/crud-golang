package main

import (
	"context"
	"html/template"
	"net/http"
	"onepage/entity"
	repo "onepage/repository"
)

var header string = "./templates/header.gohtml"
var footer string = "./templates/footer.gohtml"


func TemplateIndex(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(header, footer, "templates/index.gohtml"))
	memberRepository := repo.NewMemberRepository(repo.OpenConnection())
	ctx := context.Background()
	table, _ := memberRepository.FindAll(ctx)

	err := t.ExecuteTemplate(writer, "Index", table)
	if err != nil {
		panic(err)
	}
}

func TemplateForm(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(header, footer, "templates/form.gohtml"))
	memberRepository := repo.NewMemberRepository(repo.OpenConnection())
	ctx := context.Background()
	name := request.PostFormValue("name")
	job := request.PostFormValue("job")
	hobby := request.PostFormValue("hobby")
	member := entity.Fellow{
		Name:  name,
		Job:   job,
		Hobby: hobby,
	}
	if name != "" || job != "" || hobby != "" {
		_, err := memberRepository.Insert(ctx, member)
		if err != nil {
			panic(err)
		}
	}
	err := t.ExecuteTemplate(writer, "Form", nil)
	if err != nil {
		panic(err)
	}

}
