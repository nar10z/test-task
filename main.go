package main

import (
	"./person"
	"database/sql"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

const insertActionPrep = "INSERT INTO `actions_users` (`user_id`, `action`) VALUES(?,?)"

func main() {

	var arPersons []person.Person

	minAge := 1
	maxAge := 85

	for i := 0; i < 100; i++ {
		newPerson := new(person.Person)

		genderInt := rand.Intn(2)
		age := rand.Intn(maxAge-minAge) + minAge
		var genderRandomData int
		if genderInt == 0 {
			newPerson.Gender = string('M')
			genderRandomData = randomdata.Male
		} else {
			newPerson.Gender = string('W')
			genderRandomData = randomdata.Male
		}
		name := randomdata.FirstName(genderRandomData)
		lastName := randomdata.LastName()

		newPerson.Age = age
		newPerson.Name = name
		newPerson.LastName = lastName

		arPersons = append(arPersons, *newPerson)
	}

	db, err := sql.Open("mysql", "root:root@tcp(:3307)/tz?charset=utf8")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	addUserPrepare, err := db.Prepare("INSERT INTO `users` (`name`, `second_name`, `age`, `gender`) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer addUserPrepare.Close()

	var uArUsers = &arPersons

	for key, user := range *uArUsers {
		res, err := addUserPrepare.Exec(user.Name, user.LastName, user.Age, user.Gender)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		id, err := res.LastInsertId()

		arPersons[key].Id = int(id)
	}

	addActUser, err := db.Prepare(insertActionPrep)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer addActUser.Close()

	for _, user := range arPersons {
		actStr := user.Actions()

		fmt.Println(user.Id, " выполнил \"", actStr, "\"", user.Age, "лет")
		_, err := addActUser.Exec(user.Id, actStr)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}
	//fmt.Println(arPersons)

}
