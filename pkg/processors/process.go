package processors

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bloomingbug/depublic/configs"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/gomail.v2"
)

type process struct {
	rdb *redis.Pool
	cfg configs.Config
}

func (p *process) RunProcess() {
	fmt.Println("Background process is running...")
	pool := work.NewWorkerPool(Context{}, 10, p.cfg.Namespace.Namespace, p.rdb)
	pool.Middleware((*Context).Log)

	pool.Job("send_otp", p.sendOTPToEmail)
	pool.Job("send_token", p.sendTokenToEmail)
	pool.Start()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	pool.Stop()
}

func (p *process) sendOTPToEmail(job *work.Job) error {
	var to, otp string
	to = job.ArgString("email_address")
	if err := job.ArgError(); err != nil {
		return err
	}

	otp = job.ArgString("otp_code")
	if err := job.ArgError(); err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", p.cfg.SMTP.Sender)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "OTP Registration")
	mail.SetBody("text/plain", fmt.Sprintf("Your OTP Code is: %s", otp))

	dialer := gomail.NewDialer(p.cfg.SMTP.Host, p.cfg.SMTP.Port, p.cfg.SMTP.Username, p.cfg.SMTP.Password)

	err := dialer.DialAndSend(mail)
	if err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully")
	return nil
}

func (p *process) sendTokenToEmail(job *work.Job) error {
	var to, token, link string
	to = job.ArgString("email_address")
	if err := job.ArgError(); err != nil {
		return err
	}

	token = job.ArgString("token")
	if err := job.ArgError(); err != nil {
		return err
	}

	link = job.ArgString("link")
	if err := job.ArgError(); err != nil {
		return err
	}

	mail := gomail.NewMessage()
	mail.SetHeader("From", p.cfg.SMTP.Sender)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", "Reset Password")
	mail.SetBody("text/html", fmt.Sprintf(`
	<html>
		<body>
			<p>We have received a request to reset the password on your Account.
				</br>
				Please click this link <a href="%s?token=%s" target="_blank">here</a> to change your password.</p>
			<p>Token: %s</p>
		</body>
	</html>
	`, link, token, token))

	dialer := gomail.NewDialer(p.cfg.SMTP.Host, p.cfg.SMTP.Port, p.cfg.SMTP.Username, p.cfg.SMTP.Password)

	err := dialer.DialAndSend(mail)
	if err != nil {
		panic(err)
	}

	fmt.Println("Email sent successfully")
	return nil
}

type Process interface {
	RunProcess()
}

func NewProcess(rdb *redis.Pool, cfg configs.Config) Process {
	return &process{
		rdb: rdb,
		cfg: cfg,
	}
}

type Context struct {
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting Job: ", job.Name)
	return next()
}
