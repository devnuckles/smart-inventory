function Card({ iconClass, amount, label, className }) {
    return (
        <div className={className ? className : "col-lg-3"}>
            <div className="dashboard-card-content p-2">
                <div className="dashboard-card-icon text-center mb-2">
                    <i
                        className={`text-center bi ${
                            iconClass ? iconClass : "bi-coin"
                        }`}
                    ></i>
                </div>
                <div className="dashboard-card-details">
                    <span className="text-start w-50 d-inline-block">
                        {amount ? amount : 20}
                    </span>
                    <span className="text-end w-50 d-inline-block">
                        {label ? label : "demo label"}
                    </span>
                </div>
            </div>
        </div>
    );
}

export default Card;
