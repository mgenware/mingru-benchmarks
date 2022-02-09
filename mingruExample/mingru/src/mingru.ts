import * as mr from 'mingru';
import tables from './models/models.js';
import actions from './actions/actions.js';

(async () => {
  const dialect = new mr.MySQL();
  const builder = new mr.Builder(dialect, '../da/', {
    cleanOutDir: true, // Cleans build directory on each build.
  });

  await builder.buildAsync(async () => {
    await builder.buildActionsAsync(actions);
    await builder.buildCreateTableSQLFilesAsync(tables);
  });
})();
