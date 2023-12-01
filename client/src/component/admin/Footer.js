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
        }  max-sm:fixed bottom-0 flex h-10 border-t-2 p-2 w-full justify-between items-center border-pink-400`}
      >
        <span>Footer</span>
      </footer>
    </>
  );
};

export default Footer;
