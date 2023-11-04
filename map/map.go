package main

import "fmt"

func main() {
	user := map[string]string{
		"firstName": "Ridho",
		"lastName":  "Mahendra",
	}

	fmt.Println(user)

	fmt.Println("=====================================")

	var scores = make(map[string]int)
	fmt.Println(scores)

	scores["java"] = 85
	scores["react"] = 87
	scores["kotlin"] = 90

	fmt.Println(scores["java"])

	scores["golang"] = 100
	fmt.Println(scores)

	fmt.Println("=================== ==================")
	delete(scores, "golang")
	fmt.Println(scores)

	fmt.Println("=====================================")
	names := map[int]string{
		1: "ridho",
		2: "mahendra",
		3: "damanik",
	}
	fmt.Println(names[1])

	for _, value := range names {
		fmt.Println(value)
	}
}
