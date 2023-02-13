package main

import (
	"fmt"
	"regexp"
	"src/Assignment1/importTest"
)

func main() {
	var DB = importTest.Database{}
	var admin = importTest.Admin{Name: "admin0"}
	//admin.AddItem(&DB)
	//admin.ShowItem(&DB)

	for {
		/*
				1. Add item
				2. Show Items
				3. Give rating
				4. Registration
			 	5. Sign up
			 	6. Exit
		*/
		fmt.Println("What do you want to do? \n 1. Registration \n 2. Sign in \n 3. Exit")
		var command int
		fmt.Scan(&command)
		if command == 1 {
			var us = importTest.User{
				Name:     "",
				Login:    "",
				Password: "",
			}
			fmt.Println("Enter your name: ")
			fmt.Scan(&us.Name)
			fmt.Println("Enter your login: ")
			fmt.Scan(&us.Login)
			fmt.Println("Enter your password: ")
			fmt.Scan(&us.Password)
			DB = *DB.Registration(us)
		}
		if command == 2 {
			var us = importTest.User{
				Name:     "",
				Login:    "",
				Password: "",
			}
			fmt.Println("Enter your name: ")
			fmt.Scan(&us.Name)
			fmt.Println("Enter your login: ")
			fmt.Scan(&us.Login)
			fmt.Println("Enter your password: ")
			fmt.Scan(&us.Password)
			if DB.Auth(us) == true {
				for {
					fmt.Println("What do you want to do? \n 1. Add item \n 2. Show items \n 3. Give rating \n 4. Search items \n 5. Filter items \n 6. Exit")
					var comm int
					fmt.Scan(&comm) // command
					if comm == 1 {
						admin.AddItem(&DB)
					}
					if comm == 2 {
						admin.ShowItem(&DB)
					}
					if comm == 3 {
						us.GiveRating(&DB)
					}
					if comm == 4 {
						fmt.Println("Type string/substring what you want to search")
						var temp string
						fmt.Scan(&temp)
						var validID = regexp.MustCompile(".*" + temp + ".*")
						for i := 0; i < len(DB.Items); i++ {
							if validID.MatchString(DB.Items[i].Name) {
								fmt.Println(DB.Items[i].ID, DB.Items[i].Name, DB.Items[i].Price, DB.Items[i].Rating)
							}
						}
					}
					if comm == 5 {
						var typeFilter string
						fmt.Println("Based on which metric you want to filter? (price, rating)")
						fmt.Scan(&typeFilter)
						if typeFilter == "price" {
							fmt.Println("What should be the limit of the price?")
							var limit float64
							fmt.Scan(&limit)
							for i := 0; i < len(DB.Items); i++ {
								if DB.Items[i].Price <= limit {
									fmt.Println(DB.Items[i].ID, DB.Items[i].Name, DB.Items[i].Price, DB.Items[i].Rating)
								}
							}
						}

						if typeFilter == "rating" {
							fmt.Println("What should be the limit of the rating?")
							var limit float64
							fmt.Scan(&limit)
							for i := 0; i < len(DB.Items); i++ {
								if DB.Items[i].Rating <= limit {
									fmt.Println(DB.Items[i].ID, DB.Items[i].Name, DB.Items[i].Price, DB.Items[i].Rating)
								}
							}
						}
					}
					if comm == 6 {
						break
					}
				}
			}
		}

		if command == 3 {
			break
		}
	}
}
