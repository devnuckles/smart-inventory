import { Button } from "@mui/material";
import DynamicTable from "../../../core/common/table/table.component";
import { columns, rows } from "../supplier-table-tool";

export default function SupplierTable() {
    return (
        <>
            <div className=" row inventory-bottom-table my-5 bg-white py-4 m-0">
                <div className="inventory-table-header mb-4">
                    <div className="row">
                        <div className="col-lg-7">
                            <h2>Products</h2>
                        </div>
                        <div className="col-lg-5 inventory-table-header-button text-end">
                            <Button className="me-2" variant="contained">
                                <i className="bi bi-filter"></i>Filters
                            </Button>
                            <Button className="" variant="outlined">
                                Download all
                            </Button>
                        </div>
                    </div>
                </div>
                <div className="inventory-table text-start">
                    <DynamicTable columns={columns} rows={rows} />
                </div>
            </div>
        </>
    );
}
