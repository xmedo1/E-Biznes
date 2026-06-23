import { useState, useEffect, useContext } from 'react';
import { CartContext } from './CartContext';
import axios from 'axios';

const Platnosci = () => {
  const { getCartTotal } = useContext(CartContext);
  const [amount, setAmount] = useState("");

  useEffect(() => {
    setAmount(getCartTotal().toString());
  }, [getCartTotal]);

  const sendPayment = (e) => {
    
    e.preventDefault();

    const dane = { amount: parseFloat(amount), data: new Date() };

    axios.post('http://localhost:8080/payments', dane)
    .then(res => {
      if(res.status === 201) alert("Wyslano platnosc");
    })
    .catch(err => console.error("Blad wysylania platnosci:", err));
  };

  return (
    <div style={{ border: '1px solid #ccc', padding: '10px', margin: '10px' }}>
      <h2>Formularz Płatności</h2>
      <form onSubmit={sendPayment}>
        <input 
          type="number" 
          value={amount}
          onChange={(e) => setAmount(e.target.value)}
        />
        <button type="submit">Zapłać</button>
      </form>
    </div>
  );
};

export default Platnosci;