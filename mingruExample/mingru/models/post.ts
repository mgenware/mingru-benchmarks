import * as dd from 'dd-models';
import user from './user';

class Post extends dd.Table {
  id = dd.pk();
  title = dd.varChar(255);
  content = dd.text();
  user_id = user.id;
}

export default dd.table(Post, 'posts');
