import { Divider } from "@mui/material";

function PurchaseOverviewCard() {
    return (
        <div class="col-sm-6">
            <div class="card">
                <div class="card-body">
                    <div class="product-overview">
                        <div class="product-overview-heading dashboard-card-heading">
                            <h2>Overview</h2>
                        </div>
                        <div class="product-overview-middle">
                            <div class="row">
                                <div class="col-lg-4">
                                    <p>₹21,190</p>
                                    <p>Total Profit</p>
                                </div>
                                <div class="col-lg-4">
                                    <p>₹18,300</p>
                                    <p>Revenue</p>
                                </div>
                                <div class="col-lg-4">
                                    <p>₹17,432</p>
                                    <p>Sales</p>
                                </div>
                            </div>
                        </div>
                        <Divider />
                        <div class="product-overview-bottom">
                            <div class="row">
                                <div class="col-lg-3">
                                    <p>₹21,190</p>
                                    <p>Total Profit</p>
                                </div>
                                <div class="col-lg-3">
                                    <p>₹18,300</p>
                                    <p>Revenue</p>
                                </div>
                                <div class="col-lg-3">
                                    <p>₹17,432</p>
                                    <p>Sales</p>
                                </div>
                                <div class="col-lg-3">
                                    <p>₹17,432</p>
                                    <p>Sales</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default PurchaseOverviewCard;
