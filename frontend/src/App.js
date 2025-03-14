import React from "react";
import Login from "./components/Login/Login.js";
import Register from "./components/Register/Register.js";
import { AuthProvider } from "./contexts/AuthContext";
import "./assets/styles/App.css";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Home from "./components/Home.js";
import { getNavigateMenu } from "./components/Navigator.js";
import Backup from "./components/Backup";
import Game from "./components/Game.js"; // Updated import for Game component
import Store from "./components/Store.js"; // Updated import for Store component
import Search from "./components/Search.js";
import Cart from "./components/Cart.js"
import Wallet from "./components/Wallet.js"
import Ownership from "./components/Ownership.js";
import "./assets/styles/App.css"

const App = () => {
    return (
        <div className="app-container">
            {getNavigateMenu()}
            <div className="main-content">
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/backups" element={<Backup />} />
                    <Route path="/store" element={<Store />} />
                    <Route path="/cart" element={<Cart />} />
                    <Route path="/wallet" element={<Wallet />} />
                    <Route path="/game/:id" element={<Game />} /> {/* Updated route for Game */}
                    <Route path="/ownership" element={<Ownership />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                </Routes>
            </div>
        </div>
    );
};

const AppWrapper = () => (
    <Router>
        <AuthProvider>
            <App />
        </AuthProvider>
    </Router>
);

export default AppWrapper;
