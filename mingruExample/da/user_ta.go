 /******************************************************************************************
 * This code was automatically generated by mingru (https://github.com/mgenware/mingru)
 * Do not edit this file manually, your changes will be overwritten.
 ******************************************************************************************/

package da

import (
	"github.com/mgenware/go-packagex/dbx"
)

// TableTypeUser ...
type TableTypeUser struct {
}

// User ...
var User = &TableTypeUser{}

// ------------ Actions ------------

// UserTableSelectUsersResult ...
type UserTableSelectUsersResult struct {
	ID   uint64
	Name string
	Age  int
}

// SelectUsers ...
func (da *TableTypeUser) SelectUsers(queryable dbx.Queryable, limit, offset int) ([]*UserTableSelectUsersResult, error) {
	rows, err := queryable.Query("SELECT `id`, `name`, `age` FROM `user` LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	result := make([]*UserTableSelectUsersResult, 0, limit)
	defer rows.Close()
	for rows.Next() {
		item := &UserTableSelectUsersResult{}
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
