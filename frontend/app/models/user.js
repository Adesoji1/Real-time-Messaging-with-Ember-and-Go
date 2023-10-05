// import Model from '@ember-data/model';

// export default class UserModel extends Model {}


import Model, { attr } from '@ember-data/model';

export default class UserModel extends Model {
  @attr('string') username;
  @attr('string') email;
  @attr('string') password; // NOTE: Only use this if you're sending the password to the backend for registration or authentication. Do not store it.
  // ... Add other attributes as needed ...
}

