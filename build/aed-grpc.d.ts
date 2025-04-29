/// <reference types="node" />
/// <reference types="node" />
import { type CallOptions, ChannelCredentials, Client, type ClientOptions, type ClientUnaryCall, type handleUnaryCall, Metadata, type ServiceError, type UntypedServiceImplementation } from "@grpc/grpc-js";
import _m0 from "protobufjs/minimal";
import { AED, AEDs, Period } from "./aed";
import { Empty } from "./google/protobuf/empty";
import { Network } from "./sologenic/com-fs-utils-lib/models/metadata/metadata";
export declare const protobufPackage = "aed";
export interface AEDFilter {
    Symbol: string;
    From: Date | undefined;
    /** Apply from as equal to retrieve only this one bucket for the given time */
    SingleBucket?: boolean | undefined;
    To: Date | undefined;
    Network: Network;
    Period: Period | undefined;
    /** Indicate if the data should be backfilled with the previous period if no data is found */
    Backfill: boolean;
    /** Indicates if the data is allowed to be retrieved from the cache (default: false - no cache) */
    AllowCache: boolean;
    OrganizationID: string;
}
export interface PeriodsFilter {
    Symbol: string;
    Periods: PeriodBucket[];
}
export interface PeriodBucket {
    Period: Period | undefined;
    Timestamp: Date | undefined;
}
export declare const AEDFilter: {
    encode(message: AEDFilter, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): AEDFilter;
    fromJSON(object: any): AEDFilter;
    toJSON(message: AEDFilter): unknown;
    create<I extends {
        Symbol?: string | undefined;
        From?: Date | undefined;
        SingleBucket?: boolean | undefined;
        To?: Date | undefined;
        Network?: Network | undefined;
        Period?: {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        Backfill?: boolean | undefined;
        AllowCache?: boolean | undefined;
        OrganizationID?: string | undefined;
    } & {
        Symbol?: string | undefined;
        From?: Date | undefined;
        SingleBucket?: boolean | undefined;
        To?: Date | undefined;
        Network?: Network | undefined;
        Period?: ({
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K in Exclude<keyof I["Period"], keyof Period>]: never; }) | undefined;
        Backfill?: boolean | undefined;
        AllowCache?: boolean | undefined;
        OrganizationID?: string | undefined;
    } & { [K_1 in Exclude<keyof I, keyof AEDFilter>]: never; }>(base?: I | undefined): AEDFilter;
    fromPartial<I_1 extends {
        Symbol?: string | undefined;
        From?: Date | undefined;
        SingleBucket?: boolean | undefined;
        To?: Date | undefined;
        Network?: Network | undefined;
        Period?: {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        Backfill?: boolean | undefined;
        AllowCache?: boolean | undefined;
        OrganizationID?: string | undefined;
    } & {
        Symbol?: string | undefined;
        From?: Date | undefined;
        SingleBucket?: boolean | undefined;
        To?: Date | undefined;
        Network?: Network | undefined;
        Period?: ({
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K_2 in Exclude<keyof I_1["Period"], keyof Period>]: never; }) | undefined;
        Backfill?: boolean | undefined;
        AllowCache?: boolean | undefined;
        OrganizationID?: string | undefined;
    } & { [K_3 in Exclude<keyof I_1, keyof AEDFilter>]: never; }>(object: I_1): AEDFilter;
};
export declare const PeriodsFilter: {
    encode(message: PeriodsFilter, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): PeriodsFilter;
    fromJSON(object: any): PeriodsFilter;
    toJSON(message: PeriodsFilter): unknown;
    create<I extends {
        Symbol?: string | undefined;
        Periods?: {
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[] | undefined;
    } & {
        Symbol?: string | undefined;
        Periods?: ({
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[] & ({
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        } & {
            Period?: ({
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } & {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } & { [K in Exclude<keyof I["Periods"][number]["Period"], keyof Period>]: never; }) | undefined;
            Timestamp?: Date | undefined;
        } & { [K_1 in Exclude<keyof I["Periods"][number], keyof PeriodBucket>]: never; })[] & { [K_2 in Exclude<keyof I["Periods"], keyof {
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[]>]: never; }) | undefined;
    } & { [K_3 in Exclude<keyof I, keyof PeriodsFilter>]: never; }>(base?: I | undefined): PeriodsFilter;
    fromPartial<I_1 extends {
        Symbol?: string | undefined;
        Periods?: {
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[] | undefined;
    } & {
        Symbol?: string | undefined;
        Periods?: ({
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[] & ({
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        } & {
            Period?: ({
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } & {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } & { [K_4 in Exclude<keyof I_1["Periods"][number]["Period"], keyof Period>]: never; }) | undefined;
            Timestamp?: Date | undefined;
        } & { [K_5 in Exclude<keyof I_1["Periods"][number], keyof PeriodBucket>]: never; })[] & { [K_6 in Exclude<keyof I_1["Periods"], keyof {
            Period?: {
                Type?: import("./aed").PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            Timestamp?: Date | undefined;
        }[]>]: never; }) | undefined;
    } & { [K_7 in Exclude<keyof I_1, keyof PeriodsFilter>]: never; }>(object: I_1): PeriodsFilter;
};
export declare const PeriodBucket: {
    encode(message: PeriodBucket, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): PeriodBucket;
    fromJSON(object: any): PeriodBucket;
    toJSON(message: PeriodBucket): unknown;
    create<I extends {
        Period?: {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        Timestamp?: Date | undefined;
    } & {
        Period?: ({
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K in Exclude<keyof I["Period"], keyof Period>]: never; }) | undefined;
        Timestamp?: Date | undefined;
    } & { [K_1 in Exclude<keyof I, keyof PeriodBucket>]: never; }>(base?: I | undefined): PeriodBucket;
    fromPartial<I_1 extends {
        Period?: {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        Timestamp?: Date | undefined;
    } & {
        Period?: ({
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: import("./aed").PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K_2 in Exclude<keyof I_1["Period"], keyof Period>]: never; }) | undefined;
        Timestamp?: Date | undefined;
    } & { [K_3 in Exclude<keyof I_1, keyof PeriodBucket>]: never; }>(object: I_1): PeriodBucket;
};
export type AEDServiceService = typeof AEDServiceService;
export declare const AEDServiceService: {
    /** Store a single AED */
    readonly upsert: {
        readonly path: "/aed.AEDService/Upsert";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: AED) => Buffer;
        readonly requestDeserialize: (value: Buffer) => AED;
        readonly responseSerialize: (value: Empty) => Buffer;
        readonly responseDeserialize: (value: Buffer) => Empty;
    };
    /** Store multiple AED */
    readonly batchUpsert: {
        readonly path: "/aed.AEDService/BatchUpsert";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: AEDs) => Buffer;
        readonly requestDeserialize: (value: Buffer) => AEDs;
        readonly responseSerialize: (value: Empty) => Buffer;
        readonly responseDeserialize: (value: Buffer) => Empty;
    };
    /** Get a list of AED by from/to, symbol, period and network */
    readonly get: {
        readonly path: "/aed.AEDService/Get";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: AEDFilter) => Buffer;
        readonly requestDeserialize: (value: Buffer) => AEDFilter;
        readonly responseSerialize: (value: AEDs) => Buffer;
        readonly responseDeserialize: (value: Buffer) => AEDs;
    };
    /** Get aeds for all the given periods */
    readonly getAeDsForPeriods: {
        readonly path: "/aed.AEDService/GetAEDsForPeriods";
        readonly requestStream: false;
        readonly responseStream: false;
        readonly requestSerialize: (value: PeriodsFilter) => Buffer;
        readonly requestDeserialize: (value: Buffer) => PeriodsFilter;
        readonly responseSerialize: (value: AEDs) => Buffer;
        readonly responseDeserialize: (value: Buffer) => AEDs;
    };
};
export interface AEDServiceServer extends UntypedServiceImplementation {
    /** Store a single AED */
    upsert: handleUnaryCall<AED, Empty>;
    /** Store multiple AED */
    batchUpsert: handleUnaryCall<AEDs, Empty>;
    /** Get a list of AED by from/to, symbol, period and network */
    get: handleUnaryCall<AEDFilter, AEDs>;
    /** Get aeds for all the given periods */
    getAeDsForPeriods: handleUnaryCall<PeriodsFilter, AEDs>;
}
export interface AEDServiceClient extends Client {
    /** Store a single AED */
    upsert(request: AED, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    upsert(request: AED, metadata: Metadata, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    upsert(request: AED, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    /** Store multiple AED */
    batchUpsert(request: AEDs, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    batchUpsert(request: AEDs, metadata: Metadata, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    batchUpsert(request: AEDs, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: Empty) => void): ClientUnaryCall;
    /** Get a list of AED by from/to, symbol, period and network */
    get(request: AEDFilter, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
    get(request: AEDFilter, metadata: Metadata, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
    get(request: AEDFilter, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
    /** Get aeds for all the given periods */
    getAeDsForPeriods(request: PeriodsFilter, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
    getAeDsForPeriods(request: PeriodsFilter, metadata: Metadata, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
    getAeDsForPeriods(request: PeriodsFilter, metadata: Metadata, options: Partial<CallOptions>, callback: (error: ServiceError | null, response: AEDs) => void): ClientUnaryCall;
}
export declare const AEDServiceClient: {
    new (address: string, credentials: ChannelCredentials, options?: Partial<ClientOptions>): AEDServiceClient;
    service: typeof AEDServiceService;
    serviceName: string;
};
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P : P & {
    [K in keyof P]: Exact<P[K], I[K]>;
} & {
    [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
};
export {};
