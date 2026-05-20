import { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg] = useState("");
  const [isError, setIsError] = useState(false);
  
  const navigate = useNavigate();

  const handleRegister = (e) => {
    e.preventDefault();
    setMsg(""); 
    setIsError(false);

    const credentials = { username, password };

    fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(credentials)
    })
    .then(async (res) => {
      if (res.ok) {
        setIsError(false);
        setMsg("Zarejestrowano pomyślnie");
      } else {
        const errorMessage = await res.text();
        setIsError(true);
        setMsg(errorMessage); 
      }
    })
    .catch(err => {
      setIsError(true);
      setMsg("Błąd połączenia z serwerem.");
    });
  };

  return (
    <div style={{ maxWidth: '300px', margin: '50px auto', border: '1px solid #ccc', padding: '20px' }}>
      <h2>Rejestracja</h2>
      <form onSubmit={handleRegister}>
        <div style={{ marginBottom: '10px' }}>
          <label>Login: </label>
          <input 
            type="text" 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
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
            style={{ width: '100%' }}
          />
        </div>
        <button type="submit" style={{ width: '100%', padding: '5px', marginTop: '15px' }}>Zarejestruj</button>
      </form>
      
      {msg && <p style={{ color: isError ? 'red' : 'green', marginTop: '10px', fontWeight: 'bold' }}>{msg}</p>}

      <div style={{ marginTop: '15px', fontSize: '14px' }}>
        <Link to="/login">Zaloguj się</Link>
      </div>
    </div>
  );
};

export default Register;