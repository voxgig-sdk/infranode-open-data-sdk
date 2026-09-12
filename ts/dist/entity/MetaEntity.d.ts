import { InfranodeOpenDataEntityBase } from '../InfranodeOpenDataEntityBase';
import type { InfranodeOpenDataSDK } from '../InfranodeOpenDataSDK';
import type { Control } from '../types';
import type { Meta, MetaLoadMatch, MetaListMatch } from '../InfranodeOpenDataTypes';
declare class MetaEntity extends InfranodeOpenDataEntityBase<Meta> {
    constructor(client: InfranodeOpenDataSDK, entopts: any);
    make(this: MetaEntity): MetaEntity;
    load(this: any, reqmatch?: MetaLoadMatch, ctrl?: Control): Promise<MetaEntity>;
    list(this: any, reqmatch?: MetaListMatch, ctrl?: Control): Promise<MetaEntity[]>;
}
export { MetaEntity };
