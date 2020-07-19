package pkg

import (
	"context"
	"fmt"
	"time"
	"json-to-db/config"
	"json-to-db/models"
	"json-to-db/util"
)

func stringCheck(i interface{}) string {
	if v, ok := i.(string); ok {
		return v
	} else {
		return ""
	}
}

func LoadVotes() {
	//init DB
	collection := config.InitDB().Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	usersFromJson := util.ToJsonFromFile(".././example_1.json")
	for _, item := range usersFromJson {
		userId := item["userId"]
		firstname := item["firstName"]
		lastname := item["lastName"]

		//create model
		user := models.User{
			ID:        int(userId.(float64)),
			Firstname: stringCheck(firstname),
			Lastname:  stringCheck(lastname),
		}

		//insert into DB
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			fmt.Println(user)
		}
	}

}
