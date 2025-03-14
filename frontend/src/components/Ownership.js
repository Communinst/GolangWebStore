import React, { useEffect, useState } from "react";
import { useAuth } from "../contexts/AuthContext";
import { fetchOwnershipsByUserID, deleteOwnershipByID } from '../utils/Fetch/OwnershipF';
import '../assets/styles/Ownership.css'; // Ensure you have your styles here

const Ownership = () => {
    const [ownerships, setOwnerships] = useState([]);
    const [error, setError] = useState(null);
    const { isAuthenticated, userType, userId } = useAuth();
    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const loadOwnerships = async () => {
            try {
                const data = await fetchOwnershipsByUserID(token, userId);
                setOwnerships(data);
            } catch (error) {
                setError(error.message);
            }
        };

        if (isAuthenticated) {
            loadOwnerships();
        }
    }, [token, userId, isAuthenticated]);

    const handleDelete = async (ownershipId) => {
        try {
            await deleteOwnershipByID(token, ownershipId);
            setOwnerships(ownerships.filter(ownership => ownership.ownership_id !== ownershipId));
        } catch (error) {
            setError(error.message);
        }
    };

    return (
        <div className="ownership-container">
            <h1>Your Owned Games</h1>
            {error && <p className="error">{error}</p>}
{ownerships.length > 0 ? (
                <ul>
                    {ownerships.map((ownership) => (
                        <li key={ownership.ownership_id}>
                            <p>Game ID: {ownership.game_id}</p>
                            <p>Minutes Spent: {ownership.minutes_spent}</p>
                            <p>Receipt Date: {new Date(ownership.receipt_date).toLocaleDateString()}</p>
                            {userType === "admin" && (
                                <button onClick={() => handleDelete(ownership.ownership_id)}>Delete</button>
                            )}
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No games owned.</p>
            )}
        </div>
    );
};

export default Ownership;