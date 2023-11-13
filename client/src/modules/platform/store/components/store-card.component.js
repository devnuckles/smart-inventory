import { Button } from "@mui/material";

export default function StoreCard() {
    return (
        <div class="manage-store mt-2">
            <div class="row">
                <div class="col-lg-3">
                    <div class="manage-store-card-image">
                        <img src="../images/google-icon.png" alt="" />
                    </div>
                </div>
                <div class="col-lg-7">
                    <div class="manage-store-card-content">
                        <h2 class="dashboard-card-heading"></h2>
                        <p>1A/Krihnarajapuram, 3 rd street sulur</p>
                        <p>Coimbatore - 6313403</p>
                        <p>044- 653578</p>
                    </div>
                </div>
                <div class="col-lg-2 ms-auto text-end">
                    <div class="manage-store-card-button">
                        <Button variant="outlined">Outlined</Button>
                    </div>
                </div>
            </div>
        </div>
    );
}
