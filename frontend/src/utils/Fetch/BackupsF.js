//!done
export const backup = async (token) => {
    const response = await fetch(`/admin/dumps/create`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    });

    if (!response.ok) {
        const errorData = await response.json();
        const errorMessage = errorData.detail || "Failed to dump";
        throw new Error(errorMessage);
    }
};
//!done
export const postDumps = async (token, filename) => {
    const response = await fetch(`/admin/dumps/restore`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ 
            filename : filename }),
    });

    if (!response.ok) {
        const errorData = await response.json();
        const errorMessage = errorData.detail || "Failed to backup";
        throw new Error(errorMessage);
    }
};
//!done
export const getDumps = async (token) => {
    const response = await fetch(`/admin/dumps`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    });

    if (!response.ok) {
        const errorData = await response.json();
        const errorMessage = errorData.detail || "Failed to get backups";
        throw new Error(errorMessage);
    }

    const data = await response.json();
    console.log("API Response:", data);
    return data;
};
