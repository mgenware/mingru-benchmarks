import * as dd from 'dd-models';
import employee from './employee';

class Title extends dd.Table {
  empNo = employee.id;
  title = dd.varChar(50);
  fromDate = dd.date();
  toDate = dd.date();
}

export default dd.table(Title, 'titles');
