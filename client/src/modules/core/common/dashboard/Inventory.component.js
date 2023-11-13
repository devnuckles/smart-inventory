import React from "react";

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
                                        <p>{data.count}</p>
                                        <p>{data.days}</p>
                                    </div>
                                    <div className="col-lg-6">
                                        <p>{data.dataValue}</p>
                                        <p>{data.additionalInfo}</p>
                                    </div>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
                <div className="col-lg-12 inventory-bottom-table"></div>
            </div>
        </>
    );
}
