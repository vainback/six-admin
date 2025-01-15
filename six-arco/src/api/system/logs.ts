import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Log {
    id: number;
    ip: string;
    uid: number;
    route: string;
    route_name: string;
    request_body: string;
    response_body: string;
    latency: string;
    agent: string;
    method: string;
    tenant_id?: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
}

export const EmptyLog: Log = {
    id: 0,
    ip: '',
    uid: 0,
    route: '',
    route_name: '',
    request_body: '',
    response_body: '',
    latency: '',
    agent: '',
    method: ''
}

const root = '/logs'
export function reqLog(action: string, data: RequestParam<Log> | null)  {
    return axios.post(`${root}/${action}`, data || {});
}