package main

import (
	"fmt"
	singleton "home/singleton"
	"home/user"
)


func main() {
	ma := make(map[string]int)

	ma["m1"] = 1
	ma["m2"] = 2

	fmt.Println("ma:", ma)

	fmt.Println(ma["m2"])
	delete(ma, "m1")
	fmt.Println(ma)
	_, present := ma["m1"]
	if present {
		fmt.Println("The value is present.")
	} else {
		fmt.Println("No value is present.")
	}
	single := singleton.New()

	single["this"] = "that"

	single2 := singleton.New()

	fmt.Println("This is", single2["this"])
	users := user.Name{}
	fmt.Println(users.FullName(user.Name{ FirstName: "AbdulQuddus", LastName: "Yekeen"}))
	fmt.Println(users.EditFirstName("Adekunle"))
	fmt.Println(users.GetFirstName())
	fmt.Println(users.GetLastName())
}
