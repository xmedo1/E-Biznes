import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const UserPanel = () => {
  const [user, setUser] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (!token) return;

    fetch('http://localhost:8080/user/me', {
      headers: { 'Authorization': token }
    })
    .then(res => res.json())
    .then(data => setUser(data))
    .catch(() => localStorage.removeItem('token'));
  }, []);

  const handleLogout = () => {
    localStorage.removeItem('token');
    navigate('/login');
  };

  if (!user) return null;

  return (
    <div style={{ padding: '10px', background: '#1a1a1a', marginBottom: '20px' }}>
      <span>Zalogowano jako: <strong>{user.username}</strong></span>
      <button onClick={handleLogout} style={{ marginLeft: '10px' }}>Wyloguj</button>
    </div>
  );
};

export default UserPanel;