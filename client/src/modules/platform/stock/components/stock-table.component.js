import { DynamicTable } from "../../../core";
import { columns, rows } from "../../dashboard/dashboard-table-tool";

export default function StockTable() {
    return <DynamicTable columns={columns} rows={rows} pagination={true} />;
}
