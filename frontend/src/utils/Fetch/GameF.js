// Fetch all games
export const fetchGames = async () => {
    try {
        const response = await fetch("/api/games", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
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

// Fetch a game by ID
export const fetchGameById = async (id) => {
    try {
        const response = await fetch(`/api/games/${id}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
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

// Create a new game
export const createGame = async (gameData) => {
    try {
        const response = await fetch("/admin/games", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(gameData),
        });

        if (!response.ok) {
            throw new Error("Failed to create game.");
        }

        console.log("Game created successfully.");
    } catch (error) {
        console.error("Error creating game:", error);
    }
};

// Delete a game by ID
export const deleteGame = async (token, gameId) => {
    try {
        const response = await fetch(`/admin/games/${gameId}`, {
            method: 'DELETE',
            headers: {
                Authorization: `Bearer ${token}`,
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
                Authorization: `Bearer ${token}`,
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
