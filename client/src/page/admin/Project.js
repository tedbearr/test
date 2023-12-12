import axios from "axios";
import { useContext, useEffect, useState } from "react";
import { baseUrlAdmin } from "../../helper/Context";
import DataTable from "react-data-table-component";
import Button from "../../component/Button";
import Modal from "../../component/Modal";
import ButtonActionTable from "../../component/ButtonActionTable";
import Input from "../../component/Input";
import { useForm } from "react-hook-form";
import { toast } from "react-hot-toast";

const Project = () => {
  const [tableData, setTableData] = useState();
  let [modalStatus, setModalStatus] = useState(false);
  let [modalViewStatus, setModalViewStatus] = useState(false);
  let [isEdit, setIsEdit] = useState(false);
  let [idEdit, setIdEdit] = useState(0);
  let [viewData, setViewData] = useState({});
  let url = useContext(baseUrlAdmin);

  const {
    register,
    getValues,
    reset,
    setValue,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const View = async (id) => {
    reset();
    await axios.get(url + "/skill/" + id).then(async (res) => {
      setViewData(res.data.data);
    });
    console.log(id);
    setModalViewStatus(!modalViewStatus);
  };

  const Update = async (id) => {
    setIsEdit(true);
    setIdEdit(id);
    reset();
    await axios.get(url + "/skill/" + id).then((res) => {
      setValue("type", res.data.data.type);
      setValue("skill", res.data.data.skill);
    });
    setModalStatus(!modalStatus);
    console.log("update");
  };

  const Delete = async (id) => {
    await axios.post(url + "/skill/delete/" + id).then(async (res) => {
      // return console.log(res.data.code);
      if (res.data.code != "00") {
        toast.error(res.data.message);
      } else {
        await getData();
        toast.success("success delete");
      }
    });
    console.log("delete");
  };

  const openModalCreate = () => {
    setIsEdit(false);
    setModalStatus(!modalStatus);
  };

  const closeModal = () => {
    setIsEdit(false);
    setIdEdit(0);
    reset();
    setModalStatus(!modalStatus);
  };

  const column = [
    {
      name: "NAME",
      selector: (row) => row.name,
    },
    {
      name: "IMAGE",
      selector: (row) => row.image,
    },
    {
      name: "TECH",
      selector: (row) => row.tech,
    },
    {
      name: "LINK",
      selector: (row) => row.link,
    },
    {
      name: "ACTION",
      button: true,
      width: "300px",
      cell: (row) => (
        <>
          <div className="flex flex-wrap justify-center w-full space-x-1">
            <ButtonActionTable onclick={() => View(row.id)} icon="fa fa-eye" />
            <ButtonActionTable
              onclick={() => Update(row.id)}
              icon="fa fa-edit"
            />
            <ButtonActionTable
              onclick={() => Delete(row.id)}
              icon="fa fa-trash"
            />
          </div>
        </>
      ),
    },
  ];

  const getData = async () => {
    await axios
      .get(url + "/project/")
      .then((res) => {
        console.log(res.data.data);
        setTableData(res.data.data);
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const onSubmit = async (data) => {
    const formFile = new FormData();
    formFile.append("file", data.image[0]);
    formFile.append("name", data.name);
    formFile.append("link", data.link);
    formFile.append("tech", data.tech);
    console.log(data.name);
    let urlForm = isEdit
      ? url + "/project/update/" + idEdit
      : url + "/project/insert";
    await axios
      .post(urlForm, formFile, {
        headers: { "Content-Type": "multipart/form-data" },
      })
      .then(async (res) => {
        if (res.data.code != "00") {
          toast.error(res.data.message);
        } else {
          toast.success("success");
          await getData();
          setModalStatus(!modalStatus);
          reset();
        }
      });
  };

  const closeViewModal = () => {
    setModalViewStatus(!modalViewStatus);
    reset();
  };

  useEffect(() => {
    let asasa = process.env.REACT_APP_TEST;
    console.log(process.env);
    getData();
  }, []);

  return (
    <>
      <div className="flex flex-col space-y-5">
        <span className="text-4xl">Project</span>
        <hr></hr>
        <span className="w-full flex justify-start">
          <Button icon="fa fa-plus" onclick={openModalCreate} name="Create" />
        </span>
        <hr></hr>
        {/* <div className="max-md:bg-black"> */}
        <DataTable
          columns={column}
          data={tableData}
          pagination
          // progressPending={true}
          responsive
          subHeaderAlign="right"
          subHeaderWrap
          striped
          direction="auto"
          dense
        />
        {/* </div> */}
      </div>

      <Modal
        status={modalStatus}
        title={isEdit ? "Update" : "Create"}
        key={1}
        closeFunction={closeModal}
        saveFunction={handleSubmit((data) => onSubmit(data))}
      >
        <div className="flex w-full text-lg max-sm:text-sm">
          <form className="flex w-full px-10 flex-col space-y-4 h-auto overflow-auto">
            <div className="flex flex-row  max-sm:flex-col max-sm:space-y-1 max-sm:justify-center max-sm:items-center space-x-5">
              <label className="w-1/2 max-sm:flex max-sm:items-center max-sm:justify-center max-sm:w-full">
                Name
              </label>
              <span className="w-full max-sm:flex max-sm:w-56 max-sm:justify-center max-sm:items-center">
                <input
                  type="text"
                  placeholder="name..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("name", { required: true })}
                />
                <span className="text-red-500 max-sm:hidden max-sm:w-full">
                  {errors.name?.type === "required" && "name is required"}
                </span>
              </span>
              <span className="text-red-500 hidden max-sm:flex max-sm:w-full max-sm:justify-center max-sm:items-center">
                {errors.name?.type === "required" && "name is required"}
              </span>
            </div>
            <div className="flex flex-row  max-sm:flex-col max-sm:space-y-2 max-sm:justify-center max-sm:items-center space-x-5">
              <label className="w-1/2 max-sm:flex max-sm:items-center max-sm:justify-center max-sm:w-full">
                Image
              </label>
              <span className="w-full max-sm:flex max-sm:w-56 max-sm:justify-center max-sm:items-center">
                <input
                  type="file"
                  placeholder="skill..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("image", { required: true })}
                />
                <span className="text-red-500 max-sm:hidden max-sm:w-full">
                  {errors.image?.type === "required" && "image is required"}
                </span>
              </span>
              <span className="text-red-500 hidden max-sm:flex max-sm:w-full max-sm:justify-center max-sm:items-center">
                {errors.image?.type === "required" && "image is required"}
              </span>
            </div>
            <div className="flex flex-row  max-sm:flex-col max-sm:space-y-1 max-sm:justify-center max-sm:items-center space-x-5">
              <label className="w-1/2 max-sm:flex max-sm:items-center max-sm:justify-center max-sm:w-full">
                Tech
              </label>
              <span className="w-full max-sm:flex max-sm:w-56 max-sm:justify-center max-sm:items-center">
                <input
                  type="text"
                  placeholder="name..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("tech", { required: true })}
                />
                <span className="text-red-500 max-sm:hidden max-sm:w-full">
                  {errors.tech?.type === "required" && "tech is required"}
                </span>
              </span>
              <span className="text-red-500 hidden max-sm:flex max-sm:w-full max-sm:justify-center max-sm:items-center">
                {errors.tech?.type === "required" && "tech is required"}
              </span>
            </div>
            <div className="flex flex-row  max-sm:flex-col max-sm:space-y-1 max-sm:justify-center max-sm:items-center space-x-5">
              <label className="w-1/2 max-sm:flex max-sm:items-center max-sm:justify-center max-sm:w-full">
                Link
              </label>
              <span className="w-full max-sm:flex max-sm:w-56 max-sm:justify-center max-sm:items-center">
                <input
                  type="text"
                  placeholder="link..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("link", { required: true })}
                />
                <span className="text-red-500 max-sm:hidden max-sm:w-full">
                  {errors.link?.type === "required" && "link is required"}
                </span>
              </span>
              <span className="text-red-500 hidden max-sm:flex max-sm:w-full max-sm:justify-center max-sm:items-center">
                {errors.link?.type === "required" && "link is required"}
              </span>
            </div>
          </form>
        </div>
      </Modal>

      <Modal
        status={modalViewStatus}
        title={"View"}
        key={2}
        closeFunction={closeViewModal}
        saveFunction={"none"}
      >
        <div className="flex flex-wrap w-full text-lg">
          <form className="flex w-full flex-col space-y-4 flex-wrap flex-grow ">
            <div className="flex max-sm:flex-wrap max-sm:flex-grow flex-row w-full space-x-5">
              <label className="w-1/2 flex">
                Type <span>:</span>
              </label>
              <span className="w-full flex">
                <span> {viewData.type} </span>
              </span>
            </div>
            <div className="flex max-sm:flex-wrap max-sm:flex-grow flex-row w-full space-x-5">
              <label className="w-1/2 flex">
                Skill <span>:</span>
              </label>
              <span className="w-full">
                <span> {viewData.skill} </span>
              </span>
            </div>
            <div className="flex max-sm:flex-wrap max-sm:flex-grow flex-row w-full space-x-5">
              <label className="w-1/2">
                Created At <span>:</span>
              </label>
              <span className="w-full">
                <span> {viewData.created_at} </span>
              </span>
            </div>
          </form>
        </div>
      </Modal>
    </>
  );
};

export default Project;
