//!done
export const postRegister = async (
    login,
    navigate,
    userLogin,
    email,
    password,
    setSuccess
) => {
    const response = await fetch("/welcome/auth/sign-up", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            login: userLogin,
            email : email, 
            password : password }),
    });

    if (response.ok) {
        setSuccess("Registration successful! You can now log in.");
        await postLogin(login, navigate, email, password);
    } else {
        const errorData = await response.json();
        const errorMessage = errorData.detail || "Failed to register";
        throw new Error(errorMessage);
    }
};
//!done
export const postLogin = async (login, navigate, email, password) => {
    const response = await fetch("/welcome/auth/sign-in", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            email : email, 
            password : password }),
    });

    if (response.ok) {
        const data = await response.json();
        login(data.token, data.role, data.userId);
        navigate("/");
    } else {
        const errorData = await response.json();
        const errorMessage = errorData.detail || "Failed to login";
        throw new Error(errorMessage);
    }
};