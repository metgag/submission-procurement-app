let cart = [];
let supplierItems = [];

$(document).ready(() => {
  requireAuth();
  bindLogout();
  loadSuppliers();
  bindEvents();
  updateSubmitButtonState();
});

const updateSubmitButtonState = () => {
  const supplierId = $("#supplierSelect").val();
  const isDisabled = !supplierId || cart.length === 0;

  $("#submitOrderBtn").prop("disabled", isDisabled);
};

const loadSuppliers = () => {
  api
    .get("/suppliers")
    .then((res) => {
      res.data.forEach((s) => {
        $("#supplierSelect").append(
          `<option value="${s.id}">${s.name}</option>`,
        );
      });
    })
    .catch(() => {
      showToast({
        type: "error",
        title: "Error",
        message: "Failed to load suppliers",
      });
    });
};

const loadItemsBySupplier = (supplierId) => {
  $("#itemSelect").empty().append(`<option value="">Select Item</option>`);

  api
    .get(`/supplier-items?supplier_id=${supplierId}`)
    .then((res) => {
      supplierItems = res.data;
      res.data.forEach((si) => {
        $("#itemSelect").append(
          `<option value="${si.id}">
            ${si.item_name} (Stock: ${si.stock})
          </option>`,
        );
      });
    })
    .catch(() => {
      showToast({
        type: "error",
        title: "Error",
        message: "Failed to load items",
      });
    });
};

const bindEvents = () => {
  $("#supplierSelect").on("change", function () {
    const supplierId = $(this).val();
    cart = [];
    renderCart();
    updateSubmitButtonState();

    if (supplierId) {
      loadItemsBySupplier(supplierId);
    }
  });

  $("#addToCartBtn").click(addToCart);
  $("#submitOrderBtn").click(function (e) {
    e.preventDefault();
    submitOrder();
  });
};

const addToCart = () => {
  const supplierItemId = $("#itemSelect").val();
  const qty = parseInt($("#qtyInput").val(), 10);

  if (!supplierItemId || qty <= 0) {
    showToast({
      type: "error",
      title: "Validation Error",
      message: "Item and quantity are required",
    });
    return;
  }

  const item = supplierItems.find((i) => i.id == supplierItemId);

  if (qty > item.stock) {
    showToast({
      type: "error",
      title: "Stock Error",
      message: "Quantity exceeds available stock",
    });
    return;
  }

  const existing = cart.find((c) => c.supplier_item_id === item.id);

  if (existing) {
    existing.quantity += qty;
  } else {
    cart.push({
      supplier_item_id: item.id,
      item_name: item.item_name,
      quantity: qty,
      price: item.price,
    });
  }

  renderCart();

  $("#itemSelect").val("");
  $("#qtyInput").val("");
};

const renderCart = () => {
  const tbody = $("#cartTableBody");
  tbody.empty();

  if (cart.length === 0) {
    tbody.append(`
      <tr>
        <td colspan="5" class="py-4 text-center text-gray-500">
          Cart is empty
        </td>
      </tr>
    `);

    updateSubmitButtonState();
    return;
  }

  cart.forEach((c, index) => {
    tbody.append(`
      <tr class="text-sm hover:bg-gray-50">
        <td class="py-2">${c.item_name}</td>
        <td>${c.quantity}</td>
        <td>${formatRupiah(c.price)}</td>
        <td>${formatRupiah(c.price * c.quantity)}</td>
        <td>
          <button
            data-index="${index}"
            class="removeBtn text-red-500 text-sm cursor-pointer hover:opacity-70"
          >
            Remove
          </button>
        </td>
      </tr>
    `);
  });

  $(".removeBtn").click(function () {
    const index = $(this).data("index");
    cart.splice(index, 1);
    renderCart();
  });

  updateSubmitButtonState();
};

const submitOrder = () => {
  $("#submitOrderBtn").prop("disabled", true);
  const supplierId = $("#supplierSelect").val();

  if (!supplierId || cart.length === 0) {
    showToast({
      type: "error",
      title: "Validation Error",
      message: "Supplier and cart items are required",
    });

    $("#submitOrderBtn").prop("disabled", false);
    return;
  }

  const payload = {
    supplier_id: parseInt(supplierId),
    items: cart.map((c) => ({
      supplier_item_id: c.supplier_item_id,
      quantity: c.quantity,
    })),
  };

  api
    .post("/purchases", payload)
    .then((res) => {
      sessionStorage.setItem("purchaseResult", JSON.stringify(res.data));

      showToast({
        type: "success",
        title: "Success",
        message: res.message || "Purchase created successfully",
      });

      cart = [];
      renderCart();

      setTimeout(() => {
        window.location.href = "/invoice";
      }, 1800);
    })
    .catch((err) => {
      showToast({
        type: "error",
        title: "Purchase Failed",
        message: err.responseJSON?.message || "Server error",
      });

      $("#submitOrderBtn").prop("disabled", false);
    });
};
