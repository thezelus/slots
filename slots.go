package main

import (
	"math/rand"
	"os"
	"strconv"
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

		var err error
		numberOfSpins := 1

		if len(c.Args()) > 0 {
			numberOfSpins, err = strconv.Atoi(c.Args()[0])
		}

		if err != nil {
			println("Invalid value for arguments")
		} else {
			for i := 0; i < numberOfSpins; i++ {
				println(spinSlots())
			}
		}
	}

	app.Run(os.Args)
}

func spinSlots() string {

	symbols := []string{"\u2660", "\u2663", "\u2665", "\u2666"}

	rand.Seed(time.Now().UTC().UnixNano())

	len := len(symbols)

	first := rand.Intn(len)
	second := rand.Intn(len)
	third := rand.Intn(len)
	fourth := rand.Intn(len)

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
