import React, { useEffect, useState } from "react";
import { fetchGames, createGame, fetchGameByName, deleteGame } from "../utils/Fetch/GameF";
import { useNavigate } from "react-router-dom";
import { Alert } from "./Alert";
import { useAuth } from "../contexts/AuthContext";
import { addGenre } from "../utils/Fetch/GenreF";
import '../assets/styles/Store.css'; // Ensure you have a CSS file for styling

const Store = () => {
    const [games, setGames] = useState([]);
    const [alertMessage, setAlertMessage] = useState("");
    const [genreName, setGenreName] = useState("");
    const [genreDescription, setGenreDescription] = useState("");
    const [isLoading, setIsLoading] = useState(false);
    const [success, setSuccess] = useState("");
    const [error, setError] = useState("");
    const [gameName, setGameName] = useState("");
    const [gamePrice, setGamePrice] = useState("");
    const [gameDescription, setGameDescription] = useState("");
    const [gamePublisherId, setGamePublisherId] = useState("");
    const [searchTerm, setSearchTerm] = useState("");
    const [showAddGameModal, setShowAddGameModal] = useState(false);
    const [showAddGenreModal, setShowAddGenreModal] = useState(false);
    const [showRemoveGameModal, setShowRemoveGameModal] = useState(false);
    const [gameIdToRemove, setGameIdToRemove] = useState("");

    const navigate = useNavigate();
    const { isAuthenticated, userType } = useAuth();
    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const loadGames = async () => {
            try {
                const data = await fetchGames(token);
                console.log("Fetched games:", data); // Debugging log
                setGames(data);
            } catch (error) {
                setError(error.message);
            }
            setIsLoading(false);
        };

        setIsLoading(true);
        loadGames();
    }, [token]);

    const handleAddGenre = async () => {
        try {
            await addGenre(token, genreName, genreDescription);
            setSuccess("Genre added successfully!");
            setGenreName("");
            setGenreDescription("");
            setShowAddGenreModal(false); // Close the modal
        } catch (err) {
            setError(err.message);
        }
    };

    const handleCreateGame = async () => {
        try {
            await createGame(token, parseInt(gamePublisherId), gameName, parseInt(gamePrice), gameDescription);
            setSuccess("Game added successfully!");
            setShowAddGameModal(false); // Close the modal
            const data = await fetchGames(token);
            setGames(data);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleRemoveGame = async () => {
        try {
            await deleteGame(token, parseInt(gameIdToRemove));
            setSuccess("Game deleted successfully!");
            setShowRemoveGameModal(false); // Close the modal
            const data = await fetchGames(token);
            setGames(data);
        } catch (err) {
            setError(err.message);
        }
    };

    const handleSearch = async () => {
        try {
            const data = await fetchGameByName(token, searchTerm);
            if (data) {
                setGames([data]); // Wrap the single object in an array
            } else {
                setGames([]); // No game found
            }
            setError(null); // Clear any previous errors
        } catch (err) {
            setError(err.message);
        }
    };

    const handleGameClick = (gameId) => {
        console.log("Navigating to game ID:", gameId); // Debugging log
        navigate(`/game/${gameId}`);
    };

    const handleGoToCart = () => {
        navigate('/cart');
    };

    if (isLoading) {
        return <p>Loading...</p>;
    }

    if (!isAuthenticated) {
        return <p className="loading">You need to sign in to see this page...</p>;
    }

    return (
        <div className="store-container">
            <div className="header">
                <div className="logo">Logo and name</div>
                <div className="search-bar">
                    <input
                        type="text"
                        placeholder="Search by game name"
                        value={searchTerm}
                        onChange={(e) => setSearchTerm(e.target.value)}
                    />
                    <button onClick={handleSearch}>Search</button>
                </div>
                {(isAuthenticated && userType === "admin") && (
                    <div className="admin-buttons">
                        <button onClick={() => setShowAddGameModal(true)}>Add Game</button>
                        <button onClick={() => setShowAddGenreModal(true)}>Add Genre</button>
                        <button onClick={() => setShowRemoveGameModal(true)}>Remove Game</button>
                    </div>
                )}
                <button onClick={handleGoToCart} className="cart-button">
                    Go to Cart
                </button>
            </div>
            <div className="game-list">
                {games.length === 0 ? (
                    <p>No games available.</p>
                ) : (
                    games.map((game) => (
                        <div
                            key={game.game_id}
                            className="game-record"
                            onClick={() => handleGameClick(game.game_id)}
                            style={{ cursor: 'pointer' }}
                        >
                            <div className="game-details">
                                <div className="game-name">
                                    <h2>{game.name}</h2>
                                </div>
                                <div className="publisher-id">
                                    <p>Publisher ID: {game.publisher_id}</p>
                                </div>
                                <div className="game-price">
                                    <p>Price: {game.price}</p>
                                </div>
                            </div>
                            <div className="game-image">
                                <img
                                    src={`/images/${game.name}.jpg`} // Ensure this path is correct
                                    alt={game.name}
                                />
                            </div>
                        </div>
                    ))
                )}
            </div>
            {showAddGameModal && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close" onClick={() => setShowAddGameModal(false)}>&times;</span>
                        <h2>Add New Game</h2>
                        <input
                            type="text"
                            placeholder="Game Name"
                            value={gameName}
                            onChange={(e) => setGameName(e.target.value)}
                        />
                        <input
                            type="number"
                            placeholder="Game Price"
                            value={gamePrice}
                            onChange={(e) => setGamePrice(e.target.value)}
                        />
                        <input
                            type="number"
                            placeholder="Publisher ID"
                            value={gamePublisherId}
                            onChange={(e) => setGamePublisherId(e.target.value)}
                        />
                        <textarea
                            placeholder="Game Description"
                            value={gameDescription}
                            onChange={(e) => setGameDescription(e.target.value)}
                        ></textarea>
                        <div className="modal-buttons">
                            <button onClick={handleCreateGame}>Add Game</button>
                            <button onClick={() => setShowAddGameModal(false)}>Cancel</button>
                        </div>
                        {error && <p className="error">{error}</p>}
                        {success && <p className="success">{success}</p>}
                    </div>
                </div>
            )}
            {showAddGenreModal && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close" onClick={() => setShowAddGenreModal(false)}>&times;</span>
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
                        <div className="modal-buttons">
                            <button onClick={handleAddGenre}>Add Genre</button>
                            <button onClick={() => setShowAddGenreModal(false)}>Cancel</button>
                        </div>
                        {error && <p className="error">{error}</p>}
                        {success && <p className="success">{success}</p>}
                    </div>
                </div>
            )}
            {showRemoveGameModal && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close" onClick={() => setShowRemoveGameModal(false)}>&times;</span>
                        <h2>Remove Game</h2>
                        <input
                            type="number"
                            placeholder="Game ID"
                            value={gameIdToRemove}
                            onChange={(e) => setGameIdToRemove(e.target.value)}
                        />
                        <div className="modal-buttons">
                            <button onClick={handleRemoveGame}>Remove Game</button>
                            <button onClick={() => setShowRemoveGameModal(false)}>Cancel</button>
                        </div>
                        {error && <p className="error">{error}</p>}
                        {success && <p className="success">{success}</p>}
                    </div>
                </div>
            )}
            {alertMessage !== "" && (
                <Alert message={alertMessage} onClose={() => setAlertMessage("")} />
            )}
        </div>
    );

};

export default Store;
