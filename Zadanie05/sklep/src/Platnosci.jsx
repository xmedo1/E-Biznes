import { useState } from 'react';

const Platnosci = () => {
  const [amount, setAmount] = useState("");

  const sendPayment = (e) => {

    const dane = { amount: parseFloat(amount), data: new Date() };

    fetch('http://localhost:8080/payments', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(dane)
    })
    .then(res => {
      if(res.ok) alert("Wyslano platnosc");
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