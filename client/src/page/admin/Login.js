import { useContext, useEffect, useState } from "react";
import Button from "../../component/Button";
import axios from "axios";
import { baseUrl } from "../../helper/Context";

const Login = () => {
  const url = useContext(baseUrl);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const dataLogin = {
    username: username,
    password: password,
  };

  const login = async () => {
    await axios
      .post(`${url._currentValue}/auth/login`, dataLogin)
      .then((data) => {
        console.log(data);
      });
  };

  useEffect(() => {
    console.log(url);
    // console.log(username);
  }, []);

  return (
    <>
      <div className="w-screen h-screen flex flex-wrap justify-center items-center bg-pink-100">
        <div className="w-1/3 p-4 h-1/3 basis-1/3 rounded-md border-2 border-pink-400 bg-white">
          <div className="flex flex-wrap w-full">
            <form className="flex justify-center space-y-2 flex-wrap [&>input]:px-1 [&>input]:w-full [&>input]:h-10 [&>input]:outline-pink-400">
              <input
                type="text"
                onChange={(i) => setUsername(i.target.value)}
                placeholder="test"
              />
              <input
                type="text"
                onChange={(i) => setPassword(i.target.value)}
                placeholder="test"
              />
              <Button
                type="button"
                onclick={login}
                name={"Login"}
                title="button"
              />
            </form>
          </div>
        </div>
      </div>
    </>
  );
};

export default Login;
