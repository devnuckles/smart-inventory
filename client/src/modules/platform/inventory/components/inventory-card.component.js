import { useState } from "react";
import { Card } from "../../../core";

function InventoryCard() {
    const [data, setData] = useState({
        quantity: 825,
        recieved: 3648,
    });

    return (
        <div class="col-sm-4">
            <div class="card ">
                <div class="card-body">
                    <div className="row">
                        <h1 className="dashboard-card-heading">
                            Inventory Summary
                        </h1>
                        <Card
                            className="col-lg-6"
                            iconClass="bi-receipt-cutoff"
                            amount={data.quantity}
                            label="quantity"
                        />
                        <Card
                            className="col-lg-6"
                            iconClass="bi-bar-chart-line-fill"
                            amount={data.recieved}
                            label="recieved"
                        />
                    </div>
                </div>
            </div>
        </div>
    );
}

export default InventoryCard;
