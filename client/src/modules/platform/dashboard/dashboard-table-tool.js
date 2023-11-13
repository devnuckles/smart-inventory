const columns = [
    { id: "name", label: "Name", minWidth: 170 },
    { id: "sold", label: "Sold Quantity", minWidth: 140 },
    {
        id: "remaining",
        label: "Remaining Quantity",
        minWidth: 150,

        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "price",
        label: "Price",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
];

const rows = Array.from({ length: 3 }, (_, index) => ({
    name: "Maggi",
    sold: "₹430",
    remaining: "43 Packets",
    price: "₹430",
}));

export { columns, rows };
