- Every Package have a Collection of **Source files**.
- Each folder in your project is considered a **new package**.
- In Go packages and folders mean the same thing.
- Tha naming Convention for packages is to be **all lowercase letters** and "No Underscore and No hypens".
- Naming Convention for files is to do **snake case**
- Every Go file need a package declaraation at the top and that needs to match the *folder name*.
- In Go anything  that starts with a ***uppercase letter***  is a public declaration that can be accessed from the outside of package.
- In Go anything  that starts with a ***lowercase letter***  is a private declaration that can't be accessed from the outside of package.
- *go get github.com/fatih/color*  //example package is used for different colors 


Packages List
-----------------------------------------------------------------------
(Mostly Using Built in Packages)
1.  fmt
2.  net/http
3.  log
4.  Strings
5.  Strconv
6.  bufio
7.  os
8.  time
9.  math
10. io        // 'io' and 'io/ioutil' both pkg used for working with files
11. io/ioutil (not there this pkg after Go Veresion 1.16) // now only 'io' is the package
12. Sync
13. net/url
14. Fatal
15. encoding/json