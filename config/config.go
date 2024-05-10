package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Inirialize config file
func Config(key string) string {
	// Load config file
	err := godotenv.Load(".env")

	// Show error message if there is an error handling the config file
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

var ATSYS = Config("ATS")
var KurikulumTB = Config("KURIKULUM_TB")
var MatakuliahTB = Config("MATAKULIAH_TB")
