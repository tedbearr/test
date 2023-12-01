const Button = ({ title, name, type, onclick, icon }) => {
  return (
    <button
      onClick={onclick}
      className="px-4 outline-none hover:bg-pink-500 rounded-md flex justify-center items-center align-middle space-x-2 py-1 bg-pink-400 text-white"
      type={type}
      title={title}
    >
      {icon ? <i className={icon}></i> : ""}
      {name ? <span>{name}</span> : ""}
    </button>
  );
};

export default Button;
