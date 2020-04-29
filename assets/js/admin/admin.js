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

function scrollToAddCategory() {
    var element = document.getElementById('addNewCategory');
    element.style.display = 'block';
    element.scrollIntoView({ behavior: 'smooth', block: 'center'});
}

function submitPurchaseRequest() {
    let form = document.getElementById('purchaseForm');
    let map = [];
    let counter = 0;
    let elements = form.elements;
    let tempMap = {};
    for (let i = 0, element; element = elements[i++];) {

        switch (counter) {
            case 0:
                tempMap["id"] = element.value;
                break;
            case 1:
                tempMap["name"] = element.value;
                break;
            case 2:
                tempMap["price"] = element.value;
                break;
            case 3:
                tempMap["purchased"] = element.value;
                break;
            case 4:
                tempMap["quantity"] = element.value;
                break;
            case 5:
                tempMap["purchase_quantity"] = element.value;
                let newMap = {};
                for (let j in tempMap)
                    newMap[j] = tempMap[j];
                map.push(newMap);
                break;
        }
        if (counter < 5){
            counter++;
        } else {
            counter = 0;
        }

    }

    let newForm = document.createElement("form");
    let idElement = document.createElement("input");

    newForm.method = "POST";
    newForm.action = "";

    idElement.value = JSON.stringify(map);
    idElement.name = "payload";

    newForm.appendChild(idElement);
    newForm.style.display = 'none';
    document.body.appendChild(newForm);

    newForm.submit();
}

function sortByDate(key) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "";

    idElement.value = key;
    idElement.name = "sortKey";
    form.style.display="none";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}

function showByCategory(Category) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "";

    idElement.value = Category;
    idElement.name = "Category";
    form.style.display="none";

    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}

function editUser(id) {
    let form = document.createElement("form");
    let idElement = document.createElement("input");

    form.method = "POST";
    form.action = "/admin/user/edit";

    idElement.value = id;
    idElement.name = "id";

    form.style.display='none';
    form.appendChild(idElement);
    document.body.appendChild(form);

    form.submit();
}