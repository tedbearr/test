const ButtonClose = ({ title, name, type, onclick }) => {
  return (
    <button
      onClick={onclick}
      className="px-4 outline-none hover:bg-pink-500 rounded-md flex justify-center items-center align-middle space-x-2 py-1 bg-pink-600 text-white"
      type={type}
      title={title}
    >
      <span>{name}</span>
    </button>
  );
};

export default ButtonClose;
