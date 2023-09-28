# Design-patterns-GOLANG

## Let's discuss the representation of Observer and Strategy pattern in GO

### Observer
This pattern is implemented quite general. The main idea is about a company which provides a different courses for learning. The purpose of observer pattern in this code is that notify each time when the new course is added.

### Strategy
The initial problem is that to enroll to course the client must pass the examination in order to identify the level. There are alot of courses like language(spanish, kazakh, russian, english, etc), science(math, biology, chemistry, etc.) or programming(web, c, c#, GOLANG, etc.)
Some of them needs specific types of examinations. So  this code has tries to solve this issue by using Strategy pattern.

## Usage
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
```
bash go run .
```
