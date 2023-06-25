package model

import (
	"encoding/json"
)

type People struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

func (people *People) String() string {
	bytes, _ := json.Marshal(*people)
	return string(bytes)
}

func NewPeople(name string, sex string, age int) *People {

	return &People{
		Name: name,
		Sex:  sex,
		Age:  age,
	}

}
