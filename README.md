**Warning:** this prototype is very subject to change and will be moved
somewhere.  Joint work between sigsum, trust fabric, transparency.dev, etc.

# PROJECT-NAME - A community-maintained witness network

PROJECT-NAME is a community-maintained repository of metadata that simplifies
configuration of append-only transparency logs that use witness cosigning.

The current community maintainers are:

  - Name, organization
  - ...

## Overview

To be added.

## Objectives

  - Help witness operators discover logs that would like to be witnessed
    (automatically).
  - Help log operators discover witnesses they may collect cosignatures from
    (manually).
  - Avoid repeated configuration requests between logs and witnesses.

## Non-goals

  - Have the power to centrally override any previous configuration.
  - Say anything about which trust policy a system of logs should use.

## List of logs

The `lists` directory contains lists of logs that would like to be witnessed.
Formatting of these lists is documented separately, see [log-list format][].

There is a single list of logs right now:

  - [10qps-1Ml][]

The file name describes what performance profile a witness configuring the list
must be able to handle.  For example, `10qps-1Ml` means the list is maintained
to work for a witness that can handle 10 add-checkpoint requests (sustained on
average) with enough persistent storage to support at least one million logs.

[10qps-1Ml]: ./lists/10qps-1Ml
[log-list format]: ./log-list-format.md

TODO: signed list, not that crucial but makes sense.

## Expectations on participating witnesses

The following specifications must be supported for interoperability:

  - <https://C2SP.org/tlog-checkpoint>
  - <https://C2SP.org/tlog-cosignature>
  - <https://C2SP.org/tlog-witness>

It is optional to use <https://C2SP.org/https-bastion>.

A witness that applies configuration from a list is expected to configure *ALL*
logs, and to not remove a log from it's configuration unless abuse is detected.
PROJECT-NAME has no opinion on what a witness operator considers abuse.  This
and other relevant information should be documented in a witness about page.

**Note:** the above means a participating witness *MUST NOT* remove or update a
log's definition just because PROJECT-NAME publishes a new log list.  This
significantly decreases the amount of power that PROJECT-NAME has, i.e.,
PROJECT-NAME helps with initial discovery but have no influence after that.

TODO: key rotation support?

## Participating witnesses

Below is a table of witness operators that configure different log lists.  Log
operators that have been added to a list can pick-and-chose from witnesses that
configure them.  It is optional to collect cosignatures from a given witness.

Use the listed about pages to learn what can be expected from each witness.

  | Operator        | Configures     | About page                                                                                      |
  | --------------- | -------------- | ----------------------------------------------------------------------------------------------- |
  | Glasklar Teknik | [10qps-1Ml][]  | <https://git.glasklar.is/glasklar/services/witnessing/-/blob/main/witness.glasklar.is/about.md> |
  | ...             | ...            | ...                                                                                             |

## Frequently asked questions

### How to get my log added to a list?

[File an issue][] or send an email to MAILING-LIST.

Specify:

  - The log's verification key in [vkey format][].  Please use a schema-less URL
    and a key name that is the same as the log's [origin line][].
  - Something that convinces the maintainers the origin line makes sense, e.g.,
    that you are from `example.org` if the origin line includes `example.org`.
  - How often the log is expected to submit add-checkpoint requests.
  - Any other information you think will make the decision to include the log
    easier, e.g., remarks regarding utility vs required load.  For example, a
    log that requests cosignatures every second will be harder to get admitted
    to a list compared to a log that only requests cosignatures once per day.
  - Contact information to someone responsible for the log's operations.

For inspiration, you may look at a few previous configuration requests:

  - To be added
  - ...

**Note:** there is no guarantee that a request to be added will be granted.  The
maintainers maintain the lists of logs in good faith to keep them reliable.

[File an issue]: TO-BE-ADDED
[vkey format]: TO-BE-ADDED
[origin line]: https://C2SP.org/tlog-checkpoint#note-text

### How to get my witness added to the table?

[File an issue][] or send an email to MAILING-LIST.  Specify the information
needed to populate the table of participating witnesses.

For inspiration, you may look at a few previous configuration requests:

  - To be added
  - ...

**Note:** there is no guarantee that a request to be added will be granted.

### My log was added to a list, now what?

docdoc

### How to be removed from a list or a table

docdoc. Won't affect what logs or witnesses already configured, but removal
would avoid of course avoid additional configurations that make no sense.

### Other FAQs / things we want to point out
