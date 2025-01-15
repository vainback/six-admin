import {RequestParam} from "@/api/base";
import axios from "axios";
import { S } from "mockjs";

export interface Files {
    id: number;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    mime: String;
    type: string;
    url: string;
}

export const EmptyFile: Files = {
    id: 0,
    mime: '',
    type: '',
    url: ''
}

const root = '/files'
export function reqFiles(action: string, data: RequestParam<Files> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}