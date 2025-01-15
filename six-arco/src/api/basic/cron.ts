import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Cron {
    id: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    name: string;
    title: string;
    times: string;
    status: number;
}

export const EmptyCron: Cron = {
    id: 0,
    name: '',
    title: '',
    times: '',
    status: 0
}

const root = '/cron'
export function reqCron(action: string, data: RequestParam<Cron> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}