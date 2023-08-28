package main

import (
	"encoding/json"
	"fmt"
)

type Manager struct {
	ManagerFirstName  string  `json:"manager_first_name"`
	ManagerLastName   string  `json:"manager_last_name"`
	ManagerEmployeeID string  `json:"manager_employee_id"`
	ManagerSalary     float64 `json:"manager_salary,omitempty"`
}

type Manager1 struct {
	ManagerFirstName  string
	ManagerLastName   string
	ManagerEmployeeID string
	ManagerSalary     float64
}

func main() {
	manager := Manager{
		ManagerFirstName:  "John",
		ManagerLastName:   "Doe",
		ManagerEmployeeID: "12345",
		ManagerSalary:     75000.0,
	}

	manager1 := Manager1{
		ManagerFirstName:  "John",
		ManagerLastName:   "Doe",
		ManagerEmployeeID: "12345",
		ManagerSalary:     75000.0,
	}

	fmt.Println(manager)
	fmt.Println(manager1)

	jsonData, err := json.Marshal(manager)
	data, err := json.Marshal(manager1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
	fmt.Println(string(data))
}
