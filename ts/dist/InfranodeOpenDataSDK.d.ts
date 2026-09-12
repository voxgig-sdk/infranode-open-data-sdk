import { CityEntity } from './entity/CityEntity';
import { CompareEntity } from './entity/CompareEntity';
import { HealthEntity } from './entity/HealthEntity';
import { LiveEntity } from './entity/LiveEntity';
import { MetaEntity } from './entity/MetaEntity';
import { StationEntity } from './entity/StationEntity';
export type * from './InfranodeOpenDataTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { InfranodeOpenDataEntityBase } from './InfranodeOpenDataEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class InfranodeOpenDataSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    City(entopts?: Record<string, any>): CityEntity;
    Compare(entopts?: Record<string, any>): CompareEntity;
    Health(entopts?: Record<string, any>): HealthEntity;
    Live(entopts?: Record<string, any>): LiveEntity;
    Meta(entopts?: Record<string, any>): MetaEntity;
    Station(entopts?: Record<string, any>): StationEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): InfranodeOpenDataSDK;
    tester(testopts?: any, sdkopts?: any): InfranodeOpenDataSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof InfranodeOpenDataSDK;
export { stdutil, config, BaseFeature, InfranodeOpenDataEntityBase, InfranodeOpenDataSDK, SDK, };
