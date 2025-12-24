const isLoggedIn = () => !!getToken();

const logout = () => {
  removeToken();
  window.location.href = "/login";
};

// register
const handleRegister = () => {
  $("#registerBtn").click(() => {
    const username = $("#username").val().trim();
    const password = $("#password").val().trim();

    if (!username || !password) {
      showToast({
        type: "error",
        title: "Validation Failed",
        message:
          err.responseJSON?.message || "Username and password is required",
      });
      return;
    }

    api
      .post("/auth/register", { username, password })
      .then(() => {
        showToast({
          type: "success",
          title: "Register Success",
          message: "Account created successfully",
        });

        setTimeout(() => {
          window.location.href = "/login.html";
        }, 1500);
      })
      .catch((err) => {
        showToast({
          type: "error",
          title: "Register Failed",
          message: err.responseJSON?.message || "Register error",
        });
      });
  });
};

// login
const handleLogin = () => {
  $("#loginBtn").click(() => {
    const username = $("#username").val().trim();
    const password = $("#password").val().trim();

    if (!username || !password) {
      showToast({
        type: "error",
        title: "Validation Failed",
        message:
          err.responseJSON?.message || "Username and password is required",
      });
      return;
    }

    api
      .post("/auth/login", { username, password })
      .then((res) => {
        setToken(res.data.token);

        showToast({
          type: "success",
          title: "Login Success",
          message: "Welcome back",
        });

        setTimeout(() => (window.location.href = "/dashboard"), 2400);
      })
      .catch((err) => {
        showToast({
          type: "error",
          title: "Login Failed",
          message: err.responseJSON?.message || "Invalid username or password",
        });
      });
  });
};

// logged in user
const requireAuth = () => {
  if (!isLoggedIn()) window.location.href = "/login";
};

// logout
const bindLogout = () => {
  $("#logoutBtn").click(logout);
};
