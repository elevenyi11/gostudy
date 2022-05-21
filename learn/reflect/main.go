package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	FillNameAndAge()

}

func fillBySetting(st interface{}, settings map[string]interface{}) error {
	// func (v value) elem() value

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil.")
	}
	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

func FillNameAndAge() {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30}
	e := Employee{}
	if err := fillBySetting(&e, settings); err != nil {
		fmt.Errorf("error")
	} else {
		fmt.Printf("Employee:%v \n", e)
	}

	c := new(Customer)

	if err := fillBySetting(c, settings); err != nil {
		fmt.Println("error")
	} else {
		fmt.Printf("Customer:%v\n", c)
	}

}

type Employee struct {
	EmployeeId string
	Name       string `format:"normal"`
	Age        int
}

type Customer struct {
	CustomerId string
	Name       string
	Age        int
}
