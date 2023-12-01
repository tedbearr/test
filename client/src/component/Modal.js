import { Children } from "react";
import Button from "./Button";
import ButtonClose from "./ButtonClose";

const Modal = ({ status, title, children, saveFunction, closeFunction }) => {
  return (
    <>
      <div
        className={` ${
          status ? "block" : "hidden"
        } fixed flex max-sm:h-full max-sm:p-10 p-40 left-0 items-center top-0 w-full h-full overflow-auto z-50 bg-black bg-opacity-40`}
      >
        <div className="flex space-y-8 align-middle justify-center flex-col w-full h-auto bg-white rounded-2xl p-5">
          <div className="flex justify-start text-lg font-semibold">
            <span>{title}</span>
          </div>
          <div className="flex w-full h-full flex-col flex-wrap overflow-auto">
            {Children.map(
              children,
              (child) =>
                // console.log(child.props.children);
                child
            )}
          </div>
          <div className="flex flex-wrap space-x-1 justify-end items-center">
            {saveFunction == "none" ? (
              ""
            ) : (
              <Button name="Save" onclick={saveFunction} />
            )}

            <ButtonClose name="Close" onclick={closeFunction} />
          </div>
        </div>
      </div>
    </>
  );
};

export default Modal;
