import { PurchaseOverviewCard } from "../../purchase";
import { BestSalesCard, ProfitRevenue, BestSalesProducts } from "../../sales";

export default function Report() {
    return (
        <div className=" inventory-parent p-3">
            <div class="row mb-4">
                <PurchaseOverviewCard />
                <BestSalesCard />
            </div>
            <div class="row mb-4 bg-white m-2">
                <ProfitRevenue />
            </div>
            <div class="row m-2 ">
                <BestSalesProducts />
            </div>
        </div>
    );
}
