**Warning:** this prototype is very subject to change and will be moved
somewhere.  Joint work between sigsum, trust fabric, transparency.dev, etc.

# PROJECT-NAME - A community-maintained witness network

PROJECT-NAME is a community-maintained repository of metadata that simplifies
configuration of append-only transparency logs that use [witness cosigning][].

The current community maintainers are:

  - Name, organization
  - ...

[witness cosigning]: https://C2SP.org/tlog-witness

## Background 

Log and witness operators need to *mutually* chose each other to get a reliable
witnessing setup.  (A log that accepts any witness will be subject to DoS.  A
witness that accepts any log will be subject to DoS.)

Either the log or the witness operator needs to initiate mutual configuration.
It makes sense that the log operator is the initiator.  (The log anyway have to
manually assess which witnesses are reasonable from a trust policy perspective.)

It should be as easy as possible to operate a witness.  (It is the component
that brings trust into the system.  If it is easy to operate, then it is easier
to get a diverse set of reliable witnesses.  It is also harder to screw up.)

It is an operational burden for a witness to get asked to configure every log,
including assessments like "does it make sense to witness this log".  It is an
operational burden for a log operator to ask multiple witnesses to configure it.

## Overview

PROJECT-NAME is a central repository where log operators can request to be
witnessed.  Upon successful registration, the log operator can pick-and-chose
from participating witnesses that configure every log PROJECT-NAME accepted.

This is achieved by maintaining a list of logs that participating witnesses
configure.  The list is community maintained because it is for the community.

    TODO: figure describing this system.

PROJECT-NAME only helps with the initial discovery for mutual configuration.
Not being involved in updates or removals makes PROJECT-NAME a less juicy target
for attacks.  (It is bad if a central repository can disable all witnessing.)

## Objectives

  - Help witness operators discover logs that would like to be witnessed
    (automatically by periodically downloading a list of logs).
  - Help log operators discover witnesses they may collect cosignatures from
    (manually by registering and selecting some participating witnesses).
  - Avoid repeated configuration requests between logs and witnesses.

## Non-goals

  - Say anything about which trust policy a system of logs should use.  It is up
    to logs and users to select witnesses they find reliable and trustworthy.
  - Be the sole dictator of which logs a witness operator will configure.  Use
    of complementary log lists and other manual configuration is encouraged.
  - Have the power to centrally override any previously applied configuration.

## Participating witnesses

Below is a table of witness operators that configure logs discovered in
PROJECT-NAME's lists.  Log operators can pick-and-chose from witnesses that
configure them.  It is optional to collect cosignatures from a given witness.

  | Operator        | Configures     | About page                                                                                      |
  | --------------- | -------------- | ----------------------------------------------------------------------------------------------- |
  | Glasklar Teknik | [10qps-1Ml][]  | <https://git.glasklar.is/glasklar/services/witnessing/-/blob/main/witness.glasklar.is/about.md> |
  | ...             | ...            | ...                                                                                             |

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

Note that the above means a participating witness *MUST NOT* remove or update a
log's definition just because PROJECT-NAME publishes an updated log list.  This
significantly decreases the amount of power that PROJECT-NAME has, i.e.,
PROJECT-NAME helps with initial discovery but have no influence after that.

TODO: guideline on when an update is obviously bogus, don't apply it?

TODO: key rotation support?

## List of logs

For now, there is a single list of logs that would like to be witnessed.

  - [10qps-1Ml][]

The file name describes what performance profile a witness configuring the list
must be able to handle.  For example, `10qps-1Ml` means the list is maintained
to work for a witness that can handle 10 add-checkpoint requests (sustained on
average) with enough persistent storage to support at least one million logs.

The exact list format is documented separately, see [log-list format][].

TODO: detached signature?

[10qps-1Ml]: ./lists/10qps-1Ml
[log-list format]: ./log-list-format.md

## Register a log for witnessing

[File an issue][] or send an email to MAILING-LIST.

Specify:

  - The log's verification key in [vkey format][].  Use a schema-less URL and a
    key name that is the same as the log's [origin line][].
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

## Register a participating witness

[File an issue][] or send an email to MAILING-LIST.  Specify the information
needed to populate the table of participating witnesses.

For inspiration, you may look at a few previous configuration requests:

  - To be added
  - ...

**Note:** there is no guarantee that a request to be added will be granted.

## Frequently asked questions

To be added?
