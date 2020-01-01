// thirdparty.go connects to third party service for getting info

package thirdparty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// Doer interface defines mandatory functions
type Doer interface {
	// GetEmployee(employeeID string) Employee
	// GetEmployees() []Employee
	GetEmployeeSalary(employeeID string) string
}

// Employee struct stores all metadata for an Employee
type Employee struct {
	ID             string `json:"id"`
	EmployeeName   string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge    string `json:"employee_age"`
	ProfileImage   string `json:"profile_image"`
}

// GetEmployeeSalary returns the salary of the employee as a string.
// It calls a third party API with the employeeID, parses the response and
// returns the value found in employee_salary
// In the real world, this function would be implemented elsewhere. The third
// party library would be imported. Also, as Golang interfaces are
// implicit and we don't need to tag every implementation with interface name, we
// can simply import the third party library that implements this function and it
// will automatically be tagged with this interface.
func GetEmployeeSalary(employeeID string) string {

	// Get employee details from third party api
	var empBytes []byte // variable to hold http response body
	baseURL := "http://dummy.restapiexample.com/api/v1/employee/"
	resp, err := http.Get(baseURL + employeeID)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Error(err)
		}
		empBytes = bodyBytes
	}

	// parse json into new Employee
	var emp Employee
	json.Unmarshal([]byte(empBytes), &emp)

	return emp.EmployeeSalary
}

// GetAnnualSalary returns employees annual salary by multiplying employee
// weekly salary returned by third party api by 52 weeks. We assume the salary
// returned by third party API is weekly salary.
// To unit test this function, we'll need to use gomock to stub the GetEmployeeSalary
// function response and check if the GetAnnualSalary function is multiplying
// the salary by 52 or not
func GetAnnualSalary(d Doer, employeeID string) int {
	weeklySalaryString := d.GetEmployeeSalary(employeeID)

	weeklySalaryInt, err := strconv.Atoi(weeklySalaryString)
	if err != nil {
		log.Fatal(err)
	}
	return weeklySalaryInt * 52
}
