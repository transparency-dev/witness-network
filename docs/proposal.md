# Proposal on witness configuration network

The target audience for this proposal is: readers that are already familiar with
the transparency log specifications located at <https://C2SP.org>.

## Outline

  - A bit of background and problem statement
  - The concrete components of the proposal
    - List of logs
    - Table of witnesses
    - How to get a log into the list
    - How to get a witness into the table
  - Discussion in Q/A form that helps motivate the design

If/when we agree about the proposal at large, we likely want a shared repository
or website to maintain the list of logs, the table of witnesses, and other
related documentation that makes this easy to digest for operators that wonder
"is this for me", and if so, "what do I need to do to participate".

## Background

In the [witness protocol][], log servers collect cosignatures from witnesses by
sending synchronous `HTTP POST /add-checkpoint` requests.  A witness can have
its own public endpoint for this, or establish a long-lived connection to a
[bastion host][] that forwards requests from authorized logs to witnesses.

For a log server to collect cosignatures from a witness, it needs the witness'
URL and a public key (to verify that the returned cosignature is valid).  Note
that from a log's perspective, even a witness behind a bastion host has a URL.

For a witness to verify an `add-checkpoint` requests, it needs the log's public
key.  The witness also needs enough space to store one [checkpoint][] per log.

So, to get a functioning witnessing setup logs need to configure witnesses and
witnesses need to configure logs.  Without such *mutual configuration*, a log
will either not send a request to a witness or the witness will reject the log's
`add-checkpoint` request (both because the log's public key is unknown and
because simply cosigning any log would result in trivial denial of service).

[witness protocol]: https://C2SP.org/tlog-witness
[bastion host]: https://C2SP.org/https-bastion
[checkpoint]: https://C2SP.org/tlog-checkpoint

## Problem statement

How logs and witnesses mutually configure each other is currently adhoc - every
pair of logs and witnesses need to somehow find each other and then communicate.

The question this proposal is concerned with is: can we make discovery and
mutual configuration easier for operators of logs and witnesses?

## Solution overview

Let's establish a witness configuration network that:

  1. Helps witness operators discover logs that would like to be witnessed
     *automatically by configuring new logs from a community list*.
  2. Helps log operators discover witnesses they may collect cosignatures from
     *manually after first getting into a community list*.  Witnesses that log
     operators can pick and chose from are listed in a community table.

The list of logs and the table of witnesses can be maintained by a few trusted
community members, e.g., in a git repository or via a simple website.  There is
a process for admitting new logs and witnesses, think "send an email to a
mailing list and wait for a maintainer to process the registration request".

Witnesses will want to participate because they can configure new logs from a
living community list (as opposed to having to interact with each log operator).

Logs will want to participate because they can make one request to be included
in a community list that several participating witnesses configure.

An overview is shown in the below figure.

    TODO: figure.

## List of logs

This proposal suggests a machine-readable list of logs:

  - `10qps-1Mlogs`

The file name describes what *performance profile* a witness configuring the
list must be able to handle.  For example, `10qps-1Mlogs` means the list is
maintained to work for a witness that can handle 10 add-checkpoint requests
(sustained on average) with enough persistent storage to support 1 million logs.
The requests/s parameter is global, i.e., it applies for all logs combined.

The exact log-list format is defined separately, see [log-list-format][].

[log-list-format]: ./log-list-format.md

## Table of participating witnesses

This proposal suggests a human-readable table of participating witnesses.  The
most important field of the table is an about page, which log operators can read
to figure out whether it makes sense for them to collect cosignatures or not.

  | Operator        | About page                                                                                      |
  | --------------- | ----------------------------------------------------------------------------------------------- |
  | Glasklar Teknik | <https://git.glasklar.is/glasklar/services/witnessing/-/blob/main/witness.glasklar.is/about.md> |
  | ...             | ...                                                                                             |

A participating witness is expected to configure *ALL* logs, and to not remove
or otherwise update a log's configuration just because the community list is
updated.  The witness operator *MAY* make its own removals and configuration
updates, e.g., due to detecting abuse.  The community maintainers have no
opinion on what a witness operator considers abuse.  This and other relevant
information should be documented in the witness's linked about page.

A participating witness must configure new logs at least once per week.  A log
is "new" if none of the already configured logs have the same origin line.

## Interoperability

Logs and witnesses must support:

  - <https://C2SP.org/tlog-checkpoint>
  - <https://C2SP.org/tlog-cosignature>
  - <https://C2SP.org/tlog-witness>

It is optional to support <https://C2SP.org/https-bastion>.  If bastion host is
not supported, then the witness needs to be reachable on the public internet.

## Get a log into the community list

This proposal suggests that the process for getting a log into the list is to
send an email to a mailing list.  There should be a HOW-TO on what to include.
The HOW-TO could also link a few previous accepted registrations as examples.

Here's a dense list of the information that the log operator should provide:

  - The log's public key in [vkey format][].  The key name should be a
    schema-less URL, and be the same as the log's [origin line][].
  - Something convincing the maintainers the origin line makes sense, e.g., the
    operator is from `example.org` if the origin line contains `example.org`.
  - How often the log is expected to submit add-checkpoint requests (qpd).
  - Any other information that may make the decision to admit the log easier,
    e.g., remarks regarding utility vs required load.  (A log that requests
    cosignatures every second will be harder to get admitted to a list compared
    to a log that only requests cosignatures once per day.)
  - Contact information to someone responsible for the log's operations.

**Note:** there is no guarantee that a request to be added will be granted.  The
maintainers maintain the lists of logs in good faith to keep them reliable.

[vkey format]: https://github.com/C2SP/C2SP/pull/119/files
[origin line]: https://C2SP.org/tlog-checkpoint#note-text

## Get a witness into the community table

This proposal suggests the process for getting a witness into the community
table is to send an email to a mailing list.  There should be a HOW-TO on what
to include, and the gist is "the information needed to populate the table".

**Note:** there is no guarantee that a request to be added will be granted.  The
maintainers maintain the table of witnesses in good faith to keep it useful.

## Discussion

### Is it mandatory to participate in this witness configuration network?

No, a log or a witness that prefers to solve configuration on their own are
encouraged to do so.  Participating logs and witnesses are also encouraged to do
other complementary configuration, e.g., manually or via complementary witness
configuration networks (likely to appear eventually to strike other trade-offs,
or to depend on a different set of maintainers that use a different process).

### Is there a recommended trust policy for the participating witnesses?

No, it is not in scope to have an opinion on which trust policy to use.  This
will naturally vary depending on the log ecosystem, intended use, etc.

### How to remove or update a log in the list?

Not supported, the proposal is scoped solely for the *initial discovery that
facilitates mutual configuration between logs and witnesses*.  This ensures that
the maintainers are not in a position to disrupt already configured witnessing.

### What is the impact of a bogus list update?

A witness' previously applied configuration will not be affected.  In other
words, participating witnesses just look at the list and configure new logs.

This means the main attack vector is to inject new logs the maintainers did not
intend for.  At its worse, this may result in DoS for new logs.  For example,
the list could become full (too many logs / too high qps).  Recovery would
likely be an open discussion about what went wrong and how to correct it,
ultimately culminating in participating witnesses doing manual reconfiguration.

### What about bogus origin lines?

The maintainers *try* to verify that origin lines are reasonable.  Mistakes that
someone is unhappy about ("hey - that's my namespace") will surely happen
eventually.  However, as long as there are no origin-line collisions for real
logs the issue is quite insignificant (origin lines just need to be unique).

### Why is the list of logs not transparency logged?

Because detection of a bogus list is already pretty evident.  In other words, a
witness that configures the wrong list will either get complaints that logs are
not configured or eventually notice its performance budget is being overspent.

### Why is the list of logs not signed?

An attacker that can compromise the maintainer's distribution infrastructure
would also defeat any automatic pipeline signing.  This means each maintainer
would have to sign list updates manually for there to be much value.  While not
ruled out and doable, the overhead to manage these keys over time is likely
higher than just doing manual reconfiguration once a bogus list is detected.

In other words, we don't sign the lists because the story for authenticity is
good enough when trusting the distribution infrastructure and HTTPS.  This
assessment would be different if the impact of a bogus list was higher.  See
above why bogus lists are (by design) low impact and (by nature) easy to detect.
