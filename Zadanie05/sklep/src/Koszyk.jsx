import { useContext } from 'react';
import { CartContext } from './CartContext';

const Koszyk = () => {
  const { cart: cartItems } = useContext(CartContext);

  const sendCart = () => {
    fetch('http://localhost:8080/cart', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(cartItems)
    })
    .then(res => {
      if(res.ok) alert("Wysłano koszyk na serwer!");
    })
    .catch(err => console.error("Błąd wysyłania koszyka:", err));
  };

  return (
    <div style={{ border: '1px solid #ccc', padding: '10px', margin: '10px' }}>
      <h2>Koszyk</h2>
      {cartItems.length === 0 ? <p>Koszyk jest pusty.</p> : (
        <ul>
          {cartItems.map((item, index) => (
            <li key={index}>{item.name} - {item.price} PLN</li>
          ))}
        </ul>
      )}
      <button onClick={sendCart} disabled={cartItems.length === 0}>
        Wyślij koszyk do serwera
      </button>
    </div>
  );
};

export default Koszyk;