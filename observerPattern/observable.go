package main

import (
	"fmt"
)

type IObservable interface{
	Add(c Client)
	Notify()
	Delete(c Client)
	getNewCourses()
	addNewCourse()
}

type Course struct {
	CourseName string
	StartingDate string
	EndingDate string
}

type IObserver interface {
	Update()
}

type Client struct{
	ID int
	PhoneNumber string
	Email string
	UserName string
}

func (o *Client) Update() {
	fmt.Println("Check the updated courses")
}

type WebSite struct{
	AllClients [] Client
	CoursesList [] Course
}

func (o *WebSite) addNewCourse (c Course){
	o.CoursesList=append(o.CoursesList, c)
}

func(o WebSite) getNewCourses() []Course{
	return o.CoursesList
}

func (w *WebSite) Add(client Client){
	w.AllClients=append(w.AllClients, client)
}

func (o *WebSite) Notify(){
	for _, client := range o.AllClients {
		client.Update()
	}
}

func CheckObserver() {
	website := &WebSite{}

	client1 := Client{ID: 1, PhoneNumber: "123-456-7890", Email: "client1@example.com", UserName: "User1"}
	client2 := Client{ID: 2, PhoneNumber: "987-654-3210", Email: "client2@example.com", UserName: "User2"}


    website.Add(client1)
	website.Add(client2)

	course1 := Course{CourseName: "English", StartingDate: "2023-10-01", EndingDate: "2023-10-31"}
	course2 := Course{CourseName: "Math", StartingDate: "2023-11-01", EndingDate: "2023-11-30"}


	website.addNewCourse(course1)
	website.addNewCourse(course2)

	website.Notify()
}

func main(){
	CheckObserver()
}

