package mod

import (
	"errors"
	"net/http"

	"github.com/krum110487/go-htaccess"
	"github.com/krum110487/go-htaccess/parser"
)

type Core struct{}

func (c *Core) dirAcceptFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AcceptFilter directive
	return errors.New("AcceptFilter is not yet implemented")
}

func (c *Core) dirAcceptPathInfo(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AcceptPathInfo directive
	return errors.New("AcceptPathInfo is not yet implemented")
}

func (c *Core) dirAccessFileName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AccessFileName directive
	return errors.New("AccessFileName is not yet implemented")
}

func (c *Core) dirAddDefaultCharset(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AddDefaultCharset directive
	return errors.New("AddDefaultCharset is not yet implemented")
}

func (c *Core) dirAllowEncodedSlashes(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowEncodedSlashes directive
	return errors.New("AllowEncodedSlashes is not yet implemented")
}

func (c *Core) dirAllowOverride(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowOverride directive
	return errors.New("AllowOverride is not yet implemented")
}

func (c *Core) dirAllowOverrideList(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: AllowOverrideList directive
	return errors.New("AllowOverrideList is not yet implemented")
}

func (c *Core) dirCGIMapExtension(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIMapExtension directive
	return errors.New("CGIMapExtension is not yet implemented")
}

func (c *Core) dirCGIPassAuth(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIPassAuth directive
	return errors.New("CGIPassAuth is not yet implemented")
}

func (c *Core) dirCGIVar(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: CGIVar directive
	return errors.New("CGIVar is not yet implemented")
}

func (c *Core) dirContentDigest(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ContentDigest directive
	return errors.New("ContentDigest is not yet implemented")
}

func (c *Core) dirDefaultRuntimeDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultRuntimeDir directive
	return errors.New("DefaultRuntimeDir is not yet implemented")
}

func (c *Core) dirDefaultType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DefaultType directive
	return errors.New("DefaultType is not yet implemented")
}

func (c *Core) dirDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Define directive
	return errors.New("Define is not yet implemented")
}

func (c *Core) dirDirectory(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Directory> directive
	return errors.New("<Directory> is not yet implemented")
}

func (c *Core) dirDirectoryMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <DirectoryMatch> directive
	return errors.New("<DirectoryMatch> is not yet implemented")
}

func (c *Core) dirDocumentRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: DocumentRoot directive
	return errors.New("DocumentRoot is not yet implemented")
}

func (c *Core) dirElse(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Else> directive
	return errors.New("<Else> is not yet implemented")
}

func (c *Core) dirElseIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <ElseIf> directive
	return errors.New("<ElseIf> is not yet implemented")
}

func (c *Core) dirEnableMMAP(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableMMAP directive
	return errors.New("EnableMMAP is not yet implemented")
}

func (c *Core) dirEnableSendfile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: EnableSendfile directive
	return errors.New("EnableSendfile is not yet implemented")
}

func (c *Core) dirError(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Error directive
	return errors.New("Error is not yet implemented")
}

func (c *Core) dirErrorDocument(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorDocument directive
	return errors.New("ErrorDocument is not yet implemented")
}

func (c *Core) dirErrorLog(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorLog directive
	return errors.New("ErrorLog is not yet implemented")
}

func (c *Core) dirErrorLogFormat(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ErrorLogFormat directive
	return errors.New("ErrorLogFormat is not yet implemented")
}

func (c *Core) dirExtendedStatus(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ExtendedStatus directive
	return errors.New("ExtendedStatus is not yet implemented")
}

func (c *Core) dirFileETag(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FileETag directive
	return errors.New("FileETag is not yet implemented")
}

func (c *Core) dirFiles(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Files> directive
	return errors.New("<Files> is not yet implemented")
}

func (c *Core) dirFilesMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <FilesMatch> directive
	return errors.New("<FilesMatch> is not yet implemented")
}

func (c *Core) dirFlushMaxPipelined(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FlushMaxPipelined directive
	return errors.New("FlushMaxPipelined is not yet implemented")
}

func (c *Core) dirFlushMaxThreshold(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: FlushMaxThreshold directive
	return errors.New("FlushMaxThreshold is not yet implemented")
}

func (c *Core) dirForceType(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ForceType directive
	return errors.New("ForceType is not yet implemented")
}

func (c *Core) dirGprofDir(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: GprofDir directive
	return errors.New("GprofDir is not yet implemented")
}

func (c *Core) dirHostnameLookups(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HostnameLookups directive
	return errors.New("HostnameLookups is not yet implemented")
}

func (c *Core) dirHttpProtocolOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: HttpProtocolOptions directive
	return errors.New("HttpProtocolOptions is not yet implemented")
}

func (c *Core) dirIf(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <If> directive
	return errors.New("<If> is not yet implemented")
}

func (c *Core) dirIfDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfDefine> directive
	return errors.New("<IfDefine> is not yet implemented")
}

func (c *Core) dirIfDirective(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfDirective> directive
	return errors.New("<IfDirective> is not yet implemented")
}

func (c *Core) dirIfFile(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfFile> directive
	return errors.New("<IfFile> is not yet implemented")
}

func (c *Core) dirIfModule(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfModule> directive
	return errors.New("<IfModule> is not yet implemented")
}

func (c *Core) dirIfSection(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <IfSection> directive
	return errors.New("<IfSection> is not yet implemented")
}

func (c *Core) dirInclude(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Include directive
	return errors.New("Include is not yet implemented")
}

func (c *Core) dirIncludeOptional(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: IncludeOptional directive
	return errors.New("IncludeOptional is not yet implemented")
}

func (c *Core) dirKeepAlive(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: KeepAlive directive
	return errors.New("KeepAlive is not yet implemented")
}

func (c *Core) dirKeepAliveTimeout(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: KeepAliveTimeout directive
	return errors.New("KeepAliveTimeout is not yet implemented")
}

func (c *Core) dirLimit(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Limit> directive
	return errors.New("<Limit> is not yet implemented")
}

func (c *Core) dirLimitExcept(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <LimitExcept> directive
	return errors.New("<LimitExcept> is not yet implemented")
}

func (c *Core) dirLimitInternalRecursion(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitInternalRecursion directive
	return errors.New("LimitInternalRecursion is not yet implemented")
}

func (c *Core) dirLimitRequestBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestBody directive
	return errors.New("LimitRequestBody is not yet implemented")
}

func (c *Core) dirLimitRequestFields(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestFields directive
	return errors.New("LimitRequestFields is not yet implemented")
}

func (c *Core) dirLimitRequestFieldSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestFieldSize directive
	return errors.New("LimitRequestFieldSize is not yet implemented")
}

func (c *Core) dirLimitRequestLine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitRequestLine directive
	return errors.New("LimitRequestLine is not yet implemented")
}

func (c *Core) dirLimitXMLRequestBody(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LimitXMLRequestBody directive
	return errors.New("LimitXMLRequestBody is not yet implemented")
}

func (c *Core) dirLocation(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <Location> directive
	return errors.New("<Location> is not yet implemented")
}

func (c *Core) dirLocationMatch(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <LocationMatch> directive
	return errors.New("<LocationMatch> is not yet implemented")
}

func (c *Core) dirLogLevel(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: LogLevel directive
	return errors.New("LogLevel is not yet implemented")
}

func (c *Core) dirMaxKeepAliveRequests(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxKeepAliveRequests directive
	return errors.New("MaxKeepAliveRequests is not yet implemented")
}

func (c *Core) dirMaxRangeOverlaps(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRangeOverlaps directive
	return errors.New("MaxRangeOverlaps is not yet implemented")
}

func (c *Core) dirMaxRangeReversals(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRangeReversals directive
	return errors.New("MaxRangeReversals is not yet implemented")
}

func (c *Core) dirMaxRanges(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MaxRanges directive
	return errors.New("MaxRanges is not yet implemented")
}

func (c *Core) dirMergeSlashes(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MergeSlashes directive
	return errors.New("MergeSlashes is not yet implemented")
}

func (c *Core) dirMergeTrailers(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: MergeTrailers directive
	return errors.New("MergeTrailers is not yet implemented")
}

func (c *Core) dirMutex(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Mutex directive
	return errors.New("Mutex is not yet implemented")
}

func (c *Core) dirNameVirtualHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: NameVirtualHost directive
	return errors.New("NameVirtualHost is not yet implemented")
}

func (c *Core) dirOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Options directive
	return errors.New("Options is not yet implemented")
}

func (c *Core) dirProtocol(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Protocol directive
	return errors.New("Protocol is not yet implemented")
}

func (c *Core) dirProtocols(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: Protocols directive
	return errors.New("Protocols is not yet implemented")
}

func (c *Core) dirProtocolsHonorOrder(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ProtocolsHonorOrder directive
	return errors.New("ProtocolsHonorOrder is not yet implemented")
}

func (c *Core) dirQualifyRedirectURL(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: QualifyRedirectURL directive
	return errors.New("QualifyRedirectURL is not yet implemented")
}

func (c *Core) dirReadBufferSize(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ReadBufferSize directive
	return errors.New("ReadBufferSize is not yet implemented")
}

func (c *Core) dirRegexDefaultOptions(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RegexDefaultOptions directive
	return errors.New("RegexDefaultOptions is not yet implemented")
}

func (c *Core) dirRegisterHttpMethod(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RegisterHttpMethod directive
	return errors.New("RegisterHttpMethod is not yet implemented")
}

func (c *Core) dirRLimitCPU(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitCPU directive
	return errors.New("RLimitCPU is not yet implemented")
}

func (c *Core) dirRLimitMEM(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitMEM directive
	return errors.New("RLimitMEM is not yet implemented")
}

func (c *Core) dirRLimitNPROC(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: RLimitNPROC directive
	return errors.New("RLimitNPROC is not yet implemented")
}

func (c *Core) dirScriptInterpreterSource(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ScriptInterpreterSource directive
	return errors.New("ScriptInterpreterSource is not yet implemented")
}

func (c *Core) dirSeeRequestTail(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SeeRequestTail directive
	return errors.New("SeeRequestTail is not yet implemented")
}

func (c *Core) dirServerAdmin(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerAdmin directive
	return errors.New("ServerAdmin is not yet implemented")
}

func (c *Core) dirServerAlias(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerAlias directive
	return errors.New("ServerAlias is not yet implemented")
}

func (c *Core) dirServerName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerName directive
	return errors.New("ServerName is not yet implemented")
}

func (c *Core) dirServerPath(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerPath directive
	return errors.New("ServerPath is not yet implemented")
}

func (c *Core) dirServerRoot(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerRoot directive
	return errors.New("ServerRoot is not yet implemented")
}

func (c *Core) dirServerSignature(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerSignature directive
	return errors.New("ServerSignature is not yet implemented")
}

func (c *Core) dirServerTokens(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: ServerTokens directive
	return errors.New("ServerTokens is not yet implemented")
}

func (c *Core) dirSetHandler(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetHandler directive
	return errors.New("SetHandler is not yet implemented")
}

func (c *Core) dirSetInputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetInputFilter directive
	return errors.New("SetInputFilter is not yet implemented")
}

func (c *Core) dirSetOutputFilter(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: SetOutputFilter directive
	return errors.New("SetOutputFilter is not yet implemented")
}

func (c *Core) dirStrictHostCheck(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: StrictHostCheck directive
	return errors.New("StrictHostCheck is not yet implemented")
}

func (c *Core) dirTimeOut(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TimeOut directive
	return errors.New("TimeOut is not yet implemented")
}

func (c *Core) dirTraceEnable(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: TraceEnable directive
	return errors.New("TraceEnable is not yet implemented")
}

func (c *Core) dirUnDefine(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UnDefine directive
	return errors.New("UnDefine is not yet implemented")
}

func (c *Core) dirUseCanonicalName(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UseCanonicalName directive
	return errors.New("UseCanonicalName is not yet implemented")
}

func (c *Core) dirUseCanonicalPhysicalPort(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: UseCanonicalPhysicalPort directive
	return errors.New("UseCanonicalPhysicalPort is not yet implemented")
}

func (c *Core) dirVirtualHost(dir parser.DirectiveEntry, req *http.Request, runCtx *htaccess.Context) error {
	//TODO: <VirtualHost> directive
	return errors.New("<VirtualHost> is not yet implemented")
}
