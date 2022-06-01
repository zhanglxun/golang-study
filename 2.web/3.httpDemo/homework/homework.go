package main

import (
	"errors"
	"fmt"
)

func main() {
	values := []interface{}{0, 1, 2, 3, 5}

	newValues, err := Add(values, 4, 5)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	// 0,1,2,3,4,5
	for i, value := range newValues {
		fmt.Printf("下标：%d ,值：%d \n", i, value)
	}
}

func Add(values []interface{}, val interface{}, index int) ([]interface{}, error) {
	if index < 0 || index > len(values) {
		return nil, errors.New("index 非法")
	}
	res := []interface{}{}
	for i := 0; i < index; i++ {
		v := values[i]
		res = append(res, v)
	}
	res = append(res, val)

	for i := index; i < len(values); i++ {
		v := values[i]
		res = append(res, v)
	}

	return res, nil
}

func Delete(values []interface{}, index int) []interface{} {
	if index < 0 || index > len(values) {
		return values
	}
	res := []interface{}{}
	for i := 0; i < len(values); i++ {
		if index == i {
			continue
		}
		v := values[i]
		res = append(res, v)
	}
	return res
}
