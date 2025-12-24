const apiRequest = (method, url, data = null, options = {}) => {
  return $.ajax({
    url: BASE_API_URL + url,
    method,
    contentType: "application/json",
    data: data ? JSON.stringify(data) : null,
    headers: {
      ...authHeader(),
      ...(options.headers || {}),
    },
  });
};

const api = {
  get(url, options) {
    return apiRequest("GET", url, null, options);
  },
  post(url, data, options) {
    return apiRequest("POST", url, data, options);
  },
  patch(url, data, options) {
    return apiRequest("PATCH", url, data, options);
  },
  del(url, options) {
    return apiRequest("DELETE", url, null, options);
  },
};
