import React from "react";
import Button from "@mui/material/Button";
import OrderTable from "./order-table.component";
import { DynamicModal } from "../../../core";
import { AddProduct } from "../../product";

export default function Order() {
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
            <div className="row inventory-card p-3 m-0">
                <div className="col-lg-12 ">
                    <h2 className="dashboard-card-heading">Overall Orders</h2>
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
                                    <p className="second-data">{data.days}</p>
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
                            <h2>Orders</h2>
                        </div>
                        <div className="col-lg-5 inventory-table-header-button text-end">
                            <div className="row">
                                <div className="col-lg-4 p-0">
                                    <DynamicModal
                                        Element={<AddProduct />}
                                        label="Add Order"
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
                    <OrderTable />
                </div>
            </div>
        </div>
    );
}
