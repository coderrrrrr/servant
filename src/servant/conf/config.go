package conf

type Config struct {
	Server     Server
	Users      map[string]*User
	Commands   map[string]*Commands
	Files      map[string]*Files
	Databases  map[string]*Database
	Auth       Auth
}


type Server struct {
	Listen string
}

type Auth struct {
	Enabled       bool
	MaxTimeDelta  uint32
}

type User struct {
	Hosts     []string
	Key       string
	Allows    map[string] []string
}

type Commands struct {
	Commands map[string]*Command
}

type Command struct {
	Lang         string
	Code         string
	Timeout      uint32
	User		 string
	Lock         Lock
}

type Database struct {
	Queries  map[string]*Query
	Driver   string
	Dsn      string
}

type Query struct {
	Sqls    []string
}

type Lock struct {
	Name     string
	Timeout  uint
	Wait     bool
}

type Files struct {
	Dirs   map[string]*Dir
}

type Dir struct {
	Root     string
	Allows    []string
	Patterns  []string
}
