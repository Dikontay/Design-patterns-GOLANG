# Design-patterns-GOLANG

## Let's discuss the representation of design patterns in GO

### Abstract Factory

- **Purpose**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
- **Use Case**: Creating sets of related objects, like UI components for different operating systems, without specifying concrete classes.

### Observer
This pattern is implemented quite general. The main idea is about a company which provides a different courses for learning. The purpose of observer pattern in this code is that notify each time when the new course is added.

### Strategy
The initial problem is that to enroll to course the client must pass the examination in order to identify the level. There are a lot of courses like language(spanish, kazakh, russian, english, etc), science(math, biology, chemistry, etc.) or programming(web, c, c#, GOLANG, etc.)
Some of them needs specific types of examinations. So  this code has tried to solve this issue by using Strategy pattern.

## Example of Usage
```bash 
git clone https://github.com/Dikontay/Design-patterns-GOLANG.git
```
```bash 
cd observerPattern/
```
```bash 
go run .
```
```bash 
cd .. && cd strategyPattern/
```
```bash
go run .
```
