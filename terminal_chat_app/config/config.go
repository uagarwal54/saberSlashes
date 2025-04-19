package config

import (
	"context"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/jackc/pgx/v5"
)

type (
	// Configs struct will have all the controlls loaded into it
	Configs struct {
		DBs     []Db   `yaml:"supabase"`
		DbInUse string `yaml:"dbinUse"`
	}
	// Db struct will contain the configs of various dbs
	Db struct {
		Name       string `yaml:"name"`
		ConnString string `yaml:"conn_string"`
	}
)

// InitializeSupabase initializes the Supabase client
func InitializeSupabase(connString string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Println("Failed to connect to the Supabase database")
		return nil, err
	}

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Println("Error while fetching supabase version")
		return nil, err
	}

	log.Println("Connected to:", version)
	return conn, nil
}

// InitializeConfigStruct is the function that will load all the controlling params into Config struct
func InitializeConfigStruct(configFilePath string) (Configs, error) {
	configDataBytes, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Println("Error while reading the config file")
		return Configs{}, err
	}
	var configStruct Configs
	if err = yaml.Unmarshal(configDataBytes, &configStruct); err != nil {
		log.Println("Error unmarshalling config YAML into config struct")
		return Configs{}, err
	}

	return configStruct, nil
}
