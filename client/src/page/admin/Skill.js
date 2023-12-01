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
  let [isEdit, setIsEdit] = useState(false);
  let url = useContext(baseUrlAdmin);

  const {
    register,
    getValues,
    reset,
    setValue,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const View = () => {
    toast.success("test");
    console.log("first");
  };

  const Update = () => {
    console.log("update");
  };

  const Delete = () => {
    console.log("delete");
  };

  const openModalCreate = () => {
    setIsEdit(false);
    setModalStatus(!modalStatus);
  };

  const column = [
    {
      name: "NAME",
      selector: (row) => row.name,
    },
    {
      name: "SKILL",
      selector: (row) => row.skill,
    },
    {
      name: "ACTION",
      button: true,
      width: "300px",
      cell: () => (
        <>
          <div className="flex flex-wrap justify-center w-full space-x-1">
            <ButtonActionTable onclick={View} icon="fa fa-eye" />
            <ButtonActionTable onclick={Update} icon="fa fa-edit" />
            <ButtonActionTable onclick={Delete} icon="fa fa-trash" />
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
        closeFunction={() => setModalStatus(!modalStatus)}
      >
        <div className="flex flex-wrap w-full text-lg">
          <form className="flex w-full flex-col space-y-4 flex-wrap flex-grow ">
            <div className="flex flex-row flex-grow space-x-5">
              <label className="w-1/2">Name</label>
              <span className="w-full">
                <Input type="text" placeholder="Name..." />
              </span>
            </div>
            <div className="flex flex-row w-full space-x-5">
              <label className="w-1/2">Skill</label>
              <span className="w-full">
                <Input type="text" placeholder="Skill..." />
              </span>
            </div>
          </form>
        </div>
      </Modal>
    </>
  );
};

export default Skill;
