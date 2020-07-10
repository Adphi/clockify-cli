# Clockiy API / CLI

> WORK IN PROGRESS. Heavy changes and incomplete features.

## CLI

`WIP`

## Use Clockify API

```golang
import github.com/pkuebler/clockify-cli/pkg/clockify
```

Start Client

```golang
endpoint := "https://api.clockify.me/api/v1"
reportEndpoint := "https://reports.api.clockify.me/v1"
apiKey := "XXX"
log := logrus.NewEntry(logrus.StandardLogger())

client, err = clockify.NewAPIClient(endpoint, reportEndpoint, apiKey, nil, log)
if err != nil {
    panic(err.Error())
}

ctx := context.Background()

client.StartRatelimit(ctx)

entries, err := ctx.Client.Workspace.List()
if err != nil {
    log.Fatal(err)
}
```
