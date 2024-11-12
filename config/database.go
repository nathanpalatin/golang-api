package config

import (
	"apigo/prisma/db"

	"github.com/rs/zerolog/log"
)

func ConnectDB()(*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Connect(); err != nil {
		return nil, err
	}

	log.Info().Msg("Connected to database")
	return client, nil
}