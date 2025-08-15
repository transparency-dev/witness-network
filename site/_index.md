**Warning:** experimental work-in-progress service.

Trying to discover witnesses for your transparency log?  Or are you a witness
operator that wants to spend less time on day-to-day configuration overhead?

You've come to the right place.  *Witness Configuration* is a community effort
that helps operators of transparency logs and witnesses to collaborate with
ease.

New to witnessing? Read this [introductory post][ip] to onboard yourself.

[ip]: TODO

## Why a community for witness configuration?

In the [witness protocol][wp], logs synchronously ask witnesses to cosign every
new log state.  This log state is encoded as a [checkpoint][cp], and the request
is transferred to the witness using an HTTP POST endpoint.  A witness returns a
[cosignature][cs] for a new log state if it's append-only compared to all prior
observations of the log.

    TODO: figure.

For the witness protocol to function, the logs that initiate requests need to be
configured with which witnesses to contact.  Similarly, each contacted witness
needs to be configured with which logs to cosign, e.g., including public log
keys.  Even if public log keys could be discoverer automatically, it would
neither scale nor be light-weight to cosign and allocate storage for *all*
incoming checkpoints.

This lends itself towards a setup where each log needs to find and convince a
group of witnesses to configure it before any cosignatures can be collected,
thus creating an unfortunate barrier of overhead to be in the witnessing
ecosystem.

*Witness Configuration* is a community of logs and witnesses that work together.
The overhead to be in the witnessing ecosystem is reduced to a single request.

[wp]: https://C2SP.org/tlog-witness
[cp]: https://C2SP.org/tlog-checkpoint
[cs]: https://C2SP.org/tlog-cosignature

## How does it work?

A few community members maintain a *list of accepted logs* and a *table of
participating witnesses*.  These community members are themselves operators.

The participating witnesses periodically configure *all new logs* that appear in
the list of logs.  This means the participating witnesses depend on the
community maintainers to make a good selection of which logs to include and
exclude.

Nothing prevents a participating witness from configuring additional logs, or
even depending on another complementary witness configuration community.

An operator that got its log included in the list can select which of the
participating witnesses to collect cosignatures from.  Notably, the operator
can make this selection without ever having contacted each witness individually.

[log-list]: /log-list
[witness-table]: /witness-table

## What is needed to participate?
