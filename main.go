package main

import (
	"fmt"
	"time"
)

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

	for i, _ := range *workers {
		layout := "2006-01-02"
		var name string
		var birthdate string
		var hiredate string
		fmt.Println("Enter name")
		fmt.Scan(&name)

		fmt.Println("Enter birth date")
		fmt.Scan(&birthdate)

		fmt.Println("Enter hire date")
		fmt.Scan(&hiredate)

		(*workers)[i] = Worker{
			Name:      name,
			BirthDate: parseDateString(birthdate, layout),
			HireDate:  parseDateString(hiredate, layout)}
	}
	return workers
}

// Заполняет информацию о каждом работнике

/*
	func fillStruct(*Workers) {
		for i, _ := range *Workers {
			layout := "2006-01-02"
			var name string
			var birthdate string
			var hiredate string
			fmt.Println("Enter name")
			fmt.Scan(&name)

			fmt.Println("Enter birth date")
			fmt.Scan(&birthdate)

			fmt.Println("Enter hire date")
			fmt.Scan(&hiredate)

			(*Workers)[i] = Worker{
				Name:      name,
				BirthDate: parseDateString(birthdate, layout),
				HireDate:  parseDateString(hiredate, layout)}
		}
	}
*/
func parseDateString(dateString string, layout string) time.Time {
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}
	return parsedDate
}

func main() {
	Workers := createWorkersSlice()
	fmt.Println("Workers:")
	for _, worker := range Workers {
		fmt.Printf("Name: %s, BirthDate: %s, HireDate: %s\n", worker.Name, worker.BirthDate.Format("2006-01-02"), worker.HireDate.Format("2006-01-02"))
	}
}

/*
	var workersList []workers
	workersList = append(workersList, worker1, worker2, worker3, worker4)
	firedWorkers(workersList)
	salaryRate(workersList, 23000)
	dateCompare(workersList, comparedDate)
}

func parseDateString(dateString string, layout string) time.Time {
	parsedDate, err := time.Parse(layout, dateString)
	if err != nil {
		panic(err)
	}
	return parsedDate
}
func firedWorkers(workersList []workers) {
	for _, worker := range workersList {
		if !worker.IsEmployed {
			fmt.Println(worker.Name)
		}
	}
}
func salaryRate(workersList []workers, salary float64) {
	for _, worker := range workersList {
		if worker.Salary < salary {
			fmt.Println("Workers with salary higher than", salary, "are:", worker.Name)
		}
	}
}

func dateCompare(workersList []workers, comparedDate time.Time) {
	for _, worker := range workersList {
		if worker.HireDate.Before(comparedDate) {
			fmt.Println("Workers hired before", comparedDate, "are:", worker.Name)
		}
	}
}
*/
