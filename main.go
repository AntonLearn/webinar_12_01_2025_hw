package main

import (
	"errors"
	"fmt"
)

var (
	errorLessTwoCharacters     = errors.New("consists of less than two characters")
	errorLessThreeCharacters   = errors.New("consists of less than three characters")
	errorLessEightCharacters   = errors.New("consists of less than eight characters")
	errorMoreHundredCharacters = errors.New("consists of more than hundred characters")
	errorNotDog                = errors.New("does not contain special character '@'")
	errorCyrAlphabetCharacters = errors.New("cyrillic characters cannot be used")
	errorNotDigit              = errors.New("there must be at least one digit")
	errorNotLat                = errors.New("there must be at least one Latin character")
	errorNotSpec               = errors.New("there must be at least one special character")
)

func validateName(validatedString string) error {
	var errorTotal error
	validateStringToRunes := []rune(validatedString)
	counterCyr := 0
	for _, rune := range validateStringToRunes {
		if rune == 1025 || rune == 1105 || (rune >= 1040 && rune <= 1103) {
			counterCyr++
		}
	}
	validateStringToRunesLenght := len(validateStringToRunes)
	if counterCyr > 0 {
		errorTotal = errorCyrAlphabetCharacters
	}
	if validateStringToRunesLenght < 2 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorLessTwoCharacters)
		} else {
			errorTotal = errorLessTwoCharacters
		}
	}
	if errorTotal != nil {
		return errorTotal
	}
	return nil
}

func validateEmail(validatedString string) error {
	var errorTotal error
	validateStringToRunes := []rune(validatedString)
	counterCyr := 0
	flagDog := false
	for _, rune := range validateStringToRunes {
		if rune == 1025 || rune == 1105 || (rune >= 1040 && rune <= 1103) {
			counterCyr++
		} else if rune == 64 {
			flagDog = true
		}
	}
	validateStringToRunesLenght := len(validateStringToRunes)
	if counterCyr > 0 {
		errorTotal = errorCyrAlphabetCharacters
	}
	if validateStringToRunesLenght < 3 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorLessThreeCharacters)
		} else {
			errorTotal = errorLessThreeCharacters
		}
	}
	if validateStringToRunesLenght > 100 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorMoreHundredCharacters)
		} else {
			errorTotal = errorMoreHundredCharacters
		}
	}
	if !flagDog {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorNotDog)
		} else {
			errorTotal = errorNotDog
		}
	}
	if errorTotal != nil {
		return errorTotal
	}
	return nil
}

func validatePassword(validatedString string) error {
	var errorTotal error
	validateStringToRunes := []rune(validatedString)
	counterLat := 0
	counterDig := 0
	counterCyr := 0
	for _, rune := range validateStringToRunes {
		if rune >= 48 && rune <= 57 {
			counterDig++
		} else if (rune >= 65 && rune <= 90) || (rune >= 97 && rune <= 122) {
			counterLat++
		} else if rune == 1025 || rune == 1105 || (rune >= 1040 && rune <= 1103) {
			counterCyr++
		}
	}
	validateStringToRunesLenght := len(validateStringToRunes)
	counterSpec := validateStringToRunesLenght - (counterDig + counterLat + counterCyr)
	if counterCyr > 0 {
		errorTotal = errorCyrAlphabetCharacters
	}
	if validateStringToRunesLenght < 8 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorLessEightCharacters)
		} else {
			errorTotal = errorLessEightCharacters
		}
	}
	if len(validateStringToRunes) > 100 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorMoreHundredCharacters)
		} else {
			errorTotal = errorMoreHundredCharacters
		}
	}
	if counterDig == 0 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorNotDigit)
		} else {
			errorTotal = errorNotDigit
		}
	}
	if counterLat == 0 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorNotLat)
		} else {
			errorTotal = errorNotLat
		}
	}
	if counterSpec == 0 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorNotSpec)
		} else {
			errorTotal = errorNotSpec
		}
	}
	if errorTotal != nil {
		return errorTotal
	}
	return nil
}

func validate(mode, validatedString string) error {
	switch mode {
	case "name":
		return validateName(validatedString)
	case "email":
		return validateEmail(validatedString)
	case "password":
		return validatePassword(validatedString)
	default:
		return nil
	}
}

func main() {
	var modes = []string{"name", "email", "password"}
	var names = []string{"", "a", "ю", "aa", "aю", "a.", "Azw", "Sw_r", "Dx2", "яя", "superuser"}
	var emails = []string{
		"user@mail.com", "", "u@m", "u@", "us", "qwertyuiopasdfghjkl", "йцук", "йц@rtук", "йцукенг@",
		"йцукенгшщзхъфывапро65", "йцукен:43", "qweйцук?0@fr", "12345678@er", "456789@йф", "иasdfg@12",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxc",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcv",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm@wertyuiopasdfghjklzxc",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm@wertyuiopasdfghjklzxcv",
	}
	var passwords = []string{
		"", "qwer", "qwertyui", "qwer1", "qwertyu1", "qwerф", "йцукqwer", "йцукен", "йцукенгш", "йцу1кенг",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxc",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcv",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwerty1uiopasdfghjklzxcvbnmqwertyuiopasdfghjklzx",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwerty1uiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxc",
		"qwer!6ty", "qwer4-#t", "qw45$e%r", "qыw67e^e",
	}
	for _, mode := range modes {
		switch mode {
		case "name":
			fmt.Println("Name validation started...")
			for _, name := range names {
				err := validate(mode, name)
				if err != nil {
					fmt.Printf("Error(s) in name %s:\n%s\n", name, err)
				} else {
					fmt.Println("Name:", name, "OK")
				}
			}
			fmt.Print("Name validation is completed!", "\n\n")
		case "email":
			fmt.Println("Email validation started...")
			for _, email := range emails {
				err := validate(mode, email)
				if err != nil {
					fmt.Printf("Error(s) in email %s:\n%s\n", email, err)
				} else {
					fmt.Println("Email:", email, "OK")
				}
			}
			fmt.Print("Email validation is completed", "\n\n")
		case "password":
			fmt.Println("Password validation started...")
			for _, password := range passwords {
				err := validate(mode, password)
				if err != nil {
					fmt.Printf("Error(s) in password %s:\n%s\n", password, err)
				} else {
					fmt.Println("Email:", password, "OK")
				}
			}
			fmt.Print("Password validation is completed", "\n\n")
		}
	}
}
