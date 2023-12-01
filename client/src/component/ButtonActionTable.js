const ButtonActionTable = ({ onclick, type, title, icon }) => {
  return (
    <>
      <button
        onClick={onclick}
        className="px-2 py-2 hover:bg-pink-500 outline-none rounded-md flex justify-center items-center align-middle  bg-pink-400 text-white"
        type={type}
        title={title}
      >
        <i className={icon}></i>
      </button>
    </>
  );
};

export default ButtonActionTable;
