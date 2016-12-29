package cacahuetas

import "fmt"

type Users map[string]int
type Restrictions map[string]string

type Couple struct {
	Giver, Receiver string
}

func (c Couple) String() string {
	return fmt.Sprintf("%v:%v", c.Giver, c.Receiver)
}

var Cacahuetas []Couple
var restrictions Restrictions
var givers Users
var originals Users

func Init(u Users, r Restrictions) {
	restrictions = r
	originals = u
}

func GenerateCouples() (couples []Couple) {
	for {
		var err error
		couples, err = generateCouples()
		if err != nil {
			fmt.Printf("error generating couples: %v\n", err)
		} else {
			break
		}

	}
	return
}

func generateCouples() (couples []Couple, err error) {
	couples = []Couple{}
	givers = copy(originals)
	receivers := copy(originals)

	for {
		if len(givers) == 0 {
			break
		}
		couple, err := getCouple(givers, receivers)
		if err != nil {
			return nil, err
		}
		couples = append(couples, couple)
	}

	return
}

func copy(originalMap Users) (newMap Users) {
	newMap = make(Users, len(originalMap))
	for k, v := range originalMap {
		newMap[k] = v
	}

	return
}

func getCouple(givers, receivers Users) (couple Couple, err error) {
	couple = Couple{}
	couple.Giver = randomUser(givers)
	receiver, err := getReceiver(couple.Giver, receivers)
	if err != nil {
		return Couple{}, err
	}
	couple.Receiver = receiver

	delete(givers, couple.Giver)
	delete(receivers, couple.Receiver)
	return
}

func getReceiver(giver string, receivers Users) (receiver string, err error) {
	for {
		receiver = randomUser(receivers)
		if isRestricted(giver, receiver) && len(receivers) <= 2 {
			return "", fmt.Errorf("last %v - %v couple is restricted", giver, receiver)
		}
		if giver != receiver && !isRestricted(giver, receiver) {
			break
		}
	}
	return
}

func isRestricted(giver, receiver string) (restricted bool) {
	return giver == receiver || restrictions[receiver] == giver || restrictions[giver] == receiver
}

func randomUser(users Users) (user string) {
	for user = range users {
		break
	}
	return
}
