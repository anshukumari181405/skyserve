import React, { useState } from 'react';

const ForgotPasswordForm = ({ onResetPassword }) => {
  const [email, setEmail] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    onResetPassword(email);
    
    setEmail('');
  };

  return (
    <div>
      <h2>Forgot Password</h2>
      <form onSubmit={handleSubmit}>
        <div className="input-group">
          <input type="email" placeholder="Email Address" value={email} onChange={(e) => setEmail(e.target.value)} />
        </div>
        <div className="input-group">
          <button type="submit">Reset Password</button>
        </div>
      </form>
    </div>
  );
};

export default ForgotPasswordForm;
