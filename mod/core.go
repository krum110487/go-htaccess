package mod

import (
	"errors"
)

type Core struct{}

func (Core) DirAcceptFilter(runCtx *Context) error {
	//TODO: AcceptFilter directive
	return errors.New("AcceptFilter is not yet implemented")
}

func (Core) DirAcceptPathInfo(runCtx *Context) error {
	//TODO: AcceptPathInfo directive
	return errors.New("AcceptPathInfo is not yet implemented")
}

func (Core) DirAccessFileName(runCtx *Context) error {
	//TODO: AccessFileName directive
	return errors.New("AccessFileName is not yet implemented")
}

func (Core) DirAddDefaultCharset(runCtx *Context) error {
	//TODO: AddDefaultCharset directive
	return errors.New("AddDefaultCharset is not yet implemented")
}

func (Core) DirAllowEncodedSlashes(runCtx *Context) error {
	//TODO: AllowEncodedSlashes directive
	return errors.New("AllowEncodedSlashes is not yet implemented")
}

func (Core) DirAllowOverride(runCtx *Context) error {
	//TODO: AllowOverride directive
	return errors.New("AllowOverride is not yet implemented")
}

func (Core) DirAllowOverrideList(runCtx *Context) error {
	//TODO: AllowOverrideList directive
	return errors.New("AllowOverrideList is not yet implemented")
}

func (Core) DirCGIMapExtension(runCtx *Context) error {
	//TODO: CGIMapExtension directive
	return errors.New("CGIMapExtension is not yet implemented")
}

func (Core) DirCGIPassAuth(runCtx *Context) error {
	//TODO: CGIPassAuth directive
	return errors.New("CGIPassAuth is not yet implemented")
}

func (Core) DirCGIVar(runCtx *Context) error {
	//TODO: CGIVar directive
	return errors.New("CGIVar is not yet implemented")
}

func (Core) DirContentDigest(runCtx *Context) error {
	//TODO: ContentDigest directive
	return errors.New("ContentDigest is not yet implemented")
}

func (Core) DirDefaultRuntimeDir(runCtx *Context) error {
	//TODO: DefaultRuntimeDir directive
	return errors.New("DefaultRuntimeDir is not yet implemented")
}

func (Core) DirDefaultType(runCtx *Context) error {
	//TODO: DefaultType directive
	return errors.New("DefaultType is not yet implemented")
}

func (Core) DirDefine(runCtx *Context) error {
	//TODO: Define directive
	return errors.New("Define is not yet implemented")
}

func (Core) DirDirectory(runCtx *Context) error {
	//TODO: <Directory> directive
	return errors.New("<Directory> is not yet implemented")
}

func (Core) DirDirectoryMatch(runCtx *Context) error {
	//TODO: <DirectoryMatch> directive
	return errors.New("<DirectoryMatch> is not yet implemented")
}

func (Core) DirDocumentRoot(runCtx *Context) error {
	//TODO: DocumentRoot directive
	return errors.New("DocumentRoot is not yet implemented")
}

func (Core) DirElse(runCtx *Context) error {
	//TODO: <Else> directive
	return errors.New("<Else> is not yet implemented")
}

func (Core) DirElseIf(runCtx *Context) error {
	//TODO: <ElseIf> directive
	return errors.New("<ElseIf> is not yet implemented")
}

func (Core) DirEnableMMAP(runCtx *Context) error {
	//TODO: EnableMMAP directive
	return errors.New("EnableMMAP is not yet implemented")
}

func (Core) DirEnableSendfile(runCtx *Context) error {
	//TODO: EnableSendfile directive
	return errors.New("EnableSendfile is not yet implemented")
}

func (Core) DirError(runCtx *Context) error {
	//TODO: Error directive
	return errors.New("Error is not yet implemented")
}

func (Core) DirErrorDocument(runCtx *Context) error {
	//TODO: ErrorDocument directive
	return errors.New("ErrorDocument is not yet implemented")
}

func (Core) DirErrorLog(runCtx *Context) error {
	//TODO: ErrorLog directive
	return errors.New("ErrorLog is not yet implemented")
}

func (Core) DirErrorLogFormat(runCtx *Context) error {
	//TODO: ErrorLogFormat directive
	return errors.New("ErrorLogFormat is not yet implemented")
}

func (Core) DirExtendedStatus(runCtx *Context) error {
	//TODO: ExtendedStatus directive
	return errors.New("ExtendedStatus is not yet implemented")
}

func (Core) DirFileETag(runCtx *Context) error {
	//TODO: FileETag directive
	return errors.New("FileETag is not yet implemented")
}

func (Core) DirFiles(runCtx *Context) error {
	//TODO: <Files> directive
	return errors.New("<Files> is not yet implemented")
}

func (Core) DirFilesMatch(runCtx *Context) error {
	//TODO: <FilesMatch> directive
	return errors.New("<FilesMatch> is not yet implemented")
}

func (Core) DirFlushMaxPipelined(runCtx *Context) error {
	//TODO: FlushMaxPipelined directive
	return errors.New("FlushMaxPipelined is not yet implemented")
}

func (Core) DirFlushMaxThreshold(runCtx *Context) error {
	//TODO: FlushMaxThreshold directive
	return errors.New("FlushMaxThreshold is not yet implemented")
}

func (Core) DirForceType(runCtx *Context) error {
	//TODO: ForceType directive
	return errors.New("ForceType is not yet implemented")
}

func (Core) DirGprofDir(runCtx *Context) error {
	//TODO: GprofDir directive
	return errors.New("GprofDir is not yet implemented")
}

func (Core) DirHostnameLookups(runCtx *Context) error {
	//TODO: HostnameLookups directive
	return errors.New("HostnameLookups is not yet implemented")
}

func (Core) DirHttpProtocolOptions(runCtx *Context) error {
	//TODO: HttpProtocolOptions directive
	return errors.New("HttpProtocolOptions is not yet implemented")
}

func (Core) DirIf(runCtx *Context) error {
	//TODO: <If> directive
	return errors.New("<If> is not yet implemented")
}

func (Core) DirIfDefine(runCtx *Context) error {
	//TODO: <IfDefine> directive
	return errors.New("<IfDefine> is not yet implemented")
}

func (Core) DirIfDirective(runCtx *Context) error {
	//TODO: <IfDirective> directive
	return errors.New("<IfDirective> is not yet implemented")
}

func (Core) DirIfFile(runCtx *Context) error {
	//TODO: <IfFile> directive
	return errors.New("<IfFile> is not yet implemented")
}

func (Core) DirIfModule(runCtx *Context) error {
	//TODO: <IfModule> directive
	return errors.New("<IfModule> is not yet implemented")
}

func (Core) DirIfSection(runCtx *Context) error {
	//TODO: <IfSection> directive
	return errors.New("<IfSection> is not yet implemented")
}

func (Core) DirInclude(runCtx *Context) error {
	//TODO: Include directive
	return errors.New("Include is not yet implemented")
}

func (Core) DirIncludeOptional(runCtx *Context) error {
	//TODO: IncludeOptional directive
	return errors.New("IncludeOptional is not yet implemented")
}

func (Core) DirKeepAlive(runCtx *Context) error {
	//TODO: KeepAlive directive
	return errors.New("KeepAlive is not yet implemented")
}

func (Core) DirKeepAliveTimeout(runCtx *Context) error {
	//TODO: KeepAliveTimeout directive
	return errors.New("KeepAliveTimeout is not yet implemented")
}

func (Core) DirLimit(runCtx *Context) error {
	//TODO: <Limit> directive
	return errors.New("<Limit> is not yet implemented")
}

func (Core) DirLimitExcept(runCtx *Context) error {
	//TODO: <LimitExcept> directive
	return errors.New("<LimitExcept> is not yet implemented")
}

func (Core) DirLimitInternalRecursion(runCtx *Context) error {
	//TODO: LimitInternalRecursion directive
	return errors.New("LimitInternalRecursion is not yet implemented")
}

func (Core) DirLimitRequestBody(runCtx *Context) error {
	//TODO: LimitRequestBody directive
	return errors.New("LimitRequestBody is not yet implemented")
}

func (Core) DirLimitRequestFields(runCtx *Context) error {
	//TODO: LimitRequestFields directive
	return errors.New("LimitRequestFields is not yet implemented")
}

func (Core) DirLimitRequestFieldSize(runCtx *Context) error {
	//TODO: LimitRequestFieldSize directive
	return errors.New("LimitRequestFieldSize is not yet implemented")
}

func (Core) DirLimitRequestLine(runCtx *Context) error {
	//TODO: LimitRequestLine directive
	return errors.New("LimitRequestLine is not yet implemented")
}

func (Core) DirLimitXMLRequestBody(runCtx *Context) error {
	//TODO: LimitXMLRequestBody directive
	return errors.New("LimitXMLRequestBody is not yet implemented")
}

func (Core) DirLocation(runCtx *Context) error {
	//TODO: <Location> directive
	return errors.New("<Location> is not yet implemented")
}

func (Core) DirLocationMatch(runCtx *Context) error {
	//TODO: <LocationMatch> directive
	return errors.New("<LocationMatch> is not yet implemented")
}

func (Core) DirLogLevel(runCtx *Context) error {
	//TODO: LogLevel directive
	return errors.New("LogLevel is not yet implemented")
}

func (Core) DirMaxKeepAliveRequests(runCtx *Context) error {
	//TODO: MaxKeepAliveRequests directive
	return errors.New("MaxKeepAliveRequests is not yet implemented")
}

func (Core) DirMaxRangeOverlaps(runCtx *Context) error {
	//TODO: MaxRangeOverlaps directive
	return errors.New("MaxRangeOverlaps is not yet implemented")
}

func (Core) DirMaxRangeReversals(runCtx *Context) error {
	//TODO: MaxRangeReversals directive
	return errors.New("MaxRangeReversals is not yet implemented")
}

func (Core) DirMaxRanges(runCtx *Context) error {
	//TODO: MaxRanges directive
	return errors.New("MaxRanges is not yet implemented")
}

func (Core) DirMergeSlashes(runCtx *Context) error {
	//TODO: MergeSlashes directive
	return errors.New("MergeSlashes is not yet implemented")
}

func (Core) DirMergeTrailers(runCtx *Context) error {
	//TODO: MergeTrailers directive
	return errors.New("MergeTrailers is not yet implemented")
}

func (Core) DirMutex(runCtx *Context) error {
	//TODO: Mutex directive
	return errors.New("Mutex is not yet implemented")
}

func (Core) DirNameVirtualHost(runCtx *Context) error {
	//TODO: NameVirtualHost directive
	return errors.New("NameVirtualHost is not yet implemented")
}

func (Core) DirOptions(runCtx *Context) error {
	//TODO: Options directive
	return errors.New("Options is not yet implemented")
}

func (Core) DirProtocol(runCtx *Context) error {
	//TODO: Protocol directive
	return errors.New("Protocol is not yet implemented")
}

func (Core) DirProtocols(runCtx *Context) error {
	//TODO: Protocols directive
	return errors.New("Protocols is not yet implemented")
}

func (Core) DirProtocolsHonorOrder(runCtx *Context) error {
	//TODO: ProtocolsHonorOrder directive
	return errors.New("ProtocolsHonorOrder is not yet implemented")
}

func (Core) DirQualifyRedirectURL(runCtx *Context) error {
	//TODO: QualifyRedirectURL directive
	return errors.New("QualifyRedirectURL is not yet implemented")
}

func (Core) DirReadBufferSize(runCtx *Context) error {
	//TODO: ReadBufferSize directive
	return errors.New("ReadBufferSize is not yet implemented")
}

func (Core) DirRegexDefaultOptions(runCtx *Context) error {
	//TODO: RegexDefaultOptions directive
	return errors.New("RegexDefaultOptions is not yet implemented")
}

func (Core) DirRegisterHttpMethod(runCtx *Context) error {
	//TODO: RegisterHttpMethod directive
	return errors.New("RegisterHttpMethod is not yet implemented")
}

func (Core) DirRLimitCPU(runCtx *Context) error {
	//TODO: RLimitCPU directive
	return errors.New("RLimitCPU is not yet implemented")
}

func (Core) DirRLimitMEM(runCtx *Context) error {
	//TODO: RLimitMEM directive
	return errors.New("RLimitMEM is not yet implemented")
}

func (Core) DirRLimitNPROC(runCtx *Context) error {
	//TODO: RLimitNPROC directive
	return errors.New("RLimitNPROC is not yet implemented")
}

func (Core) DirScriptInterpreterSource(runCtx *Context) error {
	//TODO: ScriptInterpreterSource directive
	return errors.New("ScriptInterpreterSource is not yet implemented")
}

func (Core) DirSeeRequestTail(runCtx *Context) error {
	//TODO: SeeRequestTail directive
	return errors.New("SeeRequestTail is not yet implemented")
}

func (Core) DirServerAdmin(runCtx *Context) error {
	//TODO: ServerAdmin directive
	return errors.New("ServerAdmin is not yet implemented")
}

func (Core) DirServerAlias(runCtx *Context) error {
	//TODO: ServerAlias directive
	return errors.New("ServerAlias is not yet implemented")
}

func (Core) DirServerName(runCtx *Context) error {
	//TODO: ServerName directive
	return errors.New("ServerName is not yet implemented")
}

func (Core) DirServerPath(runCtx *Context) error {
	//TODO: ServerPath directive
	return errors.New("ServerPath is not yet implemented")
}

func (Core) DirServerRoot(runCtx *Context) error {
	//TODO: ServerRoot directive
	return errors.New("ServerRoot is not yet implemented")
}

func (Core) DirServerSignature(runCtx *Context) error {
	//TODO: ServerSignature directive
	return errors.New("ServerSignature is not yet implemented")
}

func (Core) DirServerTokens(runCtx *Context) error {
	//TODO: ServerTokens directive
	return errors.New("ServerTokens is not yet implemented")
}

func (Core) DirSetHandler(runCtx *Context) error {
	//TODO: SetHandler directive
	return errors.New("SetHandler is not yet implemented")
}

func (Core) DirSetInputFilter(runCtx *Context) error {
	//TODO: SetInputFilter directive
	return errors.New("SetInputFilter is not yet implemented")
}

func (Core) DirSetOutputFilter(runCtx *Context) error {
	//TODO: SetOutputFilter directive
	return errors.New("SetOutputFilter is not yet implemented")
}

func (Core) DirStrictHostCheck(runCtx *Context) error {
	//TODO: StrictHostCheck directive
	return errors.New("StrictHostCheck is not yet implemented")
}

func (Core) DirTimeOut(runCtx *Context) error {
	//TODO: TimeOut directive
	return errors.New("TimeOut is not yet implemented")
}

func (Core) DirTraceEnable(runCtx *Context) error {
	//TODO: TraceEnable directive
	return errors.New("TraceEnable is not yet implemented")
}

func (Core) DirUnDefine(runCtx *Context) error {
	//TODO: UnDefine directive
	return errors.New("UnDefine is not yet implemented")
}

func (Core) DirUseCanonicalName(runCtx *Context) error {
	//TODO: UseCanonicalName directive
	return errors.New("UseCanonicalName is not yet implemented")
}

func (Core) DirUseCanonicalPhysicalPort(runCtx *Context) error {
	//TODO: UseCanonicalPhysicalPort directive
	return errors.New("UseCanonicalPhysicalPort is not yet implemented")
}

func (Core) DirVirtualHost(runCtx *Context) error {
	//TODO: <VirtualHost> directive
	return errors.New("<VirtualHost> is not yet implemented")
}
