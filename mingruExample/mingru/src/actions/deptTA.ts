import * as mm from 'mingru-models';
import t from '../models/dept';

export class DeptTA extends mm.TableActions {
  // Selects paged departments.
  selectPage = mm.selectRows().pageMode().orderByAsc(t.id);
}

export default mm.tableActions(t, DeptTA);
