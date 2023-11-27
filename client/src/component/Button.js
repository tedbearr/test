const Button = ({ title, name, type, onclick }) => {
  return (
    <button
      onClick={onclick}
      className="px-4 rounded-md py-1 bg-pink-400 text-white"
      type={type}
      title={title}
    >
      {name}
    </button>
  );
};

export default Button;
