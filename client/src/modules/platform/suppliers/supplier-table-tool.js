const columns = [
    { id: "supplier", label: "Supplier Name", minWidth: 170 },
    { id: "product", label: "Product", minWidth: 140 },
    {
        id: "contact",
        label: "Contact Number",
        minWidth: 150,

        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "email",
        label: "Email",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "type",
        label: "Type",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
    {
        id: "status",
        label: "On the way",
        minWidth: 170,
        format: (value) => value.toLocaleString("en-US"),
    },
];

const rows = Array.from({ length: 3 }, (_, index) => ({
    supplier: "kraken",
    product: "product",
    contact: "8801403229479",
    email: "javacriptiqbal@gmailc.om",
    type: "Taking Return",
    status: 12,
}));

export { columns, rows };
