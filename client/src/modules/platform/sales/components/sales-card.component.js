import { useState } from "react";
import { Card } from "../../../core";

function SalesCard() {
    const [data, setData] = useState({
        sales: 825,
        revenue: 3648,
        profit: 2500,
        cost: 60,
    });

    return (
        <div class="col-sm-8">
            <div class="card mb-4">
                <div class="card-body">
                    <div className="row">
                        <h1 className="dashboard-card-heading">
                            Purchase Overview
                        </h1>
                        <Card
                            iconClass="bi-receipt-cutoff"
                            amount={data.sales}
                            label="Sales"
                        />
                        <Card
                            iconClass="bi-bar-chart-line-fill"
                            amount={data.revenue}
                            label="Revenue"
                        />
                        <Card
                            iconClass="bi-graph-up-arrow"
                            amount={data.profit}
                            label="Profit"
                        />
                        <Card
                            iconClass="bi-coin"
                            amount={data.cost}
                            label="Cost"
                        />
                    </div>
                </div>
            </div>
        </div>
    );
}

export default SalesCard;
