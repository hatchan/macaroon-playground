package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gopkg.in/macaroon.v1"
)

func main() {
	app := cli.NewApp()

	app.Name = "macaroon-playground"
	app.Usage = "Create and verify macaroons"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Enable debug log statements",
			EnvVar: "DEBUG",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "generate",
			Usage:  "Generate a macaroon",
			Action: generateMacaroon,
		},
		{
			Name:   "validate",
			Usage:  "Validate a macaroon",
			Action: validateMacaroon,
		},
	}

	app.Run(os.Args)
}

func generateMacaroon(c *cli.Context) {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	rootKey := []byte{'a'}
	id := "id"
	loc := "location"

	m, err := macaroon.New(rootKey, id, loc)
	if err != nil {
		log.Panic("Unable to create macaroon")
	}

	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		log.Panic("Unable to marshal macaroon")
	}

	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")
}

func validateMacaroon(c *cli.Context) {
	if c.GlobalBool("debug") {
		log.SetLevel(log.DebugLevel)
	}

	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Panic("Unable to read from stdin")
	}

	m := &macaroon.Macaroon{}

	err = json.Unmarshal(b, m)
	if err != nil {
		log.Panic("Unable to unmarshal macaroon")
	}

	rootKey := []byte{'a'}
	discharges := []*macaroon.Macaroon{}
	err = m.Verify(rootKey, verify, discharges)
	if err != nil {
		log.Panic("Unable to verify")
	}
}

func verify(caveat string) error {
	log.Error("Verify called")
	return nil
}
