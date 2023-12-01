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
      <div className="flex flex-row w-full">
        <Navbar />
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
