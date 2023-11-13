import { Button } from "@mui/material";
import ModalInput from "./ModalInput.component";

export default function InventoryModal() {
    return (
        <>
            <div className="inventory-modal p-3">
                <div className="inventory-modal-header">
                    <h2 className="dashboard-card-heading">New Product</h2>
                </div>
                <div className="inventory-modal-form">
                    <ModalInput
                        label="Product Name"
                        inputId="productName"
                        placeholder="Enter product name"
                    />

                    <ModalInput
                        label="Product ID"
                        inputId="productId"
                        placeholder="Enter product ID"
                    />

                    <ModalInput
                        label="Category"
                        inputId="productCategory"
                        placeholder="Select product category"
                    />
                    <ModalInput
                        label="Buying Price"
                        inputId="buyingPrice"
                        placeholder="Enter buying price"
                    />
                    <ModalInput
                        label="Quantity"
                        inputId="productQuantity"
                        placeholder="Enter product quantity"
                    />
                    <ModalInput
                        label="Unit"
                        inputId="productUnit"
                        placeholder="Enter product unit"
                    />
                    <ModalInput
                        label="Expiry Date"
                        inputId="productExpiryDate"
                        placeholder="Enter expiry date"
                    />
                    <ModalInput
                        label="Threshold Value"
                        inputId="productThresholdValue"
                        placeholder="Enter threshold value"
                    />

                    <div className="inventory-modal-button text-end">
                        <Button className="me-2 discard" variant="outlined">
                            Discard
                        </Button>
                        <Button className="add-product" variant="contained">
                            Add Product
                        </Button>
                    </div>
                </div>
            </div>
        </>
    );
}
