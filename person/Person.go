package person

import (
	"fmt"
	"math/rand"
)

type Person struct {
	Id,
	Age int
	Gender,
	Name,
	LastName string
}

func (p *Person) SetID(id int) {
	p.Id = id
}

func (p *Person) Actions() string {
	randActions := rand.Intn(4)
	actString := ""
	if p.Age < 2 {
		actString = p.sleep()
	} else if p.Age > 2 && p.Age < 18 {
		switch randActions {
		case 0:
			actString = p.sleep()
		case 1:
			actString = p.eat()
		case 2:
			actString = p.walk()
		default:
			actString = p.sleep()

		}
	} else {
		switch randActions {
		case 0:
			actString = p.sleep()
		case 1:
			actString = p.eat()
		case 2:
			actString = p.walk()
		case 3:
			actString = p.work()
		default:
			actString = p.sleep()

		}
	}

	return actString
}

func (p *Person) sleep() string {
	min := 6
	max := 10
	if p.Age == 1 {
		min = 8
		max = 14
	} else if p.Age == 2 {
		min = 7
		max = 13
	} else if p.Age > 2 && p.Age < 18 {
		min = 7
		max = 11
	} else {
		min = 6
		max = 10
	}

	randHourSleep := rand.Intn(max-min) + min
	str := fmt.Sprintf("Поспал %d часов", randHourSleep)
	return str
}

func (p *Person) work() string {
	workTime := rand.Intn(9-7) + 7
	str := fmt.Sprintf("Работал %d часов", workTime)
	return str
}

func (p *Person) walk() string {
	walkTime := rand.Intn(6-1) + 1

	str := fmt.Sprintf("Гулял %d часов", walkTime)
	return str
}

func (p *Person) eat() string {
	rndEat := rand.Intn(3)

	var eat string
	if rndEat == 1 {
		eat = "кашу"
	} else if rndEat == 2 {
		eat = "суп"
	} else {
		eat = "салат"
	}

	str := fmt.Sprintf("Поел %s", eat)
	return str
}
