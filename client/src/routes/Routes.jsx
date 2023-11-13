import { createBrowserRouter } from "react-router-dom";
import { Login, Dashboard, Inventory, ProductView } from "../modules/platform";
import { DashboardLayout } from "../modules/core";

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
        path: "/product-view",
        element: (
            <DashboardLayout>
                <ProductView />
            </DashboardLayout>
        ),
    },
    // {
    //     path: "/test-route",
    //     element: (
    //         <DashboardLayout>
    //             <InventoryModal />
    //         </DashboardLayout>
    //     ),
    // },
]);

export default Routes;
