import { Outlet } from "react-router-dom";
import Footer from "./Footer";
import Header from "./Header";
import Navbar from "./Navbar";

const Layout = () => {
  return (
    <>
      <div className="flex flex-row w-full">
        <Navbar />
        <div className="flex flex-col w-full">
          <Header />
          <div className="flex flex-col space-y-4 p-10 w-full ">
            <Outlet />
          </div>
          <Footer />
        </div>
      </div>
    </>
  );
};

export { Layout };
