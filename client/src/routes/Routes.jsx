import { createBrowserRouter } from "react-router-dom";
import {
    Login,
    Dashboard,
    Inventory,
    ProductView,
    Supplier,
    Order,
    Report,
    Store,
} from "../modules/platform";
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
    {
        path: "/suppliers",
        element: (
            <DashboardLayout>
                <Supplier />
            </DashboardLayout>
        ),
    },
    {
        path: "/orders",
        element: (
            <DashboardLayout>
                <Order />
            </DashboardLayout>
        ),
    },
    {
        path: "/reports",
        element: (
            <DashboardLayout>
                <Report />
            </DashboardLayout>
        ),
    },
    {
        path: "/manage-store",
        element: (
            <DashboardLayout>
                <Store />
            </DashboardLayout>
        ),
    },
]);

export default Routes;
