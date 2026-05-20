import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMsg, setErrorMsg] = useState("");
  
  const navigate = useNavigate();

  const handleLogin = (e) => {
    e.preventDefault();
    setErrorMsg("");

    const credentials = { username, password };

    fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credentials)
    })
    .then(async (res) => {
      if (res.ok) {
        const data = await res.json();
        localStorage.setItem('token', data.token); 
        navigate('/sklep');
      } else {
        const errorMessage = await res.text();
        setErrorMsg(errorMessage); 
      }
    })
    .catch(err => {
      setErrorMsg("Blad polaczenia z serwerem.");
    });
  };

  return (
    <div style={{ maxWidth: '300px', margin: '50px auto', border: '1px solid #ccc', padding: '20px' }}>
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <div style={{ marginBottom: '10px' }}>
          <label>Login: </label>
          <input 
            type="text" 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
            placeholder="admin"
            style={{ width: '100%' }}
          />
        </div>
        <div style={{ marginBottom: '10px' }}>
          <label>Hasło: </label>
          <input 
            type="password" 
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
            placeholder="password"
            style={{ width: '100%' }}
          />
        </div>
        <button type="submit" style={{ width: '100%', padding: '5px' }}>Zaloguj</button>
      </form>
      {errorMsg && <p style={{ color: 'red', marginTop: '10px', fontWeight: 'bold' }}>{errorMsg}</p>}
    </div>
  );
};

export default Login;