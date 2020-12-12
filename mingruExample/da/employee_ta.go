/******************************************************************************************
 * This file was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"fmt"
	"time"

	"github.com/mgenware/mingru-go-lib"
)

// TableTypeEmployee ...
type TableTypeEmployee struct {
}

// Employee ...
var Employee = &TableTypeEmployee{}

// ------------ Actions ------------

// DeleteByBirthDate ...
func (da *TableTypeEmployee) DeleteByBirthDate(queryable mingru.Queryable, birthDate time.Time) (int, error) {
	result, err := queryable.Exec("DELETE FROM `employees` WHERE `birth_date` = ?", birthDate)
	return mingru.GetRowsAffectedIntWithError(result, err)
}

// DeleteByID ...
func (da *TableTypeEmployee) DeleteByID(queryable mingru.Queryable, id int) error {
	result, err := queryable.Exec("DELETE FROM `employees` WHERE `emp_no` = ?", id)
	return mingru.CheckOneRowAffectedWithError(result, err)
}

// InsertUser ...
func (da *TableTypeEmployee) InsertUser(queryable mingru.Queryable, id int, firstName string, lastName string, gender string, birthDate time.Time, hireDate time.Time) error {
	_, err := queryable.Exec("INSERT INTO `employees` (`emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date`) VALUES (?, ?, ?, ?, ?, ?)", id, firstName, lastName, gender, birthDate, hireDate)
	return err
}

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
func (da *TableTypeEmployee) SelectAll(queryable mingru.Queryable, page int, pageSize int) ([]*EmployeeTableSelectAllResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := queryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
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
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}

// EmployeeTableSelectAllWithLimitResult ...
type EmployeeTableSelectAllWithLimitResult struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
	BirthDate time.Time
	HireDate  time.Time
}

// SelectAllWithLimit ...
func (da *TableTypeEmployee) SelectAllWithLimit(queryable mingru.Queryable, page int, pageSize int) ([]*EmployeeTableSelectAllWithLimitResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := queryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]*EmployeeTableSelectAllWithLimitResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			item := &EmployeeTableSelectAllWithLimitResult{}
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
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
func (da *TableTypeEmployee) SelectByID(queryable mingru.Queryable, id int) (*EmployeeTableSelectByIDResult, error) {
	result := &EmployeeTableSelectByIDResult{}
	err := queryable.QueryRow("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` WHERE `emp_no` = ?", id).Scan(&result.ID, &result.FirstName, &result.LastName, &result.Gender, &result.BirthDate, &result.HireDate)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EmployeeTableSelectPagedResult ...
type EmployeeTableSelectPagedResult struct {
	ID        int
	FirstName string
	LastName  string
	Gender    string
	BirthDate time.Time
	HireDate  time.Time
}

// SelectPaged ...
func (da *TableTypeEmployee) SelectPaged(queryable mingru.Queryable, page int, pageSize int) ([]*EmployeeTableSelectPagedResult, bool, error) {
	if page <= 0 {
		err := fmt.Errorf("Invalid page %v", page)
		return nil, false, err
	}
	if pageSize <= 0 {
		err := fmt.Errorf("Invalid page size %v", pageSize)
		return nil, false, err
	}
	limit := pageSize + 1
	offset := (page - 1) * pageSize
	max := pageSize
	rows, err := queryable.Query("SELECT `emp_no`, `first_name`, `last_name`, `gender`, `birth_date`, `hire_date` FROM `employees` ORDER BY `hire_date` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, false, err
	}
	result := make([]*EmployeeTableSelectPagedResult, 0, limit)
	itemCounter := 0
	defer rows.Close()
	for rows.Next() {
		itemCounter++
		if itemCounter <= max {
			item := &EmployeeTableSelectPagedResult{}
			err = rows.Scan(&item.ID, &item.FirstName, &item.LastName, &item.Gender, &item.BirthDate, &item.HireDate)
			if err != nil {
				return nil, false, err
			}
			result = append(result, item)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, false, err
	}
	return result, itemCounter > len(result), nil
}

// SelectSig ...
func (da *TableTypeEmployee) SelectSig(queryable mingru.Queryable, id int) (time.Time, error) {
	var result time.Time
	err := queryable.QueryRow("SELECT `birth_date` FROM `employees` WHERE `emp_no` = ?", id).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateName ...
func (da *TableTypeEmployee) UpdateName(queryable mingru.Queryable, id int, firstName string, lastName string) error {
	result, err := queryable.Exec("UPDATE `employees` SET `first_name` = ?, `last_name` = ? WHERE `emp_no` = ?", firstName, lastName, id)
	return mingru.CheckOneRowAffectedWithError(result, err)
}
