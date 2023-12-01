import { createContext } from "react";

const baseUrl = createContext("http://localhost:8070/api/v1");
const baseUrlAdmin = createContext("http://localhost:8070/admin/api/v1");
let isExpand = createContext(false);
const settingConfig = createContext();

export { baseUrl, baseUrlAdmin, isExpand, settingConfig };
