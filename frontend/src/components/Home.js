import React from "react";
import { Link } from "react-router-dom";
import { FaHome, FaShoppingCart } from "react-icons/fa";

const Home = () => {
  return (
    <div className="home-container">
      <h1>Welcome to the Game Store</h1>
      <p>Explore our collection of games and more!</p>
      <Link to="/store" className="home-link">
        <FaShoppingCart /> Go to Store
      </Link>
    </div>
  );
};

export default Home;
