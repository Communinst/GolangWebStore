import React, { useState } from "react";
import { useAuth } from "../contexts/AuthContext";
import { Link } from "react-router-dom";
import {
    FaSignInAlt,
    FaHome,
    FaSignOutAlt,
    FaUser,
    FaArrowRight,
    FaFileContract,
    FaSearch,
    FaShoppingCart
} from "react-icons/fa";
import { MdOutlineBackup } from "react-icons/md";
import { backup } from "../utils/Fetch/BackupsF";
import { MdEvent } from "react-icons/md";
import { GiShoppingBag } from "react-icons/gi";
import { FaRegMessage } from "react-icons/fa6";
import Cart from "../components/Cart"; // Import the Cart component

export const getNavigateMenu = () => {
    const { isAuthenticated, logout, userType, user } = useAuth();
    const [error, setError] = useState(null);

    const doBackUp = async () => {
        try {
            const token = localStorage.getItem("authToken");
            await backup(token);
        } catch (error) {
            setError(error.message);
        }
    };

    if (error) {
        return <p className="error">{error}</p>;
    }

    return (
        <div>
            <nav className="app-nav">
                <Link to="/" className="nav-link">
                    <FaHome /> Homepage
                </Link>
                <Link to="/store" className="nav-link">
                    <GiShoppingBag /> Store
                </Link>
                <Link to="/events" className="nav-link">
                    <MdEvent /> Events
                </Link>
                {isAuthenticated && (
                    <Link to="/contracts" className="nav-link">
                        <FaFileContract /> My contracts
                    </Link>
                )}
                {isAuthenticated && (
                    <Link to="/profile" className="nav-link">
                        <FaUser /> Profile
                    </Link>
                )}
                {isAuthenticated && (
                    <Link to="/mail" className="nav-link">
                        <FaRegMessage /> Mail
                    </Link>
                )}
                {isAuthenticated && userType === "admin" && (
                    <Link to="/backups" className="nav-link">
                        <MdOutlineBackup /> Backup List
                    </Link>
                )}
                {isAuthenticated && userType === "admin" && (
                    <Link to="/" className="nav-link" onClick={() => { doBackUp(); logout(); }}>
                        <MdOutlineBackup /> Make backup
                    </Link>
                )}
                <Link to="/search" className="nav-link">
                    <FaSearch /> Search
                </Link>
                {!isAuthenticated && (
                    <Link to="/login" className="nav-link">
                        <FaSignInAlt /> Login
                    </Link>
                )}
                {isAuthenticated && (
                    <Link to="/" className="nav-link" onClick={() => { logout(); }}>
                        <FaSignOutAlt /> Logout
                    </Link>
                )}
                {isAuthenticated && <Cart />} {/* Add the Cart component */}
            </nav>
        </div>
    );
};