import {RequestParam} from "@/api/base";
import axios from "axios";

export interface <Tpl-Module> {
    id: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    <Tpl-FieldsType>
}

export const Empty<Tpl-Module>: <Tpl-Module> = {
    id: 0,
    <Tpl-FieldsDefaultValue>
}

const root = 'admin/<Tpl-Route-Group>'
export function req<Tpl-Module>(action: string, data: RequestParam<<Tpl-Module>> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}