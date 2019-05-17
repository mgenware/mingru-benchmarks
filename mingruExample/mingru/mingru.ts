import * as mr from 'mingru';
// Import model actions
import employeeTA from './models/employeeTA';
import titleTA from './models/titleTA';

(async () => {
  const actions = [employeeTA, titleTA];
  const dialect = new mr.MySQL();
  // Build Go code to '../da/` directory
  await mr.build(actions, dialect, '../da/', { cleanBuild: true });
})();
