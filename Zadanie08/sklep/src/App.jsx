import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Produkty from './Produkty'
import Platnosci from './Platnosci'
import Login from './Login';
import Register from './Register';
import UserPanel from './UserPanel';

const Sklep = () => {
  return (
    <div>
      <UserPanel />
      <h1>Sklep</h1>
      <Produkty />
      <Platnosci />
    </div>
  );
};

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/sklep" element={<Sklep />} />
        <Route path="*" element={<Navigate to="/login" replace />} />
      </Routes>
    </Router>
  );
}

export default App