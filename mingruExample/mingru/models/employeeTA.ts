import * as dd from 'dd-models';
import t from './employee';

export class EmployeeTA extends dd.TA {
  // Select an employee by ID.
  selectByID = dd.select().byID();
  // Select all employees.
  selectAll = dd.selectRows().paginate();
}

export default dd.ta(t, EmployeeTA);
