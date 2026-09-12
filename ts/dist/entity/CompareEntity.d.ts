import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { Compare, CompareListMatch } from '../InfranodeOpenDataTypes';
declare class CompareEntity extends InfranodeOpenDataEntityBase<Compare> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: CompareEntity): CompareEntity;
    list(this: any, reqmatch?: CompareListMatch, ctrl?: Control): Promise<CompareEntity[]>;
}
export { CompareEntity };
