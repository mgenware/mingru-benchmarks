import * as dd from 'dd-models';
import t from './title';
import employee from './employee';

export class TitleTA extends dd.TA {
  // Select all titles
  selectAll = dd
    .selectAll(
      t.title,
      t.empNo,
      t.empNo.join(employee).firstName,
      t.empNo.join(employee).lastName,
      t.empNo.join(employee).birthDate,
      t.empNo.join(employee).hireDate,
    )
    .paginate();
}

export default dd.ta(t, TitleTA);
