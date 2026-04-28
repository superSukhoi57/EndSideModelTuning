export interface Machine {
    id: number;
    ip: string;
    userid: number;
    isfinsh: number;
    resultid: number | null;
    core: number | null;
    ram: number | null;
    memory: number | null;
    os: string;
    desc: string;
    createAt: string;
    updateAt: string;
}

export interface MachineCreateReq {
    ip: string;
    pwd?: string;
    core?: number;
    ram?: number;
    memory?: number;
    os?: string;
    desc?: string;
}

export interface MachineUpdateReq {
    id: number;
    ip?: string;
    pwd?: string;
    isfinsh?: number;
    resultid?: number;
    core?: number;
    ram?: number;
    memory?: number;
    os?: string;
    desc?: string;
}

export interface Parameter {
    id: number;
    userid: number;
    parameters: string;
    script: string;
    desc: string;
    createAt: string;
    updateAt: string;
}

export interface ParameterCreateReq {
    id: number;
    parameters?: string;
    script: string;
    desc?: string;
}

export interface ParameterUpdateReq {
    id: number;
    parameters?: string;
    script?: string;
    desc?: string;
}

export interface Task {
    id: number;
    paramterid: number;
    userid: number;
    desc: string;
    createAt: string;
    updateAt: string;
}

export interface TaskCreateReq {
    id: number;
    paramterid: number;
    desc?: string;
}

export interface TaskUpdateReq {
    id: number;
    desc?: string;
}

export interface Result {
    id: number;
    result: string;
    userid: number;
    machineid: number;
    desc: string;
    createAt: string;
    updateAt: string;
}

export interface ResultCreateReq {
    result: string;
    machineid: number;
    desc?: string;
}

export interface ResultUpdateReq {
    id: number;
    result?: string;
    desc?: string;
}

export interface PageResponse<T> {
    total: number;
    list: T[];
}

export interface PageParams {
    page: number;
    pageSize: number;
}
