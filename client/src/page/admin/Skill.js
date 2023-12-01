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
// import "dotenv";

const Skill = () => {
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
      console.log(res.data.data.type);
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
    // reset();
    setModalStatus(!modalStatus);
  };

  const column = [
    {
      name: "NAME",
      selector: (row) => row.type,
    },
    {
      name: "SKILL",
      selector: (row) => row.skill,
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
    await axios.get(url + "/skill/").then((res) => {
      console.log(res.data.data);
      setTableData(res.data.data);
    });
  };

  const onSubmit = async (data) => {
    console.log(data);
    let urlForm = isEdit
      ? url + "/skill/update/" + idEdit
      : url + "/skill/insert";
    await axios.post(urlForm, data).then(async (res) => {
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
        <span>Skill</span>
        <span className="w-full flex justify-end">
          <Button icon="fa fa-plus" onclick={openModalCreate} name="Create" />
        </span>
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
        <div className="flex flex-wrap w-full text-lg">
          <form className="flex w-full flex-col space-y-4 flex-wrap flex-grow ">
            <div className="flex flex-row flex-grow space-x-5">
              <label className="w-1/2">Type</label>
              <span className="w-full">
                <input
                  type="text"
                  placeholder="name..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("type", { required: true })}
                />
                <span className="text-red-500">
                  {errors.type?.type === "required" && "type is required"}
                </span>
              </span>
            </div>
            <div className="flex flex-row w-full space-x-5">
              <label className="w-1/2">Skill</label>
              <span className="w-full">
                <input
                  type="text"
                  placeholder="skill..."
                  className="outline-pink-400 px-1 border-2 w-full"
                  {...register("skill", { required: true })}
                />
                <span className="text-red-500">
                  {errors.skill?.type === "required" && "skill is required"}
                </span>
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
      >
        <div className="flex flex-wrap w-full text-lg">
          <form className="flex w-full flex-col space-y-4 flex-wrap flex-grow ">
            <div className="flex flex-row w-full space-x-5">
              <label className="w-1/2">
                Type <span>:</span>
              </label>
              <span className="w-full">
                <span> {viewData.type} </span>
              </span>
            </div>
            <div className="flex flex-row w-full space-x-5">
              <label className="w-1/2">
                Skill <span>:</span>
              </label>
              <span className="w-full">
                <span> {viewData.skill} </span>
              </span>
            </div>
            <div className="flex flex-row w-full space-x-5">
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

export default Skill;
