import * as mr from 'mingru';
// Import model actions
import userTA from './models/userTA';
import postTA from './models/postTA';

(async () => {
  const actions = [userTA, postTA];
  const dialect = new mr.MySQL();
  // Build Go code to '../da/` directory
  await mr.build(actions, dialect, '../da/');
})();
