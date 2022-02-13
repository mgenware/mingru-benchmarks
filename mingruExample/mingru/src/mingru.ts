import * as mr from 'mingru';
import source from './source.js';

(async () => {
  const dialect = new mr.MySQL();
  const builder = new mr.Builder(dialect, '../da/', {
    cleanOutDir: true, // Cleans build directory on each build.
    createTableSQL: true,
  });

  await builder.build(source);
})();
