// Function to add a game to the cart
export const addGameToCart = async(token, userId, gameId) => {
    try {
        const response = await fetch(`/api/carts/${userId}/games/${gameId}`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to add game to cart');
        }

        const data = await response.json();
        console.log(data.message);
    } catch (error) {
        console.error('Error adding game to cart:', error);
    }
}

// Function to get the cart by user ID
export const getCartByUserID = async (token, userId) => {
    try {
        const response = await fetch(`/api/carts/${userId}/`, {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to retrieve cart');
        }
        const responseData = await response.json();
        const games = Array.isArray(responseData) ? responseData : [];

        console.log('Cart:', games);
        return games;
    } catch (error) {
        console.error('Error retrieving cart:', error);
    }
}

// Function to remove a game from the cart
export const removeGameFromCart = async (token, userId, gameId) => {
    try {
        const response = await fetch(`/api/carts/${userId}/games/${gameId}`, {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json',
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Failed to remove game from cart');
        }

        //const data = await response.json();
        console.log(data.message);
    } catch (error) {
        console.error('Error removing game from cart:', error);
    }
}