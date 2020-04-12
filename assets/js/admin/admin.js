function showInventory() {
    location.href = "/admin/inventory";
}

function showVendor() {
    location.href = "/admin/vendor";
}

function showUser() {
    location.href = "/admin/user";
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

function showOrderDetail(id) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.style.display = 'none';

    form.method = "POST";
    form.action = "/admin/order/detail";

    idElement.value = id;
    idElement.name = "id";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}

function changeOwnerCredential() {
    let form = document.createElement("form");
    let ownerUsername = document.createElement("input");
    let ownerPassword = document.createElement("input");

    let ownerUsernameForm = document.getElementById("ownerUsername");
    let ownerPasswordForm = document.getElementById("ownerPassword");
    form.style.display = 'none';

    form.method = "POST";
    form.action = "";

    ownerUsername.value = ownerUsernameForm.value;
    ownerUsername.name = "ownerUsername";

    ownerPassword.value = ownerPasswordForm.value;
    ownerPassword.name = "ownerPassword";

    form.appendChild(ownerUsername);
    form.appendChild(ownerPassword);
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

function deleteData(id) {
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

function scrollToDetail() {
    var element = document.getElementById('orderedProductDetail');
    element.scrollIntoView({ behavior: 'smooth', block: 'start'});
}