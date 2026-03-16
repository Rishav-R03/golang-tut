package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	salary float64
}

type Manager struct {
	ID        int
	Name      string
	Employees []*Employee
}

func (m *Manager) AddToTeam(e *Employee) {
	m.Employees = append(m.Employees, e)
}

func (m *Manager) GetEmployee() {
	for _, emp := range m.Employees {
		fmt.Printf("EmployeeID %d Name %s Salary: %.2f", emp.ID, emp.Name, emp.GetSalary())
	}
}
func (m *Manager) GiveRaise(percent float64) error {
	if percent <= 0 || m == nil {
		return fmt.Errorf("percentage raise cannot be negative : manager cannot be nil")
	}
	for _, emp := range m.Employees {

		if emp == nil {
			continue
		}
		emp.salary += (emp.salary * percent) / 100
	}
	return nil
}

func (e *Employee) GetSalary() float64 {
	return e.salary
}

func main() {
	emp1 := &Employee{ID: 1, Name: "rishav", salary: 4000000}
	emp2 := &Employee{ID: 2, Name: "rishav", salary: 4000000}
	emp3 := &Employee{ID: 3, Name: "rishav", salary: 4000000}
	emp4 := &Employee{ID: 4, Name: "rishav", salary: 4000000}

	employees := make([]*Employee, 0, 5)
	manager := &Manager{ID: 1, Name: "rishav", Employees: employees}
	manager.AddToTeam(emp1)
	manager.AddToTeam(emp2)
	manager.AddToTeam(emp3)
	manager.AddToTeam(emp4)

	oldSalary := emp1.GetSalary()
	fmt.Printf("old salary: %.2f\n", oldSalary)
	percent := 10.0
	err := manager.GiveRaise(float64(percent))
	if err != nil {
		fmt.Println("raise failed", err)
	} else {
		fmt.Printf("annual raise result, your raise for the financial cycle is %.2f", percent)
	}
	newSalary := emp1.GetSalary()
	fmt.Printf("\nnew salary: %.2f\n", newSalary)
}
