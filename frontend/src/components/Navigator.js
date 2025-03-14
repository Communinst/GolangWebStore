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
    FaShoppingCart,
    FaWallet
} from "react-icons/fa";
import { MdOutlineBackup } from "react-icons/md";
import { backup } from "../utils/Fetch/BackupsF";
import { MdEvent } from "react-icons/md";
import { GiShoppingBag } from "react-icons/gi";
import { FaRegMessage } from "react-icons/fa6";

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
                
                {isAuthenticated && (
                    <Link to="/store" className="nav-link">
                        <GiShoppingBag /> Store
                    </Link>
                )}
                {isAuthenticated && (
                    <Link to="/cart" className="nav-link">
                        <FaShoppingCart /> Cart
                    </Link>
                )}
                {isAuthenticated && (
                    <Link to="/wallet" className="nav-link">
                        <FaWallet /> Wallet
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
                {isAuthenticated && (
                    <Link to="/ownership" className="nav-link">
                        <FaRegMessage /> Ownership
                    </Link>
                )} 
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
            </nav>
        </div>
    );
};