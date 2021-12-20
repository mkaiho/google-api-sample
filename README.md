# google-api-sample

## Links

- [My Business Notifications API Reference](https://developers.google.com/my-business/reference/notifications/rest)
- [Google API Go Client](https://github.com/googleapis/google-api-go-client#google-apis-client-library-for-go)

## Build

```.sh
$ make build
```

## Run sample binary

```.sh
$ GBP_CLIENT_ID=<Google Client ID> \
  GBP_CLIENT_SECRET=<Google Client Secret> \
  GBP_REFRESH_TOKEN=<Refresh token> \
  GBP_REDIRECT_URL=<Redirect URL> \
  ./bin/sample
```

## Run token-reciever server binary

```.sh
$ GBP_CLIENT_ID=<Google Client ID> \
  GBP_CLIENT_SECRET=<Google Client Secret> \
  GBP_REFRESH_TOKEN=<Refresh token> \
  GBP_REDIRECT_URL=<Redirect URL> \
  ./bin/token-reciever
```

## Run pubsub-pub binary

```.sh
$ GCP_CLIENT_ID=<Google Client ID> \
  GCP_CLIENT_SECRET=<Google Client Secret> \
  GCP_REFRESH_TOKEN=<Refresh token> \
  GCP_REDIRECT_URL=<Redirect URL> \
  ./bin/pubsub-pub
```

## Run pubsub-pub binary

```.sh
$ GCP_CLIENT_ID=<Google Client ID> \
  GCP_CLIENT_SECRET=<Google Client Secret> \
  GCP_REFRESH_TOKEN=<Refresh token> \
  GCP_REDIRECT_URL=<Redirect URL> \
  ./bin/pubsub-sub
```
