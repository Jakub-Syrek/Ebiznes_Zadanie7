import React, { useEffect, useState } from "react";
import axios from "axios";

function Produkty({ addToCart }) {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    axios
      .get("http://localhost:8080/api/products")
      .then((response) => {
        setProducts(response.data);
      })
      .catch((error) => {
        console.error(`There was an error retrieving the products: ${error}`);
      });
  }, []);

  return (
    <div>
      <h2>Produkty</h2>
      {products.map((product) => (
        <div key={product.id}>
          <h3>{product.name}</h3>
          <p>Cena: {product.price}</p>
          <button
            data-cy={`product-${product.id}-button`}
            onClick={() => addToCart(product)}
          >
            Dodaj do koszyka
          </button>
        </div>
      ))}
    </div>
  );
}

export default Produkty;
