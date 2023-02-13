package importTest

import (
	"fmt"
	"strconv"
)

type Admin struct {
	Name string
}

func (a Admin) AddItem(db *Database) string {
	var It = Item{
		ID:     "",
		Name:   "",
		Rating: 0,
		Price:  0,
	}
	fmt.Println("Enter the ID of the item: ")
	fmt.Scan(&It.ID)
	fmt.Println("Enter the name of the item: ")
	fmt.Scan(&It.Name)
	fmt.Println("Enter the price of the item: ")
	fmt.Scan(&It.Price)
	db.Items = append(db.Items, It)
	return "Item successfully added"
}

func (a Admin) ShowItem(db *Database) {
	for i := 0; i < len(db.Items); i++ {
		fmt.Println(strconv.Itoa(i+1) + ": " + db.Items[i].ID + ", " + db.Items[i].Name + ", " + fmt.Sprintf("%v", db.Items[i].Price) + ", " + fmt.Sprintf("%v", db.Items[i].Rating))
	}
}

func (a Admin) SearchItem(db *Database) bool {
	var Name string
	fmt.Scan(Name)
	for i := 0; i < len(db.Items); i++ {
		if db.Items[i].Name == Name {
			fmt.Println("The item is found, and it's at index " + strconv.Itoa(i+1))
			return true
		}
	}
	fmt.Println("No such item found")
	return false
}

func (a Admin) FilterItems(db *Database) {

}
