import { Link } from "react-router-dom";

const Navbar = () => {
  return (
    <>
      <div className="flex text-white text-lg text flex-wrap space-x-10 items-center justify-center w-full bg-pink-400 h-14">
        <Link to="/">Home</Link>
        <Link to="project">Project</Link>
        <Link to="experience">Experience</Link>
      </div>
    </>
  );
};

export default Navbar;
