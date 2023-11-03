function addProduct() {
    // Get input values
    const productId = document.getElementById("product-id").value;
    const productName = document.getElementById("product-name").value;
    const productQuantity = document.getElementById("product-quantity").value;
    const lastStockDate = document.getElementById("last-stock-date").value;

    // Create a new row for the table
    const table = document.getElementById("product-table");
    const newRow = table.insertRow(-1);

    // Insert cells with input values
    const idCell = newRow.insertCell(0);
    idCell.innerHTML = productId;

    const nameCell = newRow.insertCell(1);
    nameCell.innerHTML = productName;

    const quantityCell = newRow.insertCell(2);
    quantityCell.innerHTML = productQuantity;

    const dateCell = newRow.insertCell(3);
    dateCell.innerHTML = lastStockDate;

    // Create a delete button
    const deleteCell = newRow.insertCell(4);
    const deleteButton = document.createElement("button");
    deleteButton.innerText = "Delete";
    deleteButton.addEventListener("click", function() {
        deleteProduct(newRow);
    });
    deleteCell.appendChild(deleteButton);

    // Clear the input fields
    document.getElementById("product-id").value = "";
    document.getElementById("product-name").value = "";
    document.getElementById("product-quantity").value = "";
    document.getElementById("last-stock-date").value = "";
}

function deleteProduct(row) {
    const table = document.getElementById("product-table");
    table.deleteRow(row.rowIndex);
}
