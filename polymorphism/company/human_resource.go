type Employee interface {
	GetDetails() string
	GetEmployeeSalary() int
}

type Manager struct {
	Name        string
	Age         int
	Designation string
	Salary      int
}

func (mgr Manager) GetDetails() string {
	return mgr.Name + " " + mgr.Age
}

func (mgr Manager) GetEmployeeSalary() int {
	return mgr.Salary
}