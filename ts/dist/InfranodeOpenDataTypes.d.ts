export interface City {
    data: any;
    id?: string;
    meta: Record<string, any>;
}
export interface CityLoadMatch {
    id: string;
    $action?: string;
    [action: string]: any;
}
export interface CityListMatch {
    data?: any;
    id?: string;
    meta?: Record<string, any>;
}
export interface Compare {
    city: string;
    data?: Record<string, any>;
    source_status: string;
}
export interface CompareListMatch {
    city: string;
    limit?: number;
    offset?: number;
    order?: string;
    page?: number;
    resource: string;
    sort?: string;
}
export interface Health {
    redis: boolean;
    status: string;
    version: string;
}
export interface HealthLoadMatch {
    redis?: boolean;
    status?: string;
    version?: string;
}
export interface Live {
    data: any;
    meta: Record<string, any>;
}
export interface LiveLoadMatch {
    live_id: string;
    trip_id: string;
    $action?: string;
    [action: string]: any;
}
export interface Meta {
    breaker_state: string;
    enabled: boolean;
    source: string;
}
export interface MetaLoadMatch {
    breaker_state?: string;
    enabled?: boolean;
    source?: string;
}
export interface MetaListMatch {
    limit?: number;
    offset?: number;
    order?: string;
    page?: number;
    sort?: string;
}
export interface Station {
}
export interface StationLoadMatch {
    eva: string;
    $action?: string;
    [action: string]: any;
}
