package main

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "slots"
	app.Usage = "Spin the slots"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		println(spinSlots())
	}

	app.Run(os.Args)
}

func spinSlots() string {

	symbols := [4]string{"\u2660", "\u2663", "\u2665", "\u2666"}

	rand.Seed(time.Now().UTC().UnixNano())

	first := rand.Intn(4)
	second := rand.Intn(4)
	third := rand.Intn(4)
	fourth := rand.Intn(4)

	output := []string{symbols[first], symbols[second], symbols[third], symbols[fourth]}

	if (first == second) && (second == third) && (third == fourth) && (fourth == first) {
		return "Jackpot: " + strings.Join(output, " ")
	} else if (first == second) && (third == fourth) {
		return "Double trouble: " + strings.Join(output, " ")
	} else if (first == third) && (second == fourth) {
		return "Alternate reality: " + strings.Join(output, " ")
	} else if (first == fourth) && (second == third) {
		return "Sharp edges: " + strings.Join(output, " ")
	} else {
		return "Better luck next time: " + strings.Join(output, " ")
	}
}
