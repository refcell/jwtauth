package main

import (
	"fmt"
	"log"
	"os"
	"time"

	crand "crypto/rand"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "jwtauth",
		Usage: "Minimal cli to interact with JWT Auth Tokens",
		Commands: []*cli.Command{
			{
				Name:    "claims",
				Aliases: []string{"c"},
				Usage:   "generate claims for a JWT token",
				Action: func(cCtx *cli.Context) error {
					secret := cCtx.Args().First()
					if secret == "" {
						return fmt.Errorf("provide a hex encoded 64 byte secret string")
					}
					secretBytes, err := Decode(secret)
					if err != nil {
						return err
					}
					var fixedSecretBytes [32]byte
					copy(fixedSecretBytes[:], secretBytes[:32])
					s, err := NewJWTClaims(fixedSecretBytes, time.Now())
					if err != nil {
						return err
					}
					fmt.Printf("%s\n", s)
					return nil
				},
			},
			{
				Name:    "header",
				Aliases: []string{"h"},
				Usage:   "generate the http auth bearer header for a JWT token",
				Action: func(cCtx *cli.Context) error {
					secret := cCtx.Args().First()
					if secret == "" {
						return fmt.Errorf("provide a hex encoded 64 byte secret string")
					}
					secretBytes, err := Decode(secret)
					if err != nil {
						return err
					}
					var fixedSecretBytes [32]byte
					copy(fixedSecretBytes[:], secretBytes[:32])
					s, err := NewJWTClaims(fixedSecretBytes, time.Now())
					if err != nil {
						return err
					}
					fmt.Printf("Authentication Bearer %s\n", s)
					return nil
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "generate a new JWT token",
				Action: func(cCtx *cli.Context) error {
					var secret [32]byte
					if _, err := crand.Read(secret[:]); err != nil {
						return fmt.Errorf("failed to create jwt secret: %v", err)
					}
					fmt.Printf("%s\n", Encode(secret[:]))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
