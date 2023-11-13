import { Button } from "@mui/material";
import { useState } from "react";
import { ModalInput } from "../../../core";
import axios from "axios";

export default function AddProduct() {
  const [formData, setFormData] = useState({
    productName: "",
    productCategory: "",
    buyingPrice: "",
    productQuantity: "",
    productExpiryDate: "",
    productThresholdValue: "",
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: value,
    }));
  };

  const handleSubmit = async () => {
    try {
      const response = await axios.post(
        "http://localhost:3001/api/products",
        formData
      );
      console.log(response.data);
      setFormData({
        productName: "",
        productCategory: "",
        buyingPrice: "",
        productQuantity: "",
        productExpiryDate: "",
        productThresholdValue: "",
      });
    } catch (error) {
      alert("Check all inupt fields");
      console.error(error);
    }
  };

  return (
    <div className="inventory-modal p-3 bg-white">
      <div className="inventory-modal-header">
        <h2 className="dashboard-card-heading">New Product</h2>
      </div>
      <div className="inventory-modal-form">
        <ModalInput
          label="Product Name"
          inputId="productName"
          placeholder="Enter product name"
          name="productName"
          value={formData.productName}
          onChange={handleChange}
        />
        <ModalInput
          label="Category"
          inputId="productCategory"
          placeholder="Select product category"
          name="productCategory"
          value={formData.productCategory}
          onChange={handleChange}
        />
        <ModalInput
          label="Buying Price"
          inputId="buyingPrice"
          placeholder="Enter buying price"
          name="buyingPrice"
          value={formData.buyingPrice}
          onChange={handleChange}
        />
        <ModalInput
          label="Quantity"
          inputId="productQuantity"
          placeholder="Enter product quantity"
          name="productQuantity"
          value={formData.productQuantity}
          onChange={handleChange}
        />
      </div>
      <Button variant="contained" onClick={handleSubmit}>
        Submit
      </Button>
    </div>
  );
}
