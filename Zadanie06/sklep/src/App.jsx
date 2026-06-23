import { BrowserRouter, Routes, Route, Link } from 'react-router-dom';
import Produkty from './Produkty';
import Platnosci from './Platnosci';
import Koszyk from './Koszyk';

function App() {
  return (
    <BrowserRouter>
      <div>
        <h1>Sklep</h1>
        <nav>
          <Link to="/" style={{ marginRight: '10px' }}>Produkty</Link>
          <Link to="/koszyk" style={{ marginRight: '10px' }}>Koszyk</Link>
          <Link to="/platnosci">Płatności</Link>
        </nav>

        <Routes>
          <Route path="/" element={<Produkty />} />
          <Route path="/koszyk" element={<Koszyk />} />
          <Route path="/platnosci" element={<Platnosci />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;