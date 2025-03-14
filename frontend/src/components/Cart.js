import React, { useEffect, useState } from 'react';
import { addGameToCart, getCartByUserID, removeGameFromCart } from '../utils/Fetch/CartF';
import { useAuth } from '../contexts/AuthContext';
import { updateWalletBalance } from '../utils/Fetch/WalletF';
import { postOwnershipByUserId } from '../utils/Fetch/OwnershipF';

const Cart = () => {
    const [cart, setCart] = useState([]);
    const [error, setError] = useState(null);
    const { isAuthenticated, userType, userId } = useAuth();

    const token = localStorage.getItem("authToken");


    useEffect(() => { 
        const fetchCart = async () => {
            try {
                const games = await getCartByUserID(token, userId);
                setCart(games);
            } catch (err) {
                setError(err.message);
            }
        };

        if (isAuthenticated && userId) {
            fetchCart();
        }
    }, [isAuthenticated, userId, token]);

    const handleRemoveGame = async (gameId) => {
        try {
            await removeGameFromCart(token, userId, gameId);
            const updatedCart = await getCartByUserID(token, userId);
            setCart(updatedCart);
        } catch (err) {
            setError(err.message);
        }
    };

    const handlePurchase = async (price) => {
        try {
            //await updateWalletBalance(token, userId, -price)
            for (const game of cart) {
                await postOwnershipByUserId(token, parseInt(userId), parseInt(game.game_id));
            }
            const updatedCart = await getCartByUserID(token, userId)
            setCart(updatedCart)
        } catch (err) {
            setError(err.message);
        }
    };

    const totalPrice = cart.reduce((sum, game) => sum + game.price, 0);

    if (!isAuthenticated) {
        return <p className="loading">You need to sign in to see your cart...</p>;
    }

    return (
        <div className="cart-container">
            <h2>Your Cart</h2>
            {error && <p className="error">{error}</p>}
            {cart.length === 0 ? (
                <p>Your cart is empty.</p>
            ) : (
                <>
                    <ul className="cart-list">
                        {cart.map((game) => (
                            <li key={game.game_id} className="cart-item">
                                <div className="cart-item-details">
                                    <div className="cart-item-name">
                                        <h3>{game.name}</h3>
                                    </div>
                                    <div className="cart-item-price">
                                        <p>Price: ${game.price.toFixed(2)}</p>
                                    </div>
                                </div>
                                <div className="cart-item-image">
                                    <img
                                        src={`/images/${game.name}.jpg`} // Ensure this path is correct
                                        alt={game.name}
                                    />
                                </div>
                                <button
                                    className="remove-button"
                                    onClick={() => handleRemoveGame(game.game_id)}
                                >
                                    Remove
                                </button>
                            </li>
                        ))}
                    </ul>
                    <div className="cart-summary">
                        <p>Total: ${totalPrice.toFixed(2)}</p>
                        <button className="buy-button" onClick={() => handlePurchase(totalPrice.toFixed(2))}>
                            Buy
                        </button>
                    </div>
                </>
            )}
        </div>
    );
};

export default Cart;
