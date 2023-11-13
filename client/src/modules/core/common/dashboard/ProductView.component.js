import { Button } from "@mui/material";
import Stack from "@mui/material/Stack";
import ProductViewTab from "./ProductViewTab.component";

export default function ProductView() {
    return (
        <>
            <div className="row product-view-parent p-3">
                <div className=" product-view-wrapper bg-white p-3">
                    <div className="col-lg-12 product-view-header">
                        <div class="row">
                            <div class="col-lg-9 ">
                                <h2 className="dashboard-card-heading">
                                    Maggi
                                </h2>
                            </div>
                            <div class="col-lg-3 product-view-header-button text-end">
                                <Button className="me-3" variant="outlined">
                                    <i class="bi bi-pencil me-2"></i> Edit
                                </Button>
                                <Button variant="outlined">Download</Button>
                            </div>
                        </div>
                    </div>
                    <div className="col-lg-12">
                        <ProductViewTab />
                    </div>
                </div>
            </div>
        </>
    );
}
