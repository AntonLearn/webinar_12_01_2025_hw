package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	errorLessTwoCharacters        = errors.New("consists of less than two characters")
	errorLessThreeCharacters      = errors.New("consists of less than three characters")
	errorNotLatAlphabetCharacters = errors.New("does not consist only of Latin alphabet characters")
	errorMoreHundredCharacters    = errors.New("consists of more than hundred characters")
	errorNotDog                   = errors.New("does not contain special character '@'")
)

func validateName(validatedString string) error {
	var errorTotal error
	validateStringToRunes := []rune(validatedString)
	for _, rune := range validateStringToRunes {
		if rune < 65 || rune > 122 || (rune <= 96 && rune >= 91) {
			errorTotal = errorNotLatAlphabetCharacters
			break
		}
	}
	if len(validateStringToRunes) < 2 {
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
	for _, rune := range validateStringToRunes {
		if rune < 65 || rune > 122 || (rune <= 96 && rune >= 91) {
			if rune != 64 && rune != 46 {
				errorTotal = errorNotLatAlphabetCharacters
				break
			}
		}
	}
	if len(validateStringToRunes) < 3 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorLessThreeCharacters)
		} else {
			errorTotal = errorLessThreeCharacters
		}
	}
	if len(validateStringToRunes) > 100 {
		if errorTotal != nil {
			errorTotal = errors.Join(errorTotal, errorMoreHundredCharacters)
		} else {
			errorTotal = errorMoreHundredCharacters
		}
	}
	if !strings.ContainsRune(validatedString, '@') {
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

func validate(mode, validatedString string) error {
	switch mode {
	case "name":
		return validateName(validatedString)
	case "email":
		return validateEmail(validatedString)
	default:
		return nil
	}
}

func main() {
	var modes = []string{"name", "email", "password"}
	var names = []string{"", "a", "ю", "aa", "aю", "a.", "Azw", "Sw_r", "Dx2", "яя", "superuser"}
	var emails = []string{
		"user@mail.com", "", "u@m", "u@", "us", "qwertyuiopasdfghjkl",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxc",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcv",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm@wertyuiopasdfghjklzxc",
		"qwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnmqwertyuiopasdfghjklzxcvbnm@wertyuiopasdfghjklzxcv",
	}
	for _, mode := range modes {
		switch mode {
		case "name":
			fmt.Println("Name validation started...")
			for _, name := range names {
				err := validate(mode, name)
				if err != nil {
					fmt.Printf("Error(s) in name %s: %s\n", name, err)
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
					fmt.Printf("Error(s) in email %s: %s\n", email, err)
				} else {
					fmt.Println("Email:", email, "OK")
				}
			}
			fmt.Print("Email validation is completed", "\n\n")
		}
	}
}
