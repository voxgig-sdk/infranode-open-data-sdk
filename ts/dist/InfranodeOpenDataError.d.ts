import { Context } from './Context';
declare class InfranodeOpenDataError extends Error {
    isInfranodeOpenDataError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { InfranodeOpenDataError };
