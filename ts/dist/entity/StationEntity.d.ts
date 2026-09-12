import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { Station, StationLoadMatch } from '../InfranodeOpenDataTypes';
declare class StationEntity extends InfranodeOpenDataEntityBase<Station> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: StationEntity): StationEntity;
    load(this: any, reqmatch?: StationLoadMatch, ctrl?: Control): Promise<StationEntity>;
}
export { StationEntity };
