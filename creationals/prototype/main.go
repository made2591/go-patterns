package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	address *Address
	name    string
	surname string
}

type Address struct {
	street string
	number int
}

func (p *Person) toString() string {
	return fmt.Sprintf("name: %s, surname: %s, address: %s, %d", p.name, p.surname, p.address.street, p.address.number)
}

func InternalClone(a interface{}, b interface{}) (err error) {
	return nil
}

func Clone(a interface{}, b interface{}) (err error) {
	switch va := reflect.ValueOf(a); va.Kind() {
	case reflect.String:
		fmt.Println(va.String())
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		vb := reflect.ValueOf(b)
		if vb.CanSet() {
			vb.Set(va)
		}
	case reflect.Array, reflect.Chan:
		vb := reflect.ValueOf(b)
		if vb.CanSet() {
			vb.Set(va)
		}
	default:
		fmt.Printf("unhandled kind %s", va.Kind())
	}
	return err
}

func main() {
	address := &Address{street: "MyStreet", number: 42}
	person := &Person{name: "Matteo", surname: "Madeddu", address: address}
	fmt.Println(person.toString())
	fmt.Println(reflect.TypeOf(*person).NumField())

	a := 0
	var b int
	err := Clone(a, b)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)

}
