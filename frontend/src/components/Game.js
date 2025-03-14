import React, { useEffect, useState, useRef } from "react";
import { useParams } from "react-router-dom";
import { fetchGameById, updateGame } from '../utils/Fetch/GameF';
import { addGameToCart } from '../utils/Fetch/CartF'; // Import the addGameToCart function
import { useNavigate } from "react-router-dom";
import { Alert } from "./Alert";
import { useAuth } from "../contexts/AuthContext";
import { fetchReviewsByGameId, addReview, deleteReview } from '../utils/Fetch/ReviewF'
import '../assets/styles/Game.css'; // Ensure you have a CSS file for styling

const Game = () => {
    const { id } = useParams();
    const [game, setGame] = useState(null);
    const [reviews, setReviews] = useState([]);
    const [newReview, setNewReview] = useState({ message: "", recommended: true });
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const [alertMessage, setAlertMessage] = useState("");
    const navigate = useNavigate();
    const fileInputRef = useRef(null);

    const { isAuthenticated, userType, userId } = useAuth();
    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const loadGameAndReviews = async () => {
          try {
            const gameData = await fetchGameById(token, id);
            setGame(gameData);
    
            const reviewsData = await fetchReviewsByGameId(token, id);
            setReviews(reviewsData);
          } catch (error) {
            setError(error.message);
          }
          setIsLoading(false);
        };
    
        loadGameAndReviews();
      }, [id, token]);
    
      const handleReviewChange = (e) => {
        const { name, value } = e.target;
        setNewReview({ ...newReview, [name]: value });
      };
    
      const handleReviewSubmit = async () => {
        try {
          await addReview(token, id, userId, newReview);
          const updatedReviews = await fetchReviewsByGameId(token, id);
          setReviews(updatedReviews);
          setNewReview({ message: "", recommended: true });
        } catch (error) {
          setError(error.message);
        }
      };
    
      const handleDeleteReview = async (reviewId) => {
        try {
          await deleteReview(token, reviewId);
          setReviews(reviews.filter(review => review.review_id !== reviewId));
        } catch (error) {
          setError(error.message);
        }
      };

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

    const handleButtonClick = () => {
        fileInputRef.current.click();
    };

    const handleSubmit = async (e) => {
        setIsLoading(true);
        try {
            await updateGame(token, game);
            setAlertMessage("Game has been updated successfully");
        } catch (error) {
            setError(error.message);
        }
        setIsLoading(false);
    };

    const handleAddToCart = async () => {
        try {
            await addGameToCart(token, userId, game.game_id);
            setAlertMessage("Game added to cart successfully!");
        } catch (err) {
            setError(err.message);
        }
    };

    if (isLoading) {
        return <p>Loading...</p>;
    }

    if (!game) {
        return <p>Game not found!</p>;
    }

    return (
        <div className="game-container">
            {isAuthenticated && (
              <div>
                <h2>Add a Review</h2>
                <textarea
                  name="message"
                  value={newReview.message}
                  onChange={handleReviewChange}
                  placeholder="Write your review here"
                />
                <label>
                  <input
                    type="checkbox"
                    name="recommended"
                    checked={newReview.recommended}
                    onChange={() => setNewReview({ ...newReview, recommended: !newReview.recommended })}
                  />
                  Recommended
                </label>
                <button onClick={handleReviewSubmit}>Submit Review</button>
              </div>
            )}

            <h2>Reviews</h2>
            {reviews && reviews.length > 0 ? (
                <ul>
                    {reviews.map((review) => (
                        <li key={review.review_id}>
                            <p>{review.message}</p>
                            <p>Recommended: {review.recommended ? "Yes" : "No"}</p>
                            {userType === "admin" && (
                                <button onClick={() => handleDeleteReview(review.review_id)}>Delete</button>
                            )}
                        </li>
                    ))}
                </ul>
            ) : (
                <p>No reviews available.</p>
            )}
          
            {error && <p className="error">{error}</p>}
            <div className="game-header">
                <h1>{game.name}</h1>
                <div className="game-image">
                    <img src={`/images/${game.name}.jpg`} alt={game.name} />
                </div>
            </div>
            <div className="game-details">
                <div>
                    <label>Price:</label>
                    <p>{game.price}</p>
                </div>
                <div>
                    <label>Publisher ID:</label>
                    <p>{game.publisher_id}</p>
                </div>
                <div>
                    <label>Rating:</label>
                    <p>{game.rating}</p>
                </div>
                <div>
                    <label>Description:</label>
                    <p>{game.description}</p>
                </div>
                <div>
                    <label>Reviews:</label>
                    <p>{game.reviews}</p>
                </div>
            </div>
            {isAuthenticated && userType === "admin" && (
                <div className="admin-controls">
                    <input
                        type="file"
                        id="file"
                        accept="image/*"
                        onChange={handleImageChange}
                        ref={fileInputRef}
                        style={{ display: 'none' }}
                    />
                    <button onClick={handleButtonClick}>Change Photo</button>
                    <button onClick={handleSubmit}>Save Changes</button>
                </div>
            )}
            {isAuthenticated && (
                <button onClick={handleAddToCart} className="add-to-cart-button">
                    Add to Cart
                </button>
            )}
            {error && <p className="error">{error}</p>}
            {alertMessage !== "" && (
                <Alert message={alertMessage} onClose={handleAlertClose} />
            )}
        </div>
    );
};

export default Game;
