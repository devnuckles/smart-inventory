import { Button } from "@mui/material";
import { DynamicTable } from "../../../core";
import { columns, rows } from "../supplier-table-tool";
import { DynamicModal } from "../../../core";
import { AddProduct } from "../../product";

export default function SupplierTable() {
    return <DynamicTable columns={columns} rows={rows} pagination={true} />;
}
