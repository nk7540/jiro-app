package config

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

type ServerConfig struct {
	*fiber.App
	Name        string `mapstructure:"APP_NAME" yaml:"name" env:"APP_NAME" env-default:"iSend.to"`
	Url         string `mapstructure:"APP_URL" yaml:"url" env:"APP_URL" env-default:"http://localhost"`
	Host        string `mapstructure:"APP_HOST" yaml:"host" env:"APP_HOST" env-default:"localhost"`
	Port        string `mapstructure:"APP_PORT" yaml:"port" env:"APP_PORT" env-default:"8080"`
	Path        string `mapstructure:"APP_PATH" yaml:"path" env:"APP_PATH"`
	ProxyHeader string `mapstructure:"PROXY_HEADER" yaml:"PROXY_HEADER" env:"PROXY_HEADER" env-default:"*"`
	LogPath     string `mapstructure:"LOG_PATH" yaml:"log_path" env:"LOG_PATH" env-default:"storage/logs"`
	Debug       bool   `mapstructure:"APP_DEBUG" yaml:"debug" env:"APP_DEBUG" env-default:"true"`
	UploadSize  int    `mapstructure:"UPLOAD_SIZE" yaml:"upload_size" env:"UPLOAD_SIZE" env-default:"400"`
}

func (s *ServerConfig) LoadPath() {
	if s.Url == "" {
		s.Url = fmt.Sprintf("http://localhost:%s", s.Port)
	}
	path, _ := os.Getwd()
	s.Path = path
	s.LogPath = makeDir(filepath.Join(path, s.LogPath))
	s.UploadSize = s.UploadSize * 1024 * 1024
}

func (s *ServerConfig) Setup() {
	s.App = fiber.New(fiber.Config{
		Concurrency:           256 * 1024 * 1024,
		ServerHeader:          s.Name,
		BodyLimit:             s.UploadSize,
		ReduceMemoryUsage:     true,
		ErrorHandler:          ErrorHandling,
		DisableStartupMessage: true,
		ProxyHeader:           s.ProxyHeader,
	})
}

func (s *ServerConfig) ServeWithGracefulShutdown() error {
	a := s.Host + ":" + s.Port
	s.startupMessage(a, false, "")
	go func() {
		if err := s.Listen(a); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGABRT,
		syscall.SIGQUIT,
	)
	<-c
	return s.Shutdown()
}

func (s *ServerConfig) startupMessage(addr string, tls bool, processIds string) {
	// ignore child processes
	if fiber.IsChild() {
		return
	}

	var logo string
	logo += "%s"
	logo += " ┌─────────────────────────────────────────────────────┐\n"
	logo += " │ %s   │\n"
	logo += " │ %s   │\n"
	logo += " │                                                     │\n"
	logo += " │ Handlers %s  Processes %s   │\n"
	logo += " │ Prefork .%s  PID ....%s   │\n"
	logo += " └─────────────────────────────────────────────────────┘"
	logo += "%s"

	const (
		cBlack = "\u001b[90m"
		// cRed   = "\u001b[91m"
		cCyan = "\u001b[96m"
		// cGreen = "\u001b[92m"
		// cYellow  = "\u001b[93m"
		// cBlue    = "\u001b[94m"
		// cMagenta = "\u001b[95m"
		// cWhite   = "\u001b[97m"
		cReset = "\u001b[0m"
	)

	value := func(s string, width int) string {
		pad := width - len(s)
		str := ""
		for i := 0; i < pad; i++ {
			str += "."
		}
		if s == "Disabled" {
			str += " " + s
		} else {
			str += fmt.Sprintf(" %s%s%s", cCyan, s, cBlack)
		}
		return str
	}

	center := func(s string, width int) string {
		pad := strconv.Itoa((width - len(s)) / 2)
		str := fmt.Sprintf("%"+pad+"s", " ")
		str += s
		str += fmt.Sprintf("%"+pad+"s", " ")
		if len(str) < width {
			str += " "
		}
		return str
	}

	centerValue := func(s string, width int) string {
		pad := strconv.Itoa((width - len(s)) / 2)
		str := fmt.Sprintf("%"+pad+"s", " ")
		str += fmt.Sprintf("%s%s%s", cCyan, s, cBlack)
		str += fmt.Sprintf("%"+pad+"s", " ")
		if len(str)-10 < width {
			str += " "
		}
		return str
	}

	pad := func(s string, width int) (str string) {
		toAdd := width - len(s)
		str += s
		for i := 0; i < toAdd; i++ {
			str += " "
		}
		return
	}

	host, port := parseAddr(addr)
	if host == "" || host == "0.0.0.0" {
		host = "127.0.0.1"
	}
	addr = "http://" + host + ":" + port
	if tls {
		addr = "https://" + host + ":" + port
	}

	isPrefork := "Disabled"
	if s.Config().Prefork {
		isPrefork = "Enabled"
	}

	procs := strconv.Itoa(runtime.GOMAXPROCS(0))
	if !s.Config().Prefork {
		procs = "1"
	}
	routeCount := 0
	for _, route := range s.Stack() {
		routeCount += len(route)
	}
	mainLogo := fmt.Sprintf(logo,
		cBlack,
		centerValue(s.Name, 49),
		center(addr, 49),
		value(strconv.Itoa(routeCount), 14), value(procs, 12),
		value(isPrefork, 14), value(strconv.Itoa(os.Getpid()), 14),
		cReset,
	)

	var childPidsLogo string
	if s.Config().Prefork {
		var childPidsTemplate string
		childPidsTemplate += "%s"
		childPidsTemplate += " ┌───────────────────────────────────────────────────┐\n%s"
		childPidsTemplate += " └───────────────────────────────────────────────────┘"
		childPidsTemplate += "%s"

		newLine := " │ %s%s%s │"

		// Turn the `processIds` variable (in the form ",a,b,c,d,e,f,etc") into a slice of PIDs
		var pidSlice []string
		for _, v := range strings.Split(processIds, ",") {
			if v != "" {
				pidSlice = append(pidSlice, v)
			}
		}

		var lines []string
		thisLine := "Child PIDs ... "
		var itemsOnThisLine []string

		addLine := func() {
			lines = append(lines,
				fmt.Sprintf(
					newLine,
					cBlack,
					thisLine+cCyan+pad(strings.Join(itemsOnThisLine, ", "), 49-len(thisLine)),
					cBlack,
				),
			)
		}

		for _, pid := range pidSlice {
			if len(thisLine+strings.Join(append(itemsOnThisLine, pid), ", ")) > 49 {
				addLine()
				thisLine = ""
				itemsOnThisLine = []string{pid}
			} else {
				itemsOnThisLine = append(itemsOnThisLine, pid)
			}
		}

		// Add left over items to their own line
		if len(itemsOnThisLine) != 0 {
			addLine()
		}

		// Form logo
		childPidsLogo = fmt.Sprintf(childPidsTemplate,
			cBlack,
			strings.Join(lines, "\n")+"\n",
			cReset,
		)
	}

	// Combine both the child PID logo and the main Fiber logo

	// Pad the shorter logo to the length of the longer one
	splitMainLogo := strings.Split(mainLogo, "\n")
	splitChildPidsLogo := strings.Split(childPidsLogo, "\n")

	mainLen := len(splitMainLogo)
	childLen := len(splitChildPidsLogo)

	if mainLen > childLen {
		diff := mainLen - childLen
		for i := 0; i < diff; i++ {
			splitChildPidsLogo = append(splitChildPidsLogo, "")
		}
	} else {
		diff := childLen - mainLen
		for i := 0; i < diff; i++ {
			splitMainLogo = append(splitMainLogo, "")
		}
	}

	// Combine the two logos, line by line
	output := "\n"
	for i := range splitMainLogo {
		output += cBlack + splitMainLogo[i] + " " + splitChildPidsLogo[i] + "\n"
	}

	out := colorable.NewColorableStdout()
	if os.Getenv("TERM") == "dumb" || (!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd())) {
		out = colorable.NewNonColorable(os.Stdout)
	}

	_, _ = fmt.Fprintln(out, output)
}

func makeDir(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModePerm)
	}
	return path
}

func parseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
