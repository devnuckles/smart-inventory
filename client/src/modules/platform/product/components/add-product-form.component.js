// import { Formik, Form, Field, ErrorMessage } from "formik";
// import * as Yup from "yup";
// import { ModalInput } from "../../../core";
// import { Button } from "@mui/material";
// import axios from "axios";
// import { LineWave } from "react-loadingg";

// const validationSchema = Yup.object().shape({
//   productName: Yup.string().required("Product name is required"),
//   productCategory: Yup.string().required("Product category is required"),
//   buyingPrice: Yup.number()
//     .typeError("Buying price must be a number")
//     .required("Buying price is required"),
//   productQuantity: Yup.number()
//     .typeError("Product quantity must be a number")
//     .required("Product quantity is required"),
//   productExpiryDate: Yup.string().required("Expiry date is required"),
//   productThresholdValue: Yup.number()
//     .typeError("Threshold value must be a number")
//     .required("Threshold value is required"),
// });

// export default function AddProduct() {
//   const initialValues = {
//     productName: "",
//     productCategory: "",
//     buyingPrice: "",
//     productQuantity: "",
//     productExpiryDate: "",
//     productThresholdValue: "",
//   };
//   const [isLoading, setIsLoading] = useState(false);

//   const handleSubmit = async (values) => {
//     try {
//       const response = await axios.post(
//         "http://localhost:3001/api/products",
//         values
//       );
//       console.log(response.data);
//       alert("Product added successfully!");
//     } catch (error) {
//       alert("Check all input fields");
//       console.error(error);
//     }
//   };

//   return (
//     <div className="inventory-modal p-3 bg-white">
//       <div className="inventory-modal-header">
//         <h2 className="dashboard-card-heading">New Product</h2>
//       </div>
//       <Formik
//         initialValues={{
//           initialValues,
//         }}
//         validationSchema={validationSchema}
//         onSubmit={handleSubmit}
//       >
//         {(formikProps) => (
//           <Form onSubmit={formikProps.handleSubmit}>
//             <div className="input-field">
//               <label htmlFor="productName" className="product-name-label">
//                 Product Name
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="productName"
//                 name="productName"
//                 placeholder="Enter product name"
//                 className="product-name-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="productName" component="div" />
//               </div>
//             </div>

//             <div className="input-field">
//               <label
//                 htmlFor="productCategory"
//                 className="product-category-label"
//               >
//                 Product Category
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="productCategory"
//                 name="productCategory"
//                 placeholder="Enter product category"
//                 className="product-category-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="productCategory" component="div" />
//               </div>
//             </div>

//             <div className="input-field">
//               <label htmlFor="buyingPrice" className="buying-price-label">
//                 Buying Price
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="buyingPrice"
//                 name="buyingPrice"
//                 placeholder="Enter buying price"
//                 className="buying-price-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="buyingPrice" component="div" />
//               </div>
//             </div>

//             <div className="input-field">
//               <label
//                 htmlFor="productQuantity"
//                 className="product-quantity-label"
//               >
//                 Product Quantity
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="productQuantity"
//                 name="productQuantity"
//                 placeholder="Enter product quantity"
//                 className="product-quantity-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="productQuantity" component="div" />
//               </div>
//             </div>

//             <div className="input-field">
//               <label htmlFor="productExpiryDate" className="expiry-date-label">
//                 Expiry Date
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="productExpiryDate"
//                 name="productExpiryDate"
//                 placeholder="Enter expiry date"
//                 className="expiry-date-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="productExpiryDate" component="div" />
//               </div>
//             </div>

//             <div className="input-field">
//               <label
//                 htmlFor="productThresholdValue"
//                 className="threshold-value-label"
//               >
//                 Threshold Value
//               </label>
//               <br />
//               <Field
//                 style={{ marginTop: 20 }}
//                 id="productThresholdValue"
//                 name="productThresholdValue"
//                 placeholder="Enter threshold value"
//                 className="threshold-value-field"
//               />
//               <div className="formik-error-message">
//                 <ErrorMessage name="productThresholdValue" component="div" />
//               </div>
//             </div>

//             {isLoading ? (
//               <LineWave
//                 height="90"
//                 width="350"
//                 color="#43c7d9"
//                 ariaLabel="line-wave"
//                 visible={true}
//               />
//             ) : (
//               <button type="submit" className="add-product-button">
//                 Add Product
//               </button>
//             )}
//           </Form>
//         )}
//       </Formik>
//     </div>
//   );
// }

import { Button } from "@mui/material";
import { useState } from "react";
import { ModalInput } from "../../../core";
import axios from "axios";

export default function AddProduct() {
  const [formData, setFormData] = useState({
    product_name: "",
    category: "",
    buying_price: "",
    productQuantity: "",
    delivery_date: "",
    vendor_email: "",
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
        "http://localhost:3001/api/items/create",
        formData
      );
      console.log(response.data);
      setFormData({
        product_name: "",
        category: "",
        buying_price: null,
        quantity: null,
        delivery_date: null,
        vendor_email: "",
      });
    } catch (error) {
      alert("Check all input fields");
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
          inputId="product_name"
          placeholder="Enter product name"
          name="product_name"
          value={formData.product_name}
          onChange={handleChange}
        />
        <ModalInput
          label="Category"
          inputId="category"
          placeholder="Select product category"
          name="category"
          value={formData.category}
          onChange={handleChange}
        />
        <ModalInput
          label="Buying Price"
          inputId="buying_price"
          placeholder="Enter buying price"
          name="buying_price"
          value={formData.buying_price}
          onChange={handleChange}
        />
        <ModalInput
          label="Quantity"
          inputId="quantity"
          placeholder="Enter product quantity"
          name="quantity"
          value={formData.quantity}
          onChange={handleChange}
        />
        <ModalInput
          label="Delivery Date"
          inputId="delivery_date"
          placeholder="Enter Delivery Date"
          name="delivery_date"
          value={formData.delivery_date}
          onChange={handleChange}
        />
        <ModalInput
          label="Vendor Email"
          inputId="vendor_email"
          placeholder="Enter Vendor Email"
          name="vendor_email"
          value={formData.vendor_email}
          onChange={handleChange}
        />
        <Button variant="contained" onClick={handleSubmit}>
          Add Product
        </Button>
      </div>
    </div>
  );
}
