// Fetch cart items for a user
export const fetchCart = async (token, userId) => {
    try {
        const response = await fetch(`/api/cart/${userId}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to fetch cart.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching cart:", error);
        throw error;
    }
};

// Add a game to the cart
export const addGameToCart = async (token, userId, gameId) => {
    try {
        const response = await fetch(`/api/cart/${userId}/games/${gameId}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to add game to cart.");
        }
    } catch (error) {
        console.error("Error adding game to cart:", error);
        throw error;
    }
};

// Remove a game from the cart
export const removeGameFromCart = async (token, userId, gameId) => {
    try {
        const response = await fetch(`/api/cart/${userId}/games/${gameId}`, {
            method: "DELETE",
            headers: {
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to remove game from cart.");
        }
    } catch (error) {
        console.error("Error removing game from cart:", error);
        throw error;
    }
};
