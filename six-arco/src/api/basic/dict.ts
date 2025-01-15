import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Dict {
    id: number;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    type: string;
    label: string;
    value: string;
    color: string;
    is_sync: number;
}

export const EmptyDict: Dict = {
    id: 0,
    type: '',
    label: '',
    value: '',
    color: '',
    is_sync: 0,
}

const root = '/dict'
export function reqDict(action: string, data: RequestParam<Dict> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}