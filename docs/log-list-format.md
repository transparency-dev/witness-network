# Log-list format

This document describes a line-terminated log-list format.  Blank lines are
ignored.  Lines starting with `#` denote comments and are also ignored.

## Example

    #
    # List:      10qps-100klogs
    # Revision:  123
    # Generated: YYYY-MM-DD HH:MM:SS UTC
    # Other undefined debug information.
    #
    logs/v0

    # 1st list item -- foo's log
    vkey tlog.foo.org+aaaaaa+AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA
    origin none
    qpd 86400
    bastion https://bastion.tlog.foo.org/
    contact https://tlog.foo.org/contact

    # 2nd list item -- bar's log
    vkey bar.org/tlog+bbbbbbbb+BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
    origin something-not-equal-to-vkey-keyname
    qpd 24
    bastion none
    contact sysadmin (at) bar.org

## Header

The list starts with a single line declaring the list format.

    logs/v0

Zero or more logs follow after the `logs/v0` line.

## List of logs

A log is defined by the below lines.  The order of lines is significant.

    vkey VKEY
    origin ORIGIN
    bastion BASTION
    qpd QPD
    contact CONTACT

`VKEY` is the log's verification key in vkey format, see
<https://github.com/C2SP/C2SP/pull/119/files>.

`ORIGIN` is either `none` or the log's origin line, see
<https://C2SP.org/tlog-checkpoint#note-text>.  If `none`, the log's origin line
is the same as the vkey key-name.  This is recommended for newly created logs.

`BASTION` is either `none` or a bastion URL the witness may connect to, see
<https://C2SP.org/https-bastion>.  If `none`, the log operator may have fewer
witnesses to chose from because some witnesses are only reachable via bastions.

`QPD` is the number of add-checkpoint requests the log may do per day.

`CONTACT` is an arbitrary string useful for humans to reach the log operator.

Repeat the above lines to define one more log.
