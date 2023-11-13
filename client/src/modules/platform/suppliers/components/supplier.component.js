import { Button } from "@mui/material";
import { DynamicModal } from "../../../core";
import { AddProduct } from "../../product";
import SupplierTable from "./suppliers-table.component";

export default function Supplier() {
    return (
        <div className="row inventory-parent p-3">
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
                <SupplierTable />
            </div>
        </div>
    );
}
