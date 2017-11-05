package glib

/*
#cgo pkg-config: glib-2.0
#include <glib.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "github.com/electricface/go-auto-gir/util"

// Struct Variant
type Variant struct {
	Ptr unsafe.Pointer
}

func (v Variant) native() *C.GVariant {
	return (*C.GVariant)(v.Ptr)
}
func wrapVariant(p *C.GVariant) Variant {
	return Variant{unsafe.Pointer(p)}
}
func WrapVariant(p unsafe.Pointer) Variant {
	return Variant{p}
}

// VariantNewBoolean is a wrapper around g_variant_new_boolean().
func VariantNewBoolean(value bool) Variant {
	ret0 := C.g_variant_new_boolean(C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	return wrapVariant(ret0)
}

// GetStrv is a wrapper around g_variant_get_strv().
func (value Variant) GetStrv() []string {
	var length0 C.gsize
	ret0 := C.g_variant_get_strv(value.native(), &length0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetType is a wrapper around g_variant_get_type().
func (value Variant) GetType() VariantType {
	ret0 := C.g_variant_get_type(value.native())
	return wrapVariantType(ret0)
}

// IsFloating is a wrapper around g_variant_is_floating().
func (value Variant) IsFloating() bool {
	ret0 := C.g_variant_is_floating(value.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_variant_ref().
func (value Variant) Ref() Variant {
	ret0 := C.g_variant_ref(value.native())
	return wrapVariant(ret0)
}

// RefSink is a wrapper around g_variant_ref_sink().
func (value Variant) RefSink() Variant {
	ret0 := C.g_variant_ref_sink(value.native())
	return wrapVariant(ret0)
}

// Unref is a wrapper around g_variant_unref().
func (value Variant) Unref() {
	C.g_variant_unref(value.native())
}

// Struct VariantType
type VariantType struct {
	Ptr unsafe.Pointer
}

func (v VariantType) native() *C.GVariantType {
	return (*C.GVariantType)(v.Ptr)
}
func wrapVariantType(p *C.GVariantType) VariantType {
	return VariantType{unsafe.Pointer(p)}
}
func WrapVariantType(p unsafe.Pointer) VariantType {
	return VariantType{p}
}

// Struct Error
type Error struct {
	Ptr unsafe.Pointer
}

func (v Error) native() *C.GError {
	return (*C.GError)(v.Ptr)
}
func wrapError(p *C.GError) Error {
	return Error{unsafe.Pointer(p)}
}
func WrapError(p unsafe.Pointer) Error {
	return Error{p}
}

// Free is a wrapper around g_error_free().
func (error Error) Free() {
	C.g_error_free(error.native())
}

// Struct MainLoop
type MainLoop struct {
	Ptr unsafe.Pointer
}

func (v MainLoop) native() *C.GMainLoop {
	return (*C.GMainLoop)(v.Ptr)
}
func wrapMainLoop(p *C.GMainLoop) MainLoop {
	return MainLoop{unsafe.Pointer(p)}
}
func WrapMainLoop(p unsafe.Pointer) MainLoop {
	return MainLoop{p}
}

// MainLoopNew is a wrapper around g_main_loop_new().
func MainLoopNew(context MainContext, is_running bool) MainLoop {
	ret0 := C.g_main_loop_new(context.native(), C.gboolean(util.Bool2Int(is_running)) /*go:.util*/)
	return wrapMainLoop(ret0)
}

// GetContext is a wrapper around g_main_loop_get_context().
func (loop MainLoop) GetContext() MainContext {
	ret0 := C.g_main_loop_get_context(loop.native())
	return wrapMainContext(ret0)
}

// IsRunning is a wrapper around g_main_loop_is_running().
func (loop MainLoop) IsRunning() bool {
	ret0 := C.g_main_loop_is_running(loop.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Quit is a wrapper around g_main_loop_quit().
func (loop MainLoop) Quit() {
	C.g_main_loop_quit(loop.native())
}

// Ref is a wrapper around g_main_loop_ref().
func (loop MainLoop) Ref() MainLoop {
	ret0 := C.g_main_loop_ref(loop.native())
	return wrapMainLoop(ret0)
}

// Run is a wrapper around g_main_loop_run().
func (loop MainLoop) Run() {
	C.g_main_loop_run(loop.native())
}

// Unref is a wrapper around g_main_loop_unref().
func (loop MainLoop) Unref() {
	C.g_main_loop_unref(loop.native())
}

// Struct MainContext
type MainContext struct {
	Ptr unsafe.Pointer
}

func (v MainContext) native() *C.GMainContext {
	return (*C.GMainContext)(v.Ptr)
}
func wrapMainContext(p *C.GMainContext) MainContext {
	return MainContext{unsafe.Pointer(p)}
}
func WrapMainContext(p unsafe.Pointer) MainContext {
	return MainContext{p}
}

// MainContextNew is a wrapper around g_main_context_new().
func MainContextNew() MainContext {
	ret0 := C.g_main_context_new()
	return wrapMainContext(ret0)
}

// FindSourceById is a wrapper around g_main_context_find_source_by_id().
func (context MainContext) FindSourceById(source_id uint) Source {
	ret0 := C.g_main_context_find_source_by_id(context.native(), C.guint(source_id))
	return wrapSource(ret0)
}

// Iteration is a wrapper around g_main_context_iteration().
func (context MainContext) Iteration(may_block bool) bool {
	ret0 := C.g_main_context_iteration(context.native(), C.gboolean(util.Bool2Int(may_block)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Pending is a wrapper around g_main_context_pending().
func (context MainContext) Pending() bool {
	ret0 := C.g_main_context_pending(context.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_main_context_ref().
func (context MainContext) Ref() MainContext {
	ret0 := C.g_main_context_ref(context.native())
	return wrapMainContext(ret0)
}

// Unref is a wrapper around g_main_context_unref().
func (context MainContext) Unref() {
	C.g_main_context_unref(context.native())
}

// MainContextDefault is a wrapper around g_main_context_default().
func MainContextDefault() MainContext {
	ret0 := C.g_main_context_default()
	return wrapMainContext(ret0)
}

// Struct Source
type Source struct {
	Ptr unsafe.Pointer
}

func (v Source) native() *C.GSource {
	return (*C.GSource)(v.Ptr)
}
func wrapSource(p *C.GSource) Source {
	return Source{unsafe.Pointer(p)}
}
func WrapSource(p unsafe.Pointer) Source {
	return Source{p}
}

// Attach is a wrapper around g_source_attach().
func (source Source) Attach(context MainContext) uint {
	ret0 := C.g_source_attach(source.native(), context.native())
	return uint(ret0)
}

// Destroy is a wrapper around g_source_destroy().
func (source Source) Destroy() {
	C.g_source_destroy(source.native())
}

// GetCanRecurse is a wrapper around g_source_get_can_recurse().
func (source Source) GetCanRecurse() bool {
	ret0 := C.g_source_get_can_recurse(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetContext is a wrapper around g_source_get_context().
func (source Source) GetContext() MainContext {
	ret0 := C.g_source_get_context(source.native())
	return wrapMainContext(ret0)
}

// GetId is a wrapper around g_source_get_id().
func (source Source) GetId() uint {
	ret0 := C.g_source_get_id(source.native())
	return uint(ret0)
}

// GetName is a wrapper around g_source_get_name().
func (source Source) GetName() string {
	ret0 := C.g_source_get_name(source.native())
	ret := C.GoString(ret0)
	return ret
}

// GetPriority is a wrapper around g_source_get_priority().
func (source Source) GetPriority() int {
	ret0 := C.g_source_get_priority(source.native())
	return int(ret0)
}

// GetTime is a wrapper around g_source_get_time().
func (source Source) GetTime() int64 {
	ret0 := C.g_source_get_time(source.native())
	return int64(ret0)
}

// IsDestroyed is a wrapper around g_source_is_destroyed().
func (source Source) IsDestroyed() bool {
	ret0 := C.g_source_is_destroyed(source.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_source_ref().
func (source Source) Ref() Source {
	ret0 := C.g_source_ref(source.native())
	return wrapSource(ret0)
}

// SetCanRecurse is a wrapper around g_source_set_can_recurse().
func (source Source) SetCanRecurse(can_recurse bool) {
	C.g_source_set_can_recurse(source.native(), C.gboolean(util.Bool2Int(can_recurse)) /*go:.util*/)
}

// SetName is a wrapper around g_source_set_name().
func (source Source) SetName(name string) {
	name0 := C.CString(name)
	C.g_source_set_name(source.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetPriority is a wrapper around g_source_set_priority().
func (source Source) SetPriority(priority int) {
	C.g_source_set_priority(source.native(), C.gint(priority))
}

// Unref is a wrapper around g_source_unref().
func (source Source) Unref() {
	C.g_source_unref(source.native())
}

// SourceRemove is a wrapper around g_source_remove().
func SourceRemove(tag uint) bool {
	ret0 := C.g_source_remove(C.guint(tag))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SourceSetNameById is a wrapper around g_source_set_name_by_id().
func SourceSetNameById(tag uint, name string) {
	name0 := C.CString(name)
	C.g_source_set_name_by_id(C.guint(tag), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// IdleSourceNew is a wrapper around g_idle_source_new().
func IdleSourceNew() Source {
	ret0 := C.g_idle_source_new()
	return wrapSource(ret0)
}

// TimeoutSourceNew is a wrapper around g_timeout_source_new().
func TimeoutSourceNew(interval uint) Source {
	ret0 := C.g_timeout_source_new(C.guint(interval))
	return wrapSource(ret0)
}

// TimeoutSourceNewSeconds is a wrapper around g_timeout_source_new_seconds().
func TimeoutSourceNewSeconds(interval uint) Source {
	ret0 := C.g_timeout_source_new_seconds(C.guint(interval))
	return wrapSource(ret0)
}

type BookmarkFileError int

const (
	BookmarkFileErrorInvalidUri       BookmarkFileError = 0
	BookmarkFileErrorInvalidValue                       = 1
	BookmarkFileErrorAppNotRegistered                   = 2
	BookmarkFileErrorUriNotFound                        = 3
	BookmarkFileErrorRead                               = 4
	BookmarkFileErrorUnknownEncoding                    = 5
	BookmarkFileErrorWrite                              = 6
	BookmarkFileErrorFileNotFound                       = 7
)

type ChecksumType int

const (
	ChecksumTypeMd5    ChecksumType = 0
	ChecksumTypeSha1                = 1
	ChecksumTypeSha256              = 2
	ChecksumTypeSha512              = 3
)

type ConvertError int

const (
	ConvertErrorNoConversion    ConvertError = 0
	ConvertErrorIllegalSequence              = 1
	ConvertErrorFailed                       = 2
	ConvertErrorPartialInput                 = 3
	ConvertErrorBadUri                       = 4
	ConvertErrorNotAbsolutePath              = 5
	ConvertErrorNoMemory                     = 6
)

type DateDMY int

const (
	DateDMYDay   DateDMY = 0
	DateDMYMonth         = 1
	DateDMYYear          = 2
)

type DateMonth int

const (
	DateMonthBadMonth  DateMonth = 0
	DateMonthJanuary             = 1
	DateMonthFebruary            = 2
	DateMonthMarch               = 3
	DateMonthApril               = 4
	DateMonthMay                 = 5
	DateMonthJune                = 6
	DateMonthJuly                = 7
	DateMonthAugust              = 8
	DateMonthSeptember           = 9
	DateMonthOctober             = 10
	DateMonthNovember            = 11
	DateMonthDecember            = 12
)

type DateWeekday int

const (
	DateWeekdayBadWeekday DateWeekday = 0
	DateWeekdayMonday                 = 1
	DateWeekdayTuesday                = 2
	DateWeekdayWednesday              = 3
	DateWeekdayThursday               = 4
	DateWeekdayFriday                 = 5
	DateWeekdaySaturday               = 6
	DateWeekdaySunday                 = 7
)

type ErrorType int

const (
	ErrorTypeUnknown           ErrorType = 0
	ErrorTypeUnexpEof                    = 1
	ErrorTypeUnexpEofInString            = 2
	ErrorTypeUnexpEofInComment           = 3
	ErrorTypeNonDigitInConst             = 4
	ErrorTypeDigitRadix                  = 5
	ErrorTypeFloatRadix                  = 6
	ErrorTypeFloatMalformed              = 7
)

type FileError int

const (
	FileErrorExist       FileError = 0
	FileErrorIsdir                 = 1
	FileErrorAcces                 = 2
	FileErrorNametoolong           = 3
	FileErrorNoent                 = 4
	FileErrorNotdir                = 5
	FileErrorNxio                  = 6
	FileErrorNodev                 = 7
	FileErrorRofs                  = 8
	FileErrorTxtbsy                = 9
	FileErrorFault                 = 10
	FileErrorLoop                  = 11
	FileErrorNospc                 = 12
	FileErrorNomem                 = 13
	FileErrorMfile                 = 14
	FileErrorNfile                 = 15
	FileErrorBadf                  = 16
	FileErrorInval                 = 17
	FileErrorPipe                  = 18
	FileErrorAgain                 = 19
	FileErrorIntr                  = 20
	FileErrorIo                    = 21
	FileErrorPerm                  = 22
	FileErrorNosys                 = 23
	FileErrorFailed                = 24
)

type IOChannelError int

const (
	IOChannelErrorFbig     IOChannelError = 0
	IOChannelErrorInval                   = 1
	IOChannelErrorIo                      = 2
	IOChannelErrorIsdir                   = 3
	IOChannelErrorNospc                   = 4
	IOChannelErrorNxio                    = 5
	IOChannelErrorOverflow                = 6
	IOChannelErrorPipe                    = 7
	IOChannelErrorFailed                  = 8
)

type IOError int

const (
	IOErrorNone    IOError = 0
	IOErrorAgain           = 1
	IOErrorInval           = 2
	IOErrorUnknown         = 3
)

type IOStatus int

const (
	IOStatusError  IOStatus = 0
	IOStatusNormal          = 1
	IOStatusEof             = 2
	IOStatusAgain           = 3
)

type KeyFileError int

const (
	KeyFileErrorUnknownEncoding KeyFileError = 0
	KeyFileErrorParse                        = 1
	KeyFileErrorNotFound                     = 2
	KeyFileErrorKeyNotFound                  = 3
	KeyFileErrorGroupNotFound                = 4
	KeyFileErrorInvalidValue                 = 5
)

type LogWriterOutput int

const (
	LogWriterOutputHandled   LogWriterOutput = 1
	LogWriterOutputUnhandled                 = 0
)

type MarkupError int

const (
	MarkupErrorBadUtf8          MarkupError = 0
	MarkupErrorEmpty                        = 1
	MarkupErrorParse                        = 2
	MarkupErrorUnknownElement               = 3
	MarkupErrorUnknownAttribute             = 4
	MarkupErrorInvalidContent               = 5
	MarkupErrorMissingAttribute             = 6
)

type NormalizeMode int

const (
	NormalizeModeDefault        NormalizeMode = 0
	NormalizeModeNfd                          = 0
	NormalizeModeDefaultCompose               = 1
	NormalizeModeNfc                          = 1
	NormalizeModeAll                          = 2
	NormalizeModeNfkd                         = 2
	NormalizeModeAllCompose                   = 3
	NormalizeModeNfkc                         = 3
)

type OnceStatus int

const (
	OnceStatusNotcalled OnceStatus = 0
	OnceStatusProgress             = 1
	OnceStatusReady                = 2
)

type OptionArg int

const (
	OptionArgNone          OptionArg = 0
	OptionArgString                  = 1
	OptionArgInt                     = 2
	OptionArgCallback                = 3
	OptionArgFilename                = 4
	OptionArgStringArray             = 5
	OptionArgFilenameArray           = 6
	OptionArgDouble                  = 7
	OptionArgInt64                   = 8
)

type OptionError int

const (
	OptionErrorUnknownOption OptionError = 0
	OptionErrorBadValue                  = 1
	OptionErrorFailed                    = 2
)

type RegexError int

const (
	RegexErrorCompile                                  RegexError = 0
	RegexErrorOptimize                                            = 1
	RegexErrorReplace                                             = 2
	RegexErrorMatch                                               = 3
	RegexErrorInternal                                            = 4
	RegexErrorStrayBackslash                                      = 101
	RegexErrorMissingControlChar                                  = 102
	RegexErrorUnrecognizedEscape                                  = 103
	RegexErrorQuantifiersOutOfOrder                               = 104
	RegexErrorQuantifierTooBig                                    = 105
	RegexErrorUnterminatedCharacterClass                          = 106
	RegexErrorInvalidEscapeInCharacterClass                       = 107
	RegexErrorRangeOutOfOrder                                     = 108
	RegexErrorNothingToRepeat                                     = 109
	RegexErrorUnrecognizedCharacter                               = 112
	RegexErrorPosixNamedClassOutsideClass                         = 113
	RegexErrorUnmatchedParenthesis                                = 114
	RegexErrorInexistentSubpatternReference                       = 115
	RegexErrorUnterminatedComment                                 = 118
	RegexErrorExpressionTooLarge                                  = 120
	RegexErrorMemoryError                                         = 121
	RegexErrorVariableLengthLookbehind                            = 125
	RegexErrorMalformedCondition                                  = 126
	RegexErrorTooManyConditionalBranches                          = 127
	RegexErrorAssertionExpected                                   = 128
	RegexErrorUnknownPosixClassName                               = 130
	RegexErrorPosixCollatingElementsNotSupported                  = 131
	RegexErrorHexCodeTooLarge                                     = 134
	RegexErrorInvalidCondition                                    = 135
	RegexErrorSingleByteMatchInLookbehind                         = 136
	RegexErrorInfiniteLoop                                        = 140
	RegexErrorMissingSubpatternNameTerminator                     = 142
	RegexErrorDuplicateSubpatternName                             = 143
	RegexErrorMalformedProperty                                   = 146
	RegexErrorUnknownProperty                                     = 147
	RegexErrorSubpatternNameTooLong                               = 148
	RegexErrorTooManySubpatterns                                  = 149
	RegexErrorInvalidOctalValue                                   = 151
	RegexErrorTooManyBranchesInDefine                             = 154
	RegexErrorDefineRepetion                                      = 155
	RegexErrorInconsistentNewlineOptions                          = 156
	RegexErrorMissingBackReference                                = 157
	RegexErrorInvalidRelativeReference                            = 158
	RegexErrorBacktrackingControlVerbArgumentForbidden            = 159
	RegexErrorUnknownBacktrackingControlVerb                      = 160
	RegexErrorNumberTooBig                                        = 161
	RegexErrorMissingSubpatternName                               = 162
	RegexErrorMissingDigit                                        = 163
	RegexErrorInvalidDataCharacter                                = 164
	RegexErrorExtraSubpatternName                                 = 165
	RegexErrorBacktrackingControlVerbArgumentRequired             = 166
	RegexErrorInvalidControlChar                                  = 168
	RegexErrorMissingName                                         = 169
	RegexErrorNotSupportedInClass                                 = 171
	RegexErrorTooManyForwardReferences                            = 172
	RegexErrorNameTooLong                                         = 175
	RegexErrorCharacterValueTooLarge                              = 176
)

type SeekType int

const (
	SeekTypeCur SeekType = 0
	SeekTypeSet          = 1
	SeekTypeEnd          = 2
)

type ShellError int

const (
	ShellErrorBadQuoting  ShellError = 0
	ShellErrorEmptyString            = 1
	ShellErrorFailed                 = 2
)

type SliceConfig int

const (
	SliceConfigAlwaysMalloc      SliceConfig = 1
	SliceConfigBypassMagazines               = 2
	SliceConfigWorkingSetMsecs               = 3
	SliceConfigColorIncrement                = 4
	SliceConfigChunkSizes                    = 5
	SliceConfigContentionCounter             = 6
)

type SpawnError int

const (
	SpawnErrorFork        SpawnError = 0
	SpawnErrorRead                   = 1
	SpawnErrorChdir                  = 2
	SpawnErrorAcces                  = 3
	SpawnErrorPerm                   = 4
	SpawnErrorTooBig                 = 5
	SpawnError2big                   = 5
	SpawnErrorNoexec                 = 6
	SpawnErrorNametoolong            = 7
	SpawnErrorNoent                  = 8
	SpawnErrorNomem                  = 9
	SpawnErrorNotdir                 = 10
	SpawnErrorLoop                   = 11
	SpawnErrorTxtbusy                = 12
	SpawnErrorIo                     = 13
	SpawnErrorNfile                  = 14
	SpawnErrorMfile                  = 15
	SpawnErrorInval                  = 16
	SpawnErrorIsdir                  = 17
	SpawnErrorLibbad                 = 18
	SpawnErrorFailed                 = 19
)

type TestFileType int

const (
	TestFileTypeDist  TestFileType = 0
	TestFileTypeBuilt              = 1
)

type TestLogType int

const (
	TestLogTypeNone        TestLogType = 0
	TestLogTypeError                   = 1
	TestLogTypeStartBinary             = 2
	TestLogTypeListCase                = 3
	TestLogTypeSkipCase                = 4
	TestLogTypeStartCase               = 5
	TestLogTypeStopCase                = 6
	TestLogTypeMinResult               = 7
	TestLogTypeMaxResult               = 8
	TestLogTypeMessage                 = 9
	TestLogTypeStartSuite              = 10
	TestLogTypeStopSuite               = 11
)

type ThreadError int

const (
	ThreadErrorThreadErrorAgain ThreadError = 0
)

type TimeType int

const (
	TimeTypeStandard  TimeType = 0
	TimeTypeDaylight           = 1
	TimeTypeUniversal          = 2
)

type TokenType int

const (
	TokenTypeEof            TokenType = 0
	TokenTypeLeftParen                = 40
	TokenTypeRightParen               = 41
	TokenTypeLeftCurly                = 123
	TokenTypeRightCurly               = 125
	TokenTypeLeftBrace                = 91
	TokenTypeRightBrace               = 93
	TokenTypeEqualSign                = 61
	TokenTypeComma                    = 44
	TokenTypeNone                     = 256
	TokenTypeError                    = 257
	TokenTypeChar                     = 258
	TokenTypeBinary                   = 259
	TokenTypeOctal                    = 260
	TokenTypeInt                      = 261
	TokenTypeHex                      = 262
	TokenTypeFloat                    = 263
	TokenTypeString                   = 264
	TokenTypeSymbol                   = 265
	TokenTypeIdentifier               = 266
	TokenTypeIdentifierNull           = 267
	TokenTypeCommentSingle            = 268
	TokenTypeCommentMulti             = 269
)

type TraverseType int

const (
	TraverseTypeInOrder    TraverseType = 0
	TraverseTypePreOrder                = 1
	TraverseTypePostOrder               = 2
	TraverseTypeLevelOrder              = 3
)

type UnicodeBreakType int

const (
	UnicodeBreakTypeMandatory                  UnicodeBreakType = 0
	UnicodeBreakTypeCarriageReturn                              = 1
	UnicodeBreakTypeLineFeed                                    = 2
	UnicodeBreakTypeCombiningMark                               = 3
	UnicodeBreakTypeSurrogate                                   = 4
	UnicodeBreakTypeZeroWidthSpace                              = 5
	UnicodeBreakTypeInseparable                                 = 6
	UnicodeBreakTypeNonBreakingGlue                             = 7
	UnicodeBreakTypeContingent                                  = 8
	UnicodeBreakTypeSpace                                       = 9
	UnicodeBreakTypeAfter                                       = 10
	UnicodeBreakTypeBefore                                      = 11
	UnicodeBreakTypeBeforeAndAfter                              = 12
	UnicodeBreakTypeHyphen                                      = 13
	UnicodeBreakTypeNonStarter                                  = 14
	UnicodeBreakTypeOpenPunctuation                             = 15
	UnicodeBreakTypeClosePunctuation                            = 16
	UnicodeBreakTypeQuotation                                   = 17
	UnicodeBreakTypeExclamation                                 = 18
	UnicodeBreakTypeIdeographic                                 = 19
	UnicodeBreakTypeNumeric                                     = 20
	UnicodeBreakTypeInfixSeparator                              = 21
	UnicodeBreakTypeSymbol                                      = 22
	UnicodeBreakTypeAlphabetic                                  = 23
	UnicodeBreakTypePrefix                                      = 24
	UnicodeBreakTypePostfix                                     = 25
	UnicodeBreakTypeComplexContext                              = 26
	UnicodeBreakTypeAmbiguous                                   = 27
	UnicodeBreakTypeUnknown                                     = 28
	UnicodeBreakTypeNextLine                                    = 29
	UnicodeBreakTypeWordJoiner                                  = 30
	UnicodeBreakTypeHangulLJamo                                 = 31
	UnicodeBreakTypeHangulVJamo                                 = 32
	UnicodeBreakTypeHangulTJamo                                 = 33
	UnicodeBreakTypeHangulLvSyllable                            = 34
	UnicodeBreakTypeHangulLvtSyllable                           = 35
	UnicodeBreakTypeCloseParanthesis                            = 36
	UnicodeBreakTypeConditionalJapaneseStarter                  = 37
	UnicodeBreakTypeHebrewLetter                                = 38
	UnicodeBreakTypeRegionalIndicator                           = 39
	UnicodeBreakTypeEmojiBase                                   = 40
	UnicodeBreakTypeEmojiModifier                               = 41
	UnicodeBreakTypeZeroWidthJoiner                             = 42
)

type UnicodeScript int

const (
	UnicodeScriptInvalidCode           UnicodeScript = -1
	UnicodeScriptCommon                              = 0
	UnicodeScriptInherited                           = 1
	UnicodeScriptArabic                              = 2
	UnicodeScriptArmenian                            = 3
	UnicodeScriptBengali                             = 4
	UnicodeScriptBopomofo                            = 5
	UnicodeScriptCherokee                            = 6
	UnicodeScriptCoptic                              = 7
	UnicodeScriptCyrillic                            = 8
	UnicodeScriptDeseret                             = 9
	UnicodeScriptDevanagari                          = 10
	UnicodeScriptEthiopic                            = 11
	UnicodeScriptGeorgian                            = 12
	UnicodeScriptGothic                              = 13
	UnicodeScriptGreek                               = 14
	UnicodeScriptGujarati                            = 15
	UnicodeScriptGurmukhi                            = 16
	UnicodeScriptHan                                 = 17
	UnicodeScriptHangul                              = 18
	UnicodeScriptHebrew                              = 19
	UnicodeScriptHiragana                            = 20
	UnicodeScriptKannada                             = 21
	UnicodeScriptKatakana                            = 22
	UnicodeScriptKhmer                               = 23
	UnicodeScriptLao                                 = 24
	UnicodeScriptLatin                               = 25
	UnicodeScriptMalayalam                           = 26
	UnicodeScriptMongolian                           = 27
	UnicodeScriptMyanmar                             = 28
	UnicodeScriptOgham                               = 29
	UnicodeScriptOldItalic                           = 30
	UnicodeScriptOriya                               = 31
	UnicodeScriptRunic                               = 32
	UnicodeScriptSinhala                             = 33
	UnicodeScriptSyriac                              = 34
	UnicodeScriptTamil                               = 35
	UnicodeScriptTelugu                              = 36
	UnicodeScriptThaana                              = 37
	UnicodeScriptThai                                = 38
	UnicodeScriptTibetan                             = 39
	UnicodeScriptCanadianAboriginal                  = 40
	UnicodeScriptYi                                  = 41
	UnicodeScriptTagalog                             = 42
	UnicodeScriptHanunoo                             = 43
	UnicodeScriptBuhid                               = 44
	UnicodeScriptTagbanwa                            = 45
	UnicodeScriptBraille                             = 46
	UnicodeScriptCypriot                             = 47
	UnicodeScriptLimbu                               = 48
	UnicodeScriptOsmanya                             = 49
	UnicodeScriptShavian                             = 50
	UnicodeScriptLinearB                             = 51
	UnicodeScriptTaiLe                               = 52
	UnicodeScriptUgaritic                            = 53
	UnicodeScriptNewTaiLue                           = 54
	UnicodeScriptBuginese                            = 55
	UnicodeScriptGlagolitic                          = 56
	UnicodeScriptTifinagh                            = 57
	UnicodeScriptSylotiNagri                         = 58
	UnicodeScriptOldPersian                          = 59
	UnicodeScriptKharoshthi                          = 60
	UnicodeScriptUnknown                             = 61
	UnicodeScriptBalinese                            = 62
	UnicodeScriptCuneiform                           = 63
	UnicodeScriptPhoenician                          = 64
	UnicodeScriptPhagsPa                             = 65
	UnicodeScriptNko                                 = 66
	UnicodeScriptKayahLi                             = 67
	UnicodeScriptLepcha                              = 68
	UnicodeScriptRejang                              = 69
	UnicodeScriptSundanese                           = 70
	UnicodeScriptSaurashtra                          = 71
	UnicodeScriptCham                                = 72
	UnicodeScriptOlChiki                             = 73
	UnicodeScriptVai                                 = 74
	UnicodeScriptCarian                              = 75
	UnicodeScriptLycian                              = 76
	UnicodeScriptLydian                              = 77
	UnicodeScriptAvestan                             = 78
	UnicodeScriptBamum                               = 79
	UnicodeScriptEgyptianHieroglyphs                 = 80
	UnicodeScriptImperialAramaic                     = 81
	UnicodeScriptInscriptionalPahlavi                = 82
	UnicodeScriptInscriptionalParthian               = 83
	UnicodeScriptJavanese                            = 84
	UnicodeScriptKaithi                              = 85
	UnicodeScriptLisu                                = 86
	UnicodeScriptMeeteiMayek                         = 87
	UnicodeScriptOldSouthArabian                     = 88
	UnicodeScriptOldTurkic                           = 89
	UnicodeScriptSamaritan                           = 90
	UnicodeScriptTaiTham                             = 91
	UnicodeScriptTaiViet                             = 92
	UnicodeScriptBatak                               = 93
	UnicodeScriptBrahmi                              = 94
	UnicodeScriptMandaic                             = 95
	UnicodeScriptChakma                              = 96
	UnicodeScriptMeroiticCursive                     = 97
	UnicodeScriptMeroiticHieroglyphs                 = 98
	UnicodeScriptMiao                                = 99
	UnicodeScriptSharada                             = 100
	UnicodeScriptSoraSompeng                         = 101
	UnicodeScriptTakri                               = 102
	UnicodeScriptBassaVah                            = 103
	UnicodeScriptCaucasianAlbanian                   = 104
	UnicodeScriptDuployan                            = 105
	UnicodeScriptElbasan                             = 106
	UnicodeScriptGrantha                             = 107
	UnicodeScriptKhojki                              = 108
	UnicodeScriptKhudawadi                           = 109
	UnicodeScriptLinearA                             = 110
	UnicodeScriptMahajani                            = 111
	UnicodeScriptManichaean                          = 112
	UnicodeScriptMendeKikakui                        = 113
	UnicodeScriptModi                                = 114
	UnicodeScriptMro                                 = 115
	UnicodeScriptNabataean                           = 116
	UnicodeScriptOldNorthArabian                     = 117
	UnicodeScriptOldPermic                           = 118
	UnicodeScriptPahawhHmong                         = 119
	UnicodeScriptPalmyrene                           = 120
	UnicodeScriptPauCinHau                           = 121
	UnicodeScriptPsalterPahlavi                      = 122
	UnicodeScriptSiddham                             = 123
	UnicodeScriptTirhuta                             = 124
	UnicodeScriptWarangCiti                          = 125
	UnicodeScriptAhom                                = 126
	UnicodeScriptAnatolianHieroglyphs                = 127
	UnicodeScriptHatran                              = 128
	UnicodeScriptMultani                             = 129
	UnicodeScriptOldHungarian                        = 130
	UnicodeScriptSignwriting                         = 131
	UnicodeScriptAdlam                               = 132
	UnicodeScriptBhaiksuki                           = 133
	UnicodeScriptMarchen                             = 134
	UnicodeScriptNewa                                = 135
	UnicodeScriptOsage                               = 136
	UnicodeScriptTangut                              = 137
)

type UnicodeType int

const (
	UnicodeTypeControl            UnicodeType = 0
	UnicodeTypeFormat                         = 1
	UnicodeTypeUnassigned                     = 2
	UnicodeTypePrivateUse                     = 3
	UnicodeTypeSurrogate                      = 4
	UnicodeTypeLowercaseLetter                = 5
	UnicodeTypeModifierLetter                 = 6
	UnicodeTypeOtherLetter                    = 7
	UnicodeTypeTitlecaseLetter                = 8
	UnicodeTypeUppercaseLetter                = 9
	UnicodeTypeSpacingMark                    = 10
	UnicodeTypeEnclosingMark                  = 11
	UnicodeTypeNonSpacingMark                 = 12
	UnicodeTypeDecimalNumber                  = 13
	UnicodeTypeLetterNumber                   = 14
	UnicodeTypeOtherNumber                    = 15
	UnicodeTypeConnectPunctuation             = 16
	UnicodeTypeDashPunctuation                = 17
	UnicodeTypeClosePunctuation               = 18
	UnicodeTypeFinalPunctuation               = 19
	UnicodeTypeInitialPunctuation             = 20
	UnicodeTypeOtherPunctuation               = 21
	UnicodeTypeOpenPunctuation                = 22
	UnicodeTypeCurrencySymbol                 = 23
	UnicodeTypeModifierSymbol                 = 24
	UnicodeTypeMathSymbol                     = 25
	UnicodeTypeOtherSymbol                    = 26
	UnicodeTypeLineSeparator                  = 27
	UnicodeTypeParagraphSeparator             = 28
	UnicodeTypeSpaceSeparator                 = 29
)

type UserDirectory int

const (
	UserDirectoryDirectoryDesktop     UserDirectory = 0
	UserDirectoryDirectoryDocuments                 = 1
	UserDirectoryDirectoryDownload                  = 2
	UserDirectoryDirectoryMusic                     = 3
	UserDirectoryDirectoryPictures                  = 4
	UserDirectoryDirectoryPublicShare               = 5
	UserDirectoryDirectoryTemplates                 = 6
	UserDirectoryDirectoryVideos                    = 7
	UserDirectoryNDirectories                       = 8
)

type VariantClass int

const (
	VariantClassBoolean    VariantClass = 98
	VariantClassByte                    = 121
	VariantClassInt16                   = 110
	VariantClassUint16                  = 113
	VariantClassInt32                   = 105
	VariantClassUint32                  = 117
	VariantClassInt64                   = 120
	VariantClassUint64                  = 116
	VariantClassHandle                  = 104
	VariantClassDouble                  = 100
	VariantClassString                  = 115
	VariantClassObjectPath              = 111
	VariantClassSignature               = 103
	VariantClassVariant                 = 118
	VariantClassMaybe                   = 109
	VariantClassArray                   = 97
	VariantClassTuple                   = 40
	VariantClassDictEntry               = 123
)

type VariantParseError int

const (
	VariantParseErrorFailed                     VariantParseError = 0
	VariantParseErrorBasicTypeExpected                            = 1
	VariantParseErrorCannotInferType                              = 2
	VariantParseErrorDefiniteTypeExpected                         = 3
	VariantParseErrorInputNotAtEnd                                = 4
	VariantParseErrorInvalidCharacter                             = 5
	VariantParseErrorInvalidFormatString                          = 6
	VariantParseErrorInvalidObjectPath                            = 7
	VariantParseErrorInvalidSignature                             = 8
	VariantParseErrorInvalidTypeString                            = 9
	VariantParseErrorNoCommonType                                 = 10
	VariantParseErrorNumberOutOfRange                             = 11
	VariantParseErrorNumberTooBig                                 = 12
	VariantParseErrorTypeError                                    = 13
	VariantParseErrorUnexpectedToken                              = 14
	VariantParseErrorUnknownKeyword                               = 15
	VariantParseErrorUnterminatedStringConstant                   = 16
	VariantParseErrorValueExpected                                = 17
)

type AsciiType int

const (
	AsciiTypeAlnum  AsciiType = 1
	AsciiTypeAlpha            = 2
	AsciiTypeCntrl            = 4
	AsciiTypeDigit            = 8
	AsciiTypeGraph            = 16
	AsciiTypeLower            = 32
	AsciiTypePrint            = 64
	AsciiTypePunct            = 128
	AsciiTypeSpace            = 256
	AsciiTypeUpper            = 512
	AsciiTypeXdigit           = 1024
)

type FileTest int

const (
	FileTestIsRegular    FileTest = 1
	FileTestIsSymlink             = 2
	FileTestIsDir                 = 4
	FileTestIsExecutable          = 8
	FileTestExists                = 16
)

type FormatSizeFlags int

const (
	FormatSizeFlagsDefault    FormatSizeFlags = 0
	FormatSizeFlagsLongFormat                 = 1
	FormatSizeFlagsIecUnits                   = 2
)

type HookFlagMask int

const (
	HookFlagMaskActive HookFlagMask = 1
	HookFlagMaskInCall              = 2
	HookFlagMaskMask                = 15
)

type IOCondition int

const (
	IOConditionIn   IOCondition = 1
	IOConditionOut              = 4
	IOConditionPri              = 2
	IOConditionErr              = 8
	IOConditionHup              = 16
	IOConditionNval             = 32
)

type IOFlags int

const (
	IOFlagsAppend      IOFlags = 1
	IOFlagsNonblock            = 2
	IOFlagsIsReadable          = 4
	IOFlagsIsWritable          = 8
	IOFlagsIsWriteable         = 8
	IOFlagsIsSeekable          = 16
	IOFlagsMask                = 31
	IOFlagsGetMask             = 31
	IOFlagsSetMask             = 3
)

type KeyFileFlags int

const (
	KeyFileFlagsNone             KeyFileFlags = 0
	KeyFileFlagsKeepComments                  = 1
	KeyFileFlagsKeepTranslations              = 2
)

type LogLevelFlags int

const (
	LogLevelFlagsFlagRecursion LogLevelFlags = 1
	LogLevelFlagsFlagFatal                   = 2
	LogLevelFlagsLevelError                  = 4
	LogLevelFlagsLevelCritical               = 8
	LogLevelFlagsLevelWarning                = 16
	LogLevelFlagsLevelMessage                = 32
	LogLevelFlagsLevelInfo                   = 64
	LogLevelFlagsLevelDebug                  = 128
	LogLevelFlagsLevelMask                   = -4
)

type MarkupCollectType int

const (
	MarkupCollectTypeInvalid  MarkupCollectType = 0
	MarkupCollectTypeString                     = 1
	MarkupCollectTypeStrdup                     = 2
	MarkupCollectTypeBoolean                    = 3
	MarkupCollectTypeTristate                   = 4
	MarkupCollectTypeOptional                   = 65536
)

type MarkupParseFlags int

const (
	MarkupParseFlagsDoNotUseThisUnsupportedFlag MarkupParseFlags = 1
	MarkupParseFlagsTreatCdataAsText                             = 2
	MarkupParseFlagsPrefixErrorPosition                          = 4
	MarkupParseFlagsIgnoreQualified                              = 8
)

type OptionFlags int

const (
	OptionFlagsNone        OptionFlags = 0
	OptionFlagsHidden                  = 1
	OptionFlagsInMain                  = 2
	OptionFlagsReverse                 = 4
	OptionFlagsNoArg                   = 8
	OptionFlagsFilename                = 16
	OptionFlagsOptionalArg             = 32
	OptionFlagsNoalias                 = 64
)

type RegexCompileFlags int

const (
	RegexCompileFlagsCaseless         RegexCompileFlags = 1
	RegexCompileFlagsMultiline                          = 2
	RegexCompileFlagsDotall                             = 4
	RegexCompileFlagsExtended                           = 8
	RegexCompileFlagsAnchored                           = 16
	RegexCompileFlagsDollarEndonly                      = 32
	RegexCompileFlagsUngreedy                           = 512
	RegexCompileFlagsRaw                                = 2048
	RegexCompileFlagsNoAutoCapture                      = 4096
	RegexCompileFlagsOptimize                           = 8192
	RegexCompileFlagsFirstline                          = 262144
	RegexCompileFlagsDupnames                           = 524288
	RegexCompileFlagsNewlineCr                          = 1048576
	RegexCompileFlagsNewlineLf                          = 2097152
	RegexCompileFlagsNewlineCrlf                        = 3145728
	RegexCompileFlagsNewlineAnycrlf                     = 5242880
	RegexCompileFlagsBsrAnycrlf                         = 8388608
	RegexCompileFlagsJavascriptCompat                   = 33554432
)

type RegexMatchFlags int

const (
	RegexMatchFlagsAnchored        RegexMatchFlags = 16
	RegexMatchFlagsNotbol                          = 128
	RegexMatchFlagsNoteol                          = 256
	RegexMatchFlagsNotempty                        = 1024
	RegexMatchFlagsPartial                         = 32768
	RegexMatchFlagsNewlineCr                       = 1048576
	RegexMatchFlagsNewlineLf                       = 2097152
	RegexMatchFlagsNewlineCrlf                     = 3145728
	RegexMatchFlagsNewlineAny                      = 4194304
	RegexMatchFlagsNewlineAnycrlf                  = 5242880
	RegexMatchFlagsBsrAnycrlf                      = 8388608
	RegexMatchFlagsBsrAny                          = 16777216
	RegexMatchFlagsPartialSoft                     = 32768
	RegexMatchFlagsPartialHard                     = 134217728
	RegexMatchFlagsNotemptyAtstart                 = 268435456
)

type SpawnFlags int

const (
	SpawnFlagsDefault              SpawnFlags = 0
	SpawnFlagsLeaveDescriptorsOpen            = 1
	SpawnFlagsDoNotReapChild                  = 2
	SpawnFlagsSearchPath                      = 4
	SpawnFlagsStdoutToDevNull                 = 8
	SpawnFlagsStderrToDevNull                 = 16
	SpawnFlagsChildInheritsStdin              = 32
	SpawnFlagsFileAndArgvZero                 = 64
	SpawnFlagsSearchPathFromEnvp              = 128
	SpawnFlagsCloexecPipes                    = 256
)

type TestSubprocessFlags int

const (
	TestSubprocessFlagsStdin  TestSubprocessFlags = 1
	TestSubprocessFlagsStdout                     = 2
	TestSubprocessFlagsStderr                     = 4
)

type TestTrapFlags int

const (
	TestTrapFlagsSilenceStdout TestTrapFlags = 128
	TestTrapFlagsSilenceStderr               = 256
	TestTrapFlagsInheritStdin                = 512
)

type TraverseFlags int

const (
	TraverseFlagsLeaves    TraverseFlags = 1
	TraverseFlagsNonLeaves               = 2
	TraverseFlagsAll                     = 3
	TraverseFlagsMask                    = 3
	TraverseFlagsLeafs                   = 1
	TraverseFlagsNonLeafs                = 2
)
