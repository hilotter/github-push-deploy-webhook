# deployhook

deployhook is a simple deploy tool using GitHub push webhook.

## env

```
export WEBHOOK_PORT=3000
export WEBHOOK_SECRET=your_webhook_secret
export WEBHOOK_DEPLOYMENT_BRANCH=deployment/production
export WEBHOOK_DEPLOYMENT_SCRIPT_PATH=/path/to/deploy.sh
```

## write deploy script

```
#!/bin/bash

echo 'deploy'
```

## run

```
./deployhook
2020/03/03 23:09:42 listen on :3000
2020/03/03 23:10:28 deployment webhook start: deployment/production
2020/03/03 23:10:28 deploy

2020/03/03 23:10:28 deployment webhook finished
```

## public url setting

any one of them

- DNS setting
- tunnel service
  - [ngrok](https://ngrok.com/)
  - [inlets](https://github.com/inlets/inlets)

## Github Webhook Setting

- Payload URL: https://{your-url}/webhooks
- Content type: application/json
- Secret: value of WEBHOOK_SECRET
- trigger: Just the push event

![github](https://user-images.githubusercontent.com/1042519/77043545-13ec8280-6a01-11ea-91a7-32239fe007e6.jpg)
