import { createContext } from "react";

const baseUrl = createContext("http://localhost:8070/api/v1");
const baseUrlAdmin = createContext("http://localhost:8070/admin/api/v1");

export { baseUrl, baseUrlAdmin };
