import { createBrowserRouter } from "react-router-dom";
import Login from "../modules/platform/users/component/Login.component";
import DashboardHeader from "../modules/core/common/dashboard/DashboardHeader.component";
import DashboardLayout from "../modules/core/common/dashboard/layout";
import Dashboard from "../modules/core/common/dashboard/Dashboard.component";
import Inventory from "../modules/core/common/dashboard/Inventory.component";

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
]);

export default Routes;
