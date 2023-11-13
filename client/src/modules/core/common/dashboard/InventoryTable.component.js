import React from "react";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TablePagination from "@mui/material/TablePagination";
import TableRow from "@mui/material/TableRow";

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

// Placeholder rows data for testing
const rows = Array.from({ length: 10 }, (_, index) => ({
    products: "Maggi",
    buyingPrice: "₹430",
    quantity: "43 Packets",
    thresholdValue: "12 Packets",
    expiryDate: "11/12/22",
    availability: "In-stock",
    // Add more properties if needed
}));

export default function InventoryTable() {
    const [page, setPage] = React.useState(0);
    const [rowsPerPage, setRowsPerPage] = React.useState(10);

    const handleChangePage = (event, newPage) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event) => {
        setRowsPerPage(+event.target.value);
        setPage(0);
    };

    return (
        <Paper sx={{ width: "100%", overflow: "hidden" }}>
            <TableContainer sx={{ maxHeight: 440 }}>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            {columns.map((column) => (
                                <TableCell
                                    key={column.id}
                                    align={column.align}
                                    style={{ minWidth: column.minWidth }}
                                >
                                    {column.label}
                                </TableCell>
                            ))}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {rows &&
                            rows
                                .slice(
                                    page * rowsPerPage,
                                    page * rowsPerPage + rowsPerPage
                                )
                                .map((row) => (
                                    <TableRow
                                        hover
                                        role="checkbox"
                                        tabIndex={-1}
                                        key={row.code}
                                    >
                                        {columns.map((column) => {
                                            const value = row[column.id];
                                            return (
                                                <TableCell
                                                    key={column.id}
                                                    align={column.align}
                                                >
                                                    {column.format &&
                                                    typeof value === "number"
                                                        ? column.format(value)
                                                        : value}
                                                </TableCell>
                                            );
                                        })}
                                    </TableRow>
                                ))}
                    </TableBody>
                </Table>
            </TableContainer>
            <TablePagination
                rowsPerPageOptions={[10, 25, 100]}
                component="div"
                count={rows.length}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
            />
        </Paper>
    );
}
