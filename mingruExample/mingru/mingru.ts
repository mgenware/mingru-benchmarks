import * as mr from 'mingru';
// Import model actions
import userTA from './models/userTA';

(async () => {
  const actions = [userTA];
  const dialect = new mr.MySQL();
  // Build Go code to '../da/` directory
  await mr.build(actions, dialect, '../da/');
})();
