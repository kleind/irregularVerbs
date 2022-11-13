package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	irregularVerbs := [][]string{
		{"sein", "be", "was/were"},
		{"schlagen", "beat", "beat"},
		{"werden", "become", "became"},
		{"beginnen/anfangen", "begin", "began"},
		{"wetten", "bet", "bet"},
		{"beißen", "bite", "bit"},
		{"blasen/pusten", "blow", "blew"},
		{"brechen/kaputt machen", "break", "broke"},
		{"(mit)bringen", "bring", "brought"},
		{"bauen", "build", "built"},
		{"(ver)brennen", "burn", "burnt"},
		{"kaufen", "buy", "bought"},
		{"fangen", "catch", "caught"},
		{"(aus)wählen", "choose", "chose"},
		{"kommen", "come", "came"},
		{"kosten", "cost", "cost"},
		{"schneiden", "cut", "cut"},
		{"sich befassen (mit)/umgehen (mit)", "deal", "dealt"},
		{"machen/tun", "do", "did"},
		{"zeichnen/ziehen", "draw", "drew"},
		{"träumen", "dream", "dreamt"},
		{"trinken", "drink", "drank"},
		{"fahren", "drive", "drove"},
		{"essen", "eat", "ate"},
		{"(hin)fallen", "fall", "fell"},
		{"füttern/ernähren", "feed", "fed"},
		{"fühlen", "feel", "felt"},
		{"kämpfen/(sich) streiten", "fight", "fought"},
		{"finden", "find", "found"},
		{"passen", "fit", "fit"},
		{"fliegen", "fly", "flew"},
		{"vergessen", "forget", "forgot"},
		{"vergeben/verzeihen", "forgive", "forgave"},
		{"gefrieren/erstarren", "freeze", "froze"},
		{"bekommen/erhalten", "get", "got"},
		{"geben", "give", "gave"},
		{"gehen/fahren", "go", "went"},
		{"wachsen/anbauen/züchten", "grow", "grew"},
		{"hängen", "hang", "hung"},
		{"haben", "have", "had"},
		{"hören", "hear", "heard"},
		{"(sich) verstecken", "hide", "hid"},
		{"schlagen/treffen", "hit", "hit"},
		{"(fest)halten", "hold", "held"},
		{"verletzen/sich weh tun", "hurt", "hurt"},
		{"(auf)bewahren/behalten", "keep", "kept"},
		{"kennen/wissen", "know", "knew"},
	}
	fmt.Println("Es gibt", len(irregularVerbs), "Vokabeln zu lernen.")
	fmt.Println("Das hier sind alle:", irregularVerbs)
	fmt.Println("Wenn du bereit bist, drücke die Return Taste. In 10 Sekunden geht es los.")
	//time.Sleep(10000000000)
	fmt.Printf("\x1bc")

	rand.Seed(time.Now().UnixNano())

	used := []int{}
	sumCorrectAnswers := 0
	sumWrongAnswers := 0
	sumAnswers := 0

	r := rand.Intn(len(irregularVerbs))
	var infinitiveAnswer string
	var simplePastAnswer string
	var nextWord string
	for {
		if contains(used, r) {
			r = rand.Intn(len(irregularVerbs))
		} else {
			fmt.Println("Bisher richtige Antworten:", sumCorrectAnswers, "von", len(irregularVerbs))
			fmt.Println("Deutsch:", irregularVerbs[r][0])
			fmt.Print("Infinitive: ")
			fmt.Scanf("%s", &infinitiveAnswer)
			fmt.Print("Simple Past: ")
			fmt.Scanf("%s", &simplePastAnswer)
			if infinitiveAnswer == irregularVerbs[r][1] && simplePastAnswer == irregularVerbs[r][2] {
				fmt.Println(string(colorGreen) + "Das war richtig!" + string(colorReset))
				fmt.Println("")
				used = append(used, r)
				sumCorrectAnswers++
				sumAnswers++
				fmt.Println("Drücke Return wenn es weitergehen soll.")
				fmt.Scanf("%s", &nextWord)
				fmt.Printf("\x1bc")
			} else {
				fmt.Println(string(colorRed) + "Das war leider falsch. Richtig ist:" + string(colorReset))
				fmt.Println("Infinitive :" + string(colorGreen) + irregularVerbs[r][1] + string(colorReset))
				fmt.Println("Simple Past:" + string(colorGreen) + irregularVerbs[r][2] + string(colorReset))
				fmt.Println("Das Wort kommt später noch einmal." + string(colorReset))
				fmt.Println("")
				r = rand.Intn(len(irregularVerbs))
				sumWrongAnswers++
				sumAnswers++
				fmt.Println("Drücke Return wenn es weitergehen soll.")
				fmt.Scanf("%s", &nextWord)
				fmt.Printf("\x1bc")
			}
		}
		if len(used) == len(irregularVerbs) {
			percent := float64(float64(sumCorrectAnswers) / float64(sumAnswers) * 100)
			fmt.Println("Du hast alle Verben einmal bekommen.")
			fmt.Println("Du hast", sumAnswers, "Antworten gegeben. Davon waren", sumCorrectAnswers, "richtig, und", sumWrongAnswers, "falsch eingegeben.")
			fmt.Printf("%f Prozent deiner Antworten waren korrekt.\n", percent)
			os.Exit(0)
		}
	}
}
