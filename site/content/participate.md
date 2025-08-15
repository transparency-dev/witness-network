Submit your participation request by sending an email to:

    TODO: participate (at) lists.example.org    

Guidelines on what information to include are outlined below.  Refer to the
mailing list [archive][] for inspiration from previous participation requests.

The maintainers may process requests adhoc outside of the mailing list, e.g.,
when meeting the community on conferences and in various chat rooms.

[archive]: TODO

## Log guidelines

### Specify origin line

Your log is identified by a unique and never changing [origin line][].  We
recommend using a [schema-less URL][] to avoid accidental collisions.

Examples of origin lines:

  - `log.example.org`
  - `example.org/log`

[origin line]: https://github.com/C2SP/C2SP/blob/main/tlog-checkpoint.md#note-text
[schema-less URL]: https://github.com/C2SP/C2SP/blob/main/signed-note.md#signatures

## Specify public key

Your log signs [checkpoints][] before sending them to [witnesses][].  Specify
the public key in [vkey format][].  Use key type 0x01 (Ed25519).  For key name,
we recommend using the same string as the log's origin line.

Examples of public keys:

  - `log.example.org+0a72cd63+AV2yLmzeBL9B3zFJNqal0JuqxiFve0m/eqmT1+KkLVeC`
  - `example.org/log+b0201859+Aawb3XxZM7k072GrHtogNwVl3F5b5SbqFs83YYw+Yjbq`

[checkpoints]: https://C2SP.org/tlog-checkpoint
[witnesses]: https://C2SP.org/tlog-witness
[vkey format]: https://github.com/C2SP/C2SP/pull/119/files

## Specify add-checkpoint rate

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

## Specify contact information

If an issue is detected, how can we or other operators contact you?

Examples of contact information:

  - sysadmin@example.org
  - <https://www.example.org/about>

## Specify bastion host (optional)

Do you have a [bastion host][bh] that participating witnesses without a public
[add-checkpoint][ac] endpoint can connect to?  If yes, specify it.

Example:

  - <https://bastion.example.org/>

[bh]: https://C2SP.org/https-bastion

## Provide other useful information (optional)

Can you think of anything else that will make the maintainers' decision of
approving the log easier?  If yes, please specify such information.

Examples of useful information could include the purpose of the log, who it
serves, remarks about the required load vs utility, expected lifetime, etc.

## Witness guidelines

### Automate discovery of logs in the community list

The format of the [list of logs](../log-list-10qps-100klogs) is documented
[here](https://git.glasklar.is/rgdd/witness-configuration-network/-/blob/main/docs/log-list-format.md).

Ensure you're able to accommodate at least 10 [add-checkpoint][ac] requests/s on
average from all logs combined with enough persistent storage to support at
least 100,000 logs.  In other words, this is the *performance profile* that the
maintainers assume you're able to handle to maintain the log list reliably.

Ensure your witness automatically downloads the list periodically.  The exact
interval is not important, e.g., hourly, daily, or weekly.  Configure *all new
logs* that are discovered in the list.  A log is new if its origin
line is not known since before.

You *must not* remove or update an already configured log as a result of a
changed community list.  This ensures the community maintainers cannot disrupt
past configurations, which makes them less juicy targets for attack.

(You can of course witness additional logs by configuring them manually, or use
other configuration communities that strike different trade-offs than ours.)

### Specify an about page URL

Where can a log operator find configuration details and learn more about how the
witness is operated?  Specify an about page URL.

The about page must at minimum specify:

  - The witness public key in [vkey format][].  Use key type 0x04
    (cosignature/v1).  Similar to logs, we recommend using a schema-less URL for
    the name.
  - An [add-checkpoint URL][ac] (unless it is a [bastion-host][bh] only witness).
  - How often the witness is reconfigured with logs from the latest list.

The about page should list additional information that's useful for the log
operator, such as whether it is a production witness, how it's operated, etc.

Please consult the about pages of participating witnesses to get inspiration and
compose something that strikes the right balance for your intended operations.

### Specify an operator name

The name could for example be an organization or an individual.

Examples:

  - Example Company AB
  - John Doe

[wt]: /witness-table
