import * as dd from 'dd-models';
import user from './user';

const userTA = dd.actions(user);
userTA.selectAll('Users', user.id, user.name, user.age).paginate();

export default userTA;
