package main

import (
  "fmt"
  "net/smtp"
  "os"
  "os/exec"
  "strconv"
  "strings"
  "syscall"
  "github.com/jordan-wright/email"
)

func print_usage() {
  fmt.Println("USAGE: ./gotch_runner somecommand somearg1 --somearg2")
}

type MailConfig struct {
  from string
  to []string
  subject string
  text []byte
  username string
  password string
  host string
  port string
}

func try_to_get_env(key string) (val string) {
  val, ret := os.LookupEnv(key)
  if ret == false {
    envs := []string{"MAIL_FROM", "MAIL_TO", "MAIL_USERNAME", "MAIL_PASSWORD", "MAIL_HOST", "MAIL_PORT"}
    fmt.Println("Environment Variable " + key + " is not set. You must set:")
    fmt.Printf(strings.Join(envs, "\n") + "\n")
    os.Exit(2)
  }

  return val
}

func mail_setting() (config MailConfig) {
  config.from     = try_to_get_env("MAIL_FROM")
  config.to       = strings.Split(try_to_get_env("MAIL_TO"), ",")
  config.username = try_to_get_env("MAIL_USERNAME")
  config.password = try_to_get_env("MAIL_PASSWORD")
  config.host     = try_to_get_env("MAIL_HOST")
  config.port     = try_to_get_env("MAIL_PORT")

  return config
}

func send_mail(cmd string, status int, msg string, setting MailConfig) {
  e := email.NewEmail()
  e.From    = setting.from
  e.To      = setting.to
  e.Subject = cmd + " was failed."
  e.Text    = []byte(msg + "\n" + "exit code: " + string(strconv.Itoa(status)))
  
  auth := smtp.PlainAuth("", setting.username, setting.password, setting.host)

  e.Send(setting.host +":" + setting.port, auth)
  fmt.Println("mail has been sent.")
}

func exec_command(cmd string, args []string) (stdoutStderr string, status int) {
  command := exec.Command(cmd, args...)
  stdoutStderrBytes, err := command.CombinedOutput()

  status = 0
  if err != nil {
    status = 1
    if e2, ok := err.(*exec.ExitError); ok {
      if s, ok := e2.Sys().(syscall.WaitStatus); ok {
        status = s.ExitStatus()
      }
    }
  }

  stdoutStderr = string(stdoutStderrBytes)
  return stdoutStderr, status
}

func main() {
  if len(os.Args[1:]) == 0 {
    print_usage()
    os.Exit(2)
  }

  cmd  := os.Args[1]
  args := os.Args[2:]
  setting := mail_setting()

  stdoutStderr, status := exec_command(cmd, args)
  
  if len(stdoutStderr) > 0 {
    fmt.Print(stdoutStderr)
  }

  if status > 0 {
    send_mail(cmd, status, stdoutStderr, setting)
  }

  os.Exit(status)
}