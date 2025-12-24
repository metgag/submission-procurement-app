const isLoggedIn = () => !!getToken();

const logout = () => {
  removeToken();
  window.location.href = "/login";
};

// register
const handleRegister = () => {
  $("#registerForm").on("submit", (e) => {
    e.preventDefault();

    const username = $("#username").val().trim();
    const password = $("#password").val().trim();
    const btn = $("#registerBtn");

    if (!username || !password) {
      showToast({
        type: "error",
        title: "Validation Failed",
        message: "Username and password is required",
      });
      return;
    }

    if (password.length < 8) {
      showToast({
        type: "error",
        title: "Validation Failed",
        message: "Password should at least 8 characters long",
      });
      return;
    }

    btn.prop("disabled", true).text("Loading...");

    api
      .post("/auth/register", { username, password })
      .then(() => {
        setTimeout(() => {
          showToast({
            type: "success",
            title: "Register Success",
            message: "Account created successfully",
          });

          btn.text("Redirecting...");
          setTimeout(() => {
            window.location.href = "/login.html";
          }, 2800);
        }, 1600);
      })
      .catch((err) => {
        setTimeout(() => {
          showToast({
            type: "error",
            title: "Register Failed",
            message: err.responseText || "Register error",
          });

          btn.prop("disabled", false).text("Register");
        }, 1600);
      });
  });
};

// login
const handleLogin = () => {
  $("#loginForm").on("submit", (e) => {
    e.preventDefault();

    const username = $("#username").val().trim();
    const password = $("#password").val().trim();
    const btn = $("#loginBtn");

    if (!username || !password) {
      showToast({
        type: "error",
        title: "Validation Failed",
        message: "Username and password is required",
      });
      return;
    }

    btn.prop("disabled", true).text("Loading...");

    api
      .post("/auth/login", { username, password })
      .then((res) => {
        setTimeout(() => {
          setToken(res.data.token);

          showToast({
            type: "success",
            title: "Login Success",
            message: "Welcome back",
          });

          btn.text("Redirecting...");

          setTimeout(() => {
            window.location.href = "/dashboard";
          }, 2800);
        }, 1600);
      })
      .catch((err) => {
        setTimeout(() => {
          showToast({
            type: "error",
            title: "Login Failed",
            message: err.responseText || "Invalid username or password",
          });

          btn.prop("disabled", false).text("Login");
        }, 1600);
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
