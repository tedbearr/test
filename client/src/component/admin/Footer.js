import { useContext } from "react";
import { settingConfig } from "../../helper/Context";

const Footer = () => {
  let { expand } = useContext(settingConfig);
  // let expand
  return (
    <>
      <footer
        className={`${
          expand ? "" : "fixed"
        }  max-sm:fixed text-white bg-pink-400 mt-10 bottom-0 flex h-10 border-t-2 p-2 w-full justify-end items-center border-pink-400`}
      >
        <span>Footer</span>
      </footer>
    </>
  );
};

export default Footer;
