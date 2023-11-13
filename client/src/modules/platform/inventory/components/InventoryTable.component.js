import DynamicTable from "../../../core/common/table.component";
import { columns, rows } from "../inventory-table-tool";

export default function InventoryTable() {
    return <DynamicTable columns={columns} rows={rows} pagination={true} />;
}
