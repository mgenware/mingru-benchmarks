import * as mm from 'mingru-models';

class EmploymentHistoryItem extends mm.Table {
  employee_id = mm.pk(mm.int()).setDBName('emp_no');
  department_id = mm.pk(mm.char(4)).setDBName('dept_no');
  fromDate = mm.date();
  toDate = mm.date();
}

export default mm.table(EmploymentHistoryItem, { dbName: 'dept_emp' });
