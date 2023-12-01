import { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Menu } from "../../helper/MenuList";
import { settingConfig } from "../../helper/Context";

const Navbar = () => {
  let { expand, expandMobile, setExpandMobile } = useContext(settingConfig);
  const [menu, setMenu] = useState(Menu);
  // let [isExpand, setIsExpand] = useState(localStorage.getItem("isExpand"));
  let isExpand = localStorage.getItem("isExpand");
  const [location, setLocation] = useState();
  useEffect(() => {
    // setIsExpand(localStorage.getItem("isExpand"));
    setMenu(Menu);
    setLocation(document.location.pathname);
  }, []);

  return (
    <>
      <div
        className={`hidden ${
          expandMobile ? "block" : "hidden"
        } max-sm:absolute max-sm:inset-0 max-sm:bg-black/70 max-sm:backdrop-blur-sm max-sm:h-full max-sm:w-full max-sm:z-10`}
        onClick={() => setExpandMobile(false)}
      ></div>

      <aside
        className={`${expand ? "flex" : "hidden"} ${
          expandMobile
            ? "max-sm:block max-sm:-mr-40 max-sm:z-10 max-sm:w-10"
            : "max-sm:hidden"
        } space-y-10 flex-wrap text-white flex-col basis-60 bg-pink-300 h-screen max-sm:max-w-[10rem]`}
      >
        <div className="flex w-full justify-between p-2 relative">
          <span>a</span>
          <button
            onClick={() => setExpandMobile(!expandMobile)}
            className="max-sm:flex hidden"
          >
            <i className="fa fa-bars"></i>
          </button>
        </div>
        <div className="flex [&>ul]:rounded-md [&>ul]:space-x-2 [&>ul]:w-full [&>ul]:flex [&>ul]:items-baseline [&>ul]:p-2 space-y-4 justify-center items-center flex-wrap box-border w-full px-1">
          {menu.map((data) => {
            return (
              <ul
                className={
                  location == data.link ? "bg-pink-500" : "hover:bg-pink-500"
                }
                key={data.icon}
              >
                <i className={data.icon}></i>
                <Link
                  onClick={() => setLocation(data.link)}
                  className="w-full"
                  id={data.link}
                  to={data.link}
                >
                  <span>{data.name}</span>
                </Link>
              </ul>
            );
          })}
        </div>
      </aside>
    </>
  );
};

export default Navbar;
