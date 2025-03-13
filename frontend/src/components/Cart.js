import React, { useEffect, useState } from "react";
import { useAuth } from "../contexts/AuthContext";
import { fetchCart, addGameToCart, removeGameFromCart } from "../utils/Fetch/CartF";
import { FaShoppingCart } from "react-icons/fa";
import '../assets/styles/Cart.css'; // Ensure you have a CSS file for styling

const Cart = () => {
    const [cartItems, setCartItems] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [error, setError] = useState(null);
    const { isAuthenticated, user } = useAuth();
    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const loadCart = async () => {
            if (isAuthenticated && user) {
                setIsLoading(true);
                try {
                    const data = await fetchCart(token, user.user_id);
                    setCartItems(data);
                } catch (err) {
                    setError(err.message);
                }
                setIsLoading(false);
            }
        };

        loadCart();
    }, [isAuthenticated, user, token]);

    const handleAddToCart = async (gameId) => {
        try {
            await addGameToCart(token, user.user_id, gameId);
            const updatedCart = await fetchCart(token, user.user_id);
            setCartItems(updatedCart);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleRemoveFromCart = async (gameId) => {
        try {
            await removeGameFromCart(token, user.user_id, gameId);
            const updatedCart = await fetchCart(token, user.user_id);
            setCartItems(updatedCart);
        } catch (err) {
            setError(err.message);
        }
    };

    if (!isAuthenticated) {
        return null;
    }

    return (
        <div className="cart-container">
            <div className="cart-icon" onClick={() => console.log('Cart clicked')}>
                <FaShoppingCart />
                <span className="cart-count">{cartItems.length}</span>
            </div>
            {cartItems.length > 0 && (
                <div className="cart-dropdown">
                    {cartItems.map((item) => (
                        <div key={item.game_id} className="cart-item">
                            <span>{item.name}</span>
                            <button onClick={() => handleRemoveFromCart(item.game_id)}>Remove</button>
                        </div>
                    ))}
                </div>
            )}
            {error && <p className="error">{error}</p>}
        </div>
    );
};

export default Cart;
