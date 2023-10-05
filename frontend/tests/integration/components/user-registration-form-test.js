import { module, test } from 'qunit';
import { setupRenderingTest } from 'frontend/tests/helpers';
import { render } from '@ember/test-helpers';
import { hbs } from 'ember-cli-htmlbars';

module('Integration | Component | user-registration-form', function (hooks) {
  setupRenderingTest(hooks);

  test('it renders', async function (assert) {
    // Set any properties with this.set('myProperty', 'value');
    // Handle any actions with this.set('myAction', function(val) { ... });

    await render(hbs`<UserRegistrationForm />`);

    assert.dom().hasText('');

    // Template block usage:
    await render(hbs`
      <UserRegistrationForm>
        template block text
      </UserRegistrationForm>
    `);

    assert.dom().hasText('template block text');
  });
});
