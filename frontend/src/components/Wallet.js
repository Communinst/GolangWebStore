import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Alert } from "./Alert";
import { useAuth } from "../contexts/AuthContext";
import '../assets/styles/Wallet.css'; // Ensure you have a CSS file for styling
import { fetchWallet, updateWalletBalance } from "../utils/Fetch/WalletF";

const Wallet = () => {
    const [wallet, setWallet] = useState(null);
    const [isLoading, setIsLoading] = useState(true);
    const [error, setError] = useState(null);
    const [alertMessage, setAlertMessage] = useState("");
    const [showDepositModal, setShowDepositModal] = useState(false);
    const [depositAmount, setDepositAmount] = useState("");
    const navigate = useNavigate();

    const { isAuthenticated, userType, userId } = useAuth();
    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const loadBalance = async () => {
            try {
                const walletData = await fetchWallet(token, parseInt(userId));
                setWallet(walletData);
            } catch (error) {
                setError(error.message);
            }
            setIsLoading(false);
        };

        loadBalance();
    }, [token, userId]);

    const handleAlertClose = () => setAlertMessage("");

    const openDepositModal = () => {
        setShowDepositModal(true);
    };

    const closeDepositModal = () => {
        setShowDepositModal(false);
    };

    const handleSubmit = async () => {
        setIsLoading(true);
        try {
            await updateWalletBalance(token, parseInt(userId), parseFloat(depositAmount));
            setAlertMessage("Transaction complete.");
            closeDepositModal(); // Close the modal
            const updatedWallet = await fetchWallet(token, parseInt(userId));
            setWallet(updatedWallet);
        } catch (error) {
            setError(error.message);
        }
        setIsLoading(false);
    };

    if (!isAuthenticated) {
        return <p className="loading">You need to sign in to see your wallet...</p>;
    }

    if (isLoading) {
        return <p>Loading...</p>;
    }

    return (
        <div className="wallet-container">
            <div className="wallet-balance">
                <h2>Current User Balance</h2>
                <p>${wallet?.balance}</p>
            </div>
            <button onClick={openDepositModal} className="deposit-button">
                Deposit Money
            </button>
            {showDepositModal && (
                <div className="modal">
                    <div className="modal-content">
                        <span className="close" onClick={closeDepositModal}>&times;</span>
                        <h2>Deposit Money</h2>
                        <input
                            type="number"
                            placeholder="Enter amount"
                            value={depositAmount}
                            onChange={(e) => setDepositAmount(e.target.value)}
                        />
                        <button onClick={handleSubmit}>Deposit</button>
                        <button onClick={closeDepositModal}>Cancel</button>
                    </div>
                </div>
            )}
            {error && <p className="error">{error}</p>}
            {alertMessage !== "" && (
                <Alert message={alertMessage} onClose={handleAlertClose} />
            )}
        </div>
    );
};

export default Wallet;
