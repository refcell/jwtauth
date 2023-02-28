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
				Usage:   "Generate claims for a JWT token",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "expiry",
						Value: fmt.Sprintf("%ds", time.Now().Unix()),
						Usage: "Unix timestamp with units appended for the expiry of the token (eg. --expiry 1677621775s)",
					},
				},
				Action: func(cCtx *cli.Context) error {
					secret := cCtx.Args().First()
					if secret == "" {
						return fmt.Errorf("provide a hex encoded byte secret string")
					}
					secretBytes, err := Decode(secret)
					if err != nil {
						return err
					}
					var fixedSecretBytes [32]byte
					copy(fixedSecretBytes[:], secretBytes[:32])
					expiry := cCtx.String("expiry")
					if expiry == "" {
						s, err := NewJWTClaims(fixedSecretBytes, time.Now())
						if err != nil {
							return err
						}
						fmt.Printf("%s\n", s)
					} else {
						exp, err := time.ParseDuration(expiry)
						if err != nil {
							return err
						}
						s, err := NewJWTClaims(fixedSecretBytes, time.Now(), exp)
						if err != nil {
							return err
						}
						fmt.Printf("%s\n", s)
					}
					return nil
				},
			},
			{
				Name:    "header",
				Aliases: []string{"h"},
				Usage:   "Generate the http auth bearer header for a JWT token",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "expiry",
						Value: fmt.Sprintf("%ds", time.Now().Unix()),
						Usage: "Unix timestamp with units appended for the expiry of the token (eg. --expiry 1677621775s)",
					},
				},
				Action: func(cCtx *cli.Context) error {
					secret := cCtx.Args().First()
					if secret == "" {
						return fmt.Errorf("provide a hex encoded byte secret string")
					}
					secretBytes, err := Decode(secret)
					if err != nil {
						return err
					}
					var fixedSecretBytes [32]byte
					copy(fixedSecretBytes[:], secretBytes[:32])
					expiry := cCtx.String("expiry")
					if expiry == "" {
						s, err := NewJWTClaims(fixedSecretBytes, time.Now())
						if err != nil {
							return err
						}
						fmt.Printf("Authorization: Bearer %s\n", s)
					} else {
						exp, err := time.ParseDuration(expiry)
						if err != nil {
							return err
						}
						s, err := NewJWTClaims(fixedSecretBytes, time.Now(), exp)
						if err != nil {
							return err
						}
						fmt.Printf("Authorization: Bearer %s\n", s)
					}
					return nil
				},
			},
			{
				Name:    "generate",
				Aliases: []string{"g", "gen"},
				Usage:   "Generate a new hex encoded 32-byte JWT token",
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
