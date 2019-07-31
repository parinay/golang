package main

import (
	"fmt"
	"reflect"
)

type order struct {
	orderId    int
	customerID int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q).Name()
	query := fmt.Sprintf("insert into %s values(", t)
	v := reflect.ValueOf(q)
	if reflect.ValueOf(q).Kind() == reflect.Struct {

		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported Type")
				return
			}

		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return
	} else if reflect.ValueOf(q).Kind() == reflect.Int {
		query = fmt.Sprintf("%s%d", query, reflect.ValueOf(q).Int())
		query = fmt.Sprintf("%s)", query)
		return
	} else {
		fmt.Printf("unsupported type - %s value -%f\n", reflect.ValueOf(q).Kind(), reflect.ValueOf(q).Float())
	}
}

func example1Reflections() {
	o := order{
		orderId:    456,
		customerID: 1,
	}

	createQuery(o)

	e := employee{
		name:    "naveen",
		id:      565,
		address: "madras",
		salary:  90000,
		country: "India",
	}
	createQuery(e)

	i := 90
	createQuery(i)

	j := 90.90
	createQuery(j)
}
