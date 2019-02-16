 /******************************************************************************************
 * This code was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"github.com/mgenware/go-packagex/dbx"
)

// TableTypeUsers ...
type TableTypeUsers struct {
}

// Users ...
var Users = &TableTypeUsers{}

// ------------ Actions ------------

// UsersTableSelectUsersResult ...
type UsersTableSelectUsersResult struct {
	ID   uint64
	Name string
	Age  int
}

// SelectUsers ...
func (da *TableTypeUsers) SelectUsers(queryable dbx.Queryable, limit, offset int) ([]*UsersTableSelectUsersResult, error) {
	rows, err := queryable.Query("SELECT `id`, `name`, `age` FROM `users` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]*UsersTableSelectUsersResult, 0, limit)
	defer rows.Close()
	for rows.Next() {
		item := &UsersTableSelectUsersResult{}
		err = rows.Scan(&item.ID, &item.Name, &item.Age)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return result, nil
}