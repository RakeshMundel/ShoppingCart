import React, { useContext } from 'react';
import { ShopContext } from '../Context/ShopContext';

const Cart = () => {
  const { cartItems, all_product, removeFromCart } = useContext(ShopContext);

  const cartProducts = all_product.filter(product => cartItems[product.ID]);

  const totalAmount = cartProducts.reduce((total, product) => {
    return total + product.price * cartItems[product.ID];
  }, 0);

  return (
    <div style={{ padding: '2rem' }}>
      <h1>Your Cart</h1>
      {cartProducts.length === 0 ? (
        <p>Your cart is empty.</p>
      ) : (
        <>
          {cartProducts.map((product) => (
            <div key={product.ID} style={{ marginBottom: '1rem' }}>
              <h3>{product.name}</h3>
              <p>Quantity: {cartItems[product.ID]}</p>
              <p>Price per item: ₹{product.price}</p>
              <p>Total: ₹{cartItems[product.ID] * product.price}</p>
              <button onClick={() => removeFromCart(product.ID)}>Remove One</button>
              <hr />
            </div>
          ))}
          <h2>Total: ₹{totalAmount}</h2>
          <button onClick={() => alert("Checkout not implemented yet")}>
            Proceed to Checkout
          </button>
        </>
      )}
    </div>
  );
};

export default Cart;
