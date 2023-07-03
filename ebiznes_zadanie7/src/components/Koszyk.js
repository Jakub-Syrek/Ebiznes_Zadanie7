import React, { useState } from "react";

function Koszyk({cart, removeFromCart}) {
  return (
    <div>
      <h2>Koszyk</h2>
      {cart.length === 0 ? (
        <p>Koszyk jest pusty</p>
      ) : (
        cart.map(item => (
          <div key={item.id}>
            <h3>{item.name}</h3>
            <p>Cena: {item.price}</p>
            <button onClick={() => removeFromCart(item)}>Usuń</button>
          </div>
        ))
      )}
    </div>
  );
}

export default Koszyk;