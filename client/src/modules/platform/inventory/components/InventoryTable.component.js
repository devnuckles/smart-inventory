import { useState, useEffect } from "react";
import DynamicTable from "../../../core/common/table.component";
import { columns } from "../inventory-table-tool";
import axios from "axios";

export default function InventoryTable() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get("http://localhost:3001/api/items/all");
        setProducts(response.data.data.Products);
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []);

  const convertTime = (time) => {
    const timestamp = time;
    const date = new Date(timestamp);

    return date.toLocaleDateString();
  };

  const rows = products.map((product) => ({
    products: product.name,
    buyingPrice: product.buying_price,
    quantity: product.quantity,
    thresholdValue: product.threshold_value,
    expiryDate: convertTime(product.expiry_date),
    availability: product.status,
  }));
  return <DynamicTable columns={columns} rows={rows} pagination={true} />;
}
