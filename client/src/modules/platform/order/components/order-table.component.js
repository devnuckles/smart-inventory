import { DynamicTable } from "../../../core";
import { columns, rows } from "../order-table-tool";
import { useState, useEffect } from "react";
import axios from "axios";

export default function OrderTable() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get(
          "http://localhost:3001/api/orders/all"
        );
        setOrders(response.data.data.orders);
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

  const rows = orders.map((order) => ({
    id: order.order_id,
    products: order.product_name,
    value: order.buying_price,
    quantity: order.quantity,
    expection: convertTime(order.delivery_date),
  }));

  return (
    <>
      <div className="inventory-parent row inventory-bottom-table my-5 bg-white py-4 m-0">
        <div className="inventory-table text-start">
          <DynamicTable columns={columns} rows={rows} pagination={true} />
        </div>
      </div>
    </>
  );
}
