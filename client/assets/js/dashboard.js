const fetchInventory = () => {
  api
    .get("/supplier-items")
    .then((res) => {
      renderInventory(res.data);
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
