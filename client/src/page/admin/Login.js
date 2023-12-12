import { useContext, useEffect, useState } from "react";
import Button from "../../component/Button";
import axios from "axios";
import { baseUrl } from "../../helper/Context";
import { toast } from "react-hot-toast";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const url = useContext(baseUrl);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const dataLogin = {
    username: username,
    password: password,
  };

  const login = async () => {
    await axios
      .post(`${url}/auth/login`, dataLogin)
      .then(async (data) => {
        console.log(data.data.data.accessToken);
        if (data.data.code != "00") {
          toast.error(data.data.message);
        } else {
          await toast.success("success");
          localStorage.setItem("accessToken", data.data.data.accessToken);
          navigate("/admin");
        }
      })
      .catch((e) => {
        toast.error(e.message);
      });
  };

  useEffect(() => {
    console.log(url);
    // console.log(username);
  }, []);

  return (
    <>
      <div className="w-full h-screen max-sm:p-10 flex flex-wrap justify-center items-center bg-pink-100">
        <div className="px-10 max-sm:p-1 flex py-10 rounded-md border-2 border-pink-400 bg-white">
          <div className="flex w-full h-full">
            <form className="flex justify-center space-y-2 flex-wrap [&>input]:px-1 [&>input]:w-full [&>input]:h-10 [&>input]:outline-pink-400">
              <input
                type="text"
                onChange={(i) => setUsername(i.target.value)}
                placeholder="username"
              />
              <input
                type="text"
                onChange={(i) => setPassword(i.target.value)}
                placeholder="password"
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
