# `gotch_runner`

Very simple batch runner written by golang.

If batch failed, then `gotch_runner` reports it via email.

## Installation

Executable binaries are available at [releases](https://github.com/uda-cha/gotch_runner/releases).

For example, for linux x86_64,

```sh
$ wget https://github.com/uda-cha/gotch_runner/releases/download/v0.0.1-rc.1/gotch_runner_linux_amd64 -O gotch_runner
$ chmod a+x gotch_runner
```

## Usage

### Set environment variables to send email

```sh
MAIL_FROM=sender@example.com
MAIL_TO=recipient1@example.com,recipient2@example.com
MAIL_USERNAME=sender_smtp_username
MAIL_PASSWORD=sender_smtp_password
MAIL_HOST=smtp.example.com
MAIL_PORT=587
```

### run any commands you like with `gotch_runner`

```sh
gotch_runner /some/batch.sh --your_args=1
```

First argument of `gotch_runner` must be executables.

So, if you want to use `gotch_runner` with arbitrary environment variables, they must be located to before gotch_runner.

```sh
HOGE=1 ./gotch_runner ruby -e "puts ENV['HOGE']"
```
