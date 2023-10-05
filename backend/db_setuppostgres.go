// package main

// import (
// 	"database/sql"
// 	"log"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// func setupDatabase() {
// 	var err error
// 	// Open a database connection
// 	// db, err := sql.Open("postgres", "user=myuser dbname=mydb sslmode=disable")
// 	connStr := "user=adesoji password=123456 dbname=postgres sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)
// 	// db, err := sql.Open("postgres", "user=adesoji password=123456 dbname=Sprotweb_development sslmode=disable")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Read the SQL script file using os.ReadFile instead of ioutil.ReadFile
// 	script, err := os.ReadFile("create_tables.sql")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Execute the script
// 	_, err = db.Exec(string(script))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"os"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123456"
// 	dbname   = "postgres"
// )

// func setupDatabase() {
// 	// Use the same connection string format from the working script
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	// Open a database connection
// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatalf("Error opening database: %v", err)
// 	}
// 	defer db.Close()

// 	// Check the database connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Error pinging database: %v", err)
// 	}

// 	// Read the SQL script file that creates the schema
// 	script, err := os.ReadFile("create_tables.sql")
// 	if err != nil {
// 		log.Fatalf("Error reading SQL script: %v", err)
// 	}

// 	// Execute the script
// 	_, err = db.Exec(string(script))
// 	if err != nil {
// 		log.Fatalf("Error executing SQL script: %v", err)
// 	}

// 	fmt.Println("Database setup completed successfully!")
// }

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"os"
// 	"strconv"

// 	_ "github.com/lib/pq"
// )

// func getDBConnection() (*sql.DB, error) {
// 	host := os.Getenv("DB_HOST")
// 	port, _ := strconv.Atoi(os.Getenv("DB_PORT")) // Error handling omitted for brevity
// 	user := os.Getenv("DB_USER")
// 	password := os.Getenv("DB_PASSWORD")
// 	dbname := os.Getenv("DB_NAME")

// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		host, port, user, password, dbname)

// 	return sql.Open("postgres", psqlInfo)
// }

// func setupDatabase() error {
// 	db, err := getDBConnection()
// 	if err != nil {
// 		return fmt.Errorf("Error opening database: %v", err)
// 	}
// 	defer db.Close()

// 	if err = db.Ping(); err != nil {
// 		return fmt.Errorf("Error pinging database: %v", err)
// 	}

// 	script, err := os.ReadFile("create_tables.sql")
// 	if err != nil {
// 		return fmt.Errorf("Error reading SQL script: %v", err)
// 	}

// 	if _, err = db.Exec(string(script)); err != nil {
// 		return fmt.Errorf("Error executing SQL script: %v", err)
// 	}

// 	fmt.Println("Database setup completed successfully!")
// 	return nil
// }


// export DB_HOST='localhost'
// export DB_PORT='5432'
// export DB_USER='postgres'
// export DB_PASSWORD='123456'
// export DB_NAME='postgres'





package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return sql.Open("postgres", psqlInfo)
}

func setupDatabase() (*sql.DB, error) {
	db, err := getDBConnection()
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	script, err := os.ReadFile("create_tables.sql")
	if err != nil {
		return nil, fmt.Errorf("error reading SQL script: %v", err)
	}

	if _, err = db.Exec(string(script)); err != nil {
		return nil, fmt.Errorf("error executing SQL script: %v", err)
	}

	fmt.Println("Database setup completed successfully!")
	return db, nil
     

}
