package config

type AcceptFilter struct {
	Protocol     string
	AcceptFilter string
}

type AccessFileName struct {
	Filename string
}

type AddCharset struct {
	Charset string
	Exts    []string
}

type AllowOverride struct {
	DirectiveType     string
	Options           []string
	Nonfatal          string
	AllowOverrideList []string
}

type CGIMapExtension struct {
	CGIPath   string
	Extension string
}

type AcceptPathInfo struct {
	Enabled bool
	Default bool
}

type AddDefaultCharset struct {
	Enabled bool
	Charset string
}

type AllowEncodedSlashes struct {
	Enabled  bool
	NoDecode bool
}

type Define struct {
	Parameter string
	Value     string
}

type ErrorDocument struct {
	ErrorCode int
	Document  string
}

type ErrorLogFormat struct {
	Connection bool
	Request    bool
	Format     string
}

type FileETag struct {
	INode bool
	MTime bool
	Size  bool
}

type HostnameLookups struct {
	On     bool
	Double bool
}

type ErrorLog struct {
	Path string
}

type Listen struct {
	Address  string
	Protocol string
}

type MaxRangeOverlaps struct {
	Default        bool
	Unlimited      bool
	None           bool
	NumberOfRanges int
}

type Mutex struct {
	Mechanism string
	Default   bool
	Name      []string
	OmitPID   bool
}

type Option struct {
	Name   string
	Remove bool
	Add    bool
}

type Limiter struct {
	SoftLimit int
	SoftMax   bool
	HardLimit int
	HardMax   bool
}

type LogLevel struct {
	Module string
	Level  string
}

type RegexDefaultOptions struct {
	None    bool
	Options []Option
}

type ScriptInterpreterSource struct {
	Registry bool
	Script   bool
}

type ServerSignature struct {
	Active bool
	EMail  bool
}

type ServerTokens struct {
	Major       bool
	Minor       bool
	Minimal     bool
	ProductOnly bool
	OS          bool
	Full        bool
}

type TraceEnable struct {
	Active   bool
	Extended bool
}

type UseCanonicalName struct {
	Active bool
	DNS    bool
}

type CoreConfig struct {
	ModRewrite
	AcceptPathInfo           AcceptPathInfo
	AddDefaultCharset        AddDefaultCharset
	AllowEncodedSlashes      AllowEncodedSlashes
	AcceptFilter             []AcceptFilter
	AccessFileName           []string
	AddCharset               []AddCharset
	AllowOverride            AllowOverride
	CGIMapExtension          []CGIMapExtension
	DefaultRuntimeDir        string
	DefaultType              string
	Define                   []Define
	DocumentRoot             string
	EnableMMAP               bool
	EnableSendfile           bool
	ErrorDocument            []ErrorDocument
	ErrorLog                 []string
	ErrorLogFormat           []ErrorLogFormat
	ExtendedStatus           bool
	FileETag                 FileETag
	FlushMaxPipelined        int
	FlushMaxThreshold        int
	ForceType                string
	GprofDir                 string
	HostnameLookups          HostnameLookups
	Include                  []string
	IncludeOptional          []string
	KeepAlive                bool
	KeepAliveTimeout         int
	LimitInternalRecursion   int
	LimitRequestBody         int
	LimitRequestFields       int
	LimitRequestFieldSize    int
	LimitRequestLine         int
	LimitXMLRequestBody      int
	LogLevel                 []LogLevel
	MaxKeepAliveRequests     int
	MaxRangeOverlaps         MaxRangeOverlaps
	MaxRangeReversals        int
	MaxRanges                int
	MergeSlashes             bool
	MergeTrailers            bool
	Mutex                    Mutex
	NameVirtualHost          string
	Options                  []Option
	Protocols                []string
	ProtocolsHonorOrder      bool
	QualifyRedirectURL       bool
	ReadBufferSize           int
	RegexDefaultOptions      RegexDefaultOptions
	RegisterHttpMethod       []string
	RLimitCPU                Limiter
	RLimitMEM                Limiter
	RLimitNPROC              Limiter
	ScriptInterpreterSource  ScriptInterpreterSource
	SeeRequestTail           bool
	ServerAdmin              string
	ServerAlias              []string
	ServerName               string
	ServerPath               string
	ServerRoot               string
	ServerSignature          ServerSignature
	ServerTokens             ServerTokens
	SetHandler               string
	SetInputFilter           []string
	SetOutputFilter          []string
	StrictHostCheck          bool
	TimeOut                  int
	TraceEnable              TraceEnable
	UseCanonicalName         UseCanonicalName
	UseCanonicalPhysicalPort bool
}
