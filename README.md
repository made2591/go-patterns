# Go-patterns
This repo contains implementation of the GoF patterns in GoLang.

- [Go-patterns](#go-patterns)
  - [Creationals](#creationals)
    - [Abstract Factory](#abstract-factory)
      - [Example](#example)
      - [Implementation](#implementation)
    - [Builder](#builder)
      - [Example](#example-1)
      - [Implementation](#implementation-1)
    - [Factory](#factory)
      - [Example](#example-2)
      - [Implementation](#implementation-2)
    - [Prototype](#prototype)
      - [Example](#example-3)
      - [Implementation](#implementation-3)
    - [Singleton](#singleton)
      - [Example](#example-4)
      - [Implementation](#implementation-4)
  - [Structurals](#structurals)
    - [Adapter](#adapter)
      - [Example](#example-5)
      - [Implementation](#implementation-5)
    - [Bridge](#bridge)
      - [Example](#example-6)
      - [Implementation](#implementation-6)
    - [Composite](#composite)
      - [Example](#example-7)
      - [Implementation](#implementation-7)
    - [Decorator](#decorator)
      - [Example](#example-8)
      - [Implementation](#implementation-8)
    - [Decorator](#decorator-1)
      - [Example](#example-9)
      - [Implementation](#implementation-9)
    - [Facade](#facade)
      - [Example](#example-10)
      - [Implementation](#implementation-10)
    - [Flyweight](#flyweight)
      - [Example](#example-11)
      - [Implementation](#implementation-11)
    - [Proxy](#proxy)
      - [Example](#example-12)
      - [Implementation](#implementation-12)
  - [Behaviourals](#behaviourals)
    - [Chain of Responsibility](#chain-of-responsibility)
      - [Example](#example-13)
      - [Implementation](#implementation-13)
    - [Chain of Responsibility](#chain-of-responsibility-1)
      - [Example](#example-14)
      - [Implementation](#implementation-14)
    - [Command](#command)
      - [Example](#example-15)
      - [Implementation](#implementation-15)
    - [Interpreter](#interpreter)
      - [Example](#example-16)
      - [Implementation](#implementation-16)
    - [Iterator](#iterator)
      - [Example](#example-17)
      - [Implementation](#implementation-17)
    - [Mediator](#mediator)
      - [Example](#example-18)
      - [Implementation](#implementation-18)
    - [Memento](#memento)
      - [Example](#example-19)
      - [Implementation](#implementation-19)
    - [Observer](#observer)
      - [Example](#example-20)
      - [Implementation](#implementation-20)
    - [State](#state)
      - [Example](#example-21)
      - [Implementation](#implementation-21)
    - [Strategy](#strategy)
      - [Example](#example-22)
      - [Implementation](#implementation-22)
    - [Template Method](#template-method)
      - [Example](#example-23)
      - [Implementation](#implementation-23)
    - [Visitor](#visitor)
      - [Example](#example-24)
      - [Implementation](#implementation-24)

## Creationals
Creational patterns provide ways to instantiate single objects or groups of related objects. There are five such patterns.

### Abstract Factory
The `abstract factory pattern` is a design pattern that allows for the creation of groups of related objects without the requirement of specifying the exact concrete classes that will be used. One of a number of factory classes generates the object sets for you.

The `abstract factory pattern` is used to control class instantiation. More formally, it's usufull to provide a client with a set of related or dependant objects. The *family of objects* created by the factory is determined *at run-time* - at least, not during design - according to the selection of concrete factory class.

#### Example
As an example, consider a program that provide `Vehicles` and `Licenses` to drive them. The system delivers different types of vehicles. Let's suppose there are at least two different types of them: the `car` is a standard car and the respective license to drive the car is of kind `B`. The second type of vehicle provided is the `bus` is a more difficult vehicle to drive and require the driver to own a driver license of type `C`.

In the scenario we described above, there are two types of object required, a `Vehicle` object and a `License` one. With the `abstract factory pattern` we can use two factories to generate these related objects. The `NewCarFactory` may produce vehicles objects of type `car` and licenses objects of type `B`; the `NewBusFactory` creates `bus` vehicles and respective `C` licenses.

#### Implementation
Inside the [creationals/abstract-factory](https://github.com/made2591/go-patterns/tree/master/creationals/abstract-factory) folder there's an implementation of the schema described. In order, we have:

- The [Consumer](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/consumer/consumer.go) package: this represents who will *consume* our factory methods. The consumer class is also created with an interface: this specify two `Get` methods to retrieve a `Vehicle` and a `License` object. The first thing to notice is that the consumer interface doesn't specify which kind of these objects are provided. Actually, the consumer *doesn't know at all* - and it doesn't have to! - about the specific kind of object the factory methods can provide. Since it only knows about `Vechile`s and `License`s, there is a generic interface implemented by all of the `Vehicle` and `License` structs to accomplish this status. In the implementation, the `consumer` struct has two private fields that hold instances of the *abstract* vechiles and licenses. The important thing is that the actual types can remain unknown, as they will only be accessed via their base class interfaces. Thus, when a `Consumer` is instantiated, a concrete factory object is passed to the constructor and used to populate the private fields.
- The [Factory](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/factory/factory.go). This is an interface for the concrete `Factory` structs that will generate new sets of related objects with respective methods. Note that a specific `Factory` method is included for each type of object that will be instantiated. Thus, in this file we can find the interface that define which method a `Factory` struct should implement (the two real constructor `NewVechile` and `NewLicense`), and **two different implementation** of this interface for `Car` and `Bus`. Note that there could be any number of concrete factories depending upon the requirements of the software and this will not change anything of the `Consumer` side -> that it will only be able to create what any implementation of the `Factory` provide.
- The [License](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/license/license.go) and the [Vehicle](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/vehicle/vehicle.go) real object that actually implement an interface as well plus an implementation of some constructor-specific method for - respectively - specific kind of vechile and licences.
- The [main](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/main.go). In the main we can see that by only providing a different `Factory` object, a `Consumer` is able to get the objects it deside without changing anything else and - most important - without knowing anything of the object behind. It can still interact with them through the methods exposed by respective interfaces.

To the test the `abstract factory pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd creationals/abstract-factory
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

### Builder
TODO

#### Example
TODO

#### Implementation
TODO

### Factory
TODO

#### Example
TODO

#### Implementation
TODO

### Prototype
TODO

#### Example
TODO

#### Implementation
TODO

### Singleton
TODO

#### Example
TODO

#### Implementation
TODO

## Structurals
TODO

### Adapter
TODO

#### Example
TODO

#### Implementation
TODO

### Bridge
TODO

#### Example
TODO

#### Implementation
TODO

### Composite
TODO

#### Example
TODO

#### Implementation
TODO

### Decorator
TODO

#### Example
TODO

#### Implementation
TODO

### Decorator
TODO

#### Example
TODO

#### Implementation
TODO

### Facade
TODO

#### Example
TODO

#### Implementation
TODO

### Flyweight
TODO

#### Example
TODO

#### Implementation
TODO

### Proxy
TODO

#### Example
TODO

#### Implementation
TODO

## Behaviourals
Todo

### Chain of Responsibility
TODO

#### Example
TODO

#### Implementation
TODO

### Chain of Responsibility
TODO

#### Example
TODO

#### Implementation
TODO

### Command
TODO

#### Example
TODO

#### Implementation
TODO

### Interpreter
TODO

#### Example
TODO

#### Implementation
TODO

### Iterator
TODO

#### Example
TODO

#### Implementation
TODO

### Mediator
TODO

#### Example
TODO

#### Implementation
TODO

### Memento
TODO

#### Example
TODO

#### Implementation
TODO

### Observer
TODO

#### Example
TODO

#### Implementation
TODO

### State
TODO

#### Example
TODO

#### Implementation
TODO

### Strategy
TODO

#### Example
TODO

#### Implementation
TODO

### Template Method
TODO

#### Example
TODO

#### Implementation
TODO

### Visitor
TODO

#### Example
TODO

#### Implementation
TODO
