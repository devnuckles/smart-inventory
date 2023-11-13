const categoryColumns = [
    { id: "category", label: "Category", minWidth: 170 },
    { id: "turn", label: "Turn Over", minWidth: 140 },
    {
        id: "increase",
        label: "Increase By",
        minWidth: 150,

        format: (value) => value.toLocaleString("en-US"),
    },
];
const categoryRows = Array.from({ length: 3 }, (_, index) => ({
    category: "kraken",
    turn: "25965",
    increase: "35%",
}));

const productRows = Array.from({ length: 3 }, (_, index) => ({
    product: "Tomato",
    id: "15645",
    category: "Vegetable",
    quantity: "225 kg",
    turn: "1753",
    increase: "2.3%",
}));
const productColumns = [
    { id: "product", label: "Product", minWidth: 170 },
    { id: "id", label: "ID", minWidth: 140 },
    { id: "category", label: "Category", minWidth: 170 },
    { id: "quantity", label: "Quantity", minWidth: 140 },
    { id: "turn", label: "Turn Over", minWidth: 170 },
    { id: "increase", label: "Increase By", minWidth: 140 },
];

export { categoryColumns, categoryRows, productRows, productColumns };
