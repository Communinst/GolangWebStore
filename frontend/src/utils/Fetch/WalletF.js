// Fetch all games
export const fetchWallet = async (token, userId) => {
    try {
        const response = await fetch(`/api/wallets/${userId}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            const errorMessage = errorData.detail || "Failed to register";
            throw new Error(errorMessage)
        }
        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching games:", error);
        return [];
    }
};

// Update a game
export const updateWalletBalance = async (token, userId, income) => {
    try {
        const response = await fetch(`/api/wallets/${userId}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({
                balance : income }),
        });

        if (!response.ok) {
            throw new Error("Failed to fetch games.");
        }

        const data = await response.json();
        return data || [];
    } catch (error) {
        console.error("Error updating game:", error.message);
        throw new Error(error.message);
    }
};
