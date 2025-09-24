**Warning:** experimental / work-in-progress.  Please provide feedback.

Finding and configuring witnesses for your transparency log can be slow and
repetitive: log operators need to track down witnesses and exchange details, and
witness operators need to continuously decide on which new logs to cosign.

*Witness Configuration* removes this friction through a shared, community
vetted, and automated configuration process.  Join our growing community with a
single request to avoid configuration overhead proportional to `logs ×
witnesses`.

## Why a community for witness configuration?

In the [witness protocol][wp], logs synchronously ask witnesses to cosign each
new log state.  A log sends a [checkpoint][cp] to a witness, which responds with
a [cosignature][cs] if the new log state only contains appended entries (no
removals or modifications).

    cosignature ┌─────────────────────┐ cosignature
    200 OK ┌──> │         LOG         │ <──┐ 200 OK
           │    └─────────────────────┘    │
           │           │       │           │
           │           │       │           │
           │   HTTP POST /add-checkpoint   │
           │           │       │           │
           │           v       v           │
        ┌────────────────┐   ┌────────────────┐
        │  WITNESS: FOO  │   │  WITNESS: BAR  │
        │ ┌────────────┐ │ . │ ┌────────────┐ │
        │ │  verify    │ │ . │ │  verify    │ │
        │ │append-only │ │ . │ │append-only │ │
        │ └────────────┘ │   │ └────────────┘ │
        └────────────────┘   └────────────────┘

A log that collects cosignatures from multiple witnesses can convince users that
they see the same log, much like collecting testimony from several witnesses in
court can help convince the judge about a single consistent timeline of events.

For any of the above to work:

  1. Logs must know which witnesses to contact---it's a matter of configuration.
  2. Witnesses must know which logs to cosign---also a matter of configuration.

In today's emerging ecosystem, each log individually finds and persuades a set
of witnesses to apply its configuration details.  This is a slow and repetitive
process for logs and witnesses, creating an unfortunate barrier to
participation.

*Witness Configuration* is a community service that brings logs and witnesses
together, replacing many separate requests with one shared configuration step.

[wp]: https://C2SP.org/tlog-witness
[cp]: https://C2SP.org/tlog-checkpoint
[cs]: https://C2SP.org/tlog-cosignature

## How does it work?

A few community members maintain:

  - A list of approved logs (in fact, multiple lists for
    testing/staging/production).
  - A table of participating witnesses (categorized by
    testing/staging/production)

Participating witnesses periodically configure *all new logs* from the list(s)
they selected.  In other words, the participating witnesses depend on the
community maintainers' review and approval to automatically decide on new logs.

    Step 1: Ask maintainers       Step 2: Update log list
    ┌──────────────────────┐      ┌───────────────────────────────┐
    │   LOG: jellyfish     │      │ Approved logs                 │
    │   (log operator)     │      │  - jellyfish: public key, ... │
    └──────────────────────┘      │  - ...                        │
               │ motivation       └───────────────────────────────┘
               │ public key             ^                   │
               │    ...                 │                   │
               v                        │                   │
      ┌─────────────────┐               │                   │
      │    COMMUNITY    │               │                   │
      │   MAINTAINERS   │               │                   │
      │ ┌─────────────┐ │       add log │                   │
      │ │ review log  │ │ ──────────────┘                   │
      │ │ application │ │                                   │
      │ └─────────────┘ │                                   │
      └─────────────────┘                                   │
                                                            │
    Step 3: auto configuration                              │
      ┌──────────────────┐                                  │
      │     WITNESS      │                                  │
      │ ┌──────────────┐ │                    periodic pull │
      │ │ extract new  │ │ <────────────────────────────────┘
      │ │ logs; reconf │ │
      │ └──────────────┘ │
      └──────────────────┘

The process is the same to become a participating witness, except that the
maintainers populate human-readable tables instead of machine-readable lists.

Approved log operators can now choose which participating witnesses to request
cosignatures from---notably without ever contacting them individually.

## What is needed to become a participant?

Other than using the interoperable [witness protocol][wp], a community
maintainer needs to approve your log or witness.  Head over to the
[participation page][pp].

[pp]: ./participate
