 /******************************************************************************************
 * This code was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"time"

	"github.com/mgenware/go-packagex/v5/dbx"
)

// TableTypeEmployee ...
type TableTypeEmployee struct {
}

// Employee ...
var Employee = &TableTypeEmployee{}

// ------------ Actions ------------

// EmployeeTableSelectAllResult ...
type EmployeeTableSelectAllResult struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
	BirthDate time.Time
	HireDate  time.Time
}

// SelectAll ...
func (da *TableTypeEmployee) SelectAll(queryable dbx.Queryable, limit int, offset int, max int) ([]*EmployeeTableSelectAllResult, int, error) {
	rows, err := queryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, 0, err
	}
	result := make([]*EmployeeTableSelectAllResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			item := &EmployeeTableSelectAllResult{}
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, 0, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, 0, err
	}
	return result, itemCounter, nil
}

// EmployeeTableSelectByIDResult ...
type EmployeeTableSelectByIDResult struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
	BirthDate time.Time
	HireDate  time.Time
}

// SelectByID ...
func (da *TableTypeEmployee) SelectByID(queryable dbx.Queryable, id int) (*EmployeeTableSelectByIDResult, error) {
	result := &EmployeeTableSelectByIDResult{}
	err := queryable.QueryRow("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` WHERE `emp_no` = ?", id).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Gender, &result.BirthDate, &result.HireDate)
	if err != nil {
		return nil, err
	}
	return result, nil
}
