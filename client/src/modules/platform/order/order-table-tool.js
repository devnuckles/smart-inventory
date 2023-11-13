const columns = [
    { id: "products", label: "Products", minWidth: 170 },
    { id: "value", label: "Order Value", minWidth: 140 },
    {
        id: "quantity",
        label: "Quantity",
        minWidth: 150,

        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "id",
        label: "Order ID",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "expection",
        label: "Expected Delivary",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "status",
        label: "Status",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
];

const rows = Array.from({ length: 3 }, (_, index) => ({
    products: "kraken",
    value: "259",
    quantity: "46 packets",
    id: "4687",
    expection: "11/12/23",
    status: "Delayed",
}));

export { columns, rows };
