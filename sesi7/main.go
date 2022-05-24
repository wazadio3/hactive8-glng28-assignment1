package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println()

}

const (
	DB_HOST = "localhost"
	DB_POST = "5432"
	DB_USER = "juni"
	DB_PASS = "admin"
	DB_NAME = "hacktive8"
)

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	//create employee
	// emp := Employee{
	// 	Email:    "admin@noobeeid.com",
	// 	FullName: "Noobeeid",
	// 	Age:      22,
	// 	Division: "Project Manager",
	// }

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// ============== Update Employee ====================
	// isSucces, err := updateEmployee(db, 1, "Juni")
	// if err != nil {
	// 	fmt.Println("Error :", err.Error())
	// 	return
	// }
	// fmt.Println(isSucces)

	// ============== Delete Employee =====================
	// isSucces, err := deleteEmployee(db, 1)
	// if err != nil {
	// 	fmt.Println("Error :", err.Error())
	// 	return
	// }
	// fmt.Println(isSucces)

	// ============== Print All Employee ==================
	employees, err := getAllEmployees(db)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for _, employee := range *employees {
		employee.Print()
	}

	// fmt.Println("====== Get Employee by id 4 ======")
	// employee, err := getEmployeeById(db, 4)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// employee.Print()
}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_POST, DB_USER, DB_PASS, DB_NAME)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	return db, nil
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID, &employee.FullName,
			&employee.Email, &employee.Age, &employee.Division,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil

}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
		INSERT INTO employees(full_name, email, age, division)
		VALUES($1, $2, $3, $4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()

}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
		SELECT id, full_name, email, age, division
		FROM employees
		WHERE id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	var emp Employee

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func updateEmployee(db *sql.DB, id int8, full_name string) (bool, error) {
	query := `
		UPDATE employees
		SET full_name = $1
		WHERE id = $2
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_ = stmt.QueryRow(full_name, id)

	return true, nil
}

func deleteEmployee(db *sql.DB, id int8) (bool, error) {
	query := `
		DELETE FROM employees
		WHERE id = $1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_ = stmt.QueryRow(id)

	return true, nil
}
