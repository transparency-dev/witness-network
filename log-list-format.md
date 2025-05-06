# Log-list format

This document describes a line-terminated log-list format.  Lines starting with
`#` denote comments and are therefore ignored while parsing a log list.

## Example

    # list header
    logs/v0
    revision 2

    # 1st list item -- foo's public log
    origin tlog.foo.org
    vkey tlog.foo.org+aaaaaa+AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
    bastion https://bastion.tlog.foo.org/
    qpd 86400
    contact https://tlog.foo.org/contact

    # 2nd list item -- bar's serverless log
    origin bar.org/serverless-log
    vkey bar.org/serverless-log+bbbbbbbb+BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
    bastion none
    qpd 24
    contact sysadmin (at) bar.org

## Header

A log list starts with the following three-line header:

    logs/v0
    revision REVISION
    BLANK LINE

`logs/v0` is the version of this format.

`REVISION` is a monotonically increasing counter, bumped on list updates.

`BLANK LINE` is a line only containing a new line (U+000A).  This marks the end
of the header.  Zero or more logs will follow after this, see below.

## List of logs

A log is defined by the following in-order lines:

    origin ORIGIN
    vkey VKEY
    bastion BASTION
    qpd QPD
    contact CONTACT

`ORIGIN` is the log's origin line, see <https://C2SP.org/tlog-checkpoint#note-text>.

`VKEY` is the log's verification key in vkey format, see
<https://github.com/C2SP/C2SP/pull/119/files>.

`BASTION` is either `none` or a bastion URL that the witness may connect to, see
<https://C2SP.org/https-bastion>.

`QPD` is the number of add-checkpoint requests the log may do per day.

`CONTACT` is an arbitrary string useful for humans to reach the log operator.

To define another log, add a `BLANK LINE` and repeat the above definition.
