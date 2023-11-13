import React from "react";
import Button from "@mui/material/Button";
import InventoryTable from "./InventoryTable.component";

function createData(
    products,
    buyingPrice,
    quantity,
    thresholdValue,
    expiryDate,
    availability
) {
    return {
        products,
        buyingPrice,
        quantity,
        thresholdValue,
        expiryDate,
        availability,
    };
}

export default function Inventory() {
    const topCardData = [
        {
            title: "Categories",
            count: 14,
            days: "Last 7 days",
        },
        {
            title: "Total Products",
            count: 868,
            days: "Last 7 days",
            dataValue: "₹25,000",
            additionalInfo: "Revenue",
        },
        {
            title: "Top Selling",
            count: 5,
            days: "Last 7 days",
            dataValue: "₹25,000",
            additionalInfo: "Cost",
        },
        {
            title: "Low Stocks",
            count: 12,
            days: "Ordered",
            dataValue: "2",
            additionalInfo: "Not in stock",
        },
        // Add more data objects as needed
    ];

    return (
        <>
            <div className="row inventory-parent p-3">
                <div className="row inventory-card p-3 m-0">
                    <div className="col-lg-12 ">
                        <h2 className="dashboard-card-heading">
                            Overall Inventory
                        </h2>
                    </div>

                    {topCardData.map((data, index) => (
                        <div
                            key={index}
                            className="col-lg-3 inventory-top-card-column"
                        >
                            <h2>{data.title}</h2>
                            <div className="inventory-top-card-column-data">
                                <div className="row">
                                    <div className="col-lg-6">
                                        <p className="m-0 first-data">
                                            {data.count}
                                        </p>
                                        <p className="second-data">
                                            {data.days}
                                        </p>
                                    </div>
                                    <div className="col-lg-6">
                                        <p className="m-0 first-data">
                                            {data.dataValue}
                                        </p>
                                        <p className="second-data">
                                            {data.additionalInfo}
                                        </p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
                <div className=" row inventory-bottom-table my-5 bg-white py-4 m-0">
                    <div className="inventory-table-header mb-4">
                        <div className="row">
                            <div className="col-lg-7">
                                <h2>Products</h2>
                            </div>
                            <div className="col-lg-5 inventory-table-header-button text-end">
                                <Button className="me-2" variant="contained">
                                    Add Product
                                </Button>
                                <Button className="me-2" variant="outlined">
                                    <i className="bi bi-filter"></i>Filters
                                </Button>
                                <Button className="" variant="outlined">
                                    Download all
                                </Button>
                            </div>
                        </div>
                    </div>
                    <div className="inventory-table text-start">
                        <InventoryTable />
                    </div>
                </div>
            </div>
        </>
    );
}
