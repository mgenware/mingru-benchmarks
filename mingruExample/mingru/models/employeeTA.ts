import * as dd from 'dd-models';
import t from './employee';

const employeeColumns = [
  t.id,
  t.firstName,
  t.lastName,
  t.gender,
  t.birthDate,
  t.hireDate,
];

export class EmployeeTA extends dd.TA {
  // Select an employee by ID
  selectByID = dd.select(...employeeColumns).byID();
  // Select all employees
  selectAll = dd.selectAll(...employeeColumns).paginate();
}

export default dd.ta(t, EmployeeTA);
