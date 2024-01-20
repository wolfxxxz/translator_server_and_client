package services

import (
	"bufio"
	"client/internal/domain/requests"
	"client/internal/models"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/agnivade/levenshtein"
)

const (
	tapEmail     = "Your Email"
	yourName     = "Your Name"
	tapLastName  = "Your last Name"
	yourPassword = "Your Password"
	tapRole      = "Your role"
)

func scanLine() (string, error) {
	fmt.Print("       ...")
	in := bufio.NewScanner(os.Stdin)
	if in.Scan() {
		return in.Text(), nil
	}

	if err := in.Err(); err != nil {
		return "", err
	}

	return "", nil
}

func printAll(words []*models.Library) {
	for _, word := range words {
		fmt.Printf(" %v -- %v \n", word.Russian, word.English)
	}
}

func printTime(duration time.Duration) {
	minutes := int(duration.Minutes())
	seconds := int(duration.Seconds()) % 60
	fmt.Printf("Time: %d minutes %d seconds\n", minutes, seconds)
}

func compareStringsLevenshtein(quest, answer string) bool {
	quest = strings.ToLower(quest)
	answer = strings.ToLower(answer)
	mistakes := 1
	if distance := levenshtein.ComputeDistance(quest, answer); distance <= mistakes {
		return true
	}

	return false

}

func ignorSpace(s string) (c string) {
	for _, v := range s {
		if v != ' ' {
			c = c + string(v)
		}
	}

	return
}

func scanUser() *requests.CreateUserRequest {
	var email, name, lastName, password, role string
	fmt.Println(tapEmail)
	fmt.Scan(&email)
	fmt.Println(yourName)
	fmt.Scan(&name)
	fmt.Println(tapLastName)
	fmt.Scan(&lastName)
	fmt.Println(yourPassword)
	fmt.Scan(&password)
	fmt.Println(tapRole)
	fmt.Scan(&role)
	return &requests.CreateUserRequest{
		Email:    email,
		Name:     name,
		LastName: lastName,
		Password: password,
		Role:     role,
	}
}
