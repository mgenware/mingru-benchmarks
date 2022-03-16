import * as mm from 'mingru-models';
import t from '../models/dept.js';

export class DeptTA extends mm.ActionGroup {
  // Selects paged departments.
  selectPage = mm.selectRows().pageMode().orderByAsc(t.id);
}

export default mm.actionGroup(t, DeptTA);
