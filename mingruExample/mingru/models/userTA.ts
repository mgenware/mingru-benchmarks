import * as dd from "dd-models";
import user from "./user";

export class UserTA extends dd.TA {
  selectUsers = dd.selectAll(user.id, user.name, user.age).paginate();
}

export default dd.ta(user, UserTA);
