import * as mm from 'mingru-models';
import t from '../models/employee.js';

export class EmployeeTA extends mm.ActionGroup {
  // Select an employee by ID.
  selectByID = mm.selectRow().by(t.id);
  // Select all employees.
  selectAll = mm.selectRows().pageMode().orderByAsc(t.hireDate);
  // Select all employees with a limit.
  selectAllWithLimit = mm.selectRows().pageMode().orderByAsc(t.hireDate);
  // Select paged employees.
  selectPaged = mm.selectRows().pageMode().orderByAsc(t.hireDate);
  // Select an employee birth date by ID.
  selectSig = mm.selectField(t.birthDate).by(t.id);

  // Update an employee name by ID.
  updateName = mm.updateOne().setParams(t.firstName, t.lastName).by(t.id);

  // Delete an employee by ID.
  deleteByID = mm.deleteOne().by(t.id);

  // Delete all employees by some criteria.
  deleteByBirthDate = mm.deleteSome().whereSQL(t.birthDate.isEqualToParam());

  // Insert a new employee.
  insertUser = mm.insertOne().setParams();
}

export default mm.actionGroup(t, EmployeeTA);
