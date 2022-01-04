package cloud_db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"

	// we have to import the driver, but don't use it in our code
	// so we use the `_` symbol
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Courier struct {
	Name string `json: "full_name"`
	City string `json: "city"`
	WorkHours string `json: "workHours"`
}

// Exported functions/variables in Go are capitalizedbo 
func Query() []Courier {
	//Connect to AWS RDS service via the DSN (Database Source Name)
	//Note here we use a .env file and godotenv package to obfuscate the AWS RDS DSN
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	cloud_database := os.Getenv("AWS_RDS_DSN")

	db, err := sql.Open("pgx", cloud_database)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	// Setting timeout / idle connection limits
	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	// We use context-based cancellation for slow queries
	// create a parent context
	ctx := context.Background()
	// create a context from the parent context with a 300ms timeout
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)


	//querying multiple entries from the db
	rows, err := db.QueryContext(ctx, "SELECT full_name, city, \"workHours\" FROM courier_directory LIMIT 100")
	if err != nil {
		log.Fatalf("could not execute query %v", err)
	}
	
	// create a slice of couriers
	couriers := []Courier{}

	// iterate over the returned rows
	// we can go over to the next row by calling the `Next` method, which will 
	//return `false` if there are no more rows
	for rows.Next() {
		courier := Courier{}
		//create an instance of Courier and write the contents of current row to it
		if err := rows.Scan(&courier.Name, &courier.City, &courier.WorkHours); err != nil {
			log.Fatalf("could not can row: %v", err)
		}

		//append current courier instance to the slice of couriers
		couriers = append(couriers, courier)
	}
	// print the length, and all the couriers
	fmt.Printf("found %d couriers: %+v", len(couriers), couriers)
	return couriers
}

func Query_param(name string, city string, workHours string) []Courier{
	//Note here we use a .env file to obfuscate the AWS RDS DSN
	//Note here we use a .env file and godotenv package to obfuscate the AWS RDS DSN
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	cloud_database := os.Getenv("AWS_RDS_DSN")

	db, err := sql.Open("pgx", cloud_database)

	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	// Setting timeout / idle connection limits
	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	// We use context-based cancellation for slow queries
	// create a parent context
	ctx := context.Background()
	// create a context from the parent context with a 300ms timeout
	ctx, _ = context.WithTimeout(ctx, 300*time.Millisecond)

	//querying multiple entries from the db
	rows, err := db.QueryContext(ctx, "SELECT full_name, city, \"workHours\" FROM courier_directory WHERE full_name = $1 AND city = $2 AND \"workHours\" = $3 LIMIT 100", name, city, workHours)
	if err != nil {
		log.Fatalf("could not execute query %v", err)
	}
	
	// create s slice of couriers
	couriers := []Courier{}

	// iterate over the returned rows
	// we can go over to the next row by calling the `Next` method, which will 
	//return `false` if there are no more rows
	for rows.Next() {
		courier := Courier{}
		//create an instance of Courier and write the contents of current row to it
		if err := rows.Scan(&courier.Name, &courier.City, &courier.WorkHours); err != nil {
			log.Fatalf("could not can row: %v", err)
		}

		//append current courier instance to the slice of couriers
		couriers = append(couriers, courier)
	}
	// print the length, and all the couriers
	fmt.Printf("found %d couriers: %+v", len(couriers), couriers)
	return couriers
}

func Post_request(name string, city string, workHours string) {
	//Note here we use a .env file and godotenv package to obfuscate the AWS RDS DSN
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	cloud_database := os.Getenv("AWS_RDS_DSN")

	db, err := sql.Open("pgx", cloud_database)
	
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	// To verify the connection to our database instance, we can call the `Ping`
	// method. If no error is returned, we can assume a successful connection
	if err := db.Ping(); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	fmt.Println("database is reachable")

	// Setting timeout / idle connection limits
	// Maximum Idle Connections
	db.SetMaxIdleConns(5)
	// Maximum Open Connections
	db.SetMaxOpenConns(10)
	// Idle Connection Timeout
	db.SetConnMaxIdleTime(1 * time.Second)
	// Connection Lifetime
	db.SetConnMaxLifetime(30 * time.Second)

	newCourier := Courier{
		Name: name,
		City: city,
		WorkHours: workHours,
	}

	result, err := db.Exec("INSERT INTO courier_directory(full_name, city, \"workHours\") VALUES ($1, $2, $3)", newCourier.Name, newCourier.City, newCourier.WorkHours)
	if err != nil {
		log.Fatalf("could not insert row: %v", err)
	}

	// the `Result` type has special methods like `RowsAffected` which returns the
	// total number of affected rows reported by the database
	// In this case, it will tell us the number of rows that were inserted using
	// the above query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("could not get affected rows: %v", err)
	}
	// we can log how many rows were inserted
	fmt.Println("inserted", rowsAffected, "rows")
}