import { DynamicTable } from "../../../core";
import { productColumns, productRows } from "./sales.tool";

function BestSalesProducts() {
    return (
        <div class="card mb-4">
            <div class="card-body">
                <div className="row">
                    <h1 className="dashboard-card-heading">
                        Best selling category
                    </h1>
                </div>
                <DynamicTable columns={productColumns} rows={productRows} />
            </div>
        </div>
    );
}

export default BestSalesProducts;
