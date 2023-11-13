import { createBrowserRouter } from "react-router-dom";
import {
    Login,
    Dashboard,
    Inventory,
    ProductView,
    SupplierTable,
    OrderComponent,
    Report,
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
                <SupplierTable />
            </DashboardLayout>
        ),
    },
    {
        path: "/orders",
        element: (
            <DashboardLayout>
                <OrderComponent />
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
]);

export default Routes;
