# Participate

Submit your participation request by sending an email to:

    participate (at) lists.witness-network.org

Guidelines on what information to include are outlined below.  Refer to the
mailing list [archive][] for inspiration from previous participation requests.

The maintainers may process requests adhoc outside of the mailing list, e.g.,
when meeting the community on conferences and in various chat rooms like:

    #witness-network @ OFTC.net
    #witness-network @ matrix.org
    #witness-network @ transparency.dev/slack

[archive]: https://lists.witness-network.org/mailman3/hyperkitty/list/participate@lists.witness-network.org/

## Log guidelines

### Specify origin line

Your log is identified by a unique and never changing [origin line][].  We
recommend using a [schema-less URL][] to avoid accidental collisions.

Examples of origin lines:

  - `log.example.org`
  - `example.org/log`

[origin line]: https://github.com/C2SP/C2SP/blob/main/tlog-checkpoint.md#note-text
[schema-less URL]: https://github.com/C2SP/C2SP/blob/main/signed-note.md#signatures

### Specify public key

Your log signs [checkpoints][] before sending them to [witnesses][].  Specify
the public key in [vkey format][].  Use key type 0x01 (Ed25519).  For key name,
we recommend using the same string as the log's origin line.

Examples of public keys:

  - `log.example.org+0a72cd63+AV2yLmzeBL9B3zFJNqal0JuqxiFve0m/eqmT1+KkLVeC`
  - `example.org/log+b0201859+Aawb3XxZM7k072GrHtogNwVl3F5b5SbqFs83YYw+Yjbq`

[checkpoints]: https://C2SP.org/tlog-checkpoint
[witnesses]: https://C2SP.org/tlog-witness
[vkey format]: https://github.com/C2SP/C2SP/pull/119/files

### Specify add-checkpoint rate

How often will the log send [add-checkpoint][ac] requests to its witnesses?

Examples:

  - Every second
  - Every 10 seconds
  - At most 24 times per day, probably a lot less

Exceeding this limit may cause participating witnesses to rate limit or
completely block the log on sustained abuse.  The exact policy is up to the
participating witnesses---please consult their about pages for details.

A log with a small performance impact will be easier to get approved.

[ac]: https://github.com/C2SP/C2SP/blob/main/tlog-witness.md#add-checkpoint

### Specify log list

Which log list should the log be added to?  Choose between:

  - `testing`: Used for development and shorter-lived testing
  - `staging`: Used for longer-lived prototyping and dogfooding
  - `production`: Not available yet

Examples:

  - This is a dev log that will be shutdown in a while, testing list please
  - This is a real log with real usage, we're happy to dogfood staging

### Specify contact information

If an issue is detected, how can we or other operators contact you?

Examples of contact information:

  - sysadmin@example.org
  - <https://www.example.org/about>

## Provide other useful information (optional)

Can you think of anything else that will make the maintainers' decision of
approving the log easier?  If yes, please specify such information.

Examples of useful information could include the purpose of the log, who it
serves, remarks about the required load vs utility, expected lifetime, etc.

## Witness guidelines

### Specify an operator name

The name could for example be an organization or an individual.

Examples:

  - Example Company
  - John Doe

### Select log list(s)

Which [log list(s)][ll] will you be configuring?  Choose between testing,
staging, and production.  Select the largest performance profile you can
accommodate.

(A performance profile defines the number of logs your witness can keep
persistent state for, and how many add-checkpoint requests your witness can
handle for all logs combined on average.  Not applicable for testing lists.)

[ll]: /log-lists

### Automate discovery of logs in the selected list(s)

The log-list format is documented
[here](https://github.com/transparency-dev/witness-network/blob/main/log-list-format.md).

Ensure your witness automatically downloads and applies the list(s)
periodically.  Consider, e.g., hourly or daily reconfiguration; and use at most
a weekly cadence.

Configure *all new logs* that are discovered in the selected list(s).  A log is
considered new if its origin line is not known by your witness since before.

You *must not* remove or update an already configured log as a result of a
changed list.  This ensures the community maintainers cannot disrupt past
configurations, which makes them less juicy targets for attack.

In other words, a downloaded list *is not your configuration file*.  You need to
discover logs in the list, and then put them into your own configuration file.

(You can of course witness additional logs by configuring them manually, or use
other configuration communities that strike different trade-offs than ours.)

### Specify an about page URL

Where can a log operator find configuration details and learn more about how the
witness is operated?  Specify an about page URL.

The about page should at least include:

  - The witness public key in [vkey format][].  Use key type 0x04
    (cosignature/v1).  Similar to logs, we recommend using a schema-less URL for
    the name.
  - An [add-checkpoint URL][ac] (which may be referring to a
    [bastion-host][bh]).
  - Which list(s) the witness is being configured with.
  - How often the witness is reconfigured with logs from the latest list(s).

The about page should list additional information that's useful for the log
operator, such as whether it is a production witness, how it's operated, etc.

Please consult the about pages of participating witnesses to get inspiration and
compose something that strikes the right balance for your intended operations.

[bh]: https://C2SP.org/https-bastion
