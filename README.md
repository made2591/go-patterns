# go-patterns

This repo contains implementation of the GoF patterns in GoLang.
- [go-patterns](#go-patterns)
  - [Creationals](#creationals)
    - [Abstract Factory](#abstract-factory)
      - [Example](#example)
      - [Implementation](#implementation)

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
