# test-network-function-claim

`test-network-function-claim` provides a JSON Schema definition for a
[test-network-function](https://github.com/redhat-nfvpe/test-network-function) claim based on
[JSON Schema Draft-07](https://json-schema.org/draft-07/json-schema-release-notes.html).

A claim contains:
* A start time for the claim evaluation.
* A description of the Hardware System(s) Under Test
* All test configurations used by tests.
* All tests run, and their respective results.
* An end time for the claim evaluation.

## Validating a claim

### Requirements

[`jsonschema`](https://python-jsonschema.readthedocs.io/en/stable/) must be installed and made available in `$PATH`.

In order to install `jsonschema`, issue the following command:

```shell script
> pip3 install jsonschema
```

### Running the validation utility

In order to validate a given schema, issue the following command:

```shell script
> jsonschema -i <path_to_claim_file> ./claim-schema.json
```

For a successful validation of the input claim, one should expect no output and a `0` exit code.

```shell script
> jsonschema -i ./sample.json ./claim-schema.json
> echo $?
0
> 
```

If the claim does not adhere to the schema, `jsconschema` outputs an appropriate error message describing the issue.
