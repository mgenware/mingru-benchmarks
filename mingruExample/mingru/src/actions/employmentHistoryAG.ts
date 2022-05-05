import * as mm from 'mingru-models';
import t from '../models/employmentHistory.js';
import department from '../models/dept.js';
import employee from '../models/employee.js';

export class EmploymentHistoryAG extends mm.ActionGroup {
  selectAllHistory = mm
    .selectRows(
      t.fromDate,
      t.toDate,
      t.employee_id.join(employee).firstName,
      t.employee_id.join(employee).lastName,
      t.department_id.join(department).name,
    )
    .pageMode()
    .orderByAsc(t.employee_id)
    .orderByAsc(t.department_id);
}

export default mm.actionGroup(t, EmploymentHistoryAG);
