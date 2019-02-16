import * as dd from 'dd-models';
import post from './post';
import user from './user';

const postTA = dd.actions(post);
const userJoin = post.user_id.join(user);
postTA
  .selectAll(
    'Posts',
    post.id,
    post.title,
    post.content,
    post.user_id,
    userJoin.name,
    userJoin.age
  )
  .paginate();

export default postTA;
