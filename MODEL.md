# Aed Documentation

## Table of Contents

- [aed.proto](#aed)

## Overview

The Aed provides data structures and definitions for managing aed within the system.

## aed.proto {#aed}

### Package Information

- **Package Name**: `aed`
- **Go Package Path**: `github.com/sologenic/com-fs-aed-model;aed`

### Overview

The `aed.proto` file defines the Aed model.

### Messages

#### AEDs

**Field Table:**

| Field Name | Type | Number | Description |
|------------|------|--------|-------------|
| UserID | `string` | 6 |  |
| Value | `Value` | 100 |  |
| Series | `Series` | 101 |  |
| Source | `Source` | 102 |  |

#### Value

No fields defined.

### Enums

#### Source

No values defined.

#### Field

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| OSE_TIME | 13 |  |
| FIRST_PRICE | 14 |  |
| LAST_PRICE | 15 |  |

#### PeriodType

**Value Table:**

| Value Name | Number | Description |
|------------|--------|-------------|
| PERIOD_TYPE_MONTH | 5 |  |
| PERIOD_TYPE_YEAR | 6 |  |

## Version Information

This documentation corresponds to the current version of the proto files in this repository.

## Support

For more information, see:
- README.md in this repository
- Protocol Buffer documentation
