package main

import (
	"os"
	"testing"

	"github.com/how8570/URL_Shortener/database"
)

func TestMain(m *testing.M) {
	database.ConnectDB()
	defer database.CloseDB()

	code := m.Run()
	os.Exit(code)
}
