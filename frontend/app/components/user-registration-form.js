import Component from '@glimmer/component';
import { action } from '@ember/object';

export default class UserRegistrationFormComponent extends Component {
  username = '';
  password = '';
  errorMessage = null;

  @action
  async registerUser(event) {
    event.preventDefault();

    try {
      const response = await fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      });

      if (response.ok) {
        alert('Registration successful!');
        // Optionally, navigate to a different page after successful registration
      } else {
        const data = await response.json();
        this.errorMessage = `Registration failed: ${data.error}`;
        alert(this.errorMessage);
      }
    } catch (error) {
      console.error('Error registering user:', error);
      this.errorMessage = 'An error occurred while registering.';
      alert(this.errorMessage);
    }
  }
}
