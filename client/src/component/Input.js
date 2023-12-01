import { useForm } from "react-hook-form";

const Input = ({ type, placeholder, formName }) => {
  // const { register } = useForm();
  return (
    <>
      <input
        className="outline-pink-400 px-1 border-2 w-full"
        type={type}
        placeholder={placeholder}
        // {...register("name", { required: true })}
      />
    </>
  );
};

export default Input;
