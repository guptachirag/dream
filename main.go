package main

import (
	"fmt"
	"math"
)

type User struct {
	salary           float64
	age              int
	retirementAge    int
	savingFactor     float64
	retirementCorpus float64
	gainFactor       float64
	wealth           float64
}

func raiseNeededWithInvestment(user *User) (raise float64) {
	for raise = 0.01; ; raise = raise + 0.01 {
		wealth := user.wealth
		for i, time := 0.0, user.age; time < user.retirementAge; i, time = i+1, time+1 {
			wealth += (wealth * user.gain) + (user.savingPercentage * math.Pow(1+raise, i) * user.salary)
		}
		if wealth > user.retirementCorpus {
			return
		}
	}
}

func inputUserDetails(user *User) {
	fmt.Println("Current Salary")
	fmt.Scanf("%f", &user.salary)
	fmt.Println("Age")
	fmt.Scanf("%d", &user.age)
	fmt.Println("Retirement Age")
	fmt.Scanf("%d", &user.retirementAge)
	fmt.Println("Saving Percentage")
	fmt.Scanf("%f", &user.savingFactor)
	fmt.Println("Retirement Corpus")
	fmt.Scanf("%f", &user.retirementCorpus)

	var isInvesting string
	fmt.Println("Do you invest the money[y|n]")
	fmt.Scanf("%s", &isInvesting)
	if isInvesting == "y" || isInvesting == "Y" {
		fmt.Println("Gain on investment")
		fmt.Scanf("%f", &user.gainFactor)
	}
}

func main() {
	user := &User{salary: 6000000, age: 28, retirementAge: 45, savingFactor: 0.3, retirementCorpus: 300000000, gainFactor: 0.08, wealth: 1000000}
	inputUserDetails(user)
	fmt.Println(*user)
	fmt.Printf("Raise need every year %.2f Percentage\n", raiseNeededWithInvestment(user)*100)
}
