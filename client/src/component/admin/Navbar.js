import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Menu } from "../../helper/MenuList";

const Navbar = () => {
  const [menu, setMenu] = useState(Menu);

  const [location, setLocation] = useState();

  useEffect(() => {
    setMenu(Menu);
    setLocation(document.location.pathname);
  }, []);

  return (
    <>
      <aside className="flex space-y-10 flex-wrap text-white flex-col basis-60 bg-pink-300 h-screen">
        <div className="flex w-full">a</div>
        <div className="flex [&>ul]:rounded-md [&>ul]:space-x-2 [&>ul]:w-full [&>ul]:flex [&>ul]:items-baseline [&>ul]:p-2 space-y-4 justify-center items-center flex-wrap box-border w-full px-1">
          {menu.map((data) => {
            return (
              <ul
                onClick={() => setLocation(data.link)}
                className={
                  location == data.link ? "bg-pink-500" : "hover:bg-pink-500"
                }
                key={data.icon}
              >
                <i className={data.icon}></i>
                <Link className="w-full" id={data.link} to={data.link}>
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
