import { DynamicTable } from "../../../core";
import { columns, rows } from "../supplier-table-tool";
import { useState, useEffect } from "react";
import axios from "axios";
export default function SupplierTable() {
  const [suppliers, setSuppliers] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(
          "http://localhost:3001/api/suppliers/all"
        );
        setSuppliers(response.data.data.Suppliers);
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []);

  const rows = suppliers.map((supplier) => ({
    supplier: supplier.Name,
    product: supplier.Product,
    contact: supplier.Contact,
    email: supplier.Email,
  }));

  return <DynamicTable columns={columns} rows={rows} pagination={true} />;
}
