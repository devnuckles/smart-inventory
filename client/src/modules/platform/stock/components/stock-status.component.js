import { SmallChip } from "../../../core";

function StockStatus() {
    return (
        <div class="row">
            <div class="col-lg-2">
                <img
                    style={{
                        width: "50px",
                    }}
                    src="../images/product.png"
                    alt=""
                />
            </div>
            <div class="col-lg-7">
                <p className="m-0">Tata Salt</p>
                <span
                    className="m-0"
                    style={{
                        fontSize: "6px!important",
                    }}
                >
                    Remaining : 10 Packet
                </span>
            </div>
            <div class="col-lg-3">
                <SmallChip label="Low" />
            </div>
        </div>
    );
}

export default StockStatus;
