package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var colorReset string = "\033[0m"
var colorRed string = "\033[31m"
var colorGreen string = "\033[32m"
var irregularVerbs = [][]string{
	{"sein", "be", "was/were", "been"},
	{"schlagen", "beat", "beat", "beaten"},
	{"werden", "become", "became", "become"},
	{"beginnen/anfangen", "begin", "began", "begun"},
	{"wetten", "bet", "bet", "bet"},
	{"beißen", "bite", "bit", "bitten"},
	{"blasen/pusten", "blow", "blew", "blown"},
	{"brechen/kaputt machen", "break", "broke", "broken"},
	{"(mit)bringen", "bring", "brought", "brought"},
	{"bauen", "build", "built", "built"},
	{"(ver)brennen", "burn", "burnt", "burnt"},
	{"kaufen", "buy", "bought", "bought"},
	{"fangen", "catch", "caught", "caught"},
	{"(aus)wählen", "choose", "chose", "chosen"},
	{"kommen", "come", "came", "come"},
	{"kosten", "cost", "cost", "cost"},
	{"schneiden", "cut", "cut", "cut"},
	{"sich befassen (mit)/umgehen (mit)", "deal", "dealt", "dealt"},
	{"machen/tun", "do", "did", "done"},
	{"zeichnen/ziehen", "draw", "drew", "drawn"},
	{"träumen", "dream", "dreamt", "dreamt"},
	{"trinken", "drink", "drank", "drunk"},
	{"fahren", "drive", "drove", "driven"},
	{"essen", "eat", "ate", "eaten"},
	{"(hin)fallen", "fall", "fell", "fallen"},
	{"füttern/ernähren", "feed", "fed", "fed"},
	{"fühlen", "feel", "felt", "felt"},
	{"kämpfen/(sich) streiten", "fight", "fought", "fought"},
	{"finden", "find", "found", "found"},
	{"passen", "fit", "fit", "fit"},
	{"fliegen", "fly", "flew", "flown"},
	{"vergessen", "forget", "forgot", "forgotten"},
	{"vergeben/verzeihen", "forgive", "forgave", "forgiven"},
	{"gefrieren/erstarren", "freeze", "froze", "frozen"},
	{"bekommen/erhalten", "get", "got", "got"},
	{"geben", "give", "gave", "given"},
	{"gehen/fahren", "go", "went", "gone"},
	{"wachsen/anbauen/züchten", "grow", "grew", "grown"},
	{"hängen", "hang", "hung", "hung"},
	{"haben", "have", "had", "had"},
	{"hören", "hear", "heard", "heard"},
	{"(sich) verstecken", "hide", "hid", "hidden"},
	{"schlagen/treffen", "hit", "hit", "hit"},
	{"(fest)halten", "hold", "held", "held"},
	{"verletzen/sich weh tun", "hurt", "hurt", "hurt"},
	{"(auf)bewahren/behalten", "keep", "kept", "kept"},
	{"kennen/wissen", "know", "knew", "known"},
	{"(an)führen", "lead", "led", "led"},
	{"lernen", "learn", "learnt", "learnt"},
	{"(ver)lassen", "leave", "left", "left"},
	{"(ver)leihen", "lend", "lent", "lent"},
	{"lassen", "let", "let", "let"},
	{"liegen", "lie", "lay", "lain"},
	{"verlieren", "lose", "lost", "lost"},
	{"machen/tun", "make", "made", "made"},
	{"bedeuten/meinen", "mean", "meant", "meant"},
	{"treffen", "meet", "met", "met"},
	{"(be)zahlen", "pay", "paid", "paid"},
	{"legen/setzen/stellen", "put", "put", "put"},
	{"lesen", "read", "read", "read"},
	{"fahren/reiten", "ride", "rode", "ridden"},
	{"klingeln/läuten", "ring", "rang", "rung"},
	{"steigen/sich erheben", "rise", "rose", "risen"},
	{"laufen/rennen", "run", "ran", "run"},
	{"sagen", "say", "said", "said"},
	{"sehen", "see", "saw", "seen"},
	{"verkaufen", "sell", "sold", "sold"},
	{"senden/verschicken", "send", "sent", "sent"},
	{"erbauen/errichten", "setup", "setup", "setup"},
	{"scheinen/glänzen", "shine", "shone", "shone"},
	{"schießen", "shoot", "shot", "shot"},
	{"zeigen", "show", "showed", "shown"},
	{"singen", "sing", "sang", "sung"},
	{"untergehen/sinken", "sink", "sank", "sunk"},
	{"sitzen", "sit", "sat", "sat"},
	{"schlafen", "sleep", "slept", "slept"},
	{"riechen/duften", "smell", "smelt", "smelt"},
	{"sprechen", "speak", "spoke", "spoken"},
	{"buchstabieren", "spell", "spelt", "spelt"},
	{"ausgeben/verbringen", "spend", "spent", "spent"},
	{"verschütten/ausschütten", "spill", "spilt", "spilt"},
	{"stehen", "stand", "stood", "stood"},
	{"stehlen", "steal", "stole", "stolen"},
	{"stechen", "sting", "stung", "stung"},
	{"schwimmen", "swim", "swam", "swum"},
	{"nehmen", "take", "took", "taken"},
	{"unterrichten/lehren/beibringen", "teach", "taught", "taught"},
	{"erzählen", "tell", "told", "told"},
	{"(nach)denken/glauben", "think", "thought", "thought"},
	{"werfen", "throw", "threw", "thrown"},
	{"verstehen", "understand", "understood", "understood"},
	{"(auf)wachen/(auf)wecken", "wake up", "woke up", "woken up"},
	{"anhaben/tragen", "wear", "wore", "worn"},
	{"gewinnen/siegen", "win", "won", "won"},
	{"schreiben", "write", "wrote", "written"},
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func clearScreen() {
	fmt.Printf("\x1bc")
}

func main() {
	clearScreen()

	rand.Seed(time.Now().UnixNano())

	used := []int{}
	sumCorrectAnswers := 0
	sumWrongAnswers := 0
	sumAnswers := 0

	r := rand.Intn(len(irregularVerbs))
	var simplePastAnswer string
	var pastProgressiveAnswer string
	var nextWord string
	for {
		if contains(used, r) {
			r = rand.Intn(len(irregularVerbs))
		} else {
			fmt.Println("Bisher richtige Antworten:", sumCorrectAnswers, "von", len(irregularVerbs))
			fmt.Println("Deutsch:", irregularVerbs[r][0])
			fmt.Print("Infinitive: ")
			// fmt.Scanln(&infinitiveAnswer)

			inputReader := bufio.NewReader(os.Stdin)
			infinitiveAnswer, _ := inputReader.ReadString('\n')
			infinitiveAnswer = strings.TrimSpace(infinitiveAnswer)

			if infinitiveAnswer == irregularVerbs[r][1] || infinitiveAnswer == "to "+irregularVerbs[r][1] {
				fmt.Println(string(colorGreen) + "Das war richtig!" + string(colorReset))
			} else {
				fmt.Println(string(colorRed) + "Das war leider falsch. Richtig ist:" + string(colorReset))
				fmt.Println("Infinitive: " + string(colorGreen) + irregularVerbs[r][1] + string(colorReset))
			}

			fmt.Print("Simple Past: ")
			fmt.Scanf("%s", &simplePastAnswer)
			if simplePastAnswer == irregularVerbs[r][2] {
				fmt.Println(string(colorGreen) + "Das war richtig!" + string(colorReset))
			} else {
				fmt.Println(string(colorRed) + "Das war leider falsch. Richtig ist:" + string(colorReset))
				fmt.Println("Simple Past: " + string(colorGreen) + irregularVerbs[r][2] + string(colorReset))
			}

			fmt.Print("Present Perfect: ")
			fmt.Scanf("%s", &pastProgressiveAnswer)
			if pastProgressiveAnswer == irregularVerbs[r][3] {
				fmt.Println(string(colorGreen) + "Das war richtig!" + string(colorReset))
			} else {
				fmt.Println(string(colorRed) + "Das war leider falsch. Richtig ist:" + string(colorReset))
				fmt.Println("Past Progressive: " + string(colorGreen) + irregularVerbs[r][3] + string(colorReset))
			}

			if (infinitiveAnswer == irregularVerbs[r][1] || infinitiveAnswer == "to "+irregularVerbs[r][1]) && simplePastAnswer == irregularVerbs[r][2] && pastProgressiveAnswer == irregularVerbs[r][3] {
				fmt.Println("")
				used = append(used, r)
				sumCorrectAnswers++
				sumAnswers++
				fmt.Println("Drücke Return wenn es weitergehen soll.")
				fmt.Scanf("%s", &nextWord)
				clearScreen()
			} else {
				fmt.Print("\nDas Wort kommt später noch einmal.\n\n" + string(colorReset))
				r = rand.Intn(len(irregularVerbs))
				sumWrongAnswers++
				sumAnswers++
				fmt.Println("Drücke Return wenn es weitergehen soll.")
				fmt.Scanf("%s", &nextWord)
				clearScreen()
			}
		}
		if len(used) == len(irregularVerbs) {
			percent := float64(float64(sumCorrectAnswers) / float64(sumAnswers) * 100)
			fmt.Println("Du hast alle Verben einmal korrekt beantwortet.")
			fmt.Println("Du hast", sumAnswers, "Antworten gegeben. Davon waren", sumCorrectAnswers, "richtig, und", sumWrongAnswers, "falsch.")
			fmt.Printf("%.2f Prozent deiner Antworten waren korrekt.\n", percent)
			os.Exit(0)
		}
	}
}
