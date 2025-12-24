const showToast = ({ type = "success", title, message, duration = 3000 }) => {
  const templates = {
    success: `
      <div role="alert" class="toast rounded-md border border-green-500 bg-green-50 p-4 shadow-sm">
        <div class="flex items-start gap-4">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none"
            viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="-mt-0.5 size-6 text-green-700">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M9 12.75L11.25 15 15 9.75
              M21 12a9 9 0 11-18 0
              9 9 0 0118 0z"></path>
          </svg>

          <div class="flex-1">
            <strong class="block leading-tight font-medium text-green-800">
              ${title || "Success"}
            </strong>
            <p class="mt-0.5 text-sm text-green-700">
              ${message}
            </p>
          </div>
        </div>
      </div>
    `,
    error: `
      <div role="alert" class="toast rounded-md border border-red-500 bg-red-50 p-4 shadow-sm">
        <div class="flex items-start gap-4">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none"
            viewBox="0 0 24 24" stroke-width="1.5"
            stroke="currentColor" class="-mt-0.5 size-6 text-red-700">
            <path stroke-linecap="round" stroke-linejoin="round"
              d="M12 9v3.75m0 3.75h.008
              M21 12a9 9 0 11-18 0
              9 9 0 0118 0z"></path>
          </svg>

          <div class="flex-1">
            <strong class="block leading-tight font-medium text-red-800">
              ${title || "Error"}
            </strong>
            <p class="mt-0.5 text-sm text-red-700">
              ${message}
            </p>
          </div>
        </div>
      </div>
    `,
  };

  const toast = $(templates[type]);

  $("#toast-container").append(toast);

  setTimeout(() => {
    toast.fadeOut(300, () => toast.remove());
  }, duration);
};
