# Aed Documentation

## Table of Contents

- [Overview](#overview)
- [aed.proto](#aed)
  - [Messages](#messages)
    - [AEDs](#aeds)
    - [AED](#aed)
    - [Value](#value)
    - [Period](#period)
  - [Enums](#enums)
    - [Source](#source)
    - [Series](#series)
    - [Field](#field)
    - [PeriodType](#periodtype)
- [Version Information](#version-information)
- [Support](#support)

## Overview

The Aed provides a comprehensive data structure for managing aed within the system. This model supports organizational context: links items to organizations via organizationid, metadata and audit: includes metadata and audit trails for tracking changes, identification: provides unique identifiers for aed, and more. 

Key features of the aed model include:
- **Organizational Context**: Links items to organizations via OrganizationID
- **Metadata and Audit**: Includes metadata and audit trails for tracking changes
- **Identification**: Provides unique identifiers for aed

## aed.proto

### Package Information

- **Package Name**: `aed`
- **Go Package Path**: `github.com/sologenic/com-fs-aed-model;aed`

### Overview

The `aed.proto` file defines the core aed model for aed management. It provides message types for representing aed data and operations. The file integrates with external utility libraries: `metadata.proto`.

### Messages

#### AEDs {#aeds}

The `AEDs` message represents a collection of aed with pagination support for handling large result sets.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| AEDs | `AED` | Optional | AEDs field |

**Use Cases:**
- Returning paginated lists of aed from queries or searches
- Implementing pagination in aed listing APIs
- Handling large aeds efficiently

**Important Notes:**
- This message provides the aeds representation

#### AED {#aed}

The `AED` message provides aed data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| OrganizationID | `string` | Required | UUID of the organization this item belongs to |
| Symbol | `string` | Required | Denom1:Denom2 |
| Timestamp | `google.protobuf.Timestamp` | Required | Timestamp information |
| Period | `Period` | Required | Period field |
| MetaData | `metadata.MetaData` | Required | Metadata information including network and version details |
| UserID | `string` | Optional | Time series stored at user level for profit/loss, etc |
| Value | `Value` | Optional | Value field |
| Series | `Series` | Required | Series field |
| Source | `Source` | Optional | Source field |

**Use Cases:**
- Creating new aed records
- Retrieving aed information
- Updating aed data
- Associating items with specific organizations

**Important Notes:**
- The `OrganizationID` must be a valid UUID format
- The `UserID` field must match a valid identifier format

#### Value {#value}

The `Value` message provides value data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| Field | `Field` | Required | Field field |
| StringVal | `string` | Optional | String value |
| Int64Val | `int64` | Optional | Integer value |
| Float64Val | `double` | Optional | Float value |
| BoolVal | `bool` | Optional | Boolean value |

**Use Cases:**
- Creating new value records
- Retrieving value information
- Updating value data

**Important Notes:**
- This message provides the value representation

#### Period {#period}

The `Period` message provides period data and operations.

**Field Table:**

| Field Name | Type | Required/Optional | Description |
|------------|------|-------------------|-------------|
| Type | `PeriodType` | Required | Type classification for this item (see related enum) |
| Duration | `int32` | Required | The duration of the indicated period (e.g 1 minute, 3 minutes, etc) |

**Use Cases:**
- Creating new period records
- Retrieving period information
- Updating period data

**Important Notes:**
- This message provides the period representation

### Enums

#### Source {#source}

The `Source` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| SOURCE_NOT_USED | 0 | Default/unused value (protobuf convention) |
| SOURCE_EXCHANGE | 1 | Source Exchange state or type |
| SOURCE_ATS | 2 | Source Ats state or type |
| SOURCE_DEX | 3 | Source Dex state or type |

**Use Cases:**
- Setting source for items
- Filtering items by source in queries
- Enforcing business logic based on source

**Important Notes:**
- Values with `NOT_USED` prefix or number 0 follow protobuf conventions for default enum values and should not be actively used
- Only valid source values should be used in production code
- Source changes should be tracked in audit trails for compliance purposes

#### Series {#series}

The `Series` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| SERIES_NOT_USED | 0 | Default/unused value (protobuf convention) |
| INTERNAL_TRADES | 1 | Internal Trades state or type |
| MARKET_DATA_STOCKS | 2 | Market Data Stocks state or type |
| USER_PERFORMANCE | 3 | User Performance state or type |

**Use Cases:**
- Setting series for items
- Filtering items by series in queries
- Enforcing business logic based on series

**Important Notes:**
- Values with `NOT_USED` prefix or number 0 follow protobuf conventions for default enum values and should not be actively used
- Only valid series values should be used in production code
- Series changes should be tracked in audit trails for compliance purposes

#### Field {#field}

The `Field` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| FIELD_NOT_USED | 0 | Default/unused value (protobuf convention) |
| OPEN | 1 | Open state or type |
| HIGH | 2 | High state or type |
| LOW | 3 | Low state or type |
| CLOSE | 4 | Close state or type |
| VOLUME | 5 | Volume state or type |
| NUMBER_OF_TRADES | 6 | Number Of Trades state or type |
| INVERTED_VOLUME | 7 | Inverted Volume state or type |
| MARKET_CAP | 8 | Market Cap state or type |
| EPS | 9 | Eps state or type |
| PE_RATIO | 10 | Pe Ratio state or type |
| YIELD | 11 | Yield state or type |
| OPEN_TIME | 12 | Open Time state or type |
| CLOSE_TIME | 13 | Close Time state or type |
| FIRST_PRICE | 14 | First Price state or type |
| LAST_PRICE | 15 | Last Price state or type |

**Use Cases:**
- Setting field for items
- Filtering items by field in queries
- Enforcing business logic based on field

**Important Notes:**
- Values with `NOT_USED` prefix or number 0 follow protobuf conventions for default enum values and should not be actively used
- Only valid field values should be used in production code
- Field changes should be tracked in audit trails for compliance purposes

#### PeriodType {#periodtype}

The `PeriodType` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| PERIOD_TYPE_DO_NOT_USE | 0 | Default/unused value (protobuf convention) |
| PERIOD_TYPE_MINUTE | 1 | Period Type Minute state or type |
| PERIOD_TYPE_HOUR | 2 | Period Type Hour state or type |
| PERIOD_TYPE_DAY | 3 | Period Type Day state or type |
| PERIOD_TYPE_WEEK | 4 | Period Type Week state or type |
| PERIOD_TYPE_MONTH | 5 | Period Type Month state or type |
| PERIOD_TYPE_YEAR | 6 | Period Type Year state or type |

**Use Cases:**
- Setting periodtype for items
- Filtering items by periodtype in queries
- Enforcing business logic based on periodtype

**Important Notes:**
- Values with `NOT_USED` prefix or number 0 follow protobuf conventions for default enum values and should not be actively used
- Only valid periodtype values should be used in production code
- PeriodType changes should be tracked in audit trails for compliance purposes

## Version Information

This documentation corresponds to the Protocol Buffer definitions in `aed.proto`. The proto file(s) use `proto3` syntax. When referencing this documentation, ensure that the version of the proto files matches the version of the generated code and API implementations you are using.

## Support

For additional information and support:
- See `README.md` for project setup, installation, and usage instructions
- Refer to the Protocol Buffer definitions in `aed.proto` for the authoritative source of truth
- Check the imported utility libraries for details on related types:
  - `sologenic/com-fs-utils-lib/models/metadata/metadata.proto`
