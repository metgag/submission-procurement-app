const BASE_API_URL = "/api/v1";

const getToken = () => {
  return localStorage.getItem("token");
};

const setToken = (token) => {
  localStorage.setItem("token", token);
};

const removeToken = () => {
  localStorage.removeItem("token");
};

const authHeader = () => {
  return {
    Authorization: "Bearer " + getToken(),
  };
};
