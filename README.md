# boilr-webhook

This is a [boilr template](https://github.com/tmrts/boilr) for creating a webhook extension. Use a webhook extension to receive global webhooks events from the Drone server and perform an action. Get started by cloning the project and installing the template:

```console
$ cd boilr-webhook
$ boilr template save . drone-webhook -f
```

create a project in directory my-webhooks:

```console
$ boilr template use drone-webhook my-webhooks
```

enter the docker registry name for this project:

```text
[?] Please choose a value for "DockerRepository" [default: "foo/bar"]:
```

enter the go module name:

```text
[?] Please choose a value for "GoModule" [default: "github.com/foo/bar":
```
