package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"time"
)

const layout = "2006-01-02"

type Worker struct {
	Name      string
	BirthDate time.Time
	HireDate  time.Time
}

// запрашивает количество работников и создает слайс с структут с заданным количеством
func createWorkersSlice() []Worker {
	fmt.Println("Enter ammount of workers")
	var workersAmmount int
	fmt.Scan(&workersAmmount)

	workers := make([]Worker, workersAmmount)
	return workers
}

func fillStruct(worker *Worker) {

	for {
		var name string
		fmt.Println("Enter name")
		fmt.Scan(&name)

		// Проверка на пустую строку
		if len(name) == 0 {
			fmt.Println("Name cannot be empty, enter name again")
			continue
		}
		worker.Name = name
		break
	}
	for {

		var birthdate string
		fmt.Println("Enter birth date")
		fmt.Scan(&birthdate)

		parsedDate, err := parseDateString(birthdate, layout)
		if err != nil {
			fmt.Println("Invalid date format, enter birth date again ")
			continue
		}
		worker.BirthDate = parsedDate
		break
	}

	for {
		var hiredate string
		fmt.Println("Enter hire date")
		fmt.Scan(&hiredate)
		parsedDate, err := parseDateString(hiredate, layout)
		if err != nil {
			fmt.Println("Invalid date format, enter hire date again ")
			continue
		}
		worker.HireDate = parsedDate
		break
	}
}

func parseDateString(dateString string, layout string) (time.Time, error) {
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}
func dateCompare(workersList []Worker, comparedDate time.Time) {
	for _, worker := range workersList {
		if worker.HireDate.Before(comparedDate) {
			fmt.Println("Workers hired before", comparedDate, "are:", worker.Name)
		}
	}
}
func createWorkersTable(db *sql.DB) error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS workers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        birthdate TEXT,
        hiredate TEXT
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("error creating workers table: %v", err)
	}
	return nil
}

func addWorkerToDatabase(db *sql.DB, workers []Worker) {
	for _, worker := range workers {
		_, err := db.Exec("INSERT INTO workers (name, birthdate, hiredate) VALUES (?, ?, ?)",
			worker.Name, worker.BirthDate, worker.HireDate)

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	Workers := createWorkersSlice()
	for i := 0; i < len(Workers); i++ {
		fillStruct(&Workers[i])
	}

	fmt.Println("Workers:")
	for _, worker := range Workers {
		fmt.Printf("Name: %s, BirthDate: %s, HireDate: %s\n", worker.Name, worker.BirthDate.Format(layout), worker.HireDate.Format(layout))
	}
	comparedDate, err := parseDateString("2007-01-01", layout)
	if err != nil {
		fmt.Println("Error parsing compared date:", err)
		return
	}
	dateCompare(Workers, comparedDate)

	db, err := sql.Open("sqlite", "database_workers")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	createWorkersTable(db)
	addWorkerToDatabase(db, Workers)
}
