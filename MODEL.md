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

#### Field {#field}

The `Field` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| OSE_TIME | 13 | Ose Time state or type |
| FIRST_PRICE | 14 | First Price state or type |
| LAST_PRICE | 15 | Last Price state or type |

**Use Cases:**
- Setting field for items
- Filtering items by field in queries
- Enforcing business logic based on field

**Important Notes:**
- Only valid field values should be used in production code
- Field changes should be tracked in audit trails for compliance purposes

#### PeriodType {#periodtype}

The `PeriodType` enum defines the possible states or types for aed, allowing for classification and state management.

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| PERIOD_TYPE_MONTH | 5 | Period Type Month state or type |
| PERIOD_TYPE_YEAR | 6 | Period Type Year state or type |

**Use Cases:**
- Setting periodtype for items
- Filtering items by periodtype in queries
- Enforcing business logic based on periodtype

**Important Notes:**
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
