package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Welcome to the epic fairy tale generator! Let's create a magical and enchanting story.")
	var hero, heroine string

	fmt.Print("Enter the name of the hero: ")
	fmt.Scanln(&hero)
	fmt.Print("Enter the name of the heroine: ")
	fmt.Scanln(&heroine)

	fairyTale(hero, heroine)
}

func getRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10) + 1
}

func fairyTale(heroName, heroineName string) {
	fmt.Printf("Once upon a time, in the enchanted kingdom of %s and %s,\n", heroName, heroineName)
	time.Sleep(3 * time.Second)

	if getRandom() == 10 {
		fortunateEvent(heroName)
		return
	}

	questResult := embarkOnQuest(heroName, heroineName)

	if getRandom() == 10 {
		fortunateEvent(heroName)
		return
	}

	switch questResult {
	case 1:
		fmt.Printf("Oh no! Despite their courage, %s and %s faced a formidable challenge.\n", heroName, heroineName)
		time.Sleep(3 * time.Second)
		fmt.Printf("The duo encountered a mystical creature in the dark forest, testing the limits of their bravery.\n")
		time.Sleep(4 * time.Second)
		fmt.Printf("%s and %s, determined to overcome the challenge, fought valiantly, but alas, it was not enough.\n", heroName, heroineName)
		time.Sleep(5 * time.Second)
		fmt.Printf("This unfortunate event affected their lives. %s found themselves constantly worried, unable to sleep or work peacefully.\n", heroName)
		time.Sleep(5 * time.Second)
		fmt.Printf("%s and %s, you must rise again! Face this challenge once more and prove the strength of your bond!\n", heroName, heroineName)
		time.Sleep(10 * time.Second)
		fairyTale(heroName, heroineName)

	case 2:
		fmt.Printf("The journey was tough, but %s and %s managed to overcome their fears and complete the quest.\n", heroName, heroineName)
		time.Sleep(3 * time.Second)
		fmt.Printf("Along the way, they discovered a hidden treasure that brought prosperity to their kingdom.\n")
		time.Sleep(4 * time.Second)
		fmt.Printf("The enchanted land flourished, and %s and %s ruled justly, bringing joy to all.\n", heroName, heroineName)

	case 3:
		fmt.Printf("A mysterious portal appeared before %s and %s, leading them to a magical realm.\n", heroName, heroineName)
		time.Sleep(3 * time.Second)
		fmt.Printf("In this realm, they met wise wizards who bestowed upon them mystical powers.\n")
		time.Sleep(4 * time.Second)
		fmt.Printf("%s and %s returned to their kingdom as powerful sorcerers, bringing prosperity and magic.\n", heroName, heroineName)

	default:
		fmt.Printf("Hooray! %s and %s are true heroes! They successfully completed their quest, learning to control their fears along the way.\n", heroName, heroineName)
		time.Sleep(4 * time.Second)
		fmt.Printf("Their bond grew stronger with each challenge, and the magical creatures they encountered bestowed upon them wisdom and courage.\n")
		time.Sleep(5 * time.Second)
		fmt.Printf("With newfound strength, %s and %s returned to their kingdom, greeted by cheers and celebrations.\n", heroName, heroineName)
		time.Sleep(4 * time.Second)
		fmt.Printf("The enchanted land flourished, and %s and %s ruled justly, bringing prosperity to all.\n", heroName, heroineName)
	}
}

func embarkOnQuest(heroName, heroineName string) int {
	fmt.Printf("%s and %s, in the midst of this enchanted realm, you might be feeling a bit anxious.\n", heroName, heroineName)
	time.Sleep(3 * time.Second)
	var answer string

	for {
		fmt.Printf("Did you ever embark on a magical quest together (yes/no), %s? ", heroName)
		fmt.Scanln(&answer)
		answer = toLower(answer)

		if answer == "yes" || answer == "no" {
			break
		}
		fmt.Printf("Take your time and answer correctly (yes/no)\n")
	}

	if answer == "yes" {
		fmt.Printf("It's wonderful that you've faced magical quests before!\n")
		time.Sleep(3 * time.Second)
		var treatment string

		for {
			fmt.Printf("%s, have you sought the advice of the wise wizard Gandalf for guidance (yes/no)? ", heroName)
			fmt.Scanln(&treatment)
			treatment = toLower(treatment)

			if treatment == "yes" || treatment == "no" {
				break
			}
			fmt.Printf("I know it may not be easy to share, but focus and tell me (yes/no)\n")
		}

		if treatment == "yes" {
			return getRandom()
		} else if treatment == "no" {
			index := getRandom() - 1
			if index >= 0 && index < 5 {
				return []int{1, 2, 3, 4, 5}[index]
			}
		}
	} else if answer == "no" {
		fmt.Printf("Unfortunately, challenges can arise even in the most magical of lands, but it is never too late to face them.\n")
		time.Sleep(4 * time.Second)
		var feeling string

		for {
			fmt.Printf("%s, do you feel uneasy right now (yes/no)? ", heroName)
			fmt.Scanln(&feeling)
			feeling = toLower(feeling)

			if feeling == "yes" || feeling == "no" {
				break
			}
			fmt.Printf("I am deeply concerned about your well-being, so tell me, do you feel uneasy (yes/no)?\n")
		}

		if feeling == "yes" {
			return getRandom()
		} else if feeling == "no" {
			return getRandom()
		}
	}
	return 0
}

func fortunateEvent(name string) {
	fmt.Printf("Congratulations, %s! Today is your lucky day! A magical event has unfolded, and all your worries are swept away.\n", name)
}

func toLower(s string) string {
	return strings.ToLower(s)
}
