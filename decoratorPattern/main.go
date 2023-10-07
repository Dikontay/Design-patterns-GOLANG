package main

import "fmt"

type Course interface {
	getDescription() string
	getPrice() float64
}

type CourseDecorator struct {
	decoratedCourse Course
}

type English struct {
	Course CourseDecorator
}

type Math struct {
	Course CourseDecorator
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

func (c *CourseDecorator) getDescription() string {
	return c.decoratedCourse.getDescription()
}

func (c *CourseDecorator) getPrice() float64 {

	return c.decoratedCourse.getPrice()
}

func main() {
	englishCourse := &English{}
	mathCourse := &Math{}

	// Decorate the English course
	decoratedEnglish := &CourseDecorator{decoratedCourse: englishCourse}

	// Decorate the Math course
	decoratedMath := &CourseDecorator{decoratedCourse: mathCourse}

	fmt.Println("English Course Description:", decoratedEnglish.getDescription())
	fmt.Println("English Course Price: $", decoratedEnglish.getPrice())

	fmt.Println("Math Course Description:", decoratedMath.getDescription())
	fmt.Println("Math Course Price: $", decoratedMath.getPrice())
}
