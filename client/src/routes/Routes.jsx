import { createBrowserRouter } from "react-router-dom";
import Login from "../modules/platform/users/component/Login.component";
import DashboardLayout from "../modules/core/common/dashboard/layout";
import Dashboard from "../modules/core/common/dashboard/Dashboard.component";
import Inventory from "../modules/core/common/dashboard/Inventory.component";
import InventoryModal from "../modules/core/common/dashboard/InventoryModal.component";
import ProductView from "../modules/core/common/dashboard/ProductView.component";

const Routes = createBrowserRouter([
    {
        path: "/",
        element: <Login />,
    },
    {
        path: "/dashboard",
        element: (
            <DashboardLayout>
                <Dashboard />
            </DashboardLayout>
        ),
    },
    {
        path: "/inventory",
        element: (
            <DashboardLayout>
                <Inventory />
            </DashboardLayout>
        ),
    },
    {
        path: "/inventory-modal",
        element: (
            <DashboardLayout>
                <InventoryModal />
            </DashboardLayout>
        ),
    },
    {
        path: "/product-view",
        element: (
            <DashboardLayout>
                <ProductView />
            </DashboardLayout>
        ),
    },
]);

export default Routes;
