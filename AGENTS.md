# AGENTS.md

This file provides guidance to AI coding agents when working with code in this repository.

## Repository Overview

`certsuite-claim` provides the JSON Schema definition and GoLang language bindings for a [redhat-best-practices-for-k8s](https://github.com/redhat-best-practices-for-k8s/certsuite) claim file.

A claim is an attestation of tests performed by the certsuite, containing:
- The certsuite version used for testing
- Start and end times of the claim evaluation
- Description of the Hardware System(s) Under Test
- All test configurations used by tests
- All tests run with their respective results (pass/fail/skip states, durations, failure details)
- Node information from the OpenShift cluster

The claim schema is defined in `schemas/claim.schema.json` using [JSON Schema Draft-07](https://json-schema.org/draft-07/json-schema-release-notes.html). This JSON Schema is the authoritative definition - in case of ambiguity between the schema and any language binding, the JSON Schema rules.

## Build Commands

All commands should be run from the repository root:

```bash
make                          # Run all targets: generate schema, format, lint, build, test
make build                    # Build all Go packages
make generate-golang-schema   # Generate pkg/claim/schema.go from JSON Schema
make fmt                      # Format Go code with gofmt
make lint                     # Run golangci-lint (5 minute timeout)
make tests                    # Run unit tests with coverage
make coverage-html            # Generate and open HTML coverage report
make validate-example         # Validate the example claim file against the schema
make vet                      # Run go vet on all packages
```

## Test Commands

```bash
make tests                    # Run all tests with coverage output to cover.out
go test -v ./...              # Run tests with verbose output
make coverage-html            # View HTML coverage report in browser
```

## Code Organization

```
certsuite-claim/
├── cmd/
│   └── generate/
│       └── generate.go       # CLI tool to generate GoLang schema from JSON Schema
├── internal/
│   └── generate/
│       ├── config.go         # Override configuration parsing (override.json)
│       └── generate.go       # Schema generation with type override support
├── pkg/
│   └── claim/
│       ├── schema.go         # AUTO-GENERATED - GoLang structs from JSON Schema
│       ├── catalog.go        # TestCaseDescription and helper functions
│       ├── schema_test.go    # Unit tests for schema marshaling/unmarshaling
│       └── testdata/         # Test JSON files (valid and invalid claims)
├── schemas/
│   ├── claim.schema.json     # Authoritative JSON Schema definition
│   └── claim.example.json    # Example valid claim file
├── override.json             # Type override configuration for code generation
└── .golangci.yml             # Linter configuration
```

## Key Data Structures

The main GoLang types in `pkg/claim/schema.go` (auto-generated):

- **Root**: Top-level structure containing a single `Claim`
- **Claim**: Contains metadata, versions, configurations, nodes, and results
- **Metadata**: Start and end times of test execution
- **Versions**: certsuite version, claim format version, OCP/K8s versions
- **Result**: Individual test case result with state, duration, failure details
- **Identifier**: Unique test case identifier (id, suite, tags)
- **CatalogInfo**: Test description, remediation, exception process, best practice reference
- **CategoryClassification**: Mandatory/optional classification per scenario (Telco, FarEdge, Extended, NonTelco)

The `pkg/claim/catalog.go` provides:
- **TestCaseDescription**: Describes a test case with identifier, description, remediation, and classification
- **BuildTestCaseDescription()**: Helper function to construct test case metadata

## Key Dependencies

From `go.mod`:

| Dependency | Purpose |
|------------|---------|
| `github.com/sebrandon1/generate` | JSON Schema to GoLang code generator |
| `github.com/stretchr/testify` | Testing assertions |

## Development Guidelines

### Go Version
This repository uses Go 1.26.0.

### Modifying the Claim Schema

1. Edit `schemas/claim.schema.json` as needed
2. Run `make` to regenerate the GoLang bindings and run all checks
3. Ensure backwards compatibility when possible; breaking changes require a semantic major version bump
4. Update `schemas/claim.example.json` if adding new required fields

### Type Override System

The `override.json` file allows remapping JSON Schema "object" types that should remain as generic `map[string]interface{}` rather than being converted to specific GoLang structs. This is used for properties like "nodes" and "results" where the structure is defined elsewhere or needs flexibility.

Example override configuration:
```json
{
  "properties": [
    {
      "path": ["claim", "nodes"],
      "overrideType": "map[string]interface{}"
    }
  ]
}
```

### Generated Code

**WARNING**: `pkg/claim/schema.go` is auto-generated. Do not edit it manually. Instead:
1. Modify `schemas/claim.schema.json`
2. Update `override.json` if type remapping is needed
3. Run `make generate-golang-schema`

The generated code includes custom `MarshalJSON()` and `UnmarshalJSON()` methods that enforce required fields.

### Linting

The project uses golangci-lint with extensive configuration in `.golangci.yml`:
- Function length limit: 50 lines, 25 statements
- Cyclomatic complexity limit: 15
- Line length limit: 200 characters
- Magic number detection enabled
- Generated code is excluded from linting

Run linting:
```bash
make lint                     # Uses 5 minute timeout
golangci-lint run --timeout 5m0s
```

### Validating Claim Files

To validate a claim file against the schema:

```bash
# Install jsonschema validator
pip3 install jsonschema

# Validate a claim file
jsonschema -i <path_to_claim_file> ./schemas/claim.schema.json

# Validate the example (should produce no output and exit 0)
make validate-example
```

### Test Data

Test files in `pkg/claim/testdata/` include:
- `claim-valid.json`: A complete valid claim file
- `claim-invalid-additional-property.json`: Tests additional property rejection
- `claim-invalid-bool-results.json`: Tests type validation
- `invalid-json.json`: Tests JSON parsing errors
- `missing-claim.json`: Tests required field validation

### Adding New Fields

1. Add the field to `schemas/claim.schema.json` with:
   - Type definition
   - Description
   - Add to `required` array if mandatory
2. Run `make generate-golang-schema` to regenerate Go code
3. Update `schemas/claim.example.json` with the new field
4. Add test cases covering the new field
5. Run `make` to verify all checks pass

## Common Workflows

### Regenerating Schema After JSON Changes
```bash
make generate-golang-schema   # Regenerate schema.go
make fmt                      # Format the generated code
make lint                     # Verify linting passes
make tests                    # Run tests
```

### Full Validation Cycle
```bash
make                          # Runs all: generate, fmt, lint, build, tests
make validate-example         # Additionally validate the example file
```

### Viewing Test Coverage
```bash
make coverage-html            # Opens coverage report in browser
```
