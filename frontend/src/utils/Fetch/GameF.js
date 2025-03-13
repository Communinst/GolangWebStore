// Fetch all games
export const fetchGames = async (token) => {
    try {
        const response = await fetch("/api/games", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to fetch games.");
        }

        const data = await response.json();
        return data || [];
    } catch (error) {
        console.error("Error fetching games:", error);
        return [];
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
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to add game to cart.");
        }
    } catch (error) {
        console.error("Error adding game to cart:", error);
        throw error;
    }
};


// Fetch a game by ID
export const fetchGameById = async (token, id) => {
    try {
        const response = await fetch(`/api/games/${id}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to fetch game.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching game:", error);
        return null;
    }
};
// Fetch a game by name
export const fetchGameByName = async (token, name) => {
    try {
        const response = await fetch(`/api/games/name/${name}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to fetch game.");
        }

        const data = await response.json();
        return data; // This will be a single game object
    } catch (error) {
        console.error("Error fetching game by name:", error);
        throw error; // Re-throw the error to handle it in the component
    }
};


// Create a new game
export const createGame = async (token, publisherId, name, price, description) => {
    try {
        const response = await fetch("/admin/games", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({
                publisher_id: publisherId,
                name: name,
                price: price,
                description: description,
            }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.detail || "Failed to create game.");
        }

        console.log("Game created successfully.");
    } catch (error) {
        console.error("Error creating game:", error);
        throw error;
    }
};

// Delete a game by ID
export const deleteGame = async (token, gameId) => {
    try {
        const response = await fetch(`/admin/games/${gameId}`, {
            method: 'DELETE',
            headers: {
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.detail || "Failed to delete game";
            throw new Error(errorMessage);
        }
    } catch (error) {
        console.error("Error deleting game:", error.message);
        throw new Error(error.message);
    }
};

// Update a game
export const updateGame = async (token, game) => {
    try {
        const response = await fetch(`/admin/games/${game.id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify(game),
        });

        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.detail || "Failed to update game";
            throw new Error(errorMessage);
        }
    } catch (error) {
        console.error("Error updating game:", error.message);
        throw new Error(error.message);
    }
};
