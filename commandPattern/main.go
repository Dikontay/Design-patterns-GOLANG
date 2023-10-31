package main

import (
	observer "design/observerPattern"
	"fmt"
)

type Command interface {
	Execute()
}

type Subscription struct {
	subscriptionManager map[*observer.Course]bool
}

func (s *Subscription) Activate(course *observer.Course) {
	s.subscriptionManager[course] = true
	fmt.Printf("You have activated your subscription for %s\n", course.CourseName)
}

func (s *Subscription) Dectivate(course *observer.Course) {
	delete(s.subscriptionManager, course)
	fmt.Printf("You have deactivated your subscription for %s\n", course.CourseName)
}

type SubscribeCommand struct {
	subscription *Subscription
	course       *observer.Course
}

func (s *SubscribeCommand) Execute() {
	s.subscription.Activate(s.course)
}

type CancelSubscriptionCommand struct {
	subscription *Subscription
	course       *observer.Course
}

func (c *CancelSubscriptionCommand) Execute() {
	c.subscription.Dectivate(c.course)
}

func main() {
	subscription := &Subscription{
		subscriptionManager: make(map[*observer.Course]bool),
	}

	course := &observer.Course{CourseName: "Math 101"}

	subscribeCommand := &SubscribeCommand{
		subscription: subscription,
		course:       course,
	}

	cancelCommand := &CancelSubscriptionCommand{
		subscription: subscription,
		course:       course,
	}


	subscribeCommand.Execute()

	cancelCommand.Execute()
}