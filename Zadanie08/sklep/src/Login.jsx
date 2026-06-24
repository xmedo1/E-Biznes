import { useState, useEffect } from 'react';
import { useNavigate, Link, useSearchParams } from 'react-router-dom';

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [errorMsg, setErrorMsg] = useState("");
  
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();

  useEffect(() => {
    const token = searchParams.get("token");
    const error = searchParams.get("error");
    
    if (token) {
      localStorage.setItem('token', token); 
      navigate('/sklep');
    } else if (error) {
      setErrorMsg("Błąd logowania przez dostawcę: " + error);
    }
  }, [searchParams, navigate]);

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
        <button type="submit" style={{ width: '100%', padding: '5px', marginTop: '15px' }}>Zaloguj</button>
      </form>

      <div style={{ marginTop: '20px' }}>
        <a href="http://localhost:8080/auth/google/login" style={{ textDecoration: 'none' }}>
          <button type="button" style={{ width: '100%', padding: '8px', backgroundColor: '#db4437', color: 'white', border: 'none', cursor: 'pointer', fontWeight: 'bold' }}>
            Zaloguj przez Google
          </button>
        </a>
      </div>

      <div style={{ marginTop: '10px' }}>
        <a href="http://localhost:8080/auth/github/login" style={{ textDecoration: 'none' }}>
          <button type="button" style={{ width: '100%', padding: '8px', backgroundColor: '#110639', color: 'white', border: 'none', cursor: 'pointer', fontWeight: 'bold' }}>
            Zaloguj przez GitHub
          </button>
        </a>
      </div>

      {errorMsg && <p style={{ color: 'red', marginTop: '10px', fontWeight: 'bold' }}>{errorMsg}</p>}

      <div style={{ marginTop: '15px', fontSize: '14px' }}>
        <Link to="/register">Zarejestruj się</Link>
      </div>
    </div>
  );
};

export default Login;