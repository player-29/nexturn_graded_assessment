// A company wants to manage its employees' data in memory. Each employee has an ID,
// name, age, and department. You need to build a small application that performs the
// following:
// 1. Add Employee: Accept input for employee details and store them in an array of
// structs. Validate the input:
// o ID must be unique.
// o Age should be greater than 18. If validation fails, return custom error
// messages.
// 2. Search Employee: Search for an employee by ID or name using conditions.
// Return the details if found, or return an error if not found.
// 3. List Employees by Department: Use loops to filter and display all employees in
// a given department.
// 4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and
// display the count of employees in that department.


// Bonus:
// Refactor the repetitive code using functions, and add error handling for invalid
// operations like searching for a non-existent employee.










package main

import(
	"fmt"
	"errors"
	"strings"
)
type Employee struct {
	ID         string
	Name       string
	Age        int
	Department string
}

const (
	HR = "HR"
	IT = "IT"
)

var employees []Employee

func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees by Department")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEmployee()
		case 2:
			searchEmployee()
		case 3:
			listEmployeesByDepartment()
		case 4:
			countEmployeesByDepartment()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}


func addEmployee() {
	var id, name, department string
	var age int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&age)
	fmt.Print("Enter Employee Department: ")
	fmt.Scan(&department)

	if err := validateEmployeeInput(id, age); err != nil {
		fmt.Println("Error:", err)
		return
	}

	employee := Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
	}
	employees = append(employees, employee)
	fmt.Println("Employee added successfully!")
}

func validateEmployeeInput(id string, age int) error {
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}
	return nil
}

func searchEmployee() {
	var searchKey string
	fmt.Print("Enter Employee ID or Name to search: ")
	fmt.Scan(&searchKey)

	employee, err := findEmployee(searchKey)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Employee Found: %+v\n", *employee)
}

func findEmployee(searchKey string) (*Employee, error) {
	for _, emp := range employees {
		if emp.ID == searchKey || strings.EqualFold(emp.Name, searchKey) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func listEmployeesByDepartment() {
	var department string
	fmt.Print("Enter Department Name: ")
	fmt.Scan(&department)

	fmt.Println("Employees in", department, "department:")
	found := false
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			fmt.Printf("%+v\n", emp)
			found = true
		}
	}
	if !found {
		fmt.Println("No employees found in this department.")
	}
}

func countEmployeesByDepartment() {
	departments := []string{HR, IT}
	for _, dept := range departments {
		count := 0
		for _, emp := range employees {
			if strings.EqualFold(emp.Department, dept) {
				count++
			}
		}
		fmt.Printf("Number of employees in %s: %d\n", dept, count)
	}
}
