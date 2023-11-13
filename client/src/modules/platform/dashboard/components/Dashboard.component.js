import { InventoryCard } from "../../inventory";
import { ProductCard } from "../../product";
import { PurchaseCard } from "../../purchase";
import { SalesCard } from "../../sales";
import ProductSummary from "../chart/product-summary";
import SalesPurchase from "../chart/sales-purchase";

export default function Dashboard() {
    return (
        <>
            <div class="row mb-4">
                <SalesCard />
                <InventoryCard />
                <PurchaseCard />
                <ProductCard />
            </div>
            <div className="row">
                <div className="col-lg-8">
                    <div class="card mb-4">
                        <div class="card-body">
                            <SalesPurchase />{" "}
                        </div>
                    </div>
                </div>
                <div className="col-lg-4">
                    <div class="card mb-4">
                        <div class="card-body">
                            <ProductSummary />
                        </div>
                    </div>
                </div>
            </div>
        </>
    );
}
