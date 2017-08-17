package database

import (
	"log"
)

var (
	// Database info
	databases Database
)

const (
	// TypeFile is FileInJSON
	TypeFile string = "File"
	// TypeMySQL is MySQL
	TypeMySQL string = "MySQL"
)

type Database struct {
	// Database type
	Type string
	// MySQL info if used
	MySQL MySQLDatabase
	// Bolt info if used
	File FileDatabase
}

// MySQLInfo is the details for the database connection
type MySQLDatabase struct {
	Username  string
	Password  string
	Name      string
	Hostname  string
	Port      int
	Parameter string
}

// FileInfo is the details for the database connection
type FileDatabase struct {
	Path string
}

// Connect to the database
func Connect(d Database) {
	// Store the config
	databases = d

	switch d.Type {
	case TypeMySQL:
		//Connect to MySQL
	case TypeFile:
		// Connect to File
	default:
		log.Println("No registered database in config")
	}
}

// ReadConfig returns the database information
func ReadConfig() Database {
	return databases
}
