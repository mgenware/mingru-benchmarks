import * as mm from 'mingru-models';
import t from '../models/employmentHistory';
import department from '../models/dept';
import employee from '../models/employee';

export class EmploymentHistoryTA extends mm.TableActions {
  selectAllHistory = mm
    .selectPage(
      t.fromDate,
      t.toDate,
      t.employee_id.join(employee).firstName,
      t.employee_id.join(employee).lastName,
      t.department_id.join(department).name,
    )
    .orderByAsc(t.employee_id)
    .orderByAsc(t.department_id);
}

export default mm.tableActions(t, EmploymentHistoryTA);
