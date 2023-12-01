import { useContext, useState } from "react";
import { settingConfig } from "../../helper/Context";

const Header = () => {
  let { expand, setExpand, expandMobile, setExpandMobile } =
    useContext(settingConfig);
  const Expand = () => {
    console.log(expand);
    setExpand(!expand);
    localStorage.setItem("isExpand", expand);
  };

  const Expand2 = () => {
    setExpandMobile(!expandMobile);
  };
  return (
    <>
      <div
        className={`h-10 items-center w-full flex justify-between p-2 border-b-2 border-pink-400`}
      >
        <span className="max-sm:flex hidden">
          <button onClick={Expand2}>
            <i className="fa fa-bars"></i>
          </button>
        </span>
        <span className="max-sm:hidden">
          <button onClick={Expand}>
            <i className="fa fa-bars"></i>
          </button>
        </span>
        <span>Header</span>
      </div>
    </>
  );
};

export default Header;
