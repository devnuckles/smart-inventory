const columns = [
    { id: "products", label: "Products", minWidth: 170 },
    { id: "buyingPrice", label: " Buying Price", minWidth: 140 },
    {
        id: "quantity",
        label: "Quantity",
        minWidth: 150,

        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "thresholdValue",
        label: "Threshold Value",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "expiryDate",
        label: "Expiry Date",
        minWidth: 150,
        format: (value) => value.toFixed(2),
    },
    {
        id: "availability",
        label: " Availability",
        minWidth: 170,
        format: (value) => value.toFixed(2),
    },
];

const rows = Array.from({ length: 10 }, (_, index) => ({
    products: "Maggi",
    buyingPrice: "₹430",
    quantity: "43 Packets",
    thresholdValue: "12 Packets",
    expiryDate: "11/12/22",
    availability: "In-stock",
}));

export { columns, rows };
