package main

import "github.com/google/uuid"

type Project struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Members     []Member  `json:"members"`
	Tasks       []Task    `json:"tasks"`
}

type Member struct {
	Email      string
	Name       string
	Lastname   string
	Permission int
}

type Task struct {
	Name           string
	Descripciption string
	Assignes       []Member
	State          int
}

func CreateNewProject(title string, descripion string, members []Member, tasks []Task) Project {
	id := uuid.New()
	newProject := Project{
		Id:          id,
		Title:       title,
		Description: descripion,
		Members:     append(members),
		Tasks:       append(tasks),
	}
	return newProject
}
