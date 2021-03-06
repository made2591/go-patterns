# Go-patterns
This repo contains implementation of the GoF patterns in GoLang.

[![Build Status](https://travis-ci.org/made2591/go-patterns.svg?branch=master)](https://travis-ci.org/made2591/go-patterns)
[![codebeat badge](https://codebeat.co/badges/53c8e4e9-5bed-485f-9a18-570bce089e1b)](https://codebeat.co/projects/github-com-made2591-go-patterns-master)
[![Coverage Status](https://coveralls.io/repos/github/made2591/go-patterns/badge.svg?branch=master)](https://coveralls.io/github/made2591/go-patterns?branch=master)
[![License](https://img.shields.io/github/license/made2591/go-patterns.svg)](https://opensource.org/licenses/MIT)

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

- The [consumer](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/consumer/consumer.go) package: this represents who will *consume* our factory methods. The consumer class is also created with an interface: this specify two `Get` methods to retrieve a `Vehicle` and a `License` object. The first thing to notice is that the consumer interface doesn't specify which kind of these objects are provided. Actually, the consumer *doesn't know at all* - and it doesn't have to! - about the specific kind of object the factory methods can provide. Since it only knows about `Vechile`s and `License`s, there is a generic interface implemented by all of the `Vehicle` and `License` structs to accomplish this status. In the implementation, the `consumer` struct has two private fields that hold instances of the *abstract* vechiles and licenses. The important thing is that the actual types can remain unknown, as they will only be accessed via their base class interfaces. Thus, when a `Consumer` is instantiated, a concrete factory object is passed to the constructor and used to populate the private fields.
- The [factory](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/factory/factory.go). This is an interface for the concrete `Factory` structs that will generate new sets of related objects with respective methods. Note that a specific `Factory` method is included for each type of object that will be instantiated. Thus, in this file we can find the interface that define which method a `Factory` struct should implement (the two real constructor `NewVechile` and `NewLicense`), and **two different implementation** of this interface for `Car` and `Bus`. Note that there could be any number of concrete factories depending upon the requirements of the software and this will not change anything of the `Consumer` side -> that it will only be able to create what any implementation of the `Factory` provide.
- The [license](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/license/license.go) and the [Vehicle](https://github.com/made2591/go-patterns/blob/master/creationals/abstract-factory/vehicle/vehicle.go) real object that actually implement an interface as well plus an implementation of some constructor-specific method for - respectively - specific kind of vechile and licences.
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
The `builder pattern` is used to control class instantiation. More precisely, it's define to create complex objects with constituent parts that must be created in the same order or using a specific algorithm. An external class, known as the `director`, controls the construction algorithm.

#### Example
Consider a vehicle factory system. It may be that each time a car vehicle is prepared, it can include one or more wheels, one control system, one engine and some additional security feature. The employees working in the factory select the types of action required and the `builder pattern` is used to determine exactly what goes into each vehicle.

Using this model, when the employee selects a "Volkswagen Polo", the building system determines that the vehicle will include a four wheels, one steering wheel, one car engine and two standard airbags. If the customer orders a "Volkswagen Golf", they receive a car made by four special wheels, one steering wheel, one car engine and four special airbags. In each case, the constituent parts are the same and are created in the same order. The final product is quite different.

#### Implementation
Inside the [creationals/builder](https://github.com/made2591/go-patterns/tree/master/creationals/builder) folder there's an implementation of the schema described. In order, we have:

- The [director](https://github.com/made2591/go-patterns/blob/master/creationals/builder/director/director.go) package: this represents who will *direct* our building process, by controlling the algorithm that generates the final product object, or the vehicle in this case. In the normal situation, a `director` object is instantiated and its construct method - in this example, `MakeVehicle` method - is called by the user. The method includes a parameter to capture the specific concrete builder object that is to be used to generate the product. The director then calls methods of the concrete builder (see them in a while) in the correct order to generate the product object. On completion of the process, the getProduct method - in this example, `GetVehicle` method - of the builder object can be used to return the product.
- The [builder](https://github.com/made2591/go-patterns/blob/master/creationals/builder/builder/builder.go) package: this package define both an interface called `Builder` and two different implementation of *different product Builder*s. It defines all of the steps that must be taken in order to correctly create a product. Each step is generally defined in the interface as the actual functionality of the builder is carried out in the concrete implementation. As already said, the `GetVehicle` method is used to return the final product. The concrectes builders that implement this interface are thew `PoloBuilder` and the `GolfBuilder` and both contain the functionality to create a particular complex product.
- The [vehicle](https://github.com/made2591/go-patterns/blob/master/creationals/builder/vehicle/vehicle.go) package: it contains the type of the complex object that is to be generated by the builder pattern.
- The [main](https://github.com/made2591/go-patterns/blob/master/creationals/builder/main.go). In the main we can see that by only providing a different `Builder` object, the `Director` is able to order the production of the objects it desires without changing anything else and - most important - without knowing anything of the object behind. It can still interact with them through the methods exposed by respective interfaces.

To the test the `builder pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd creationals/builder
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

### Factory
The `factory pattern` is used to control class instantiation. The `factory pattern` is used to replace class constructors, abstracting the process of object generation so that the type of the object instantiated can be determined at run-time.

#### Example
A good example of the use of the factory pattern is when creating a type of vehicle that will be selected by the end-user using a graphical interface. In this case, an *abstract* struct named `Vehicle` - like the one defined for other creational pattern samples - may be created that defines the base functionality of all vehicle. Many concrete structs may be created, perhaps `PoloVehicle`, `GolfVehicle`, etc, each with specific functionality for a different type of vehicle. At run-time, a further struct, perhaps named `VehicleFactory`, generates objects of the correct concrete `Vehicle` struct based upon a parameter passed to the factory method.

#### Implementation
- The [factory](https://github.com/made2591/go-patterns/blob/master/creationals/factory/factory/factory.go) package: this is an interface for the concrete factory structs that will actually generate new objects. This interface should contain the signature for the factory method. In simple situations the factory method may be implemented directly, rather than being declared as a required method to implement through an interface.
- The [vehicle](https://github.com/made2591/go-patterns/blob/master/creationals/factory/vehicle/vehicle.go) package: this package define an interface called `Vehicle`. The two different implementation of *different vehicle*s are defined in the [bmw](https://github.com/made2591/go-patterns/blob/master/creationals/factory/bmw/bmw.go) and [volkswagen](https://github.com/made2591/go-patterns/blob/master/creationals/factory/volkswagen/volkswagen.go) packages. The `Vehicle` interface defines a couple of methods to set and retrieve specific model type.
- The [bmw](https://github.com/made2591/go-patterns/blob/master/creationals/factory/bmw/bmw.go)/[volkswagen](https://github.com/made2591/go-patterns/blob/master/creationals/factory/volkswagen/volkswagen.go) packages: they both contain an implementation of a specific `Vechile` and `Factory` constructor. From both these packages, the only methods that are visible from outside are the `New[Bmw/Volkswagen]Factory`, the respective `CreateVehicle` that can be called over the factory and the methods exposed by the `Vehicle` interface to respect the contract, or `GetModel`, `SetModel` and `PrintDetails`.
- The [main](https://github.com/made2591/go-patterns/blob/master/creationals/factory/main.go). In the main we can see that it is possible to create different model of the same vehicle throught respective factory, by passing one parameter to respective `CreateVehicle` method called over the factory implementation. This parameter could have been selected by an external user at run-time. When the object's underlying type is outputted, we can see that the correct car model was selected.

To the test the `factory pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd creationals/factory
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

### Prototype
The `prototype pattern` is used to instantiate a class by copying, or cloning, the properties of an existing object. The new object is an exact copy of the prototype but permits modification without altering the original. This practise is particularly useful when the construction of a brand new object, using the new operator, is inefficient.

When an object is cloned the new object is either a *shallow* or *deep copy*. A shallow copy duplicates all of the object's properties. If any property contains a reference type, the reference is copied. This means that changes to the referenced object are visible in both the clone and the original object. A deep copy clones the main object and all child objects. Any properties of reference types are also cloned, giving a truly independent copy. The prototype pattern usually generates deep copies, though this is dependant upon the situation.

#### Example
WIP

#### Implementation
WIP

### Singleton
The `singleton` is used to control class instantiation. The pattern ensures that only one object of a particular class is ever created. All further references to objects of the singleton class refer to the same underlying instance.

The singleton pattern is useful when a single, global point of access to a limited resource is required. It is more appropriate than creating a global variable as this may be copied, leading to multiple access points and the risk that the duplicates become out of step with the original.

An example of the use of a singleton class is a logger instance. Since logging is not something that changes the behaviour of the application, creating multiple instances of the logging object is generally not required. By forcing a single, global instance to be used, only one underlying writer is created.

#### Example
Look at the `singleton` package and level instance.

#### Implementation
- The [logger](https://github.com/made2591/go-patterns/blob/master/creationals/singleton/logger/logger.go) package: this is an interface for the concrete logger struct that is implemented throught the use of the `sync`. The pattern used leverages the `Once` struct defined in the `sync` package: this struct simply contains a mutex:

```
// Once is an object that will perform exactly one action.
type Once struct {
	m    Mutex
	done uint32
}
```

and force a function `func()` to be executed only once thanks to the `Do(f func() {...})` method. Let's have a look at the method:

```
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

As stated in the code, this means that given `var once Once` if `once.Do(f)` is called multiple times, only the first call will invoke `f`, even if `f` has a different value in each invocation. A new instance of Once is required for each function to execute.

- The [main](https://github.com/made2591/go-patterns/blob/master/creationals/singleton/main.go). In the main, we can see that it is possible to create different many logger instances, but actually always the same is returned.

To the test the `singleton pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd creationals/singleton
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

## Structurals
The second type of design pattern is the structural pattern. Structural patterns provide a manner to define relationships between classes or objects.

### Adapter
The `adapter pattern` is used to provide a link between two otherwise incompatible types by wrapping the "adaptee" with a class that supports the interface required by the client.

#### Example
Imagine an application for a vehicle reseller that displays a list of vechiles. The list is intented to be provided as a CSV string - a column based comma separated string of values, with a new row for each vehicle. It may be that the internal vehicle system includes a method that permit the retrieval of the vehicle available in the catalogue, including their model type, as an array of Vehicles. The solution interact with the vehicle catalogue throught an interface that include a method to get the data, but this method is not compatible with the one provided by the internal vehicle system (in this case, because a specific format is required but it could be due to some different presentation logic behind).

By following the `adapter` pattern, I created a new struct to be an adapter. This struct provide the reseller desired method (by implementing the reseller interface) and - at the same time - hold an object of the type required by the application used by the reseller. When the main application ask for the list of vehicles, this request is passed and processed by the internal vehicle system. The response from the latter will then be translated to a format that can be used by the reseller thanks to the adapter: the latter both exposes a method implemented by the interface of the reseller and keep an instance of the internal vehicle system.

#### Implementation
- The [vehicle](https://github.com/made2591/go-patterns/blob/master/structurals/adapter/vehicle/vehicle.go) package: this is an interface to provide different vehicle recycled from the implementation given in the creationals pattern `abstrac-factory`. I only added a method called `GetAvailableVehicle()` to return the list of vehicles available. Thus, this package play the role of the internal vehicle system This is the dummy use case build to act as the interval vehicle system mentioned before;
- The [reseller](https://github.com/made2591/go-patterns/blob/master/structurals/adapter/reseller/reseller.go) package: this is the interface implemented by the adapter that will expose the method expected by the client
- The [adapter](https://github.com/made2591/go-patterns/blob/master/structurals/adapter/adapater/adapater.go) package: this package provides the adapter that will handle the internal vehicle system properly and translate the requests in the desidered format transparently, by keeping internal system.
- The [main](https://github.com/made2591/go-patterns/blob/master/structurals/adapter/main.go) package: it represents the console application. As you can see, we first create a `InternalVehicleSystem` object. We then create a new `Adapter` that holds a reference to the `InternalVehicleSystem` object. Finally, we can instantiate the `Reseller` object, passing the adapter to use as its source of vehicles list information. The call to `GetAvailableVehicles()` is intercepted by the `Adapter` so that the data can be retrieved from the incompatible `InternalVehicleSystem` object and correctly formatted for output.


To the test the `adapter pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd structurals/adapter
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

### Bridge
The `bridge pattern` is a a structural pattern as it defines a manner for creating relationships between classes. The pattern is used to separate the abstract elements of a class from the implementation details. For example, the abstract elements may contain logic depending on a specific operating system. If we imagine a software to dispatch installer utilities, someone would be able to create it without any knowledge of the implementation details of interoperability with the operating system. The pattern provides the means to replace the implementation details without modifying the abstraction. This permits, for example, changing operating systems, or databases, etc. with no impact to the logic of setup.

#### Example
In the example build, there's the need to provide a generic installer utility to accomplish some setup steps across multiple platform to which the software managed has to been distributed. An interface that represents the installer could or could not include some properties to hold the details of the installation process. Several struct that implements this interface could then be included to permit the complention of the steps required depending on the operating system. The implementation of the setup to be used could then be selected according to the underlying architecture automatically - here, is simulated throught manual imposition.

#### Implementation
- The [abstractsystems](https://github.com/made2591/go-patterns/blob/master/structurals/bridge/abstractsystems/system.go) package: this represents the interface base for the implementation of different setup method. This interface defines a single method signature that will be used to launch a setup. The method accept a parameter just as an example;
- The [systems](https://github.com/made2591/go-patterns/blob/master/structurals/bridge/systems/systems.go) package: this represents two initial concrete setup implementations for two different systems. As we are using the bridge pattern, further system implementations could be added as required in the future.
- The [abstractinstaller](https://github.com/made2591/go-patterns/blob/master/structurals/bridge/abstractinstaller/installer.go) package: this package describes an installer with the required parameter (dummy, for example). The installe holds also a reference to an implementation object - an abstract systems. Furthermore, **the installer includes also a `Setup` method that calls the `Installer` method of the referenced `System` implementation**.
- The [main](https://github.com/made2591/go-patterns/blob/master/structurals/bridge/main.go) package: it represents the software that pack the install steps. The main code creates an installer and launch the setup using each type of system setups.


To the test the `bridge pattern`:

```
git clone https://github.com/made2591/go-patterns/
cd structurals/bridge
go run main.go
```

To run the tests, at the same folder level run:

```
go test -v ./...
```

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
