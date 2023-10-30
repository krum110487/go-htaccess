package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Core struct{}

func (c *Core) DirAcceptFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AcceptFilter directive
	return errors.New("AcceptFilter is not yet implemented")
}

func (c *Core) DirAcceptPathInfo(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AcceptPathInfo directive
	return errors.New("AcceptPathInfo is not yet implemented")
}

func (c *Core) DirAccessFileName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AccessFileName directive
	return errors.New("AccessFileName is not yet implemented")
}

func (c *Core) DirAddDefaultCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddDefaultCharset directive
	return errors.New("AddDefaultCharset is not yet implemented")
}

func (c *Core) DirAllowEncodedSlashes(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowEncodedSlashes directive
	return errors.New("AllowEncodedSlashes is not yet implemented")
}

func (c *Core) DirAllowOverride(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowOverride directive
	return errors.New("AllowOverride is not yet implemented")
}

func (c *Core) DirAllowOverrideList(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowOverrideList directive
	return errors.New("AllowOverrideList is not yet implemented")
}

func (c *Core) DirCGIMapExtension(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIMapExtension directive
	return errors.New("CGIMapExtension is not yet implemented")
}

func (c *Core) DirCGIPassAuth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIPassAuth directive
	return errors.New("CGIPassAuth is not yet implemented")
}

func (c *Core) DirCGIVar(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIVar directive
	return errors.New("CGIVar is not yet implemented")
}

func (c *Core) DirContentDigest(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ContentDigest directive
	return errors.New("ContentDigest is not yet implemented")
}

func (c *Core) DirDefaultRuntimeDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultRuntimeDir directive
	return errors.New("DefaultRuntimeDir is not yet implemented")
}

func (c *Core) DirDefaultType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultType directive
	return errors.New("DefaultType is not yet implemented")
}

func (c *Core) DirDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Define directive
	return errors.New("Define is not yet implemented")
}

func (c *Core) DirDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Directory> directive
	return errors.New("<Directory> is not yet implemented")
}

func (c *Core) DirDirectoryMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <DirectoryMatch> directive
	return errors.New("<DirectoryMatch> is not yet implemented")
}

func (c *Core) DirDocumentRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DocumentRoot directive
	return errors.New("DocumentRoot is not yet implemented")
}

func (c *Core) DirElse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Else> directive
	return errors.New("<Else> is not yet implemented")
}

func (c *Core) DirElseIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <ElseIf> directive
	return errors.New("<ElseIf> is not yet implemented")
}

func (c *Core) DirEnableMMAP(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableMMAP directive
	return errors.New("EnableMMAP is not yet implemented")
}

func (c *Core) DirEnableSendfile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableSendfile directive
	return errors.New("EnableSendfile is not yet implemented")
}

func (c *Core) DirError(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Error directive
	return errors.New("Error is not yet implemented")
}

func (c *Core) DirErrorDocument(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorDocument directive
	return errors.New("ErrorDocument is not yet implemented")
}

func (c *Core) DirErrorLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorLog directive
	return errors.New("ErrorLog is not yet implemented")
}

func (c *Core) DirErrorLogFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorLogFormat directive
	return errors.New("ErrorLogFormat is not yet implemented")
}

func (c *Core) DirExtendedStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExtendedStatus directive
	return errors.New("ExtendedStatus is not yet implemented")
}

func (c *Core) DirFileETag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FileETag directive
	return errors.New("FileETag is not yet implemented")
}

func (c *Core) DirFiles(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Files> directive
	return errors.New("<Files> is not yet implemented")
}

func (c *Core) DirFilesMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <FilesMatch> directive
	return errors.New("<FilesMatch> is not yet implemented")
}

func (c *Core) DirFlushMaxPipelined(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FlushMaxPipelined directive
	return errors.New("FlushMaxPipelined is not yet implemented")
}

func (c *Core) DirFlushMaxThreshold(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FlushMaxThreshold directive
	return errors.New("FlushMaxThreshold is not yet implemented")
}

func (c *Core) DirForceType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ForceType directive
	return errors.New("ForceType is not yet implemented")
}

func (c *Core) DirGprofDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: GprofDir directive
	return errors.New("GprofDir is not yet implemented")
}

func (c *Core) DirHostnameLookups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HostnameLookups directive
	return errors.New("HostnameLookups is not yet implemented")
}

func (c *Core) DirHttpProtocolOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HttpProtocolOptions directive
	return errors.New("HttpProtocolOptions is not yet implemented")
}

func (c *Core) DirIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <If> directive
	return errors.New("<If> is not yet implemented")
}

func (c *Core) DirIfDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfDefine> directive
	return errors.New("<IfDefine> is not yet implemented")
}

func (c *Core) DirIfDirective(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfDirective> directive
	return errors.New("<IfDirective> is not yet implemented")
}

func (c *Core) DirIfFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfFile> directive
	return errors.New("<IfFile> is not yet implemented")
}

func (c *Core) DirIfModule(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfModule> directive
	return errors.New("<IfModule> is not yet implemented")
}

func (c *Core) DirIfSection(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfSection> directive
	return errors.New("<IfSection> is not yet implemented")
}

func (c *Core) DirInclude(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Include directive
	return errors.New("Include is not yet implemented")
}

func (c *Core) DirIncludeOptional(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IncludeOptional directive
	return errors.New("IncludeOptional is not yet implemented")
}

func (c *Core) DirKeepAlive(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: KeepAlive directive
	return errors.New("KeepAlive is not yet implemented")
}

func (c *Core) DirKeepAliveTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: KeepAliveTimeout directive
	return errors.New("KeepAliveTimeout is not yet implemented")
}

func (c *Core) DirLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Limit> directive
	return errors.New("<Limit> is not yet implemented")
}

func (c *Core) DirLimitExcept(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <LimitExcept> directive
	return errors.New("<LimitExcept> is not yet implemented")
}

func (c *Core) DirLimitInternalRecursion(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitInternalRecursion directive
	return errors.New("LimitInternalRecursion is not yet implemented")
}

func (c *Core) DirLimitRequestBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestBody directive
	return errors.New("LimitRequestBody is not yet implemented")
}

func (c *Core) DirLimitRequestFields(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestFields directive
	return errors.New("LimitRequestFields is not yet implemented")
}

func (c *Core) DirLimitRequestFieldSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestFieldSize directive
	return errors.New("LimitRequestFieldSize is not yet implemented")
}

func (c *Core) DirLimitRequestLine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestLine directive
	return errors.New("LimitRequestLine is not yet implemented")
}

func (c *Core) DirLimitXMLRequestBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitXMLRequestBody directive
	return errors.New("LimitXMLRequestBody is not yet implemented")
}

func (c *Core) DirLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Location> directive
	return errors.New("<Location> is not yet implemented")
}

func (c *Core) DirLocationMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <LocationMatch> directive
	return errors.New("<LocationMatch> is not yet implemented")
}

func (c *Core) DirLogLevel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LogLevel directive
	return errors.New("LogLevel is not yet implemented")
}

func (c *Core) DirMaxKeepAliveRequests(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxKeepAliveRequests directive
	return errors.New("MaxKeepAliveRequests is not yet implemented")
}

func (c *Core) DirMaxRangeOverlaps(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRangeOverlaps directive
	return errors.New("MaxRangeOverlaps is not yet implemented")
}

func (c *Core) DirMaxRangeReversals(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRangeReversals directive
	return errors.New("MaxRangeReversals is not yet implemented")
}

func (c *Core) DirMaxRanges(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRanges directive
	return errors.New("MaxRanges is not yet implemented")
}

func (c *Core) DirMergeSlashes(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MergeSlashes directive
	return errors.New("MergeSlashes is not yet implemented")
}

func (c *Core) DirMergeTrailers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MergeTrailers directive
	return errors.New("MergeTrailers is not yet implemented")
}

func (c *Core) DirMutex(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Mutex directive
	return errors.New("Mutex is not yet implemented")
}

func (c *Core) DirNameVirtualHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: NameVirtualHost directive
	return errors.New("NameVirtualHost is not yet implemented")
}

func (c *Core) DirOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Options directive
	return errors.New("Options is not yet implemented")
}

func (c *Core) DirProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Protocol directive
	return errors.New("Protocol is not yet implemented")
}

func (c *Core) DirProtocols(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Protocols directive
	return errors.New("Protocols is not yet implemented")
}

func (c *Core) DirProtocolsHonorOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProtocolsHonorOrder directive
	return errors.New("ProtocolsHonorOrder is not yet implemented")
}

func (c *Core) DirQualifyRedirectURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: QualifyRedirectURL directive
	return errors.New("QualifyRedirectURL is not yet implemented")
}

func (c *Core) DirReadBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReadBufferSize directive
	return errors.New("ReadBufferSize is not yet implemented")
}

func (c *Core) DirRegexDefaultOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RegexDefaultOptions directive
	return errors.New("RegexDefaultOptions is not yet implemented")
}

func (c *Core) DirRegisterHttpMethod(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RegisterHttpMethod directive
	return errors.New("RegisterHttpMethod is not yet implemented")
}

func (c *Core) DirRLimitCPU(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitCPU directive
	return errors.New("RLimitCPU is not yet implemented")
}

func (c *Core) DirRLimitMEM(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitMEM directive
	return errors.New("RLimitMEM is not yet implemented")
}

func (c *Core) DirRLimitNPROC(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitNPROC directive
	return errors.New("RLimitNPROC is not yet implemented")
}

func (c *Core) DirScriptInterpreterSource(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptInterpreterSource directive
	return errors.New("ScriptInterpreterSource is not yet implemented")
}

func (c *Core) DirSeeRequestTail(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SeeRequestTail directive
	return errors.New("SeeRequestTail is not yet implemented")
}

func (c *Core) DirServerAdmin(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerAdmin directive
	return errors.New("ServerAdmin is not yet implemented")
}

func (c *Core) DirServerAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerAlias directive
	return errors.New("ServerAlias is not yet implemented")
}

func (c *Core) DirServerName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerName directive
	return errors.New("ServerName is not yet implemented")
}

func (c *Core) DirServerPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerPath directive
	return errors.New("ServerPath is not yet implemented")
}

func (c *Core) DirServerRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerRoot directive
	return errors.New("ServerRoot is not yet implemented")
}

func (c *Core) DirServerSignature(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerSignature directive
	return errors.New("ServerSignature is not yet implemented")
}

func (c *Core) DirServerTokens(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerTokens directive
	return errors.New("ServerTokens is not yet implemented")
}

func (c *Core) DirSetHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetHandler directive
	return errors.New("SetHandler is not yet implemented")
}

func (c *Core) DirSetInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetInputFilter directive
	return errors.New("SetInputFilter is not yet implemented")
}

func (c *Core) DirSetOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetOutputFilter directive
	return errors.New("SetOutputFilter is not yet implemented")
}

func (c *Core) DirStrictHostCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StrictHostCheck directive
	return errors.New("StrictHostCheck is not yet implemented")
}

func (c *Core) DirTimeOut(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TimeOut directive
	return errors.New("TimeOut is not yet implemented")
}

func (c *Core) DirTraceEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TraceEnable directive
	return errors.New("TraceEnable is not yet implemented")
}

func (c *Core) DirUnDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UnDefine directive
	return errors.New("UnDefine is not yet implemented")
}

func (c *Core) DirUseCanonicalName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UseCanonicalName directive
	return errors.New("UseCanonicalName is not yet implemented")
}

func (c *Core) DirUseCanonicalPhysicalPort(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UseCanonicalPhysicalPort directive
	return errors.New("UseCanonicalPhysicalPort is not yet implemented")
}

func (c *Core) DirVirtualHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <VirtualHost> directive
	return errors.New("<VirtualHost> is not yet implemented")
}
