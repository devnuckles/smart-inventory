import DynamicTable from "../../../core/common/table/table.component";
import { columns, rows } from "../../dashboard/dashboard-table-tool";

export default function StockTable() {
    return <DynamicTable columns={columns} rows={rows} />;
}
