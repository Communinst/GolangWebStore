import React, { useEffect, useState } from "react";
import { getDumps, postDumps } from "../utils/Fetch/BackupsF";
import { useAuth } from "../contexts/AuthContext";

const Backup = () => {
    const [backUps, setBackUps] = useState([]);
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);
    const { isAuthenticated, logout, userType } = useAuth();

    const token = localStorage.getItem("authToken");

    useEffect(() => {
        const fetchBackUps = async () => {
            try {
                const data = await getDumps(token);
                setBackUps(data);
            } catch (error) {
                setError(error.message);
            } finally {
                setLoading(false);
            }
        };
        fetchBackUps();
    }, [token]);

    const handleClick = async (filename) => {
        try {
            await postDumps(token, filename);
        } catch (err) {
            setError(err.message);
        }
    };

    if (error) {
        return <p className="error">{error}</p>;
    }

    if (loading) {
        return <p className="loading">Loading...</p>;
    }

    return (
        <div>
            {isAuthenticated && userType === "admin" && (
                <div className="back-page">
                    {backUps.map((backUp, index) => (
                        <button key={index} onClick={() => handleClick(backUp.filename)}>
                            {backUp.filename}
                        </button>
                    ))}

                    {backUps.length === 0 && (
                        <p className="no-backup">No backups available</p>
                    )}
                </div>
            )}
            {(!isAuthenticated || userType !== "admin") && (
                <div className="back-page">
                    <p className="loading">You have no rights to see this page...</p>
                </div>
            )}
        </div>
    );
};

export default Backup;
