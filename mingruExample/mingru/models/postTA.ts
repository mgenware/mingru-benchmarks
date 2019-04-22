import * as dd from "dd-models";
import post from "./post";
import user from "./user";

const userJoin = post.user_id.join(user);
export class PostTA extends dd.TA {
  selectPosts = dd
    .selectAll(
      post.id,
      post.title,
      post.content,
      post.user_id,
      userJoin.name,
      userJoin.age
    )
    .paginate();
}

export default dd.ta(post, PostTA);
