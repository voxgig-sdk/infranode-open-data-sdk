import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { Health, HealthLoadMatch } from '../InfranodeOpenDataTypes';
declare class HealthEntity extends InfranodeOpenDataEntityBase<Health> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: HealthEntity): HealthEntity;
    load(this: any, reqmatch?: HealthLoadMatch, ctrl?: Control): Promise<HealthEntity>;
}
export { HealthEntity };
