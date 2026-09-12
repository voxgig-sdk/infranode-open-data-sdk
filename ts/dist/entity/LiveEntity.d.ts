import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { Live, LiveLoadMatch } from '../InfranodeOpenDataTypes';
declare class LiveEntity extends InfranodeOpenDataEntityBase<Live> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: LiveEntity): LiveEntity;
    load(this: any, reqmatch?: LiveLoadMatch, ctrl?: Control): Promise<LiveEntity>;
}
export { LiveEntity };
