import React from "react";

const Home = () => {
  return (
    <>
      <div className="flex flex-col space-y-10 flex-wrap w-full h-full p-10">
        <div
          title="aboutMe"
          className="flex w-full flex-wrap basis-full flex-col "
        >
          <span className="flex">Tentang Saya</span>
          <div className="flex flex-wrap">Saya adalah...</div>
        </div>
        <div
          title="skills"
          className="flex space-y-5 w-full flex-wrap flex-col"
        >
          <span className="flex">Skill</span>
          <div className="flex flex-wrap flex-col space-y-2">
            <span className="flex">Front End</span>
            <span className="flex">Lorem ipsum, Lorem ipsum, Lorem ipsum,</span>
          </div>
          <div className="flex flex-wrap flex-col space-y-2">
            <span className="flex">Backend End</span>
            <span className="flex">Lorem ipsum, Lorem ipsum, Lorem ipsum,</span>
          </div>
          <div className="flex flex-wrap flex-col space-y-2">
            <span className="flex">Language</span>
            <span className="flex">Lorem ipsum, Lorem ipsum, Lorem ipsum,</span>
          </div>
        </div>
      </div>
    </>
  );
};

export default Home;
