// SignUpForm.js
import React, { useState } from 'react';
import '../CSS/signup.css';

function SignUpForm({ onSignUp }) {
  const [newUsername, setNewUsername] = useState('');
  const [newPassword, setNewPassword] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSignUp(newUsername, newPassword);
  };

  return (
    <div>
      <h2>Sign Up</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input type="text" value={newUsername} onChange={(e) => setNewUsername(e.target.value)} />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" value={newPassword} onChange={(e) => setNewPassword(e.target.value)} />
        </div>
        <button type="submit">Sign Up</button>
      </form>
    </div>
  );
}

export default SignUpForm;
