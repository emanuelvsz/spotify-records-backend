package main

import (
	"flag"
	cfg "module/src/app/api/config"
	"module/src/infra/postgres"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

const defaultEnvFilePath = "./src/app/api/.env"

// NewAPI
// @title SPOTIFY RECORDS API
// @version 1.0
// @description Aplicação de artistas do spotify
// @contact.name DIT - IFAL
// @contact.email evs10@aluno.ifal.edu.br
// @BasePath /api
// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	loadEnvIntoSystemFromFile()

	setupPostgres()
	serve(
		cfg.Env().Server.Host,
		cfg.Env().Server.Port,
	)
}

func setupPostgres() {
	err := postgres.SetUpCredentials(
		cfg.Env().Postgres.User,
		cfg.Env().Postgres.Password,
		cfg.Env().Postgres.Name,
		cfg.Env().Postgres.Host,
		cfg.Env().Postgres.Port,
		cfg.Env().Postgres.SSLMode,
	)
	if err != nil {
		panic(err)
	}

	err = postgres.MigrateUp()
	if err != nil {
		panic(err)
	}
}

func loadEnvIntoSystemFromFile() {
	log.Info().Msg("loading environment variables...")

	envPathPtr := flag.String("env", "", "override environment variable path")
	flag.Parse()
	if *envPathPtr == "" {
		*envPathPtr = defaultEnvFilePath
	}

	err := godotenv.Load(*envPathPtr)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to load config from env")
	}
}
