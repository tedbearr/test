import axios from "axios";
import { useContext, useEffect, useState } from "react";
import { baseUrlAdmin } from "../../helper/Context";
import DataTable from "react-data-table-component";
import Button from "../../component/Button";

const Skill = () => {
  let [tableData, setTableData] = useState();
  let url = useContext(baseUrlAdmin);
  const column = [
    {
      name: "test",
      selector: (row) => row.name,
    },
  ];

  const getData = async () => {
    await axios.get(url + "/skill/").then((data) => {
      setTableData(data.data);
    });
  };

  useEffect(() => {
    getData();
  }, []);
  return (
    <>
      <span>Skill</span>
      <span className="w-full flex justify-end">
        <Button name="Create" />
      </span>
      <div>
        <DataTable columns={column} data={tableData} />
      </div>
    </>
  );
};

export default Skill;
