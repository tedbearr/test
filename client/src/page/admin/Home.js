import axios from "axios";
import { useContext, useEffect } from "react";
import { baseUrlAdmin } from "../../helper/Context";

const Home = () => {
  const url = useContext(baseUrlAdmin);

  const getData = async () => {
    await axios
      .get(url + "/skill/")
      .then((data) => {
        console.log(data);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  useEffect(() => {
    getData();
  });

  return (
    <>
      <span className="w-full border-2 p-4 rounded-lg border-pink-400">
        Welcome!
      </span>
    </>
  );
};

export default Home;
