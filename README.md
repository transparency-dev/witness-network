**Warning:** this prototype is very subject to change and will be moved - WIP.
Joint work between sigsum / trust fabric / transparency.dev / filippo.

# PROJECT-NAME - A community-maintained witness network

PROJECT-NAME is a community-maintained repository that simplifies configuration
of transparency logs with [witness cosigning][].

## Community maintainers

  - Name, organization
  - ...

[witness cosigning]: https://C2SP.org/tlog-witness

## Background 

Log and witness operators need to mutually chose each other to get a reliable
witnessing setup.  (A log that accepts any witness will be subject to DoS.  A
witness that accepts any log will be subject to DoS.)

Either the log or the witness operator needs to initiate mutual configuration.
It makes sense that the log operator is the initiator.  (The log anyway have to
manually assess which witnesses are reasonable from a trust policy perspective.)

It should be as easy as possible to operate a witness.  (It is the component
that brings trust into the system.  If it is easy to operate, then it is easier
to get a diverse set of reliable witnesses.)

It is an operational burden for a witness to configure every log on request.
(Each request requires subjective assessments like "does it make sense to
witness this log".)

It is an operational burden for a log operator to ask multiple witnesses to
configure it.  It is also hard to discover which witnesses may be asked.

## Overview

PROJECT-NAME is a central repository where log operators can request to be
witnessed.  Witnesses that participate in PROJECT-NAME are committed to
configure all new logs that appear in a list maintained by PROJECT-NAME.  So,
admitted log operators can collect cosignatures from any participating witness.

    TODO: figure describing this system.

## Objectives

**High-level goals:**

  - Help witness operators discover logs that would like to be witnessed
    (automatically by configuring new logs from a community list).
  - Help log operators discover witnesses they may collect cosignatures from
    (manually by registering and selecting some participating witnesses).
  - Avoid repeated configuration requests between logs and witnesses.

**To be avoided:**

  - Get the power to automatically override any previously applied
    configuration.  (It is bad if a central party can disrupt witnessing.)
  - Become the dictator of which logs witness operators configure.  (Use of
    complementary log lists and any other manual configuration is encouraged.)
  - Say anything about which trust policy a log system should use.  (It is up to
    logs and users to select witnesses they find reliable and trustworthy.)

## List of logs

Right now there is a single list of logs that would like to be witnessed.

  - [10qps-1Mlogs][]

The file name describes what *performance profile* a witness configuring the
list must be able to handle.  For example, `10qps-1Mlogs` means the list is
maintained to work for a witness that can handle 10 add-checkpoint requests
(sustained on average) with enough persistent storage to support 1 million logs.
The requests/s parameter is global, i.e., it applies for all logs combined.

The exact list format is documented separately, see [log-list format][].

[10qps-1Mlogs]: ./lists/10qps-1Mlogs
[log-list format]: ./log-list-format.md

## Participating witnesses

Below is a table of witnesses that admitted log operators can collect
cosignatures from.  It is optional to collect cosignatures from a witness.

  | Operator        | Configures       | About page                                                                                      |
  | --------------- | ---------------- | ----------------------------------------------------------------------------------------------- |
  | Glasklar Teknik | [10qps-1Mlogs][] | <https://git.glasklar.is/glasklar/services/witnessing/-/blob/main/witness.glasklar.is/about.md> |
  | ...             | ...              | ...                                                                                             |

## Expectations on participating witnesses

The following specifications must be supported for interoperability:

  - <https://C2SP.org/tlog-checkpoint>
  - <https://C2SP.org/tlog-cosignature>
  - <https://C2SP.org/tlog-witness>

It is optional to support <https://C2SP.org/https-bastion>.  If bastion host is
not supported, then the witness needs to be reachable on the public internet.

A witness that applies configuration from a list is expected to configure *ALL*
logs, and to not remove or otherwise update a log's configuration just because
PROJECT-NAME publishes a new list.  The witness operator may make its own
removals and configuration updates, e.g., due to detecting abuse.  PROJECT-NAME
has no opinion on what a witness operator considers abuse.  This and other
relevant information should be documented in a witness about page, e.g., so that
informed decisions about (not) depending on a given witness can be formed.

A participating witness must configure new logs at least once per week.  A log
is "new" if none of the already configured logs have the same origin line.

## Register a participating witness

[File an issue][] or send an email to MAILING-LIST.  Specify the information
needed to populate the table of participating witnesses.

For inspiration, you may look at a few previous configuration requests:

  - To be added
  - ...

**Note:** there is no guarantee that a request to be added will be granted.

## Register a log for witnessing

[File an issue][] or send an email to MAILING-LIST.

Specify:

  - The log's public key in [vkey format][].  Use a schema-less URL and a key
    name that is the same as the log's [origin line][].
  - Something that convinces the maintainers the origin line makes sense, e.g.,
    that you are from `example.org` if the origin line contains `example.org`.
  - How often the log is expected to submit add-checkpoint requests.
  - Any other information you think will make the decision to admit the log
    easier, e.g., remarks regarding utility vs required load.  A log that
    requests cosignatures every second will be harder to get admitted to a list
    compared to a log that only requests cosignatures once per day.
  - Contact information to someone responsible for the log's operations.

For inspiration, you may look at a few previous configuration requests:

  - To be added
  - ...

**Note:** there is no guarantee that a request to be added will be granted.  The
maintainers maintain the lists of logs in good faith to keep them reliable.

[File an issue]: TO-BE-ADDED
[vkey format]: TO-BE-ADDED
[origin line]: https://C2SP.org/tlog-checkpoint#note-text

## Frequently asked questions

### How do rotate a log's key?
### What is the impact of a bogus list update?
### What about bogus origin lines?
### Why are the lists not signed?
### Why are the lists not transparency logged?
