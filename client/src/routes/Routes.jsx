import { createBrowserRouter } from "react-router-dom";
import {
    Login,
    Dashboard,
    Inventory,
    ProductView,
    SupplierTable,
    OrderComponent,
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
]);

export default Routes;
