import {RequestParam} from "@/api/base";
import axios from "axios";

export interface Tests {
    id: number;
    create_time?: string;
    update_time?: string;
    delete_time?: string;
    te_str: string;
    te_int: number;
    te_bigint: number;
    te_float: number;
    te_decimal: number;
    te_select: string;
    te_select_many: string[];
    te_radio: string;
    te_checkbox: string[];
    te_switch: boolean;
    te_timepicker: string | null;
    te_datepicker: string | null;
    te_datetimepicker: string | null;
    te_image_one: string[];
    te_image_many: string[];
    te_video: string[];
    te_file: string[];
    te_text: string;
    te_editor: string;
}

export const EmptyTests: Tests = {
    id: 0,
    te_str: '',
    te_int: 0,
    te_bigint: 0,
    te_float: 0,
    te_decimal: 0,
    te_select: '',
    te_select_many: [],
    te_radio: '',
    te_checkbox: [],
    te_switch: false,
    te_timepicker: null,
    te_datepicker: null,
    te_datetimepicker: null,
    te_image_one: [],
    te_image_many: [],
    te_video: [],
    te_file: [],
    te_text: '',
    te_editor: '',
}

const root = 'admin/tests'
export function reqTests(action: string, data: RequestParam<Tests> | null)  {
    return axios.post(`${root}/${action}`, data ||{});
}