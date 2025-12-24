let currentPage = 1;
const pageSize = 10;

const fetchInventory = (page = 1) => {
  currentPage = page;

  api
    .get(`/supplier-items?page=${currentPage}&page_size=${pageSize}`)
    .then((res) => {
      renderInventory(res.data);
      renderPagination(res.meta);
    })
    .catch(() => {
      showToast({
        type: "error",
        title: "Load Inventory Failed",
        message: "Unable to load inventory",
      });
    });
};

const renderInventory = (items) => {
  const tbody = $("#inventoryTableBody");
  tbody.empty();

  if (!items || items.length < 1) {
    tbody.append(`
        <tr>
            <td colspan="2" class="px-3 py-2 text-center">
                No inventory data available
            </td>
        </tr>
    `);
    return;
  }

  items.forEach((item) => {
    tbody.append(`
        <tr class="*:text-gray-900 *:first:font-medium">
            <td class="px-3 py-2 whitespace-nowrap">
                ${item.item_name}
            </td>
            <td class="px-3 py-2 whitespace-nowrap">
                ${item.stock}
            </td>
            <td class="px-3 py-2 whitespace-nowrap">
                ${item.supplier_name}
            </td>
        </tr>
    `);
  });
};

const renderPagination = (meta) => {
  const pagination = $("#pagination");
  pagination.empty();

  if (!meta || meta.total_pages <= 1) return;

  pagination.append(`
    <li>
      <a href="#" class="px-3 py-1 rounded hover:opacity-70 hover:bg-gray-300" ${meta.page <= 1 ? "disabled" : ""} data-page="${meta.page - 1}"><</a>
    </li>
  `);

  for (let i = 1; i <= meta.total_pages; i++) {
    pagination.append(`
      <li>
        <a href="#" class="px-3 py-1 rounded ${
          i === meta.page
            ? "bg-indigo-600 text-white border-indigo-600 cursor-not-allowed"
            : "hover:opacity-70 hover:bg-gray-300"
        }" data-page="${i}">${i}
        </a>
      </li>
    `);
  }

  pagination.append(`
    <li>
      <a href="#" class="px-3 py-1 rounded hover:opacity-70 hover:bg-gray-300" ${meta.page >= meta.total_pages ? "disabled" : ""} data-page="${meta.page + 1}">></a>
    </li>
  `);

  $("#pagination a").click(function (e) {
    e.preventDefault();
    const page = $(this).data("page");
    if (page && page !== currentPage) {
      fetchInventory(page);
    }
  });
};
