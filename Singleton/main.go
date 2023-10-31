package main

import (
	observer "design/observerPattern"
	"fmt"
	"sync"
)

type CourseManager struct {
	UniqueCourses map[string]*observer.Course
}

var (
	instance *CourseManager
	mu       sync.Mutex
)

func GetCourseManagerInstance() (*CourseManager, error) {
	if instance== nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &CourseManager{
				UniqueCourses: make(map[string]*observer.Course),
			}
		}
		return instance, nil
	} else {
		return nil, fmt.Errorf("The CourseManager instance is already created.\nYou cannot create another instance")
	}

}

func main() {
	//create first course instance
	cm, err := GetCourseManagerInstance()
	if err != nil {
		fmt.Println(err)
		return
	}
	cm.UniqueCourses["math101"] = &observer.Course{CourseName: "Math 101", StartingDate: "25.10.2023", EndingDate: "25.11.2023"}
	cm.UniqueCourses["compSci101"] = &observer.Course{CourseName: "Computer Science 101", StartingDate: "19.11.2023", EndingDate: "23.02.2024"}
	fmt.Println("Course 1:", cm.UniqueCourses["math101"].CourseName)
	//cannot create second course instance
	cm2, err := GetCourseManagerInstance()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println("Course 1:", cm.UniqueCourses["compSci101"].CourseName)
	fmt.Println("Course 2:", cm2.UniqueCourses["compSci101"].CourseName)
}
