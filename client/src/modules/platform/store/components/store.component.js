import { Card } from "@mui/material";
import { DynamicModal } from "../../../core";
import { AddProduct } from "../../product";
import StoreCard from "./store-card.component";

export default function Store() {
    const topCardData = [
        {
            title: "Total Orders",
            count: 14,
            days: "Last 7 days",
        },
        {
            title: "Total Recieved",
            count: 868,
            days: "Last 7 days",
            dataValue: "₹25,000",
            additionalInfo: "Cost",
        },
        {
            title: "Total Returned",
            count: 5,
            days: "Last 7 days",
            dataValue: "₹25,000",
            additionalInfo: "Ordered",
        },
        {
            title: "On the way",
            count: 12,
            days: "Ordered",
            dataValue: "2",
            additionalInfo: "Cost",
        },
        // Add more data objects as needed
    ];

    return (
        <div className="row inventory-parent p-3">
            <div className=" row inventory-bottom-table my-5 bg-white py-4 m-0">
                <div className="inventory-table-header mb-4">
                    <div className="row">
                        <div className="col-lg-7">
                            <h2>Manage Store</h2>
                        </div>
                        <div className="col-lg-5 inventory-table-header-button text-end">
                            <div className="row">
                                <div className="col-lg-4 p-0 ms-auto">
                                    <DynamicModal
                                        Element={<AddProduct />}
                                        label="Add Store"
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="row">
                        <Card className="mt-4">
                            <StoreCard />
                        </Card>
                        <Card className="mt-4">
                            <StoreCard />
                        </Card>
                        <Card className="mt-4">
                            <StoreCard />
                        </Card>
                    </div>
                </div>
            </div>
        </div>
    );
}
