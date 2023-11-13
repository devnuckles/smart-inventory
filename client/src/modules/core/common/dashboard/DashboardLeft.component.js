import DashboardLeftNav from "../dashboard/DashboardLeftNav.component";
export default function DashboardLeft() {
    return (
        <>
            <div className="row">
                <div className="col-lg-12 dashboard-left p-3">
                    <div className="dashboard-left-header text-center">
                        <img className="me-2" src="../images/logo.png"></img>
                        <h2 className="d-inline-block">KANBAN</h2>
                    </div>
                    <div className="dashboard-left-menu">
                        <DashboardLeftNav />
                    </div>
                    <div className="dashboard-left-bottom-buttons ">
                        <a className="d-inline-block w-100 text-decoration-none my-2">
                            <i class="bi bi-gear me-2"></i>
                            Settings
                        </a>
                        <a className="d-inline-block w-100 text-decoration-none my-2">
                            <i class="bi bi-box-arrow-right me-2"></i> Logout
                        </a>
                    </div>
                </div>
            </div>
        </>
    );
}
