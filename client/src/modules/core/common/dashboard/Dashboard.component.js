import React from "react";

const DashboardCard = ({ iconClass, amount, label }) => (
    <div className="dashboard-card-content p-2">
        <div className="dashboard-card-icon text-center">
            <i className={`bi ${iconClass}`}></i>
        </div>
        <div className="dashboard-card-details">
            <span className="text-start w-50 d-inline-block">{amount}</span>
            <span className="text-end w-50 d-inline-block">{label}</span>
        </div>
    </div>
);

export default function Dashboard() {
    const cardData = [
        { icon: "bi-receipt-cutoff", amount: "₹ 832", label: "Sales" },
        {
            icon: "bi-bar-chart-line-fill",
            amount: "₹ 18,300",
            label: "Revenue",
        },
        { icon: "bi-graph-up-arrow", amount: "₹ 868", label: "Profit" },
        { icon: "bi-coin", amount: "₹ 17,432", label: "Cost" },
        // Add more card data as needed
    ];

    const inventoryData = [
        { icon: "bi-box2", amount: "868", label: "Quantity in Hand" },
        { icon: " bi-geo-alt", amount: "200", label: "To be received" },
        // Add more inventory data as needed
    ];

    const purchaseData = [
        { icon: "bi-cart", amount: "82", label: "Purchase" },
        { icon: "bi-coin", amount: "₹ 13,573", label: "Cost" },
        { icon: "bi-x-circle", amount: "5", label: "Cancel" },
        { icon: " bi-graph-down-arrow", amount: "₹17,432", label: "Return" },
        // Add more purchase data as needed
    ];

    const productSummary = [
        {
            icon: "bi-person-circle",
            amount: "31",
            label: "Number of Suppliers",
        },
        {
            icon: "bi-list-columns-reverse",
            amount: "21",
            label: "Number of Categories",
        },
        // Add more inventory data as needed
    ];

    const renderDashboardCards = (data, columnSize) => (
        <div className={`col-lg-${columnSize} bg-white p-0`}>
            <h2 className="dashboard-card-heading">{data.title}</h2>
            <div className="row">
                {data.cards.map((card, index) => (
                    <div
                        className={`col-lg-${data.cardSize} dashboard-inventory-summary mb-4`}
                        key={index}
                    >
                        <DashboardCard
                            iconClass={`bi ${card.icon}`}
                            amount={card.amount}
                            label={card.label}
                        />
                    </div>
                ))}
            </div>
        </div>
    );

    return (
        <>
            <div className="dashboard-parent p-4">
                <div className="row dashboard-content-top ">
                    {renderDashboardCards(
                        {
                            title: "Sales Overview",
                            cards: cardData,
                            cardSize: 3,
                        },
                        8
                    )}
                    {renderDashboardCards(
                        {
                            title: "Inventory Summary",
                            cards: inventoryData,
                            cardSize: 6,
                        },
                        4
                    )}
                </div>

                <div className="row dashboard-content-middle mt-4 ">
                    {renderDashboardCards(
                        {
                            title: "Purchase Overview",
                            cards: purchaseData,
                            cardSize: 3,
                        },
                        8
                    )}
                    {renderDashboardCards(
                        {
                            title: "Product Summary",
                            cards: productSummary,
                            cardSize: 6,
                        },
                        4
                    )}
                </div>
            </div>
        </>
    );
}
