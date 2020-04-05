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

function addProduct(){
    let form = document.getElementById('productForm');
    let image = document.getElementById("image").files;

    if (image.length) {

        // Begin file upload
        console.log("Uploading file to Imgur..");

        // Replace ctrlq with your own API key
        var apiUrl = 'https://api.imgur.com/3/image';
        var apiKey = 'fca6e3aed6b3058';

        var settings = {
            async: false,
            crossDomain: true,
            processData: false,
            contentType: false,
            type: 'POST',
            url: apiUrl,
            headers: {
                Authorization: 'Client-ID ' + apiKey,
                Accept: 'application/json'
            },
            mimeType: 'multipart/form-data'
        };

        var formData = new FormData();
        formData.append("image", image[0]);
        settings.data = formData;

        // Response contains stringified JSON
        // Image URL available at response.data.link
        $.ajax(settings).done(function(response) {
            console.log(response);
            let responseObject = JSON.parse(response);
            let imageUrl = document.getElementById('imageUrl');
            let link = responseObject.data.link;
            imageUrl.value = link;
        });

    }
}

$("document").ready(function() {

    $('input[type=file]').on("change", function() {

    });
});