package main

import "fmt"

type user struct {
	userName string
	password string
}

func userData() {
	uName := "Enter the user name: "
	pWord := "Enter the password: "
	name := getValue(uName)
	pass := getValue(pWord)
	constructor(name, pass)
}

func constructor(name, pass string) {
	var data user
	data.userName = name
	data.password = pass
	display(data)
}
func getValue(str string) string {
	fmt.Println(str)
	var dummy string
	fmt.Scan(&dummy)
	return dummy
}
func display(data user) {
	fmt.Printf("\nUserName:%s\nPassword:%s", data.userName, data.password)
	mapData()
}
func mapData() {
	mapData := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	fmt.Print("\n", len(mapData))
}
