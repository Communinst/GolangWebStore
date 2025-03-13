import React, { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import { fetchGameById, updateGame } from '../utils/Fetch/GameF';
import { useNavigate } from "react-router-dom";
import { Alert } from "./Alert";
import { useAuth } from "../contexts/AuthContext";

const Game = () => {
    const { id } = useParams();
    const [game, setGame] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate();
    const [alertMessage, setAlertMessage] = useState("");

    useEffect(() => {
        const loadGame = async () => {
            const gameData = await fetchGameById(id);
            setGame(gameData);
            setIsLoading(false);
        };

        loadGame();
    }, [id]);

    const handleAlertClose = () => setAlertMessage("");

    const handleChange = (field, value) => {
        setGame({ ...game, [field]: value });
    };

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onloadend = () => {
                const imageBase64 = reader.result.split(",")[1];
                setGame((prevGame) => ({
                    ...prevGame,
                    image: imageBase64,
                }));
            };
            reader.onerror = () => setError("Error reading file.");
            reader.readAsDataURL(file);
        }
    };

    const fileInputRef = useRef(null);

    const handleButtonClick = () => {
        fileInputRef.current.click();
    };

    const { isAuthenticated, userType } = useAuth();
    if (!isAuthenticated)
        return <p className="loading">You need to sign in to see this page...</p>;

    const handleSubmit = async (e) => {
        setIsLoading(true);
        const token = localStorage.getItem("authToken");
        try {
            await updateGame(token, game);
            setAlertMessage("Game has been updated successfully");
        } catch (error) {
            setError(error.message);
        }
        setIsLoading(false);
    };

    if (isLoading) {
        return <p>Loading...</p>;
    }

    if (!game) {
        return <p>Game not found!</p>;
    }

    return (
        <div>
            {isAuthenticated && userType === "admin" ? (
                <h1>Edit Game</h1>
            ) : (
                <h1>View Game</h1>
            )}

            <div>
                <div>
                    <label>ID:</label>
                    <p>{game.id}</p>
                </div>
                <div>
                    <label>Name:</label>
                    <input
                        type="text"
                        value={game.name || ''}
                        onChange={(e) => handleChange('name', e.target.value)}
                        disabled={!isAuthenticated || userType !== "admin"}
                    />
                </div>
                <div>
                    <label>Genre:</label>
                    <input
                        type="text"
                        value={game.genre || ''}
                        onChange={(e) => handleChange('genre', e.target.value)}
                        disabled={!isAuthenticated || userType !== "admin"}
                    />
                </div>
                <div>
                    <label>Price:</label>
                    <input
                        type="text"
                        value={game.price || ''}
                        onChange={(e) => handleChange('price', e.target.value)}
                        disabled={!isAuthenticated || userType !== "admin"}
                    />
                </div>
                <div>
                    <input
                        type="file"
                        id="file"
                        accept="image/*"
                        onChange={handleImageChange}
                        ref={fileInputRef}
                        style={{ display: 'none' }}
                        disabled={!isAuthenticated || userType !== "admin"}
                    />
                    {(isAuthenticated && userType === "admin") &&
                        <button onClick={handleButtonClick} disabled={!isAuthenticated || userType !== "admin"}>
                            {game.image ? "Change photo" : "Choose a photo"}
                        </button>
                    }
                </div>
                {game.image && <img src={`data:image/png;base64,${game.image}`} alt="game" />}
                <div>
                    {(isAuthenticated && userType === "admin") &&
                        <button onClick={handleSubmit} disabled={!isAuthenticated || userType !== "admin"}>
                            Save Changes
                        </button>
                    }
                </div>
                {error && <p className="error">{error}</p>}
                {alertMessage !== "" && (
                    <Alert message={alertMessage} onClose={handleAlertClose} />
                )}
            </div>
        </div>
    );
};

export default Game;
