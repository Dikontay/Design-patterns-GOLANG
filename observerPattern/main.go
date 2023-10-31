package observer

import (
	"fmt"
)

type IWebsite interface{
	Add(o IClientAccount)
	Notify()
	Delete(o IClientAccount)
	getNewCourses()
	addNewCourse()
}

type Course struct {
	CourseName string
	StartingDate string
	EndingDate string
}

type IClientAccount interface {
	Update()
}

type Client struct{
	ID int
	PhoneNumber string
	Email string
	UserName string
}


func (o *Client) Update() {
	fmt.Printf("Check the updated courses %s\n", o.UserName)
}

type WebSite struct{
	AllClients [] IClientAccount
	CoursesList [] Course
}

func (o *WebSite) addNewCourse (c Course){
	o.CoursesList=append(o.CoursesList, c)
	fmt.Printf("There is new course: %s, staring date: %s, ending date: %s\n",c.CourseName, c.StartingDate, c.EndingDate )
}

func(w *WebSite) Delete(o IClientAccount) {
	var isDeleted bool
	var deletedUser IClientAccount
	for i, client := range w.AllClients {
        if client == o {
            w.AllClients = append(w.AllClients[:i], w.AllClients[i+1:]...)
			isDeleted = true
			deletedUser = client
            break
        }
    }
	if !isDeleted {
		fmt.Println("the user is not in list\n")
	} else {
		fmt.Printf("The user %d was removed from system\n", deletedUser)
	}

}

func(o WebSite) getNewCourses()[]Course{
	return o.CoursesList
}

func (w *WebSite) Add(client IClientAccount){
	w.AllClients=append(w.AllClients, client)
	
}

func (o *WebSite) Notify(){
	for _, client := range o.AllClients {
		client.Update()
	}
}

func CheckObserver() {
	website := &WebSite{}

	client1 := &Client{ID: 1, PhoneNumber: "123-456-7890", Email: "client1@example.com", UserName: "User1"}
	client2 := &Client{ID: 2, PhoneNumber: "987-654-3210", Email: "client2@example.com", UserName: "User2"}
	client3 := &Client{ID: 3, PhoneNumber: "987-654-3210", Email: "client2@example.com", UserName: "User3"}

    website.Add(client1)
	website.Add(client2)
	website.Add(client3)

	course1 := Course{CourseName: "English", StartingDate: "2023-10-01", EndingDate: "2023-10-31"}
	course2 := Course{CourseName: "Math", StartingDate: "2023-11-01", EndingDate: "2023-11-30"}


	website.addNewCourse(course1)
	website.addNewCourse(course2)

	website.Notify()
	course3 := Course{CourseName: "Web-Devolopment", StartingDate: "2023-11-01", EndingDate: "2023-11-30"}
	website.addNewCourse(course3)
	website.Notify()
	website.Delete(client3)

	course4 := Course{CourseName: "Go-lang development", StartingDate: "2023-11-01", EndingDate: "2023-11-30"}
	website.addNewCourse(course4)
	website.Notify()
}

// func main(){
// 	CheckObserver()
// }

