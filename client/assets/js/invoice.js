$(document).ready(() => {
  const raw = sessionStorage.getItem("purchaseResult");

  if (!raw) {
    window.location.href = "/purchase";
    return;
  }

  const data = JSON.parse(raw);

  $("#purchaseId").text(data.id);
  $("#grandTotal").text(formatRupiah(data.grand_total));

  const tbody = $("#itemsTable");
  tbody.empty();

  data.items.forEach((item) => {
    tbody.append(`
      <tr class="border-t">
        <td class="p-2">${item.item_id}</td>
        <td class="p-2 text-center">${item.quantity}</td>
        <td class="p-2 text-right">${formatRupiah(item.subtotal)}</td>
      </tr>
    `);
  });

  sessionStorage.removeItem("purchaseResult");
});
