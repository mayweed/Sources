package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/**
 * Complete the hackathon before your opponent by following the principles of Green IT
 **/

type Application struct {
	objectType               string
	id                       int
	trainingNeeded           int
	codingNeeded             int
	dailyRoutineNeeded       int
	taskPrioritizationNeeded int
	architectureStudyNeeded  int
	continuousDeliveryNeeded int
	codeReviewNeeded         int
	refactoringNeeded        int
}

type Player struct {
	playerLocation                        int
	playerScore                           int
	playerPermanentDailyRoutineCards      int
	playerPermanentArchitectureStudyCards int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	for {
		// gamePhase: can be MOVE, GIVE_CARD, THROW_CARD, PLAY_CARD or RELEASE
		var gamePhase string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &gamePhase)

		var applicationsCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &applicationsCount)

		var apps []Application

		for i := 0; i < applicationsCount; i++ {
			// trainingNeeded: number of TRAINING skills needed to release this application
			// codingNeeded: number of CODING skills needed to release this application
			// dailyRoutineNeeded: number of DAILY_ROUTINE skills needed to release this application
			// taskPrioritizationNeeded: number of TASK_PRIORITIZATION skills needed to release this application
			// architectureStudyNeeded: number of ARCHITECTURE_STUDY skills needed to release this application
			// continuousDeliveryNeeded: number of CONTINUOUS_DELIVERY skills needed to release this application
			// codeReviewNeeded: number of CODE_REVIEW skills needed to release this application
			// refactoringNeeded: number of REFACTORING skills needed to release this application
			var objectType string
			var id, trainingNeeded, codingNeeded, dailyRoutineNeeded, taskPrioritizationNeeded, architectureStudyNeeded, continuousDeliveryNeeded, codeReviewNeeded, refactoringNeeded int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &objectType, &id, &trainingNeeded, &codingNeeded, &dailyRoutineNeeded, &taskPrioritizationNeeded, &architectureStudyNeeded, &continuousDeliveryNeeded, &codeReviewNeeded, &refactoringNeeded)
			apps = append(apps, Application{objectType, id, trainingNeeded, codingNeeded, dailyRoutineNeeded, taskPrioritizationNeeded, architectureStudyNeeded, continuousDeliveryNeeded, codeReviewNeeded, refactoringNeeded})
		}

		var players []Player
		for i := 0; i < 2; i++ {
			// playerLocation: id of the zone in which the player is located
			// playerPermanentDailyRoutineCards: number of DAILY_ROUTINE the player has played. It allows them to take cards from the adjacent zones
			// playerPermanentArchitectureStudyCards: number of ARCHITECTURE_STUDY the player has played. It allows them to draw more cards
			var playerLocation, playerScore, playerPermanentDailyRoutineCards, playerPermanentArchitectureStudyCards int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &playerLocation, &playerScore, &playerPermanentDailyRoutineCards, &playerPermanentArchitectureStudyCards)
			players = append(players, Player{playerLocation, playerScore, playerPermanentDailyRoutineCards, playerPermanentArchitectureStudyCards})
		}

		var cardLocationsCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &cardLocationsCount)

		var myHand []int
		var myDraw []int
		var myDiscard []int

		for i := 0; i < cardLocationsCount; i++ {
			// cardsLocation: the location of the card list. It can be HAND, DRAW, DISCARD or OPPONENT_CARDS (AUTOMATED and OPPONENT_AUTOMATED will appear in later leagues)
			var cardsLocation string
			var training, coding, dailyRoutine, taskPrioritization, architectureStudy, continuousDelivery, codeReview, refactoring, bonus, technicalDebt int
			scanner.Scan()
			fmt.Sscan(scanner.Text(), &cardsLocation, &training, &coding, &dailyRoutine, &taskPrioritization, &architectureStudy, &continuousDelivery, &codeReview, &refactoring, &bonus, &technicalDebt)
			switch cardsLocation {
			case "HAND":
				myHand = append(myHand, training, coding, dailyRoutine, taskPrioritization, architectureStudy, continuousDelivery, codeReview, refactoring, bonus, technicalDebt)
			case "DRAW":
				myDraw = append(myDraw, training, coding, dailyRoutine, taskPrioritization, architectureStudy, continuousDelivery, codeReview, refactoring, bonus, technicalDebt)
			case "DISCARD":
				myDiscard = append(myDiscard, training, coding, dailyRoutine, taskPrioritization, architectureStudy, continuousDelivery, codeReview, refactoring, bonus, technicalDebt)

			}
		}

		var possibleMovesCount int
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &possibleMovesCount)

		for i := 0; i < possibleMovesCount; i++ {
			scanner.Scan()
			possibleMove := scanner.Text()
			_ = possibleMove // to avoid unused error
		}
		log.Println(gamePhase)
		log.Println(cardLocationsCount, myHand, myDraw)

		// In the first league: RANDOM | MOVE <zoneId> | RELEASE <applicationId> | WAIT;
		fmt.Println("RANDOM")
	}
}
