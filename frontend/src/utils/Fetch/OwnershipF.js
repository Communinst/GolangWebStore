export const fetchOwnershipsByUserID = async (token, userId) => {
    try {
        const response = await fetch(`/api/ownerships/user/${userId}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to fetch ownerships.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching ownerships:", error);
        return [];
    }
};

// Delete an ownership by ID
export const deleteOwnershipByID = async (token, ownershipId) => {
    try {
        const response = await fetch(`/admin/ownerships/${ownershipId}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to delete ownership.");
        }
    } catch (error) {
        console.error("Error deleting ownership:", error);
        throw error;
    }
};

export const postOwnershipByUserId = async (token, userId, gameId) => {
    try {
        const response = await fetch(`/api/ownerships/${userId}/games/${gameId}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({
                minutes_spent : 0,
            }),
        });

        if (!response.ok) {
            throw new Error("Failed to fetch ownerships.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching ownerships:", error);
        return [];
    }
};