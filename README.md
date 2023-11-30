# Design-patterns-GOLANG

## Let's discuss the representation of design patterns in GO

### Abstract Factory

- **Purpose**: Provides an interface for creating families of related or dependent objects without specifying their concrete classes.
- **Use Case**: Creating sets of related objects, like UI components for different operating systems, without specifying concrete classes.

### Adapter

- **Purpose**: Allows incompatible interfaces to work together by wrapping itself around an object and exposing a standard interface to interact with that object.
- **Use Case**: Integrating new functionality or systems without changing existing code, like using a new API with an old system.

### Command

- **Purpose**: Encapsulates a request as an object, thereby allowing for parameterization of clients with different requests, queue or log requests, and support undoable operations.
- **Use Case**: Implementing things like menu actions, where actions can be triggered in various ways and undone.

### Decorator

- **Purpose**: Dynamically adds additional responsibilities to an object without modifying its structure.
- **Use Case**: Adding new features to objects without altering the class, such as GUI component enhancements.

### Facade

- **Purpose**: Provides a simplified interface to a complex subsystem.
- **Use Case**: Simplifying the interaction with a complex library or API, like a set of network operations.

### Factory Method

- **Purpose**: Defines an interface for creating an object but lets subclasses alter the type of objects that will be created.
- **Use Case**: Frameworks where library code needs to create objects but doesn't know the specific types.


### Observer

- **Purpose**: Allows a set of observer objects to watch and react to events or changes in a subject object.
- **Use Case**: Implementing event handling systems, like UI event listeners.

## Singleton

- **Purpose**: Ensures a class has only one instance and provides a global point of access to it.
- **Use Case**: Managing a shared resource, such as a database connection.

## Strategy

- **Purpose**: Defines a family of algorithms, encapsulates each one, and makes them interchangeable, allowing the algorithm to vary independently from clients that use it.
- **Use Case**: Supporting multiple variations of an algorithm, such as different sorting or compression algorithms.

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
