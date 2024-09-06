# metal-sweeper

[![Experimental](https://img.shields.io/badge/Stability-Experimental-red.svg)](https://github.com/equinix-labs/equinix-labs/blob/main/experimental-statement.md#experimental-statement)

Experimental GitHub Action for managing [Equinix Metal](https://metal.equinix.com) Projects.

> :bulb: See also:
>
> - [equinix-metal-project](https://github.com/equinix-labs/metal-project-action) action
> - [equinix-metal-examples](https://github.com/equinix-labs/metal-actions-example) examples

Given a Equinix Metal User API Token, a server name (specified as runnerName), and a Project ID, the server in the project is deleted.

Create a project with the [Equinix Metal Project Action](https://github.com/equinix-labs/metal-project-action).

See the [Equinix Metal Actions Example](https://github.com/equinix-labs/metal-actions-example) for usage examples.

## Input

| With          | Description                                                                                             |
| ------------- | ------------------------------------------------------------------------------------------------------- |
| `authToken`   | (required) A Equinix Metal User API Token                                                               |
| `projectID`   | (required) Project ID that will be deleted.                                                             |
| `runnerName`  | the name of the server will be deleted.                                                                 |

## Output

There are no outputs.

## Support

This repository is [Experimental](https://github.com/equinix-labs/equinix-labs/blob/main/experimental-statement.md) meaning that it's based on untested ideas or techniques and not yet established or finalized or involves a radically new and innovative style! This means that support is best effort (at best!) and we strongly encourage you to NOT use this in production.
