import { module, test } from 'qunit';
import { setupRenderingTest } from 'frontend/tests/helpers';
import { render } from '@ember/test-helpers';
import { hbs } from 'ember-cli-htmlbars';

module('Integration | Component | user-login-form', function (hooks) {
  setupRenderingTest(hooks);

  test('it renders', async function (assert) {
    // Set any properties with this.set('myProperty', 'value');
    // Handle any actions with this.set('myAction', function(val) { ... });

    await render(hbs`<UserLoginForm />`);

    assert.dom().hasText('');

    // Template block usage:
    await render(hbs`
      <UserLoginForm>
        template block text
      </UserLoginForm>
    `);

    assert.dom().hasText('template block text');
  });
});
