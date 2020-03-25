function showInventory() {
    location.href = "/admin/inventory";
}

function showVendor() {
    location.href = "/admin/vendor";
}

function showTask() {
    location.href = "/admin/task";
}

function showOrder() {
    location.href = "/admin/order"
}

function deleteProduct(id) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "";

    idElement.value = id;
    idElement.name = "id";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}

function showVendorDetail(id) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "/admin/vendor/detail";

    idElement.value = id;
    idElement.name = "id";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}

function showProductDetail(id) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "/admin/inventory/detail";

    idElement.value = id;
    idElement.name = "id";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}