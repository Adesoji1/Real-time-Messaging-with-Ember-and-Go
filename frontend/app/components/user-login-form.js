import Component from '@glimmer/component';
import { action } from '@ember/object';

export default class UserLoginFormComponent extends Component {
  username = '';
  password = '';
  errorMessage = null;

  @action
  async loginUser(event) {
    event.preventDefault();

    try {
      const response = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          username: this.username,
          password: this.password,
        }),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('token', data.token);
        alert('Login successful!');
        // Optionally, navigate to a different page after successful login
      } else {
        const data = await response.json();
        this.errorMessage = `Login failed: ${data.error}`;
        alert(this.errorMessage);
      }
    } catch (error) {
      console.error('Error logging in:', error);
      this.errorMessage = 'An error occurred while logging in.';
      alert(this.errorMessage);
    }
  }
}
