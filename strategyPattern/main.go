package main

import "fmt"

//identifying level by testing the clients is our main task.
//There are a lot of courses but some of them needs a specific type of test
//For example if the client chooses the english, spanish or other language courses
//the type of testing for all languages are the same. However if at runtime the client chooses math or programming or other
//courses we cannot use the test type of language courses. we need to provide a specific test in some cases
//(Python programming or Go programming for both of them there are quiz and contester for testing basic knowledge )

// interface strategy
type EnterCourse interface {
	PassTest() string
}

// concrete strategies
type EnglishTest struct {
	Speaking  string
	Reading   string
	Writing   string
	Listening string
}
type SpanishTest struct {
	Speaking  string
	Reading   string
	Writing   string
	Listening string
}
type MathTest struct {
	Algebra  string
	Geometry string
}

type ProgrammingTest struct {
	Contester string
	Quiz      string
}

// context
type Course struct {
	CourseName    string
	EnterStrategy EnterCourse
}

func (e *EnglishTest) PassTest() string {
	return fmt.Sprintf("Pass the English test - Speaking: %s, Reading: %s, Writing: %s, Listening: %s", e.Speaking, e.Reading, e.Writing, e.Listening)
}

func (m *MathTest) PassTest() string {
	return fmt.Sprintf("Pass the Math test - Algebra: %s, Geometry: %s", m.Algebra, m.Geometry)
}

func (p *ProgrammingTest) PassTest() string {
	return fmt.Sprintf("Pass the Programming test - Contester: %s, Quiz: %s", p.Contester, p.Quiz)
}

func checkStrategy() {
	englishCourse := &Course{
		CourseName:    "English Course",
		EnterStrategy: &EnglishTest{Speaking: "Good", Reading: "Excellent", Writing: "Very Good", Listening: "Average"},
	}

	spanishCourse := &Course{
		CourseName:    "Spanish Course",
		EnterStrategy: &EnglishTest{Speaking: "santa mierda", Reading: "Bad", Writing: "Not good", Listening: "excelente"},
	}

	mathCourse := &Course{
		CourseName:    "Math Course",
		EnterStrategy: &MathTest{Algebra: "Pass", Geometry: "Excellent"},
	}

	programmingCourse := &Course{
		CourseName:    "Programming Course",
		EnterStrategy: &ProgrammingTest{Contester: "Pass", Quiz: "Very Good"},
	}
	fmt.Println("For English Course:", englishCourse.EnterStrategy.PassTest())
	fmt.Println("For Math Course:", mathCourse.EnterStrategy.PassTest())
	fmt.Println("For Programming Course:", programmingCourse.EnterStrategy.PassTest())
	fmt.Println("For Spanish Course:", spanishCourse.EnterStrategy.PassTest())

}

func main() {
	checkStrategy()
}
