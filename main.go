package main

import "fmt"

type Users map[string]int

type Couple struct {
	Giver, Receiver string
}

func main() {
	givers := Users{
		"Julien":   0,
		"Céline":   1,
		"Ludivine": 2,
		"Jérémie":  3,
		"Anthony":  4,
		"Elisa":    5,
	}

	couples := []Couple{}

	receivers := copy(givers)

	for {
		if len(givers) == 0 {
			break
		}
		couple := getCouple(givers, receivers)
		couples = append(couples, couple)
	}

	fmt.Printf("couples: %v", couples)

}

func copy(originalMap Users) (newMap Users) {
	newMap = make(Users, len(originalMap))
	for k, v := range originalMap {
		newMap[k] = v
	}

	return
}

func getCouple(givers, receivers Users) Couple {
	couple := Couple{}
	couple.Giver = randomUser(givers)
	couple.Receiver = getReceiver(couple.Giver, receivers)
	delete(givers, couple.Giver)
	delete(receivers, couple.Receiver)
	return couple
}

func getReceiver(giver string, receivers Users) (receiver string) {
	for {
		receiver = randomUser(receivers)
		if giver != receiver {
			break
		}
	}
	return
}

func randomUser(users Users) (user string) {
	for user = range users {
		break
	}
	return
}
