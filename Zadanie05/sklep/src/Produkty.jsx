import { useState, useEffect, useContext } from 'react';
import { CartContext } from './CartContext';

const Produkty = () => {
  const [items, setItems] = useState([]);
  const { addToCart } = useContext(CartContext);

  useEffect(() => {
    fetch('http://localhost:8080/products')
      .then(response => response.json())
      .then(data => setItems(data))
      .catch(error => console.error('Błąd:', error));
  }, []);

  return (
    <div style={{ border: '1px solid #ccc', padding: '10px', margin: '10px' }}>
      <h2>Lista Produktów</h2>
      {items.length === 0 ? <p>Brak produktów.</p> : (
        <ul>
          {items.map(p => (
            <li key={p.id}>
              {p.name} - {p.price} PLN
              <button onClick={() => addToCart(p)} style={{ marginLeft: '10px' }}>Dodaj do koszyka</button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Produkty;