const Input = ({ type, placeholder }) => {
  return (
    <>
      <input
        className="outline-pink-400 px-1 border-2 w-full"
        type={type}
        placeholder={placeholder}
      />
    </>
  );
};

export default Input;
