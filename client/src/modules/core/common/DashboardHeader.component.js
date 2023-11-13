import { Search } from "lucide-react";

export default function DashboardHeader() {
  return (
    <>
      <div className="row">
        <div className="col-lg-12 dashboard-right-header py-5">
          <div className="row">
            <div className="col-lg-10 dashboard-right-header-left">
              <div className="dashboard-right-header-search">
                <i class="bi bi-search me-2"></i>

                <input
                  type="search"
                  className="form-control d-inline"
                  id="searchbar"
                  aria-describedby="searchHelp"
                  placeholder="Search product, supplier, order"
                />
              </div>
            </div>
            <div className="col-lg-2 dashboard-right-header-right text-end">
              <i class="bi bi-bell mx-4"></i>
              <img src="../images/profile-image.png"></img>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
