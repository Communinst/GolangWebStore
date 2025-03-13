import React, { useEffect, useState } from "react";
import { fetchGames, deleteGame, createGame } from "../utils/Fetch/GameF";
import { useNavigate } from "react-router-dom";
import { Alert } from "./Alert";
import { useAuth } from "../contexts/AuthContext";
import { addGenre } from "../utils/Fetch/GenreF";

const Store = () => {
    const [games, setGames] = useState([]);
    const navigate = useNavigate();
    const [alertMessage, setAlertMessage] = useState("");
    const [genreName, setGenreName] = useState("");
    const [genreDescription, setGenreDescription] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [success, setSuccess] = useState("");
    const [error, setError] = useState("");
    const [gameIdToDelete, setGameIdToDelete] = useState("");

    useEffect(() => {
        const loadGames = async () => {
            const data = await fetchGames();
            setGames(data);
        };
        setIsLoading(true);
        loadGames();
        setIsLoading(false);
    }, []);

    const handleCreateGame = async () => {
        await createGame({ name: "New Game", genre: "Default", price: 0 });
        const loadGames = async () => {
            const data = await fetchGames();
            setGames(data);
        };
        setIsLoading(true);
        loadGames();
        setIsLoading(false);
    };

    const handleAddGenre = async () => {
        const token = localStorage.getItem("authToken");
        try {
            await addGenre(token, genreName, genreDescription);
            setSuccess("Genre added successfully!");
            setGenreName("");
            setGenreDescription("");
        } catch (err) {
            setError(err.message);
        }
    };

    const handleDeleteGame = async () => {
        const token = localStorage.getItem("authToken");
        try {
            await deleteGame(token, gameIdToDelete);
            setGames(games.filter(game => game.id !== gameIdToDelete));
            setAlertMessage("Game deleted successfully!");
            setGameIdToDelete("");
        } catch (e) {
            console.error("Error deleting game:", e.message);
            setAlertMessage(e.message);
        }
    };

    if (isLoading) {
        return <p>Loading...</p>;
    }

    const handleGameClick = (gameId) => {
        navigate(`/game/${gameId}`);
    };

    const { isAuthenticated, userType } = useAuth();
    if (!isAuthenticated)
        return <p className="loading">You need to sign in to see this page...</p>;

    return (
        <div className="store-container">
            {(isAuthenticated && userType === "admin") && (
                <div>
                    <button
                        className="add-game-button"
                        onClick={handleCreateGame}
                    >
                        Add Default Game
                    </button>
                    <div>
                        <input
                            type="text"
                            placeholder="Game ID to Delete"
                            value={gameIdToDelete}
                            onChange={(e) => setGameIdToDelete(e.target.value)}
                        />
                        <button onClick={handleDeleteGame}>Delete Game</button>
                    </div>
                    <div>
                        <h2>Add New Genre</h2>
                        <input
                            type="text"
                            placeholder="Genre Name"
                            value={genreName}
                            onChange={(e) => setGenreName(e.target.value)}
                        />
                        <input
                            type="text"
                            placeholder="Genre Description"
                            value={genreDescription}
                            onChange={(e) => setGenreDescription(e.target.value)}
                        />
                        <button onClick={handleAddGenre}>Add Genre</button>
                        {error && <p className="error">{error}</p>}
                        {success && <p className="success">{success}</p>}
                    </div>
                </div>
            )}
            <div>
                <h1>Games</h1>
                {games.length === 0 ? (
                    <p>No games available.</p>
                ) : (
                    <div className="game-list">
                        {games.map((game) => (
                            <div key={game.id} className="game-card">
                                <h2>
                                    <button
                                        className="game-id-button"
                                        onClick={() => handleGameClick(game.id)}
                                    >
                                        {game.name}
                                    </button>
                                </h2>
                                <div className="game-info">
                                    <img
                                        src={`data:image/png;base64,${game.image}`}
                                        alt="No game image"
                                        className="game-image"
                                    />
                                    <p><strong>Genre:</strong> {game.genre}</p>
                                    <p><strong>Price:</strong> {game.price}</p>
                                    <p><strong>ID:</strong> {game.id}</p>
                                </div>
                            </div>
                        ))}
                    </div>
                )}
            </div>
            {alertMessage !== "" && (
                <Alert message={alertMessage} onClose={() => setAlertMessage("")} />
            )}
        </div>
    );
};

export default Store;
