import { createBrowserRouter } from "react-router-dom";
import Login from "../modules/platform/users/component/Login.component";
import DashboardLayout from "../modules/core/common/dashboard/layout";
import Dashboard from "../modules/core/common/dashboard/Dashboard.component";
import Inventory from "../modules/core/common/dashboard/Inventory.component";
import InventoryModal from "../modules/core/common/dashboard/InventoryModal.component";

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
]);

export default Routes;
