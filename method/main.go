package main

import (
	"fmt"
)

func main() {

	kiran := User{"kiran", "kiran@ki.com", true, 23}
	fmt.Println(kiran)
	fmt.Printf("kiran details are : %+v\n", kiran)
	//fmt.Printf("name is %v and email is %v \n", kiran.Name, kiran.Email)
	kiran.GetStatus()
	kiran.NewMail()
	fmt.Printf("name is %v and email is %v \n", kiran.Name, kiran.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println("is user active :  ", u.Status)

}
func (u User) NewMail() {
	u.Email = "test@t.com"
	fmt.Println("new email of user is:", u.Email)
}
