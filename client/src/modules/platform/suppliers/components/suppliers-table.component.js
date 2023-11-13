import { Button } from "@mui/material";
import { DynamicTable } from "../../../core";
import { columns, rows } from "../supplier-table-tool";
import { DynamicModal } from "../../../core";
import { AddProduct } from "../../product";

export default function SupplierTable() {
    return (
        <div className=" row inventory-bottom-table my-5 bg-white py-4 m-0">
            <div className="inventory-table-header mb-4">
                <div className="row">
                    <div className="col-lg-7">
                        <h2>Suppliers</h2>
                    </div>
                    <div className="col-lg-5 inventory-table-header-button text-end">
                        <div className="row">
                            <div className="col-lg-4 p-0">
                                <DynamicModal
                                    Element={<AddProduct />}
                                    label="Add Supplier"
                                />
                            </div>
                            <div className="col-lg-3 p-0">
                                <Button
                                    className="me-2 custom-font-size text-dark"
                                    variant="outlined"
                                >
                                    <i className="bi bi-filter"></i>
                                    Filters
                                </Button>
                            </div>
                            <div className="col-lg-5 p-0">
                                <Button
                                    className="me-2 custom-font-size text-dark"
                                    variant="outlined"
                                >
                                    Download all
                                </Button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div className="inventory-table text-start">
                <DynamicTable columns={columns} rows={rows} pagination={true} />
            </div>
        </div>
    );
}
