export const fetchReviewsByGameId = async (token, gameId) => {
    try {
        const response = await fetch(`/api/reviews/game/${gameId}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            throw new Error("Failed to fetch reviews.");
        }

        const data = await response.json();
        return data;
    } catch (error) {
        console.error("Error fetching reviews:", error);
        return [];
    }
};

// Add a review for a game
export const addReview = async (token, gameId, userId, review) => {
    try {
        const response = await fetch(`/api/reviews/${userId}/games/${gameId}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            body: JSON.stringify({
                message : review.message,
                recommended : review.recommended,
                user_id : userId,
                game_id : gameId
            }),
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to add review.");
        }
    } catch (error) {
        console.error("Error adding review:", error);
        throw error;
    }
};

// Delete a review
export const deleteReview = async (token, reviewId) => {
    try {
        const response = await fetch(`/admin/reviews/${reviewId}`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || "Failed to delete review.");
        }
    } catch (error) {
        console.error("Error deleting review:", error);
        throw error;
    }
};