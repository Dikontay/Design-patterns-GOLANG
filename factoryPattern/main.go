package main

import "fmt"

type Course interface {
	getDescription() string
	getPrice() float64
}

type English struct {
	course Course
}

type Math struct {
	Course Course
}

func (m *Math) getDescription() string {
	return "Geometry for preparing SAT"
}

func (m *Math) getPrice() float64 {
	return 42500.0
}

func (e *English) getDescription() string {
	return "English course for intermediate level"
}

func (e *English) getPrice() float64 {
	return 15000.0
}

type CourseFactory interface {
	createCourse() Course
}

type EnglisCourseFactory struct {
	Location     string
	StartingDate string
	EndingDate   string
}

func (ef *EnglisCourseFactory) createCourse() Course {
	return &English{}
}

type MathCourseFactory struct {
	Location     string
	StartingDate string
	EndingDate   string
}

func (m *MathCourseFactory) createCourse() Course {
	return &Math{}
}

func main() {
	// Create an English course factory
	englishFactory := &EnglisCourseFactory{
		Location:     "Online",
		StartingDate: "2023-10-15",
		EndingDate:   "2023-12-15",
	}

	// Use the factory to create an English course
	englishCourse := englishFactory.createCourse()

	// Get the description and price of the English course
	fmt.Println("English Course Description:", englishCourse.getDescription())
	fmt.Printf("English Course Price: $%.2f\n", englishCourse.getPrice())

	// Create a Math course factory
	mathFactory := &MathCourseFactory{
		Location:     "In-Person",
		StartingDate: "2023-11-01",
		EndingDate:   "2023-12-31",
	}

	// Use the factory to create a Math course
	mathCourse := mathFactory.createCourse()

	// Get the description and price of the Math course
	fmt.Println("Math Course Description:", mathCourse.getDescription())
	fmt.Printf("Math Course Price: $%.2f\n", mathCourse.getPrice())
}
