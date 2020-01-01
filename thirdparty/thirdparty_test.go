// thirdparty_test.go contains unit tests for package thirdparty

package thirdparty

import (
	"testing"

	"github.com/golang/mock/gomock"
	"gitlab.com/mayankkapoor/go-rest-mux-app/mocks"
)

func TestGetAnnualSalary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDoer := mocks.NewMockDoer(ctrl)

	// Test data
	testEmployeeID := "2"

	// Stub the GetEmployeeSalary function to return dummy value of "100" when
	// function is called with employeeID of "2"
	mockDoer.
		EXPECT().
		GetEmployeeSalary(gomock.Eq(testEmployeeID)).
		Return("100").
		AnyTimes()

	// Call the GetAnnualSalary function and check if it returns expected value
	// of $100/week * 52weeks = 5200
	expected := 5200
	actual := GetAnnualSalary(mockDoer, testEmployeeID)

	if actual != expected {
		t.Errorf("Annual Salary was incorrect, got: %d, expected: %d.", actual, expected)
	}

}
