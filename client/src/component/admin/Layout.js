import { useContext, useState } from "react";
import { Outlet } from "react-router-dom";
import NotifToast from "../Toaster";
import Footer from "./Footer";
import Header from "./Header";
import Navbar from "./Navbar";

const Layout = () => {
  return (
    <>
      <NotifToast />
      <div className="flex flex-row w-full h-full">
        <div className="flex bg-pink-300 max-sm:bg-pink-300">
          <Navbar />
        </div>

        {/* <div className="w-1/4"></div> */}
        <div className={`flex flex-col w-full`}>
          <Header />
          <div className="flex flex-col p-10 w-full h-full ">
            <Outlet />
          </div>
          <Footer />
        </div>
      </div>
    </>
  );
};

export { Layout };
