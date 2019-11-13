import * as mm from 'mingru-models';
import t from '../models/dept';

export class DeptTA extends mm.TableActions {
  // Selects paged departments.
  selectPage = mm.selectPage().orderByAsc(t.id);
}

export default mm.tableActions(t, DeptTA);
