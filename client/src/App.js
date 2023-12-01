import Home from "./page/Home";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./component/Navbar";
import Footer from "./component/Footer";
import NotFound from "./component/NotFound";
import Project from "./page/Project";
import Experience from "./page/Experience";
import * as context from "./helper/Context";
import Layout from "./component/Layout";
import * as AdminLayout from "./component/admin/Layout";
import AdminHome from "./page/admin/Home";
import Login from "./page/admin/Login";
import Skill from "./page/admin/Skill";
import { useState } from "react";

function App() {
  const [expand, setExpand] = useState(true);
  const [expandMobile, setExpandMobile] = useState(false);
  return (
    <div className="flex flex-wrap flex-col w-full h-full p-0 m-0 font-serif box-border">
      <context.settingConfig.Provider
        value={{ expand, setExpand, expandMobile, setExpandMobile }}
      >
        <Router>
          <Routes>
            <Route element={<Layout />}>
              <Route path="/" element={<Home />} />\
              <Route path="project" element={<Project />} />
              <Route path="experience" element={<Experience />} />
            </Route>
            <Route path="/admin" element={<AdminLayout.Layout />}>
              <Route path="/admin/" element={<AdminHome />} />
              <Route path="/admin/project" element={<Project />} />
              <Route path="/admin/skills" element={<Skill />} />
            </Route>
            <Route path="/admin/Login" element={<Login />} />
            <Route path="*" element={<NotFound />} />
          </Routes>
        </Router>
      </context.settingConfig.Provider>
    </div>
  );
}

export default App;
