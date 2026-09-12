import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { City, CityLoadMatch, CityListMatch } from '../InfranodeOpenDataTypes';
declare class CityEntity extends InfranodeOpenDataEntityBase<City> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: CityEntity): CityEntity;
    load(this: any, reqmatch?: CityLoadMatch, ctrl?: Control): Promise<CityEntity>;
    list(this: any, reqmatch?: CityListMatch, ctrl?: Control): Promise<CityEntity[]>;
}
export { CityEntity };
