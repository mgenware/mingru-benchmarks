import * as mm from 'mingru-models';
import t from '../models/employee';

export class EmployeeTA extends mm.TableActions {
  // Select an employee by ID.
  selectByID = mm.select().byID();
  // Select all employees.
  selectAll = mm.selectPage().orderByAsc(t.hireDate);
  // Select all employees with a limit.
  selectAllWithLimit = mm.selectPage().orderByAsc(t.hireDate);
  // Select paged employees.
  selectPaged = mm.selectPage().orderByAsc(t.hireDate);
  // Select an employee birth date by ID.
  selectSig = mm.selectField(t.birthDate).byID();

  // Update an employee name by ID.
  updateName = mm
    .updateOne()
    .setInputs(t.firstName, t.lastName)
    .byID();

  // Delete an employee by ID.
  deleteByID = mm.deleteOne().byID();

  // Delete all employees by some criteria.
  deleteByBirthDate = mm.deleteSome().where(t.birthDate.isEqualToInput());

  // Insert a new employee.
  insertUser = mm.insertOne().setInputs();
}

export default mm.tableActions(t, EmployeeTA);
