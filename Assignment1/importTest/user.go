package importTest

import (
	"fmt"
	"strconv"
)

type User struct {
	Name     string
	Login    string
	Password string
}

func (u User) GiveRating(db *Database) string {
	fmt.Println("Which of these items you want to rate?")
	for i := 0; i < len(db.Items); i++ {
		fmt.Println(strconv.Itoa(i+1) + ": " + db.Items[i].ID + ", " + db.Items[i].Name + ", " + fmt.Sprintf("%v", db.Items[i].Price) + ", " + fmt.Sprintf("%v", db.Items[i].Rating))
	}
	var index int
	fmt.Scan(&index)
	fmt.Println("Which rating you want to give? (1-10)")
	var rating float64
	fmt.Scan(&rating)
	db.Items[index].RatedPeople++
	if rating <= 10 {
		db.Items[index-1].Rating = (db.Items[index-1].Rating + rating) / float64(db.Items[index].RatedPeople)
		return "Successfully gave rating"
	} else {
		return "Your rating crosses the limit"
	}
}
