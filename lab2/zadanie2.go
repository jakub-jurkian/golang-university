// Uczestnik
// Imie - string,
// Repertuar - lista stringow(utworow),
// Macierz ocen - lista w liscie obiektow intow(ocena),

// uczestnik gra utwor i ten utwor zostaje oceniony przez jurorow
// jurorzy wystawiaja ocene od 0 do 25 za kazdy utwor kazdego uczestnika
// pozniej nastepuje ewentualna korekcja ocen jurorow
// na koniec brana jest srednia ocen kazdego utworu kazdego uczestnika i wpisywana jako jego finalna ocena
// wyznaczanie uczestnika z najwyższą liczbą punktów uzyskaną za konkretny utwór

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
)

type Participant struct {
	Name       string
	Repertoire []string
	Scores     [][]int
	Averages   []float64
	FinalScore float64
}

type SongWinner struct {
	Song   string
	Points float64
	Winner Participant
}

func assignNotes(p Participant, juryAmount int) Participant {
	newP := p
	newP.Scores = make([][]int, len(p.Repertoire))

	for i := range len(newP.Repertoire) {
		newP.Scores[i] = make([]int, juryAmount)
		for j := range juryAmount {
			note := rand.Intn(26)
			newP.Scores[i][j] = note
		}
	}
	return newP
}

func correctScores(score int, average float64, phase int) float64 {
	newScore := float64(score)
	threshold := 2.0

	if phase == 1 {
		threshold = 3.0
	}

	if newScore > average+threshold {
		newScore = average + threshold
	}
	if newScore < average-threshold {
		newScore = average - threshold
	}

	return newScore
}

func averageScores(p Participant, phase int) Participant {
	newP := p
	newP.Averages = make([]float64, len(newP.Repertoire))

	for i := range len(newP.Scores) {
		// Raw AVG
		sum := 0
		for j := range len(newP.Scores[i]) {
			sum += newP.Scores[i][j]
		}
		avg := float64(sum) / float64(len(newP.Scores[i]))

		// Korekta i nowa srednia
		correctedAvg := 0.0
		for j := range newP.Scores[i] {
			correctedAvg += correctScores(newP.Scores[i][j], avg, phase)
		}
		correctedAvg /= float64(len(newP.Scores[i]))

		newP.Averages[i] = correctedAvg
	}

	return newP
}

func finalScore(p Participant) Participant {
	newP := p
	avg := 0.0
	for i := range len(newP.Averages) {
		avg += float64(newP.Averages[i])
	}
	if len(newP.Averages) > 0 {
		newP.FinalScore = math.Round((avg/float64(len(newP.Averages)))*100) / 100
	}

	return newP
}

func songWinners(participants []Participant, repertuar []string) []SongWinner {
	var songWinners []SongWinner

	for i := range len(repertuar) {
		maxNote := 0.0
		bestP := participants[0]
		for j := range len(participants) {
			if float64(maxNote) < float64(participants[j].Averages[i]) {
				maxNote = float64(participants[j].Averages[i])
				bestP = participants[j]
			}
		}
		songWinners = append(songWinners, SongWinner{Song: repertuar[i], Points: maxNote, Winner: bestP})
	}

	return songWinners
}

func main() {
	currentPhase := 1

	repertuar := []string{"A", "B", "C"}

	participants := []Participant{
		{
			Name:       "Jan",
			Repertoire: repertuar,
		},
		{
			Name:       "Adam",
			Repertoire: repertuar,
		},
		{
			Name:       "Zosia",
			Repertoire: repertuar,
		},
	}

	for i := range len(participants) {
		newP := participants[i]
		newP = assignNotes(newP, 5)
		newP = averageScores(newP, currentPhase)
		newP = finalScore(newP)

		participants[i] = newP
	}

	fmt.Printf("%+v\n", participants[0])

	sort.Slice(participants, func(i, j int) bool {
		return participants[i].FinalScore > participants[j].FinalScore
	})

	winners := songWinners(participants, repertuar)

	fmt.Printf("%+v\n", winners)

	// Results
	for i, p := range participants {
		fmt.Printf("%d. %-12s | Wynik końcowy: %.2f | Średnie za utwory: %v\n",
			i+1, p.Name, p.FinalScore, p.Averages)
	}

	for _, sw := range winners {
		fmt.Printf("Utwór %-2s: Najlepszy wykonawca -> %-10s (%.2f pkt)\n",
			sw.Song, sw.Winner.Name, sw.Points)
	}
}
