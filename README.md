# Public Action by Brian Morrison II for THAT TX '24

This Action was published in front of a live studio audience.

## Example

Below is a simple workflow that can be used to say hello to your favorite camper :D

```yaml
name: Saying hello!

on:
  [workflow_dispatch]

jobs:
  say_hello:
    runs-on: ubuntu-latest
    steps:
      - name: Why hello there...
        uses: brianmmdev-that-actions/public-action@v1
        with:
          name: Brian
```

## Input variables

### Required

`name` - The name of the camper.

## Outputs

This Action has no output variables.

