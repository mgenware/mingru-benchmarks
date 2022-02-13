import * as mm from 'mingru-models';

class Dept extends mm.Table {
  id = mm.pk(mm.char(4)).setDBName('dept_no');
  name = mm.varChar(40).setDBName('dept_name');
}

export default mm.table(Dept, { dbName: 'departments' });
