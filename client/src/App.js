import Home from "./page/Home";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./component/Navbar";
import Footer from "./component/Footer";
import NotFound from "./component/NotFound";
import Project from "./page/Project";
import Experience from "./page/Experience";
import * as context from "./helper/Context";

function App() {
  return (
    <div className="flex flex-wrap flex-col w-full h-full p-0 m-0 font-serif bg-pink-100">
      <context.baseUrl.Provider>
        <Router>
          <Navbar />
          <Routes>
            <Route path="/" element={<Home />} />\
            <Route path="project" element={<Project />} />
            <Route path="experience" element={<Experience />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
          <Footer />
        </Router>
      </context.baseUrl.Provider>
    </div>
  );
}

export default App;
