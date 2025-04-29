import _m0 from "protobufjs/minimal";
import { MetaData } from "./sologenic/com-fs-utils-lib/models/metadata/metadata";
export declare const protobufPackage = "aed";
export declare enum Series {
    SERIES_NOT_USED = 0,
    /** INTERNAL_TRADES - Source: dex trades, Usage: general trade graphs, supports TradeView-like graphing tools (ohlc) */
    INTERNAL_TRADES = 1,
    /** MARKET_DATA_STOCKS - Source: external market data provider for the general stock market (can be more than 1 provider, but only a single provider per stock) */
    MARKET_DATA_STOCKS = 2,
    UNRECOGNIZED = -1
}
export declare function seriesFromJSON(object: any): Series;
export declare function seriesToJSON(object: Series): string;
export declare enum Field {
    FIELD_NOT_USED = 0,
    OPEN = 1,
    HIGH = 2,
    LOW = 3,
    CLOSE = 4,
    VOLUME = 5,
    NUMBER_OF_TRADES = 6,
    INVERTED_VOLUME = 7,
    MARKET_CAP = 8,
    EPS = 9,
    PE_RATIO = 10,
    YIELD = 11,
    OPEN_TIME = 12,
    CLOSE_TIME = 13,
    UNRECOGNIZED = -1
}
export declare function fieldFromJSON(object: any): Field;
export declare function fieldToJSON(object: Field): string;
export declare enum PeriodType {
    PERIOD_TYPE_DO_NOT_USE = 0,
    PERIOD_TYPE_MINUTE = 1,
    PERIOD_TYPE_HOUR = 2,
    PERIOD_TYPE_DAY = 3,
    PERIOD_TYPE_WEEK = 4,
    UNRECOGNIZED = -1
}
export declare function periodTypeFromJSON(object: any): PeriodType;
export declare function periodTypeToJSON(object: PeriodType): string;
export interface AEDs {
    AEDs: AED[];
}
export interface AED {
    OrganizationID: string;
    Symbol: string;
    Timestamp: Date | undefined;
    Period: Period | undefined;
    MetaData: MetaData | undefined;
    Value: Value[];
    Series: Series;
}
export interface Value {
    Field: Field;
    /** String value */
    StringVal?: string | undefined;
    /** Integer value */
    Int64Val?: number | undefined;
    /** Float value */
    Float64Val?: number | undefined;
}
export interface Period {
    Type: PeriodType;
    /** The duration of the indicated period (e.g 1 minute, 3 minutes, etc) */
    Duration: number;
}
export declare const AEDs: {
    encode(message: AEDs, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): AEDs;
    fromJSON(object: any): AEDs;
    toJSON(message: AEDs): unknown;
    create<I extends {
        AEDs?: {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[] | undefined;
    } & {
        AEDs?: ({
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[] & ({
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        } & {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: ({
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } & {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } & { [K in Exclude<keyof I["AEDs"][number]["Period"], keyof Period>]: never; }) | undefined;
            MetaData?: ({
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } & {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } & { [K_1 in Exclude<keyof I["AEDs"][number]["MetaData"], keyof MetaData>]: never; }) | undefined;
            Value?: ({
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] & ({
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            } & {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            } & { [K_2 in Exclude<keyof I["AEDs"][number]["Value"][number], keyof Value>]: never; })[] & { [K_3 in Exclude<keyof I["AEDs"][number]["Value"], keyof {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[]>]: never; }) | undefined;
            Series?: Series | undefined;
        } & { [K_4 in Exclude<keyof I["AEDs"][number], keyof AED>]: never; })[] & { [K_5 in Exclude<keyof I["AEDs"], keyof {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[]>]: never; }) | undefined;
    } & { [K_6 in Exclude<keyof I, "AEDs">]: never; }>(base?: I | undefined): AEDs;
    fromPartial<I_1 extends {
        AEDs?: {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[] | undefined;
    } & {
        AEDs?: ({
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[] & ({
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        } & {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: ({
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } & {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } & { [K_7 in Exclude<keyof I_1["AEDs"][number]["Period"], keyof Period>]: never; }) | undefined;
            MetaData?: ({
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } & {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } & { [K_8 in Exclude<keyof I_1["AEDs"][number]["MetaData"], keyof MetaData>]: never; }) | undefined;
            Value?: ({
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] & ({
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            } & {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            } & { [K_9 in Exclude<keyof I_1["AEDs"][number]["Value"][number], keyof Value>]: never; })[] & { [K_10 in Exclude<keyof I_1["AEDs"][number]["Value"], keyof {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[]>]: never; }) | undefined;
            Series?: Series | undefined;
        } & { [K_11 in Exclude<keyof I_1["AEDs"][number], keyof AED>]: never; })[] & { [K_12 in Exclude<keyof I_1["AEDs"], keyof {
            OrganizationID?: string | undefined;
            Symbol?: string | undefined;
            Timestamp?: Date | undefined;
            Period?: {
                Type?: PeriodType | undefined;
                Duration?: number | undefined;
            } | undefined;
            MetaData?: {
                Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
                UpdatedAt?: Date | undefined;
                CreatedAt?: Date | undefined;
                UpdatedByAccount?: string | undefined;
            } | undefined;
            Value?: {
                Field?: Field | undefined;
                StringVal?: string | undefined;
                Int64Val?: number | undefined;
                Float64Val?: number | undefined;
            }[] | undefined;
            Series?: Series | undefined;
        }[]>]: never; }) | undefined;
    } & { [K_13 in Exclude<keyof I_1, "AEDs">]: never; }>(object: I_1): AEDs;
};
export declare const AED: {
    encode(message: AED, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): AED;
    fromJSON(object: any): AED;
    toJSON(message: AED): unknown;
    create<I extends {
        OrganizationID?: string | undefined;
        Symbol?: string | undefined;
        Timestamp?: Date | undefined;
        Period?: {
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        MetaData?: {
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } | undefined;
        Value?: {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[] | undefined;
        Series?: Series | undefined;
    } & {
        OrganizationID?: string | undefined;
        Symbol?: string | undefined;
        Timestamp?: Date | undefined;
        Period?: ({
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K in Exclude<keyof I["Period"], keyof Period>]: never; }) | undefined;
        MetaData?: ({
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } & {
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } & { [K_1 in Exclude<keyof I["MetaData"], keyof MetaData>]: never; }) | undefined;
        Value?: ({
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[] & ({
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        } & {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        } & { [K_2 in Exclude<keyof I["Value"][number], keyof Value>]: never; })[] & { [K_3 in Exclude<keyof I["Value"], keyof {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[]>]: never; }) | undefined;
        Series?: Series | undefined;
    } & { [K_4 in Exclude<keyof I, keyof AED>]: never; }>(base?: I | undefined): AED;
    fromPartial<I_1 extends {
        OrganizationID?: string | undefined;
        Symbol?: string | undefined;
        Timestamp?: Date | undefined;
        Period?: {
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } | undefined;
        MetaData?: {
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } | undefined;
        Value?: {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[] | undefined;
        Series?: Series | undefined;
    } & {
        OrganizationID?: string | undefined;
        Symbol?: string | undefined;
        Timestamp?: Date | undefined;
        Period?: ({
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } & {
            Type?: PeriodType | undefined;
            Duration?: number | undefined;
        } & { [K_5 in Exclude<keyof I_1["Period"], keyof Period>]: never; }) | undefined;
        MetaData?: ({
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } & {
            Network?: import("./sologenic/com-fs-utils-lib/models/metadata/metadata").Network | undefined;
            UpdatedAt?: Date | undefined;
            CreatedAt?: Date | undefined;
            UpdatedByAccount?: string | undefined;
        } & { [K_6 in Exclude<keyof I_1["MetaData"], keyof MetaData>]: never; }) | undefined;
        Value?: ({
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[] & ({
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        } & {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        } & { [K_7 in Exclude<keyof I_1["Value"][number], keyof Value>]: never; })[] & { [K_8 in Exclude<keyof I_1["Value"], keyof {
            Field?: Field | undefined;
            StringVal?: string | undefined;
            Int64Val?: number | undefined;
            Float64Val?: number | undefined;
        }[]>]: never; }) | undefined;
        Series?: Series | undefined;
    } & { [K_9 in Exclude<keyof I_1, keyof AED>]: never; }>(object: I_1): AED;
};
export declare const Value: {
    encode(message: Value, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Value;
    fromJSON(object: any): Value;
    toJSON(message: Value): unknown;
    create<I extends {
        Field?: Field | undefined;
        StringVal?: string | undefined;
        Int64Val?: number | undefined;
        Float64Val?: number | undefined;
    } & {
        Field?: Field | undefined;
        StringVal?: string | undefined;
        Int64Val?: number | undefined;
        Float64Val?: number | undefined;
    } & { [K in Exclude<keyof I, keyof Value>]: never; }>(base?: I | undefined): Value;
    fromPartial<I_1 extends {
        Field?: Field | undefined;
        StringVal?: string | undefined;
        Int64Val?: number | undefined;
        Float64Val?: number | undefined;
    } & {
        Field?: Field | undefined;
        StringVal?: string | undefined;
        Int64Val?: number | undefined;
        Float64Val?: number | undefined;
    } & { [K_1 in Exclude<keyof I_1, keyof Value>]: never; }>(object: I_1): Value;
};
export declare const Period: {
    encode(message: Period, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Period;
    fromJSON(object: any): Period;
    toJSON(message: Period): unknown;
    create<I extends {
        Type?: PeriodType | undefined;
        Duration?: number | undefined;
    } & {
        Type?: PeriodType | undefined;
        Duration?: number | undefined;
    } & { [K in Exclude<keyof I, keyof Period>]: never; }>(base?: I | undefined): Period;
    fromPartial<I_1 extends {
        Type?: PeriodType | undefined;
        Duration?: number | undefined;
    } & {
        Type?: PeriodType | undefined;
        Duration?: number | undefined;
    } & { [K_1 in Exclude<keyof I_1, keyof Period>]: never; }>(object: I_1): Period;
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
