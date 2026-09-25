# Package Study Reference

## Package Information

use it for? 3. Exported package-level functions. Each function is shown
on its own line for easy reading.

## Priority Levels

important / Strong working knowledge 🟡 YELLOW = Important / Working
knowledge 🔵 BLUE = Useful / Learn as needed 🟢 GREEN = Specialized /
Learn when required

## Notes

standard-library documentation available in the installed Go
toolchain. - “Functions” means exported package-level functions
(including constructors). - Methods belonging to exported types are not
listed as package-level functions. - Some packages primarily expose
types, constants, variables, or interfaces and therefore may have few or
no package-level functions. - A plain .txt file cannot reliably display
actual font colors, so colored-square priority labels are used.

============================================================ 🔴 RED

### 1. `fmt`

What is this package? scanf. The format ‘verbs’ are derived from C’s but
are simpler.

What purpose do we use this package for? scanf. The format ‘verbs’ are
derived from C’s but are simpler.

Functions: - func Append(b []byte, a …any) []byte - func Appendf(b
[]byte, format string, a …any) []byte - func Appendln(b []byte, a …any)
[]byte - func Errorf(format string, a …any) error - func
FormatString(state State, verb rune) string - func Fprint(w io.Writer, a
…any) (n int, err error) - func Fprintf(w io.Writer, format string, a
…any) (n int, err error) - func Fprintln(w io.Writer, a …any) (n int,
err error) - func Fscan(r io.Reader, a …any) (n int, err error) - func
Fscanf(r io.Reader, format string, a …any) (n int, err error) - func
Fscanln(r io.Reader, a …any) (n int, err error) - func Print(a …any) (n
int, err error) - func Printf(format string, a …any) (n int, err
error) - func Println(a …any) (n int, err error) - func Scan(a …any) (n
int, err error) - func Scanf(format string, a …any) (n int, err error) -
func Scanln(a …any) (n int, err error) - func Sprint(a …any) string -
func Sprintf(format string, a …any) string - func Sprintln(a …any)
string - func Sscan(str string, a …any) (n int, err error) - func
Sscanf(str string, format string, a …any) (n int, err error) - func
Sscanln(str string, a …any) (n int, err error)

### 2. `os`

What is this package? functionality. The design is Unix-like, although
the error handling is Go-like; failing calls return values of type error
rather than error numbers. Often, more information is available within
the error. For example, if a call that takes a file name fails, such as
Open or Stat, the error will include the failing file name when printed
and will be of type *PathError, which may be unpacked for more
information.

What purpose do we use this package for? functionality. The design is
Unix-like, although the error handling is Go-like; failing calls return
values of type error rather than error numbers. Often, more information
is available within the error. For example, if a call that takes a file
name fails, such as Open or Stat, the error will include the failing
file name when printed and will be of type *PathError, which may be
unpacked for more information.

Functions: - func Chdir(dir string) error - func Chmod(name string, mode
FileMode) error - func Chown(name string, uid, gid int) error - func
Chtimes(name string, atime time.Time, mtime time.Time) error - func
Clearenv() - func CopyFS(dir string, fsys fs.FS) error - func DirFS(dir
string) fs.FS - func Environ() []string - func Executable() (string,
error) - func Exit(code int) - func Expand(s string, mapping
func(string) string) string - func ExpandEnv(s string) string - func
Getegid() int - func Getenv(key string) string - func Geteuid() int -
func Getgid() int - func Getgroups() ([]int, error) - func Getpagesize()
int - func Getpid() int - func Getppid() int - func Getuid() int - func
Getwd() (dir string, err error) - func Hostname() (name string, err
error) - func IsExist(err error) bool - func IsNotExist(err error)
bool - func IsPathSeparator(c uint8) bool - func IsPermission(err error)
bool - func IsTimeout(err error) bool - func Lchown(name string, uid,
gid int) error - func Link(oldname, newname string) error - func
LookupEnv(key string) (string, bool) - func Mkdir(name string, perm
FileMode) error - func MkdirAll(path string, perm FileMode) error - func
MkdirTemp(dir, pattern string) (string, error) - func
NewSyscallError(syscall string, err error) error - func Pipe() (r File,
w File, err error) - func ReadFile(name string) ([]byte, error) - func
Readlink(name string) (string, error) - func Remove(name string) error -
func RemoveAll(path string) error - func Rename(oldpath, newpath string)
error - func SameFile(fi1, fi2 FileInfo) bool - func Setenv(key, value
string) error - func Symlink(oldname, newname string) error - func
TempDir() string - func Truncate(name string, size int64) error - func
Unsetenv(key string) error - func UserCacheDir() (string, error) - func
UserConfigDir() (string, error) - func UserHomeDir() (string, error) -
func WriteFile(name string, data []byte, perm FileMode) error

### 3. `io`

What is this package? wrap existing implementations of such primitives,
such as those in package os, into shared public interfaces that abstract
the functionality, plus some other related primitives.

What purpose do we use this package for? wrap existing implementations
of such primitives, such as those in package os, into shared public
interfaces that abstract the functionality, plus some other related
primitives.

Functions: - func Copy(dst Writer, src Reader) (written int64, err
error) - func CopyBuffer(dst Writer, src Reader, buf []byte) (written
int64, err error) - func CopyN(dst Writer, src Reader, n int64) (written
int64, err error) - func Pipe() (PipeReader, PipeWriter) - func
ReadAll(r Reader) ([]byte, error) - func ReadAtLeast(r Reader, buf
[]byte, min int) (n int, err error) - func ReadFull(r Reader, buf
[]byte) (n int, err error) - func WriteString(w Writer, s string) (n
int, err error)

### 4. `bufio`

What is this package? object, creating another object (Reader or Writer)
that also implements the interface but provides buffering and some help
for textual I/O.

What purpose do we use this package for? object, creating another object
(Reader or Writer) that also implements the interface but provides
buffering and some help for textual I/O.

Functions: - func ScanBytes(data []byte, atEOF bool) (advance int, token
[]byte, err error) - func ScanLines(data []byte, atEOF bool) (advance
int, token []byte, err error) - func ScanRunes(data []byte, atEOF bool)
(advance int, token []byte, err error) - func ScanWords(data []byte,
atEOF bool) (advance int, token []byte, err error)

### 5. `bytes`

What is this package? analogous to the facilities of the strings
package.

What purpose do we use this package for? analogous to the facilities of
the strings package.

Functions: - func Clone(b []byte) []byte - func Compare(a, b []byte)
int - func Contains(b, subslice []byte) bool - func ContainsAny(b
[]byte, chars string) bool - func ContainsFunc(b []byte, f func(rune)
bool) bool - func ContainsRune(b []byte, r rune) bool - func Count(s,
sep []byte) int - func Cut(s, sep []byte) (before, after []byte, found
bool) - func CutPrefix(s, prefix []byte) (after []byte, found bool) -
func CutSuffix(s, suffix []byte) (before []byte, found bool) - func
Equal(a, b []byte) bool - func EqualFold(s, t []byte) bool - func
Fields(s []byte) [][]byte - func FieldsFunc(s []byte, f func(rune) bool)
[][]byte - func HasPrefix(s, prefix []byte) bool - func HasSuffix(s,
suffix []byte) bool - func Index(s, sep []byte) int - func IndexAny(s
[]byte, chars string) int - func IndexByte(b []byte, c byte) int - func
IndexFunc(s []byte, f func(r rune) bool) int - func IndexRune(s []byte,
r rune) int - func Join(s [][]byte, sep []byte) []byte - func
LastIndex(s, sep []byte) int - func LastIndexAny(s []byte, chars string)
int - func LastIndexByte(s []byte, c byte) int - func LastIndexFunc(s
[]byte, f func(r rune) bool) int - func Map(mapping func(r rune) rune, s
[]byte) []byte - func Repeat(b []byte, count int) []byte - func
Replace(s, old, new []byte, n int) []byte - func ReplaceAll(s, old, new
[]byte) []byte - func Runes(s []byte) []rune - func Split(s, sep []byte)
[][]byte - func SplitAfter(s, sep []byte) [][]byte - func SplitAfterN(s,
sep []byte, n int) [][]byte - func SplitN(s, sep []byte, n int)
[][]byte - func Title(s []byte) []byte - func ToLower(s []byte) []byte -
func ToLowerSpecial(c unicode.SpecialCase, s []byte) []byte - func
ToTitle(s []byte) []byte - func ToTitleSpecial(c unicode.SpecialCase, s
[]byte) []byte - func ToUpper(s []byte) []byte - func ToUpperSpecial(c
unicode.SpecialCase, s []byte) []byte - func ToValidUTF8(s, replacement
[]byte) []byte - func Trim(s []byte, cutset string) []byte - func
TrimFunc(s []byte, f func(r rune) bool) []byte - func TrimLeft(s []byte,
cutset string) []byte - func TrimLeftFunc(s []byte, f func(r rune) bool)
[]byte - func TrimPrefix(s, prefix []byte) []byte - func TrimRight(s
[]byte, cutset string) []byte - func TrimRightFunc(s []byte, f func(r
rune) bool) []byte - func TrimSpace(s []byte) []byte - func
TrimSuffix(s, suffix []byte) []byte

### 6. `strings`

What is this package? For information about UTF-8 strings in Go, see
https://blog.golang.org/strings.

What purpose do we use this package for? For information about UTF-8
strings in Go, see https://blog.golang.org/strings.

Functions: - func Clone(s string) string - func Compare(a, b string)
int - func Contains(s, substr string) bool - func ContainsAny(s, chars
string) bool - func ContainsFunc(s string, f func(rune) bool) bool -
func ContainsRune(s string, r rune) bool - func Count(s, substr string)
int - func Cut(s, sep string) (before, after string, found bool) - func
CutPrefix(s, prefix string) (after string, found bool) - func
CutSuffix(s, suffix string) (before string, found bool) - func
EqualFold(s, t string) bool - func Fields(s string) []string - func
FieldsFunc(s string, f func(rune) bool) []string - func HasPrefix(s,
prefix string) bool - func HasSuffix(s, suffix string) bool - func
Index(s, substr string) int - func IndexAny(s, chars string) int - func
IndexByte(s string, c byte) int - func IndexFunc(s string, f func(rune)
bool) int - func IndexRune(s string, r rune) int - func Join(elems
[]string, sep string) string - func LastIndex(s, substr string) int -
func LastIndexAny(s, chars string) int - func LastIndexByte(s string, c
byte) int - func LastIndexFunc(s string, f func(rune) bool) int - func
Map(mapping func(rune) rune, s string) string - func Repeat(s string,
count int) string - func Replace(s, old, new string, n int) string -
func ReplaceAll(s, old, new string) string - func Split(s, sep string)
[]string - func SplitAfter(s, sep string) []string - func SplitAfterN(s,
sep string, n int) []string - func SplitN(s, sep string, n int)
[]string - func Title(s string) string - func ToLower(s string) string -
func ToLowerSpecial(c unicode.SpecialCase, s string) string - func
ToTitle(s string) string - func ToTitleSpecial(c unicode.SpecialCase, s
string) string - func ToUpper(s string) string - func ToUpperSpecial(c
unicode.SpecialCase, s string) string - func ToValidUTF8(s, replacement
string) string - func Trim(s, cutset string) string - func TrimFunc(s
string, f func(rune) bool) string - func TrimLeft(s, cutset string)
string - func TrimLeftFunc(s string, f func(rune) bool) string - func
TrimPrefix(s, prefix string) string - func TrimRight(s, cutset string)
string - func TrimRightFunc(s string, f func(rune) bool) string - func
TrimSpace(s string) string - func TrimSuffix(s, suffix string) string

### 7. `strconv`

What is this package? basic data types.

What purpose do we use this package for? basic data types.

Functions: - func AppendBool(dst []byte, b bool) []byte - func
AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte -
func AppendInt(dst []byte, i int64, base int) []byte - func
AppendQuote(dst []byte, s string) []byte - func AppendQuoteRune(dst
[]byte, r rune) []byte - func AppendQuoteRuneToASCII(dst []byte, r rune)
[]byte - func AppendQuoteRuneToGraphic(dst []byte, r rune) []byte - func
AppendQuoteToASCII(dst []byte, s string) []byte - func
AppendQuoteToGraphic(dst []byte, s string) []byte - func AppendUint(dst
[]byte, i uint64, base int) []byte - func Atoi(s string) (int, error) -
func CanBackquote(s string) bool - func FormatBool(b bool) string - func
FormatComplex(c complex128, fmt byte, prec, bitSize int) string - func
FormatFloat(f float64, fmt byte, prec, bitSize int) string - func
FormatInt(i int64, base int) string - func FormatUint(i uint64, base
int) string - func IsGraphic(r rune) bool - func IsPrint(r rune) bool -
func Itoa(i int) string - func ParseBool(str string) (bool, error) -
func ParseComplex(s string, bitSize int) (complex128, error) - func
ParseFloat(s string, bitSize int) (float64, error) - func ParseInt(s
string, base int, bitSize int) (i int64, err error) - func ParseUint(s
string, base int, bitSize int) (uint64, error) - func Quote(s string)
string - func QuoteRune(r rune) string - func QuoteRuneToASCII(r rune)
string - func QuoteRuneToGraphic(r rune) string - func QuoteToASCII(s
string) string - func QuoteToGraphic(s string) string - func
QuotedPrefix(s string) (string, error) - func Unquote(s string) (string,
error) - func UnquoteChar(s string, quote byte) (value rune, multibyte
bool, tail string, err error)

### 8. `errors`

What is this package? The New function creates errors whose only content
is a text message.

What purpose do we use this package for? The New function creates errors
whose only content is a text message.

Functions: - func As(err error, target any) bool - func Is(err, target
error) bool - func Join(errs …error) error - func New(text string)
error - func Unwrap(err error) error

### 9. `time`

What is this package? The calendrical calculations always assume a
Gregorian calendar, with no leap seconds.

What purpose do we use this package for? The calendrical calculations
always assume a Gregorian calendar, with no leap seconds.

Functions: - func After(d Duration) <-chan Time - func Sleep(d
Duration) - func Tick(d Duration) <-chan Time

### 10. `context`

What is this package? signals, and other request-scoped values across
API boundaries and between processes.

What purpose do we use this package for? signals, and other
request-scoped values across API boundaries and between processes.

Functions: - func AfterFunc(ctx Context, f func()) (stop func() bool) -
func Cause(c Context) error - func WithCancel(parent Context) (ctx
Context, cancel CancelFunc) - func WithCancelCause(parent Context) (ctx
Context, cancel CancelCauseFunc) - func WithDeadline(parent Context, d
time.Time) (Context, CancelFunc) - func WithDeadlineCause(parent
Context, d time.Time, cause error) (Context, CancelFunc) - func
WithTimeout(parent Context, timeout time.Duration) (Context,
CancelFunc) - func WithTimeoutCause(parent Context, timeout
time.Duration, cause error) (Context, CancelFunc)

### 11. `sync`

What is this package? locks. Other than the Once and WaitGroup types,
most are intended for use by low-level library routines. Higher-level
synchronization is better done via channels and communication.

What purpose do we use this package for? locks. Other than the Once and
WaitGroup types, most are intended for use by low-level library
routines. Higher-level synchronization is better done via channels and
communication.

Functions: - func OnceFunc(f func()) func() - func OnceValueT any func()
T - func OnceValuesT1, T2 any func() (T1, T2)

### 12. `sync/atomic`

What is this package? implementing synchronization algorithms.

What purpose do we use this package for? implementing synchronization
algorithms.

Functions: - func AddInt32(addr int32, delta int32) (new int32) - func
AddInt64(addr int64, delta int64) (new int64) - func AddUint32(addr
uint32, delta uint32) (new uint32) - func AddUint64(addr uint64, delta
uint64) (new uint64) - func AddUintptr(addr uintptr, delta uintptr) (new
uintptr) - func AndInt32(addr int32, mask int32) (old int32) - func
AndInt64(addr int64, mask int64) (old int64) - func AndUint32(addr
uint32, mask uint32) (old uint32) - func AndUint64(addr uint64, mask
uint64) (old uint64) - func AndUintptr(addr uintptr, mask uintptr) (old
uintptr) - func CompareAndSwapInt32(addr int32, old, new int32) (swapped
bool) - func CompareAndSwapInt64(addr int64, old, new int64) (swapped
bool) - func CompareAndSwapPointer(addr unsafe.Pointer, old, new
unsafe.Pointer) (swapped bool) - func CompareAndSwapUint32(addr uint32,
old, new uint32) (swapped bool) - func CompareAndSwapUint64(addr uint64,
old, new uint64) (swapped bool) - func CompareAndSwapUintptr(addr
uintptr, old, new uintptr) (swapped bool) - func LoadInt32(addr int32)
(val int32) - func LoadInt64(addr int64) (val int64) - func
LoadPointer(addr unsafe.Pointer) (val unsafe.Pointer) - func
LoadUint32(addr uint32) (val uint32) - func LoadUint64(addr uint64) (val
uint64) - func LoadUintptr(addr uintptr) (val uintptr) - func
OrInt32(addr int32, mask int32) (old int32) - func OrInt64(addr int64,
mask int64) (old int64) - func OrUint32(addr uint32, mask uint32) (old
uint32) - func OrUint64(addr uint64, mask uint64) (old uint64) - func
OrUintptr(addr uintptr, mask uintptr) (old uintptr) - func
StoreInt32(addr int32, val int32) - func StoreInt64(addr int64, val
int64) - func StorePointer(addr unsafe.Pointer, val unsafe.Pointer) -
func StoreUint32(addr uint32, val uint32) - func StoreUint64(addr
uint64, val uint64) - func StoreUintptr(addr uintptr, val uintptr) -
func SwapInt32(addr int32, new int32) (old int32) - func SwapInt64(addr
int64, new int64) (old int64) - func SwapPointer(addr unsafe.Pointer,
new unsafe.Pointer) (old unsafe.Pointer) - func SwapUint32(addr uint32,
new uint32) (old uint32) - func SwapUint64(addr uint64, new uint64) (old
uint64) - func SwapUintptr(addr *uintptr, new uintptr) (old uintptr)

### 13. `net`

What is this package? UDP, domain name resolution, and Unix domain
sockets.

What purpose do we use this package for? UDP, domain name resolution,
and Unix domain sockets.

Functions: - func Dial(network, address string) (Conn, error) - func
DialIP(network string, laddr, raddr IPAddr) (IPConn, error) - func
DialTCP(network string, laddr, raddr TCPAddr) (TCPConn, error) - func
DialTimeout(network, address string, timeout time.Duration) (Conn,
error) - func DialUDP(network string, laddr, raddr UDPAddr) (UDPConn,
error) - func DialUnix(network string, laddr, raddr UnixAddr) (UnixConn,
error) - func FileConn(f os.File) (c Conn, err error) - func
FileListener(f os.File) (ln Listener, err error) - func FilePacketConn(f
os.File) (c PacketConn, err error) - func InterfaceAddrs() ([]Addr,
error) - func InterfaceByIndex(index int) (Interface, error) - func
InterfaceByName(name string) (Interface, error) - func Interfaces()
([]Interface, error) - func JoinHostPort(host, port string) string -
func Listen(network, address string) (Listener, error) - func
ListenIP(network string, laddr IPAddr) (IPConn, error) - func
ListenMulticastUDP(network string, ifi Interface, gaddr UDPAddr)
(UDPConn, error) - func ListenPacket(network, address string)
(PacketConn, error) - func ListenTCP(network string, laddr TCPAddr)
(TCPListener, error) - func ListenUDP(network string, laddr UDPAddr)
(UDPConn, error) - func ListenUnix(network string, laddr UnixAddr)
(UnixListener, error) - func ListenUnixgram(network string, laddr
UnixAddr) (UnixConn, error) - func LookupAddr(addr string) (names
[]string, err error) - func LookupCNAME(host string) (cname string, err
error) - func LookupHost(host string) (addrs []string, err error) - func
LookupIP(host string) ([]IP, error) - func LookupMX(name string) ([]MX,
error) - func LookupNS(name string) ([]NS, error) - func
LookupPort(network, service string) (port int, err error) - func
LookupSRV(service, proto, name string) (cname string, addrs []SRV, err
error) - func LookupTXT(name string) ([]string, error) - func
ParseCIDR(s string) (IP, IPNet, error) - func ParseMAC(s string) (hw
HardwareAddr, err error) - func Pipe() (Conn, Conn) - func
ResolveIPAddr(network, address string) (IPAddr, error) - func
ResolveTCPAddr(network, address string) (TCPAddr, error) - func
ResolveUDPAddr(network, address string) (UDPAddr, error) - func
ResolveUnixAddr(network, address string) (UnixAddr, error) - func
SplitHostPort(hostport string) (host, port string, err error)

### 14. `net/http`

What is this package? Get, Head, Post, and PostForm make HTTP (or HTTPS)
requests:

What purpose do we use this package for? Get, Head, Post, and PostForm
make HTTP (or HTTPS) requests:

Functions: - func CanonicalHeaderKey(s string) string - func
DetectContentType(data []byte) string - func Error(w ResponseWriter,
error string, code int) - func Get(url string) (resp Response, err
error) - func Handle(pattern string, handler Handler) - func
HandleFunc(pattern string, handler func(ResponseWriter, Request)) - func
Head(url string) (resp Response, err error) - func ListenAndServe(addr
string, handler Handler) error - func ListenAndServeTLS(addr, certFile,
keyFile string, handler Handler) error - func MaxBytesReader(w
ResponseWriter, r io.ReadCloser, n int64) io.ReadCloser - func
NewRequest(method, url string, body io.Reader) (Request, error) - func
NewRequestWithContext(ctx context.Context, method, url string, body
io.Reader) (Request, error) - func NotFound(w ResponseWriter, r
Request) - func ParseCookie(line string) ([]Cookie, error) - func
ParseHTTPVersion(vers string) (major, minor int, ok bool) - func
ParseSetCookie(line string) (Cookie, error) - func ParseTime(text
string) (t time.Time, err error) - func Post(url, contentType string,
body io.Reader) (resp Response, err error) - func PostForm(url string,
data url.Values) (resp Response, err error) - func
ProxyFromEnvironment(req Request) (url.URL, error) - func
ProxyURL(fixedURL url.URL) func(Request) (url.URL, error) - func
ReadRequest(b bufio.Reader) (Request, error) - func ReadResponse(r
bufio.Reader, req Request) (Response, error) - func Redirect(w
ResponseWriter, r Request, url string, code int) - func Serve(l
net.Listener, handler Handler) error - func ServeContent(w
ResponseWriter, req Request, name string, modtime time.Time, …) - func
ServeFile(w ResponseWriter, r Request, name string) - func ServeFileFS(w
ResponseWriter, r Request, fsys fs.FS, name string) - func ServeTLS(l
net.Listener, handler Handler, certFile, keyFile string) error - func
SetCookie(w ResponseWriter, cookie *Cookie) - func StatusText(code int)
string

### 15. `net/url`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func JoinPath(base string, elem …string) (result string,
err error) - func PathEscape(s string) string - func PathUnescape(s
string) (string, error) - func QueryEscape(s string) string - func
QueryUnescape(s string) (string, error)

### 16. `encoding/json`

What is this package? The mapping between JSON and Go values is
described in the documentation for the Marshal and Unmarshal functions.

What purpose do we use this package for? The mapping between JSON and Go
values is described in the documentation for the Marshal and Unmarshal
functions.

Functions: - func Compact(dst bytes.Buffer, src []byte) error - func
HTMLEscape(dst bytes.Buffer, src []byte) - func Indent(dst
*bytes.Buffer, src []byte, prefix, indent string) error - func Marshal(v
any) ([]byte, error) - func MarshalIndent(v any, prefix, indent string)
([]byte, error) - func Unmarshal(data []byte, v any) error - func
Valid(data []byte) bool

### 17. `path/filepath`

What is this package? a way compatible with the target operating
system-defined file paths.

What purpose do we use this package for? a way compatible with the
target operating system-defined file paths.

Functions: - func Abs(path string) (string, error) - func Base(path
string) string - func Clean(path string) string - func Dir(path string)
string - func EvalSymlinks(path string) (string, error) - func Ext(path
string) string - func FromSlash(path string) string - func Glob(pattern
string) (matches []string, err error) - func HasPrefix(p, prefix string)
bool - func IsAbs(path string) bool - func IsLocal(path string) bool -
func Join(elem …string) string - func Localize(path string) (string,
error) - func Match(pattern, name string) (matched bool, err error) -
func Rel(basepath, targpath string) (string, error) - func Split(path
string) (dir, file string) - func SplitList(path string) []string - func
ToSlash(path string) string - func VolumeName(path string) string - func
Walk(root string, fn WalkFunc) error - func WalkDir(root string, fn
fs.WalkDirFunc) error

### 18. `regexp`

What is this package? The syntax of the regular expressions accepted is
the same general syntax used by Perl, Python, and other languages. More
precisely, it is the syntax accepted by RE2 and described at
https://golang.org/s/re2syntax, except for . For an overview of the
syntax, see the regexp/syntax package.

What purpose do we use this package for? The syntax of the regular
expressions accepted is the same general syntax used by Perl, Python,
and other languages. More precisely, it is the syntax accepted by RE2
and described at https://golang.org/s/re2syntax, except for . For an
overview of the syntax, see the regexp/syntax package.

Functions: - func Match(pattern string, b []byte) (matched bool, err
error) - func MatchReader(pattern string, r io.RuneReader) (matched
bool, err error) - func MatchString(pattern string, s string) (matched
bool, err error) - func QuoteMeta(s string) string

### 19. `sort`

What is this package? collections.

What purpose do we use this package for? collections.

Functions: - func Find(n int, cmp func(int) int) (i int, found bool) -
func Float64s(x []float64) - func Float64sAreSorted(x []float64) bool -
func Ints(x []int) - func IntsAreSorted(x []int) bool - func
IsSorted(data Interface) bool - func Search(n int, f func(int) bool)
int - func SearchFloat64s(a []float64, x float64) int - func
SearchInts(a []int, x int) int - func SearchStrings(a []string, x
string) int - func Slice(x any, less func(i, j int) bool) - func
SliceIsSorted(x any, less func(i, j int) bool) bool - func SliceStable(x
any, less func(i, j int) bool) - func Sort(data Interface) - func
Stable(data Interface) - func Strings(x []string) - func
StringsAreSorted(x []string) bool

### 20. `slices`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func AllSlice ~[]E, E any iter.Seq2[int, E] - func
AppendSeqSlice ~[]E, E any Slice - func BackwardSlice ~[]E, E any
iter.Seq2[int, E] - func BinarySearchS ~[]E, E cmp.Ordered (int, bool) -
func BinarySearchFuncS ~[]E, E, T any (int, bool) - func ChunkSlice
~[]E, E any iter.Seq[Slice] - func ClipS ~[]E, E any S - func CloneS
~[]E, E any S - func CollectE any []E - func CompactS ~[]E, E comparable
S - func CompactFuncS ~[]E, E any S - func CompareS ~[]E, E cmp.Ordered
int - func CompareFuncS1 ~[]E1, S2 ~[]E2, E1, E2 any int - func ConcatS
~[]E, E any S - func ContainsS ~[]E, E comparable bool - func
ContainsFuncS ~[]E, E any bool - func DeleteS ~[]E, E any S - func
DeleteFuncS ~[]E, E any S - func EqualS ~[]E, E comparable bool - func
EqualFuncS1 ~[]E1, S2 ~[]E2, E1, E2 any bool - func GrowS ~[]E, E any
S - func IndexS ~[]E, E comparable int - func IndexFuncS ~[]E, E any
int - func InsertS ~[]E, E any S - func IsSortedS ~[]E, E cmp.Ordered
bool - func IsSortedFuncS ~[]E, E any bool - func MaxS ~[]E, E
cmp.Ordered E - func MaxFuncS ~[]E, E any E - func MinS ~[]E, E
cmp.Ordered E - func MinFuncS ~[]E, E any E - func RepeatS ~[]E, E any
S - func ReplaceS ~[]E, E any S - func ReverseS ~[]E, E any - func SortS
~[]E, E cmp.Ordered - func SortFuncS ~[]E, E any - func SortStableFuncS
~[]E, E any - func SortedE cmp.Ordered []E - func SortedFuncE any []E -
func SortedStableFuncE any []E - func ValuesSlice ~[]E, E any
iter.Seq[E]

### 21. `maps`

What is this package? This package does not have any special handling
for non-reflexive keys (keys k where k != k), such as floating-point
NaNs.

What purpose do we use this package for? This package does not have any
special handling for non-reflexive keys (keys k where k != k), such as
floating-point NaNs.

Functions: - func AllMap ~map[K]V, K comparable, V any iter.Seq2[K, V] -
func CloneM ~map[K]V, K comparable, V any M - func CollectK comparable,
V any map[K]V - func CopyM1 ~map[K]V, M2 ~map[K]V, K comparable, V any -
func DeleteFuncM ~map[K]V, K comparable, V any - func EqualM1, M2
~map[K]V, K, V comparable bool - func EqualFuncM1 ~map[K]V1, M2
~map[K]V2, K comparable, V1, V2 any bool - func InsertMap ~map[K]V, K
comparable, V any - func KeysMap ~map[K]V, K comparable, V any
iter.Seq[K] - func ValuesMap ~map[K]V, K comparable, V any iter.Seq[V]

### 22. `database/sql`

What is this package? The sql package must be used in conjunction with a
database driver. See https://golang.org/s/sqldrivers for a list of
drivers.

What purpose do we use this package for? The sql package must be used in
conjunction with a database driver. See https://golang.org/s/sqldrivers
for a list of drivers.

Functions: - func Drivers() []string - func Register(name string, driver
driver.Driver)

### 23. `testing`

What is this package? It is intended to be used in concert with the “go
test” command, which automates execution of any function of the form

What purpose do we use this package for? It is intended to be used in
concert with the “go test” command, which automates execution of any
function of the form

Functions: - func AllocsPerRun(runs int, f func()) (avg float64) - func
CoverMode() string - func Coverage() float64 - func Init() - func
Main(matchString func(pat, str string) (bool, error), tests
[]InternalTest, …) - func RegisterCover(c Cover) - func
RunBenchmarks(matchString func(pat, str string) (bool, error), …) - func
RunExamples(matchString func(pat, str string) (bool, error), examples
[]InternalExample) (ok bool) - func RunTests(matchString func(pat, str
string) (bool, error), tests []InternalTest) (ok bool) - func Short()
bool - func Testing() bool - func Verbose() bool

### 24. `net/http/httptest`

What is this package? const DefaultRemoteAddr = “1.2.3.4”

What purpose do we use this package for? const DefaultRemoteAddr =
“1.2.3.4”

Functions: - func NewRequest(method, target string, body io.Reader)
http.Request - func NewRequestWithContext(ctx context.Context, method,
target string, body io.Reader) http.Request

### 25. `log`

What is this package? with methods for formatting output. It also has a
predefined ‘standard’ Logger accessible through helper functions
Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than
creating a Logger manually. That logger writes to standard error and
prints the date and time of each logged message. Every log message is
output on a separate line: if the message being printed does not end in
a newline, the logger will add one. The Fatal functions call os.Exit(1)
after writing the log message. The Panic functions call panic after
writing the log message.

What purpose do we use this package for? with methods for formatting
output. It also has a predefined ‘standard’ Logger accessible through
helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are
easier to use than creating a Logger manually. That logger writes to
standard error and prints the date and time of each logged message.
Every log message is output on a separate line: if the message being
printed does not end in a newline, the logger will add one. The Fatal
functions call os.Exit(1) after writing the log message. The Panic
functions call panic after writing the log message.

Functions: - func Fatal(v …any) - func Fatalf(format string, v …any) -
func Fatalln(v …any) - func Flags() int - func Output(calldepth int, s
string) error - func Panic(v …any) - func Panicf(format string, v
…any) - func Panicln(v …any) - func Prefix() string - func Print(v
…any) - func Printf(format string, v …any) - func Println(v …any) - func
SetFlags(flag int) - func SetOutput(w io.Writer) - func SetPrefix(prefix
string) - func Writer() io.Writer

### 26. `log/slog`

What is this package? message, a severity level, and various other
attributes expressed as key-value pairs.

What purpose do we use this package for? message, a severity level, and
various other attributes expressed as key-value pairs.

Functions: - func Debug(msg string, args …any) - func DebugContext(ctx
context.Context, msg string, args …any) - func Error(msg string, args
…any) - func ErrorContext(ctx context.Context, msg string, args …any) -
func Info(msg string, args …any) - func InfoContext(ctx context.Context,
msg string, args …any) - func Log(ctx context.Context, level Level, msg
string, args …any) - func LogAttrs(ctx context.Context, level Level, msg
string, attrs …Attr) - func NewLogLogger(h Handler, level Level)
log.Logger - func SetDefault(l Logger) - func Warn(msg string, args
…any) - func WarnContext(ctx context.Context, msg string, args …any)

============================================================ 🟠 ORANGE

### 27. `flag`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func Arg(i int) string - func Args() []string - func
Bool(name string, value bool, usage string) bool - func BoolFunc(name,
usage string, fn func(string) error) - func BoolVar(p bool, name string,
value bool, usage string) - func Duration(name string, value
time.Duration, usage string) time.Duration - func DurationVar(p
time.Duration, name string, value time.Duration, usage string) - func
Float64(name string, value float64, usage string) float64 - func
Float64Var(p float64, name string, value float64, usage string) - func
Func(name, usage string, fn func(string) error) - func Int(name string,
value int, usage string) int - func Int64(name string, value int64,
usage string) int64 - func Int64Var(p int64, name string, value int64,
usage string) - func IntVar(p int, name string, value int, usage
string) - func NArg() int - func NFlag() int - func Parse() - func
Parsed() bool - func PrintDefaults() - func Set(name, value string)
error - func String(name string, value string, usage string) string -
func StringVar(p string, name string, value string, usage string) - func
TextVar(p encoding.TextUnmarshaler, name string, value
encoding.TextMarshaler, …) - func Uint(name string, value uint, usage
string) uint - func Uint64(name string, value uint64, usage string)
uint64 - func Uint64Var(p uint64, name string, value uint64, usage
string) - func UintVar(p uint, name string, value uint, usage string) -
func UnquoteUsage(flag Flag) (name string, usage string) - func
Var(value Value, name string, usage string) - func Visit(fn
func(Flag)) - func VisitAll(fn func(*Flag))

### 28. `os/exec`

What is this package? to remap stdin and stdout, connect I/O with pipes,
and do other adjustments.

What purpose do we use this package for? to remap stdin and stdout,
connect I/O with pipes, and do other adjustments.

Functions: - func LookPath(file string) (string, error)

### 29. `io/fs`

What is this package? provided by the host operating system but also by
other packages.

What purpose do we use this package for? provided by the host operating
system but also by other packages.

Functions: - func FormatDirEntry(dir DirEntry) string - func
FormatFileInfo(info FileInfo) string - func Glob(fsys FS, pattern
string) (matches []string, err error) - func ReadFile(fsys FS, name
string) ([]byte, error) - func ValidPath(name string) bool - func
WalkDir(fsys FS, root string, fn WalkDirFunc) error

### 30. `encoding/csv`

What is this package? kinds of CSV files; this package supports the
format described in RFC 4180, except that Writer uses LF instead of CRLF
as newline character by default.

What purpose do we use this package for? kinds of CSV files; this
package supports the format described in RFC 4180, except that Writer
uses LF instead of CRLF as newline character by default.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 31. `encoding/xml`

What is this package? const Header =  + “” var HTMLAutoClose []string =
htmlAutoClose var HTMLEntity map[string]string = htmlEntity

What purpose do we use this package for? const Header =  + “” var
HTMLAutoClose []string = htmlAutoClose var HTMLEntity map[string]string
= htmlEntity

Functions: - func Escape(w io.Writer, s []byte) - func EscapeText(w
io.Writer, s []byte) error - func Marshal(v any) ([]byte, error) - func
MarshalIndent(v any, prefix, indent string) ([]byte, error) - func
Unmarshal(data []byte, v any) error

### 32. `encoding/base64`

What is this package? const StdPadding rune = ‘=’ … var RawStdEncoding =
StdEncoding.WithPadding(NoPadding) var RawURLEncoding =
URLEncoding.WithPadding(NoPadding) var StdEncoding =
NewEncoding(“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”)
var URLEncoding =
NewEncoding(“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_“)

What purpose do we use this package for? const StdPadding rune = ‘=’ …
var RawStdEncoding = StdEncoding.WithPadding(NoPadding) var
RawURLEncoding = URLEncoding.WithPadding(NoPadding) var StdEncoding =
NewEncoding(“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/”)
var URLEncoding =
NewEncoding(“ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_“)

Functions: - func NewDecoder(enc Encoding, r io.Reader) io.Reader - func
NewEncoder(enc Encoding, w io.Writer) io.WriteCloser

### 33. `encoding/hex`

What is this package? var ErrLength = errors.New(“encoding/hex: odd
length hex string”)

What purpose do we use this package for? var ErrLength =
errors.New(“encoding/hex: odd length hex string”)

Functions: - func AppendDecode(dst, src []byte) ([]byte, error) - func
AppendEncode(dst, src []byte) []byte - func Decode(dst, src []byte)
(int, error) - func DecodeString(s string) ([]byte, error) - func
DecodedLen(x int) int - func Dump(data []byte) string - func Dumper(w
io.Writer) io.WriteCloser - func Encode(dst, src []byte) int - func
EncodeToString(src []byte) string - func EncodedLen(n int) int - func
NewDecoder(r io.Reader) io.Reader - func NewEncoder(w io.Writer)
io.Writer

### 34. `encoding/binary`

What is this package? and encoding and decoding of varints.

What purpose do we use this package for? and encoding and decoding of
varints.

Functions: - func Append(buf []byte, order ByteOrder, data any) ([]byte,
error) - func AppendUvarint(buf []byte, x uint64) []byte - func
AppendVarint(buf []byte, x int64) []byte - func Decode(buf []byte, order
ByteOrder, data any) (int, error) - func Encode(buf []byte, order
ByteOrder, data any) (int, error) - func PutUvarint(buf []byte, x
uint64) int - func PutVarint(buf []byte, x int64) int - func Read(r
io.Reader, order ByteOrder, data any) error - func ReadUvarint(r
io.ByteReader) (uint64, error) - func ReadVarint(r io.ByteReader)
(int64, error) - func Size(v any) int - func Uvarint(buf []byte)
(uint64, int) - func Varint(buf []byte) (int64, int) - func Write(w
io.Writer, order ByteOrder, data any) error

### 35. `crypto/rand`

What is this package? var Reader io.Reader

What purpose do we use this package for? var Reader io.Reader

Functions: - func Int(rand io.Reader, max big.Int) (n big.Int, err
error) - func Prime(rand io.Reader, bits int) (*big.Int, error) - func
Read(b []byte) (n int, err error)

### 36. `crypto/sha256`

What is this package? FIPS 180-4.

What purpose do we use this package for? FIPS 180-4.

Functions: - func New() hash.Hash - func New224() hash.Hash - func
Sum224(data []byte) [Size224]byte - func Sum256(data []byte) [Size]byte

### 37. `crypto/sha512`

What is this package? hash algorithms as defined in FIPS 180-4.

What purpose do we use this package for? hash algorithms as defined in
FIPS 180-4.

Functions: - func New() hash.Hash - func New384() hash.Hash - func
New512_224() hash.Hash - func New512_256() hash.Hash - func Sum384(data
[]byte) [Size384]byte - func Sum512(data []byte) [Size]byte - func
Sum512_224(data []byte) [Size224]byte - func Sum512_256(data []byte)
[Size256]byte

### 38. `crypto/hmac`

What is this package? defined in U.S. Federal Information Processing
Standards Publication 198. An HMAC is a cryptographic hash that uses a
key to sign a message. The receiver verifies the hash by recomputing it
using the same key.

What purpose do we use this package for? defined in U.S. Federal
Information Processing Standards Publication 198. An HMAC is a
cryptographic hash that uses a key to sign a message. The receiver
verifies the hash by recomputing it using the same key.

Functions: - func Equal(mac1, mac2 []byte) bool - func New(h func()
hash.Hash, key []byte) hash.Hash

### 39. `crypto/aes`

What is this package? Federal Information Processing Standards
Publication 197.

What purpose do we use this package for? Federal Information Processing
Standards Publication 197.

Functions: - func NewCipher(key []byte) (cipher.Block, error)

### 40. `crypto/cipher`

What is this package? be wrapped around low-level block cipher
implementations. See
https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html and NIST
Special Publication 800-38A.

What purpose do we use this package for? be wrapped around low-level
block cipher implementations. See
https://csrc.nist.gov/groups/ST/toolkit/BCM/current_modes.html and NIST
Special Publication 800-38A.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 41. `crypto/tls`

What is this package? as specified in RFC 8446.

What purpose do we use this package for? as specified in RFC 8446.

Functions: - func CipherSuiteName(id uint16) string - func
Listen(network, laddr string, config Config) (net.Listener, error) -
func NewListener(inner net.Listener, config Config) net.Listener - func
VersionName(version uint16) string

### 42. `crypto/x509`

What is this package? It allows parsing and generating certificates,
certificate signing requests, certificate revocation lists, and encoded
public and private keys. It provides a certificate verifier, complete
with a chain builder.

What purpose do we use this package for? It allows parsing and
generating certificates, certificate signing requests, certificate
revocation lists, and encoded public and private keys. It provides a
certificate verifier, complete with a chain builder.

Functions: - func CreateCertificate(rand io.Reader, template, parent
Certificate, pub, priv any) ([]byte, error) - func
CreateCertificateRequest(rand io.Reader, template CertificateRequest,
priv any) (csr []byte, err error) - func CreateRevocationList(rand
io.Reader, template RevocationList, issuer Certificate, …) ([]byte,
error) - func DecryptPEMBlock(b pem.Block, password []byte) ([]byte,
error) - func EncryptPEMBlock(rand io.Reader, blockType string, data,
password []byte, alg PEMCipher) (pem.Block, error) - func
IsEncryptedPEMBlock(b pem.Block) bool - func MarshalECPrivateKey(key
ecdsa.PrivateKey) ([]byte, error) - func MarshalPKCS1PrivateKey(key
rsa.PrivateKey) []byte - func MarshalPKCS1PublicKey(key rsa.PublicKey)
[]byte - func MarshalPKCS8PrivateKey(key any) ([]byte, error) - func
MarshalPKIXPublicKey(pub any) ([]byte, error) - func ParseCRL(crlBytes
[]byte) (pkix.CertificateList, error) - func ParseDERCRL(derBytes
[]byte) (pkix.CertificateList, error) - func ParseECPrivateKey(der
[]byte) (ecdsa.PrivateKey, error) - func ParsePKCS1PrivateKey(der
[]byte) (rsa.PrivateKey, error) - func ParsePKCS1PublicKey(der []byte)
(rsa.PublicKey, error) - func ParsePKCS8PrivateKey(der []byte) (key any,
err error) - func ParsePKIXPublicKey(derBytes []byte) (pub any, err
error) - func SetFallbackRoots(roots CertPool)

### 43. `encoding/pem`

What is this package? Enhanced Mail. The most common use of PEM encoding
today is in TLS keys and certificates. See RFC 1421.

What purpose do we use this package for? Enhanced Mail. The most common
use of PEM encoding today is in TLS keys and certificates. See RFC 1421.

Functions: - func Encode(out io.Writer, b Block) error - func
EncodeToMemory(b Block) []byte

### 44. `mime`

What is this package? const BEncoding = WordEncoder(‘b’) … var
ErrInvalidMediaParameter = errors.New(“mime: invalid media parameter”)

What purpose do we use this package for? const BEncoding =
WordEncoder(‘b’) … var ErrInvalidMediaParameter = errors.New(“mime:
invalid media parameter”)

Functions: - func AddExtensionType(ext, typ string) error - func
ExtensionsByType(typ string) ([]string, error) - func FormatMediaType(t
string, param map[string]string) string - func ParseMediaType(v string)
(mediatype string, params map[string]string, err error) - func
TypeByExtension(ext string) string

### 45. `mime/multipart`

What is this package? The implementation is sufficient for HTTP (RFC
2388) and the multipart bodies generated by popular browsers.

What purpose do we use this package for? The implementation is
sufficient for HTTP (RFC 2388) and the multipart bodies generated by
popular browsers.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 46. `html/template`

What is this package? HTML output safe against code injection. It
provides the same interface as text/template and should be used instead
of text/template whenever the output is HTML.

What purpose do we use this package for? HTML output safe against code
injection. It provides the same interface as text/template and should be
used instead of text/template whenever the output is HTML.

Functions: - func HTMLEscape(w io.Writer, b []byte) - func
HTMLEscapeString(s string) string - func HTMLEscaper(args …any) string -
func IsTrue(val any) (truth, ok bool) - func JSEscape(w io.Writer, b
[]byte) - func JSEscapeString(s string) string - func JSEscaper(args
…any) string - func URLQueryEscaper(args …any) string

### 47. `text/template`

What is this package? To generate HTML output, see html/template, which
has the same interface as this package but automatically secures HTML
output against certain attacks.

What purpose do we use this package for? To generate HTML output, see
html/template, which has the same interface as this package but
automatically secures HTML output against certain attacks.

Functions: - func HTMLEscape(w io.Writer, b []byte) - func
HTMLEscapeString(s string) string - func HTMLEscaper(args …any) string -
func IsTrue(val any) (truth, ok bool) - func JSEscape(w io.Writer, b
[]byte) - func JSEscapeString(s string) string - func JSEscaper(args
…any) string - func URLQueryEscaper(args …any) string

### 48. `embed`

What is this package? Go source files that import “embed” can use the
//go:embed directive to initialize a variable of type string, []byte, or
FS with the contents of files read from the package directory or
subdirectories at compile time.

What purpose do we use this package for? Go source files that import
“embed” can use the //go:embed directive to initialize a variable of
type string, []byte, or FS with the contents of files read from the
package directory or subdirectories at compile time.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

============================================================ 🟡 YELLOW

### 49. `cmp`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func CompareT Ordered int - func LessT Ordered bool - func
OrT comparable T

### 50. `crypto/rsa`

What is this package? RSA is a single, fundamental operation that is
used in this package to implement either public-key encryption or
public-key signatures.

What purpose do we use this package for? RSA is a single, fundamental
operation that is used in this package to implement either public-key
encryption or public-key signatures.

Functions: - func DecryptOAEP(hash hash.Hash, random io.Reader, priv
PrivateKey, ciphertext []byte, …) ([]byte, error) - func
DecryptPKCS1v15(random io.Reader, priv PrivateKey, ciphertext []byte)
([]byte, error) - func DecryptPKCS1v15SessionKey(random io.Reader, priv
PrivateKey, ciphertext []byte, key []byte) error - func EncryptOAEP(hash
hash.Hash, random io.Reader, pub PublicKey, msg []byte, label []byte)
([]byte, error) - func EncryptPKCS1v15(random io.Reader, pub PublicKey,
msg []byte) ([]byte, error) - func SignPKCS1v15(random io.Reader, priv
PrivateKey, hash crypto.Hash, hashed []byte) ([]byte, error) - func
SignPSS(rand io.Reader, priv PrivateKey, hash crypto.Hash, digest
[]byte, …) ([]byte, error) - func VerifyPKCS1v15(pub PublicKey, hash
crypto.Hash, hashed []byte, sig []byte) error - func VerifyPSS(pub
PublicKey, hash crypto.Hash, digest []byte, sig []byte, opts PSSOptions)
error

### 51. `crypto/ecdsa`

What is this package? as defined in FIPS 186-4 and SEC 1, Version 2.0.

What purpose do we use this package for? as defined in FIPS 186-4 and
SEC 1, Version 2.0.

Functions: - func Sign(rand io.Reader, priv PrivateKey, hash []byte) (r,
s big.Int, err error) - func SignASN1(rand io.Reader, priv PrivateKey,
hash []byte) ([]byte, error) - func Verify(pub PublicKey, hash []byte,
r, s big.Int) bool - func VerifyASN1(pub PublicKey, hash, sig []byte)
bool

### 52. `crypto/ed25519`

What is this package? https://ed25519.cr.yp.to/.

What purpose do we use this package for? https://ed25519.cr.yp.to/.

Functions: - func GenerateKey(rand io.Reader) (PublicKey, PrivateKey,
error) - func Sign(privateKey PrivateKey, message []byte) []byte - func
Verify(publicKey PublicKey, message, sig []byte) bool - func
VerifyWithOptions(publicKey PublicKey, message, sig []byte, opts
*Options) error

### 53. `math`

What is this package? This package does not guarantee bit-identical
results across architectures.

What purpose do we use this package for? This package does not guarantee
bit-identical results across architectures.

Functions: - func Abs(x float64) float64 - func Acos(x float64)
float64 - func Acosh(x float64) float64 - func Asin(x float64) float64 -
func Asinh(x float64) float64 - func Atan(x float64) float64 - func
Atan2(y, x float64) float64 - func Atanh(x float64) float64 - func
Cbrt(x float64) float64 - func Ceil(x float64) float64 - func
Copysign(f, sign float64) float64 - func Cos(x float64) float64 - func
Cosh(x float64) float64 - func Dim(x, y float64) float64 - func Erf(x
float64) float64 - func Erfc(x float64) float64 - func Erfcinv(x
float64) float64 - func Erfinv(x float64) float64 - func Exp(x float64)
float64 - func Exp2(x float64) float64 - func Expm1(x float64) float64 -
func FMA(x, y, z float64) float64 - func Float32bits(f float32) uint32 -
func Float32frombits(b uint32) float32 - func Float64bits(f float64)
uint64 - func Float64frombits(b uint64) float64 - func Floor(x float64)
float64 - func Frexp(f float64) (frac float64, exp int) - func Gamma(x
float64) float64 - func Hypot(p, q float64) float64 - func Ilogb(x
float64) int - func Inf(sign int) float64 - func IsInf(f float64, sign
int) bool - func IsNaN(f float64) (is bool) - func J0(x float64)
float64 - func J1(x float64) float64 - func Jn(n int, x float64)
float64 - func Ldexp(frac float64, exp int) float64 - func Lgamma(x
float64) (lgamma float64, sign int) - func Log(x float64) float64 - func
Log10(x float64) float64 - func Log1p(x float64) float64 - func Log2(x
float64) float64 - func Logb(x float64) float64 - func Max(x, y float64)
float64 - func Min(x, y float64) float64 - func Mod(x, y float64)
float64 - func Modf(f float64) (int float64, frac float64) - func NaN()
float64 - func Nextafter(x, y float64) (r float64) - func Nextafter32(x,
y float32) (r float32) - func Pow(x, y float64) float64 - func Pow10(n
int) float64 - func Remainder(x, y float64) float64 - func Round(x
float64) float64 - func RoundToEven(x float64) float64 - func Signbit(x
float64) bool - func Sin(x float64) float64 - func Sincos(x float64)
(sin, cos float64) - func Sinh(x float64) float64 - func Sqrt(x float64)
float64 - func Tan(x float64) float64 - func Tanh(x float64) float64 -
func Trunc(x float64) float64 - func Y0(x float64) float64 - func Y1(x
float64) float64 - func Yn(n int, x float64) float64

### 54. `math/big`

What is this package? following numeric types are supported:

What purpose do we use this package for? following numeric types are
supported:

Functions: - func Jacobi(x, y Int) int - func ParseFloat(s string, base
int, prec uint, mode RoundingMode) (f Float, b int, err error)

### 55. `math/bits`

What is this package? predeclared unsigned integer types.

What purpose do we use this package for? predeclared unsigned integer
types.

Functions: - func Add(x, y, carry uint) (sum, carryOut uint) - func
Add32(x, y, carry uint32) (sum, carryOut uint32) - func Add64(x, y,
carry uint64) (sum, carryOut uint64) - func Div(hi, lo, y uint) (quo,
rem uint) - func Div32(hi, lo, y uint32) (quo, rem uint32) - func
Div64(hi, lo, y uint64) (quo, rem uint64) - func LeadingZeros(x uint)
int - func LeadingZeros16(x uint16) int - func LeadingZeros32(x uint32)
int - func LeadingZeros64(x uint64) int - func LeadingZeros8(x uint8)
int - func Len(x uint) int - func Len16(x uint16) (n int) - func Len32(x
uint32) (n int) - func Len64(x uint64) (n int) - func Len8(x uint8)
int - func Mul(x, y uint) (hi, lo uint) - func Mul32(x, y uint32) (hi,
lo uint32) - func Mul64(x, y uint64) (hi, lo uint64) - func OnesCount(x
uint) int - func OnesCount16(x uint16) int - func OnesCount32(x uint32)
int - func OnesCount64(x uint64) int - func OnesCount8(x uint8) int -
func Rem(hi, lo, y uint) uint - func Rem32(hi, lo, y uint32) uint32 -
func Rem64(hi, lo, y uint64) uint64 - func Reverse(x uint) uint - func
Reverse16(x uint16) uint16 - func Reverse32(x uint32) uint32 - func
Reverse64(x uint64) uint64 - func Reverse8(x uint8) uint8 - func
ReverseBytes(x uint) uint - func ReverseBytes16(x uint16) uint16 - func
ReverseBytes32(x uint32) uint32 - func ReverseBytes64(x uint64) uint64 -
func RotateLeft(x uint, k int) uint - func RotateLeft16(x uint16, k int)
uint16 - func RotateLeft32(x uint32, k int) uint32 - func RotateLeft64(x
uint64, k int) uint64 - func RotateLeft8(x uint8, k int) uint8 - func
Sub(x, y, borrow uint) (diff, borrowOut uint) - func Sub32(x, y, borrow
uint32) (diff, borrowOut uint32) - func Sub64(x, y, borrow uint64)
(diff, borrowOut uint64) - func TrailingZeros(x uint) int - func
TrailingZeros16(x uint16) int - func TrailingZeros32(x uint32) int -
func TrailingZeros64(x uint64) int - func TrailingZeros8(x uint8) int

### 56. `path`

What is this package? The path package should only be used for paths
separated by forward slashes, such as the paths in URLs. This package
does not deal with Windows paths with drive letters or backslashes; to
manipulate operating system paths, use the path/filepath package.

What purpose do we use this package for? The path package should only be
used for paths separated by forward slashes, such as the paths in URLs.
This package does not deal with Windows paths with drive letters or
backslashes; to manipulate operating system paths, use the path/filepath
package.

Functions: - func Base(path string) string - func Clean(path string)
string - func Dir(path string) string - func Ext(path string) string -
func IsAbs(path string) bool - func Join(elem …string) string - func
Match(pattern, name string) (matched bool, err error) - func Split(path
string) (dir, file string)

### 57. `runtime`

What is this package? such as functions to control goroutines. It also
includes the low-level type information used by the reflect package; see
reflect’s documentation for the programmable interface to the run-time
type system.

What purpose do we use this package for? such as functions to control
goroutines. It also includes the low-level type information used by the
reflect package; see reflect’s documentation for the programmable
interface to the run-time type system.

Functions: - func BlockProfile(p []BlockProfileRecord) (n int, ok
bool) - func Breakpoint() - func CPUProfile() []byte - func Caller(skip
int) (pc uintptr, file string, line int, ok bool) - func Callers(skip
int, pc []uintptr) int - func GC() - func GOMAXPROCS(n int) int - func
GOROOT() string - func Goexit() - func GoroutineProfile(p []StackRecord)
(n int, ok bool) - func Gosched() - func KeepAlive(x any) - func
LockOSThread() - func MemProfile(p []MemProfileRecord, inuseZero bool)
(n int, ok bool) - func MutexProfile(p []BlockProfileRecord) (n int, ok
bool) - func NumCPU() int - func NumCgoCall() int64 - func
NumGoroutine() int - func ReadMemStats(m *MemStats) - func ReadTrace()
[]byte - func SetBlockProfileRate(rate int) - func SetCPUProfileRate(hz
int) - func SetCgoTraceback(version int, traceback, context, symbolizer
unsafe.Pointer) - func SetFinalizer(obj any, finalizer any) - func
SetMutexProfileFraction(rate int) int - func Stack(buf []byte, all bool)
int - func StartTrace() error - func StopTrace() - func
ThreadCreateProfile(p []StackRecord) (n int, ok bool) - func
UnlockOSThread() - func Version() string

### 58. `runtime/debug`

What is this package? are running.

What purpose do we use this package for? are running.

Functions: - func FreeOSMemory() - func PrintStack() - func
ReadGCStats(stats GCStats) - func SetCrashOutput(f os.File, opts
CrashOptions) error - func SetGCPercent(percent int) int - func
SetMaxStack(bytes int) int - func SetMaxThreads(threads int) int - func
SetMemoryLimit(limit int64) int64 - func SetPanicOnFault(enabled bool)
bool - func SetTraceback(level string) - func Stack() []byte - func
WriteHeapDump(fd uintptr)

### 59. `runtime/pprof`

What is this package? visualization tool.

What purpose do we use this package for? visualization tool.

Functions: - func Do(ctx context.Context, labels LabelSet, f
func(context.Context)) - func ForLabels(ctx context.Context, f func(key,
value string) bool) - func Label(ctx context.Context, key string)
(string, bool) - func SetGoroutineLabels(ctx context.Context) - func
StartCPUProfile(w io.Writer) error - func StopCPUProfile() - func
WithLabels(ctx context.Context, labels LabelSet) context.Context - func
WriteHeapProfile(w io.Writer) error

### 60. `net/mail`

What is this package? For the most part, this package follows the syntax
as specified by RFC 5322 and extended by RFC 6532. Notable
divergences: - Obsolete address formats are not parsed, including
addresses with embedded route information. - The full range of spacing
(the CFWS syntax element) is not supported, such as breaking addresses
across lines. - No unicode normalization is performed. - A leading From
line is permitted, as in mbox format (RFC 4155).

What purpose do we use this package for? For the most part, this package
follows the syntax as specified by RFC 5322 and extended by RFC 6532.
Notable divergences: - Obsolete address formats are not parsed,
including addresses with embedded route information. - The full range of
spacing (the CFWS syntax element) is not supported, such as breaking
addresses across lines. - No unicode normalization is performed. - A
leading From line is permitted, as in mbox format (RFC 4155).

Functions: - func ParseDate(date string) (time.Time, error)

### 61. `testing/fstest`

What is this package? systems.

What purpose do we use this package for? systems.

Functions: - func TestFS(fsys fs.FS, expected …string) error

============================================================ 🔵 BLUE

### 62. `archive/tar`

What is this package? Tape archives (tar) are a file format for storing
a sequence of files that can be read and written in a streaming manner.
This package aims to cover most variations of the format, including
those produced by GNU and BSD tar tools.

What purpose do we use this package for? Tape archives (tar) are a file
format for storing a sequence of files that can be read and written in a
streaming manner. This package aims to cover most variations of the
format, including those produced by GNU and BSD tar tools.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 63. `archive/zip`

What is this package? See the ZIP specification for details.

What purpose do we use this package for? See the ZIP specification for
details.

Functions: - func RegisterCompressor(method uint16, comp Compressor) -
func RegisterDecompressor(method uint16, dcomp Decompressor)

### 64. `compress/gzip`

What is this package? as specified in RFC 1952.

What purpose do we use this package for? as specified in RFC 1952.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 65. `compress/zlib`

What is this package? as specified in RFC 1950.

What purpose do we use this package for? as specified in RFC 1950.

Functions: - func NewReader(r io.Reader) (io.ReadCloser, error) - func
NewReaderDict(r io.Reader, dict []byte) (io.ReadCloser, error)

### 66. `compress/flate`

What is this package? 1951. The gzip and zlib packages implement access
to DEFLATE-based file formats.

What purpose do we use this package for? 1951. The gzip and zlib
packages implement access to DEFLATE-based file formats.

Functions: - func NewReader(r io.Reader) io.ReadCloser - func
NewReaderDict(r io.Reader, dict []byte) io.ReadCloser

### 67. `container/heap`

What is this package? heap.Interface. A heap is a tree with the property
that each node is the minimum-valued node in its subtree.

What purpose do we use this package for? heap.Interface. A heap is a
tree with the property that each node is the minimum-valued node in its
subtree.

Functions: - func Fix(h Interface, i int) - func Init(h Interface) -
func Pop(h Interface) any - func Push(h Interface, x any) - func
Remove(h Interface, i int) any

### 68. `container/list`

What is this package? To iterate over a list (where l is a *List):

What purpose do we use this package for? To iterate over a list (where l
is a *List):

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 69. `container/ring`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 70. `image`

What is this package? The fundamental interface is called Image. An
Image contains colors, which are described in the image/color package.

What purpose do we use this package for? The fundamental interface is
called Image. An Image contains colors, which are described in the
image/color package.

Functions: - func RegisterFormat(name, magic string, decode
func(io.Reader) (Image, error), …)

### 71. `image/color`

What is this package? var Black = Gray16{ … } …

What purpose do we use this package for? var Black = Gray16{ … } …

Functions: - func CMYKToRGB(c, m, y, k uint8) (uint8, uint8, uint8) -
func RGBToCMYK(r, g, b uint8) (uint8, uint8, uint8, uint8) - func
RGBToYCbCr(r, g, b uint8) (uint8, uint8, uint8) - func YCbCrToRGB(y, cb,
cr uint8) (uint8, uint8, uint8)

### 72. `image/png`

What is this package? The PNG specification is at
https://www.w3.org/TR/PNG/.

What purpose do we use this package for? The PNG specification is at
https://www.w3.org/TR/PNG/.

Functions: - func Decode(r io.Reader) (image.Image, error) - func
DecodeConfig(r io.Reader) (image.Config, error) - func Encode(w
io.Writer, m image.Image) error

### 73. `image/jpeg`

What is this package? JPEG is defined in ITU-T T.81:
https://www.w3.org/Graphics/JPEG/itu-t81.pdf.

What purpose do we use this package for? JPEG is defined in ITU-T T.81:
https://www.w3.org/Graphics/JPEG/itu-t81.pdf.

Functions: - func Decode(r io.Reader) (image.Image, error) - func
DecodeConfig(r io.Reader) (image.Config, error) - func Encode(w
io.Writer, m image.Image, o *Options) error

### 74. `image/gif`

What is this package? The GIF specification is at
https://www.w3.org/Graphics/GIF/spec-gif89a.txt.

What purpose do we use this package for? The GIF specification is at
https://www.w3.org/Graphics/GIF/spec-gif89a.txt.

Functions: - func Decode(r io.Reader) (image.Image, error) - func
DecodeConfig(r io.Reader) (image.Config, error) - func Encode(w
io.Writer, m image.Image, o Options) error - func EncodeAll(w io.Writer,
g GIF) error

============================================================ 🟢 GREEN

### 75. `go/ast`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func FileExports(src File) bool - func FilterDecl(decl
Decl, f Filter) bool - func FilterFile(src File, f Filter) bool - func
FilterPackage(pkg Package, f Filter) bool - func Fprint(w io.Writer,
fset token.FileSet, x any, f FieldFilter) error - func Inspect(node
Node, f func(Node) bool) - func IsExported(name string) bool - func
IsGenerated(file File) bool - func NotNilFilter(_ string, v
reflect.Value) bool - func PackageExports(pkg Package) bool - func
Preorder(root Node) iter.Seq[Node] - func Print(fset token.FileSet, x
any) error - func SortImports(fset token.FileSet, f *File) - func Walk(v
Visitor, node Node)

### 76. `go/parser`

What is this package? a variety of forms (see the various Parse
functions); the output is an abstract syntax tree (AST) representing the
Go source. The parser is invoked through one of the Parse functions.

What purpose do we use this package for? a variety of forms (see the
various Parse* functions); the output is an abstract syntax tree (AST)
representing the Go source. The parser is invoked through one of the
Parse* functions.

Functions: - func ParseDir(fset token.FileSet, path string, filter
func(fs.FileInfo) bool, mode Mode) (pkgs map[string]ast.Package, first
error) - func ParseExpr(x string) (ast.Expr, error) - func
ParseExprFrom(fset token.FileSet, filename string, src any, mode Mode)
(expr ast.Expr, err error) - func ParseFile(fset token.FileSet, filename
string, src any, mode Mode) (f *ast.File, err error)

### 77. `go/token`

What is this package? programming language and basic operations on
tokens (printing, predicates).

What purpose do we use this package for? programming language and basic
operations on tokens (printing, predicates).

Functions: - func IsExported(name string) bool - func IsIdentifier(name
string) bool - func IsKeyword(name string) bool

### 78. `go/types`

What is this package? type-checking of Go packages. Use Config.Check to
invoke the type checker for a package. Alternatively, create a new type
checker with NewChecker and invoke it incrementally by calling
Checker.Files.

What purpose do we use this package for? type-checking of Go packages.
Use Config.Check to invoke the type checker for a package.
Alternatively, create a new type checker with NewChecker and invoke it
incrementally by calling Checker.Files.

Functions: - func AssertableTo(V Interface, T Type) bool - func
AssignableTo(V, T Type) bool - func CheckExpr(fset token.FileSet, pkg
Package, pos token.Pos, expr ast.Expr, info Info) (err error) - func
Comparable(T Type) bool - func ConvertibleTo(V, T Type) bool - func
DefPredeclaredTestFuncs() - func ExprString(x ast.Expr) string - func
Id(pkg Package, name string) string - func Identical(x, y Type) bool -
func IdenticalIgnoreTags(x, y Type) bool - func Implements(V Type, T
Interface) bool - func IsInterface(t Type) bool - func ObjectString(obj
Object, qf Qualifier) string - func Satisfies(V Type, T Interface)
bool - func SelectionString(s Selection, qf Qualifier) string - func
TypeString(typ Type, qf Qualifier) string - func WriteExpr(buf
bytes.Buffer, x ast.Expr) - func WriteSignature(buf bytes.Buffer, sig
Signature, qf Qualifier) - func WriteType(buf bytes.Buffer, typ Type, qf
Qualifier)

### 79. `go/format`

What is this package? Note that formatting of Go source code changes
over time, so tools relying on consistent formatting should execute a
specific version of the gofmt binary instead of using this package. That
way, the formatting will be stable, and the tools won’t need to be
recompiled each time gofmt changes.

What purpose do we use this package for? Note that formatting of Go
source code changes over time, so tools relying on consistent formatting
should execute a specific version of the gofmt binary instead of using
this package. That way, the formatting will be stable, and the tools
won’t need to be recompiled each time gofmt changes.

Functions: - func Node(dst io.Writer, fset *token.FileSet, node any)
error - func Source(src []byte) ([]byte, error)

### 80. `debug/dwarf`

What is this package? from executable files, as defined in the DWARF 2.0
Standard at http://dwarfstd.org/doc/dwarf-2.0.0.pdf.

What purpose do we use this package for? from executable files, as
defined in the DWARF 2.0 Standard at
http://dwarfstd.org/doc/dwarf-2.0.0.pdf.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 81. `debug/elf`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - func NewFile(r io.ReaderAt) (File, error) - func Open(name
string) (File, error) - func R_INFO(sym, typ uint32) uint64 - func
R_INFO32(sym, typ uint32) uint32 - func R_SYM32(info uint32) uint32 -
func R_SYM64(info uint64) uint32 - func R_TYPE32(info uint32) uint32 -
func R_TYPE64(info uint64) uint32 - func ST_INFO(bind SymBind, typ
SymType) uint8

### 82. `debug/macho`

What is this package? A Go standard-library package.

What purpose do we use this package for? Use this package for the
functionality provided by its standard-library API.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 83. `debug/pe`

What is this package? files.

What purpose do we use this package for? files.

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 84. `plugin`

What is this package? A plugin is a Go main package with exported
functions and variables that has been built with:

What purpose do we use this package for? A plugin is a Go main package
with exported functions and variables that has been built with:

Functions: - No exported package-level functions listed; this package
primarily exposes types, methods, constants, or variables.

### 85. `syscall`

What is this package? primitives. The details vary depending on the
underlying system, and by default, godoc will display the syscall
documentation for the current system. If you want godoc to display
syscall documentation for another system, set $GOOS and $GOARCH to the
desired system. For example, if you want to view documentation for
freebsd/arm on linux/amd64, set $GOOS to freebsd and $GOARCH to arm. The
primary use of syscall is inside other packages that provide a more
portable interface to the system, such as “os”, “time” and “net”. Use
those packages rather than this one if you can. For details of the
functions and data types in this package consult the manuals for the
appropriate operating system. These calls return err == nil to indicate
success; otherwise err is an operating system error describing the
failure. On most systems, that error has type Errno.

What purpose do we use this package for? primitives. The details vary
depending on the underlying system, and by default, godoc will display
the syscall documentation for the current system. If you want godoc to
display syscall documentation for another system, set $GOOS and $GOARCH
to the desired system. For example, if you want to view documentation
for freebsd/arm on linux/amd64, set $GOOS to freebsd and $GOARCH to arm.
The primary use of syscall is inside other packages that provide a more
portable interface to the system, such as “os”, “time” and “net”. Use
those packages rather than this one if you can. For details of the
functions and data types in this package consult the manuals for the
appropriate operating system. These calls return err == nil to indicate
success; otherwise err is an operating system error describing the
failure. On most systems, that error has type Errno.

Functions: - func Accept(fd int) (nfd int, sa Sockaddr, err error) -
func Accept4(fd int, flags int) (nfd int, sa Sockaddr, err error) - func
Access(path string, mode uint32) (err error) - func Acct(path string)
(err error) - func Adjtimex(buf Timex) (state int, err error) - func
AllThreadsSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err
Errno) - func AllThreadsSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr)
(r1, r2 uintptr, err Errno) - func AttachLsf(fd int, i []SockFilter)
error - func Bind(fd int, sa Sockaddr) (err error) - func
BindToDevice(fd int, device string) (err error) - func
BytePtrFromString(s string) (byte, error) - func ByteSliceFromString(s
string) ([]byte, error) - func Chdir(path string) (err error) - func
Chmod(path string, mode uint32) (err error) - func Chown(path string,
uid int, gid int) (err error) - func Chroot(path string) (err error) -
func Clearenv() - func Close(fd int) (err error) - func CloseOnExec(fd
int) - func CmsgLen(datalen int) int - func CmsgSpace(datalen int) int -
func Connect(fd int, sa Sockaddr) (err error) - func Creat(path string,
mode uint32) (fd int, err error) - func DetachLsf(fd int) error - func
Dup(oldfd int) (fd int, err error) - func Dup2(oldfd int, newfd int)
(err error) - func Dup3(oldfd int, newfd int, flags int) (err error) -
func Environ() []string - func EpollCreate(size int) (fd int, err
error) - func EpollCreate1(flag int) (fd int, err error) - func
EpollCtl(epfd int, op int, fd int, event EpollEvent) (err error) - func
EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) -
func Exec(argv0 string, argv []string, envv []string) (err error) - func
Exit(code int) - func Faccessat(dirfd int, path string, mode uint32,
flags int) (err error) - func Fallocate(fd int, mode uint32, off int64,
len int64) (err error) - func Fchdir(fd int) (err error) - func
Fchmod(fd int, mode uint32) (err error) - func Fchmodat(dirfd int, path
string, mode uint32, flags int) error - func Fchown(fd int, uid int, gid
int) (err error) - func Fchownat(dirfd int, path string, uid int, gid
int, flags int) (err error) - func FcntlFlock(fd uintptr, cmd int, lk
Flock_t) error - func Fdatasync(fd int) (err error) - func Flock(fd int,
how int) (err error) - func ForkExec(argv0 string, argv []string, attr
ProcAttr) (pid int, err error) - func Fstat(fd int, stat Stat_t) (err
error) - func Fstatfs(fd int, buf Statfs_t) (err error) - func Fsync(fd
int) (err error) - func Ftruncate(fd int, length int64) (err error) -
func Futimes(fd int, tv []Timeval) (err error) - func Futimesat(dirfd
int, path string, tv []Timeval) (err error) - func Getcwd(buf []byte) (n
int, err error) - func Getdents(fd int, buf []byte) (n int, err error) -
func Getegid() (egid int) - func Getenv(key string) (value string, found
bool) - func Geteuid() (euid int) - func Getgid() (gid int) - func
Getgroups() (gids []int, err error) - func Getpagesize() int - func
Getpeername(fd int) (sa Sockaddr, err error) - func Getpgid(pid int)
(pgid int, err error) - func Getpgrp() (pid int) - func Getpid() (pid
int) - func Getppid() (ppid int) - func Getpriority(which int, who int)
(prio int, err error) - func Getrlimit(resource int, rlim Rlimit) (err
error) - func Getrusage(who int, rusage Rusage) (err error) - func
Getsockname(fd int) (sa Sockaddr, err error) - func
GetsockoptICMPv6Filter(fd, level, opt int) (ICMPv6Filter, error) - func
GetsockoptIPMreq(fd, level, opt int) (IPMreq, error) - func
GetsockoptIPMreqn(fd, level, opt int) (IPMreqn, error) - func
GetsockoptIPv6MTUInfo(fd, level, opt int) (IPv6MTUInfo, error) - func
GetsockoptIPv6Mreq(fd, level, opt int) (IPv6Mreq, error) - func
GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error) -
func GetsockoptInt(fd, level, opt int) (value int, err error) - func
GetsockoptUcred(fd, level, opt int) (Ucred, error) - func Gettid() (tid
int) - func Gettimeofday(tv Timeval) (err error) - func Getuid() (uid
int) - func Getwd() (wd string, err error) - func Getxattr(path string,
attr string, dest []byte) (sz int, err error) - func InotifyAddWatch(fd
int, pathname string, mask uint32) (watchdesc int, err error) - func
InotifyInit() (fd int, err error) - func InotifyInit1(flags int) (fd
int, err error) - func InotifyRmWatch(fd int, watchdesc uint32) (success
int, err error) - func Ioperm(from int, num int, on int) (err error) -
func Iopl(level int) (err error) - func Kill(pid int, sig Signal) (err
error) - func Klogctl(typ int, buf []byte) (n int, err error) - func
Lchown(path string, uid int, gid int) (err error) - func Link(oldpath
string, newpath string) (err error) - func Listen(s int, n int) (err
error) - func Listxattr(path string, dest []byte) (sz int, err error) -
func LsfSocket(ifindex, proto int) (int, error) - func Lstat(path
string, stat Stat_t) (err error) - func Madvise(b []byte, advice int)
(err error) - func Mkdir(path string, mode uint32) (err error) - func
Mkdirat(dirfd int, path string, mode uint32) (err error) - func
Mkfifo(path string, mode uint32) (err error) - func Mknod(path string,
mode uint32, dev int) (err error) - func Mknodat(dirfd int, path string,
mode uint32, dev int) (err error) - func Mlock(b []byte) (err error) -
func Mlockall(flags int) (err error) - func Mmap(fd int, offset int64,
length int, prot int, flags int) (data []byte, err error) - func
Mount(source string, target string, fstype string, flags uintptr, data
string) (err error) - func Mprotect(b []byte, prot int) (err error) -
func Munlock(b []byte) (err error) - func Munlockall() (err error) -
func Munmap(b []byte) (err error) - func Nanosleep(time Timespec,
leftover Timespec) (err error) - func NetlinkRIB(proto, family int)
([]byte, error) - func Open(path string, mode int, perm uint32) (fd int,
err error) - func Openat(dirfd int, path string, flags int, mode uint32)
(fd int, err error) - func ParseDirent(buf []byte, max int, names
[]string) (consumed int, count int, newnames []string) - func
ParseNetlinkMessage(b []byte) ([]NetlinkMessage, error) - func
ParseNetlinkRouteAttr(m NetlinkMessage) ([]NetlinkRouteAttr, error) -
func ParseSocketControlMessage(b []byte) ([]SocketControlMessage,
error) - func ParseUnixCredentials(m SocketControlMessage) (Ucred,
error) - func ParseUnixRights(m SocketControlMessage) ([]int, error) -
func Pause() (err error) - func Pipe(p []int) error - func Pipe2(p
[]int, flags int) error - func PivotRoot(newroot string, putold string)
(err error) - func Pread(fd int, p []byte, offset int64) (n int, err
error) - func PtraceAttach(pid int) (err error) - func PtraceCont(pid
int, signal int) (err error) - func PtraceDetach(pid int) (err error) -
func PtraceGetEventMsg(pid int) (msg uint, err error) - func
PtraceGetRegs(pid int, regsout PtraceRegs) (err error) - func
PtracePeekData(pid int, addr uintptr, out []byte) (count int, err
error) - func PtracePeekText(pid int, addr uintptr, out []byte) (count
int, err error) - func PtracePokeData(pid int, addr uintptr, data
[]byte) (count int, err error) - func PtracePokeText(pid int, addr
uintptr, data []byte) (count int, err error) - func PtraceSetOptions(pid
int, options int) (err error) - func PtraceSetRegs(pid int, regs
PtraceRegs) (err error) - func PtraceSingleStep(pid int) (err error) -
func PtraceSyscall(pid int, signal int) (err error) - func Pwrite(fd
int, p []byte, offset int64) (n int, err error) - func RawSyscall(trap,
a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno) - func RawSyscall6(trap,
a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) - func
Read(fd int, p []byte) (n int, err error) - func ReadDirent(fd int, buf
[]byte) (n int, err error) - func Readlink(path string, buf []byte) (n
int, err error) - func Reboot(cmd int) (err error) - func Recvfrom(fd
int, p []byte, flags int) (n int, from Sockaddr, err error) - func
Recvmsg(fd int, p, oob []byte, flags int) (n, oobn int, recvflags int,
from Sockaddr, err error) - func Removexattr(path string, attr string)
(err error) - func Rename(oldpath string, newpath string) (err error) -
func Renameat(olddirfd int, oldpath string, newdirfd int, newpath
string) (err error) - func Rmdir(path string) error - func Seek(fd int,
offset int64, whence int) (off int64, err error) - func Select(nfd int,
r FdSet, w FdSet, e FdSet, timeout Timeval) (n int, err error) - func
Sendfile(outfd int, infd int, offset int64, count int) (written int, err
error) - func Sendmsg(fd int, p, oob []byte, to Sockaddr, flags int)
(err error) - func SendmsgN(fd int, p, oob []byte, to Sockaddr, flags
int) (n int, err error) - func Sendto(fd int, p []byte, flags int, to
Sockaddr) (err error) - func SetLsfPromisc(name string, m bool) error -
func SetNonblock(fd int, nonblocking bool) (err error) - func
Setdomainname(p []byte) (err error) - func Setegid(egid int) (err
error) - func Setenv(key, value string) error - func Seteuid(euid int)
(err error) - func Setfsgid(gid int) (err error) - func Setfsuid(uid
int) (err error) - func Setgid(gid int) (err error) - func
Setgroups(gids []int) (err error) - func Sethostname(p []byte) (err
error) - func Setpgid(pid int, pgid int) (err error) - func
Setpriority(which int, who int, prio int) (err error) - func
Setregid(rgid, egid int) (err error) - func Setresgid(rgid, egid, sgid
int) (err error) - func Setresuid(ruid, euid, suid int) (err error) -
func Setreuid(ruid, euid int) (err error) - func Setrlimit(resource int,
rlim Rlimit) error - func Setsid() (pid int, err error) - func
SetsockoptByte(fd, level, opt int, value byte) (err error) - func
SetsockoptICMPv6Filter(fd, level, opt int, filter ICMPv6Filter) error -
func SetsockoptIPMreq(fd, level, opt int, mreq IPMreq) (err error) -
func SetsockoptIPMreqn(fd, level, opt int, mreq IPMreqn) (err error) -
func SetsockoptIPv6Mreq(fd, level, opt int, mreq IPv6Mreq) (err error) -
func SetsockoptInet4Addr(fd, level, opt int, value [4]byte) (err
error) - func SetsockoptInt(fd, level, opt int, value int) (err error) -
func SetsockoptLinger(fd, level, opt int, l Linger) (err error) - func
SetsockoptString(fd, level, opt int, s string) (err error) - func
SetsockoptTimeval(fd, level, opt int, tv Timeval) (err error) - func
Settimeofday(tv Timeval) (err error) - func Setuid(uid int) (err
error) - func Setxattr(path string, attr string, data []byte, flags int)
(err error) - func Shutdown(fd int, how int) (err error) - func
SlicePtrFromStrings(ss []string) ([]byte, error) - func Socket(domain,
typ, proto int) (fd int, err error) - func Socketpair(domain, typ, proto
int) (fd [2]int, err error) - func Splice(rfd int, roff int64, wfd int,
woff int64, len int, flags int) (n int64, err error) - func
StartProcess(argv0 string, argv []string, attr ProcAttr) (pid int,
handle uintptr, err error) - func Stat(path string, stat Stat_t) (err
error) - func Statfs(path string, buf Statfs_t) (err error) - func
StringBytePtr(s string) byte - func StringByteSlice(s string) []byte -
func StringSlicePtr(ss []string) []byte - func Symlink(oldpath string,
newpath string) (err error) - func Sync() - func SyncFileRange(fd int,
off int64, n int64, flags int) (err error) - func Syscall(trap, a1, a2,
a3 uintptr) (r1, r2 uintptr, err Errno) - func Syscall6(trap, a1, a2,
a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno) - func Sysinfo(info
Sysinfo_t) (err error) - func Tee(rfd int, wfd int, len int, flags int)
(n int64, err error) - func Tgkill(tgid int, tid int, sig Signal) (err
error) - func Time(t Time_t) (tt Time_t, err error) - func Times(tms
Tms) (ticks uintptr, err error) - func TimespecToNsec(ts Timespec)
int64 - func TimevalToNsec(tv Timeval) int64 - func Truncate(path
string, length int64) (err error) - func Umask(mask int) (oldmask int) -
func Uname(buf Utsname) (err error) - func UnixCredentials(ucred Ucred)
[]byte - func UnixRights(fds …int) []byte - func Unlink(path string)
error - func Unlinkat(dirfd int, path string) error - func
Unmount(target string, flags int) (err error) - func Unsetenv(key
string) error - func Unshare(flags int) (err error) - func Ustat(dev
int, ubuf Ustat_t) (err error) - func Utime(path string, buf Utimbuf)
(err error) - func Utimes(path string, tv []Timeval) (err error) - func
UtimesNano(path string, ts []Timespec) (err error) - func Wait4(pid int,
wstatus WaitStatus, options int, rusage *Rusage) (wpid int, err error) -
func Write(fd int, p []byte) (n int, err error)

============================================================ END
