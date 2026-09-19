# Proposal: Dynamic witness capacity allocation

This is a proposal I made to increase witness utilization, originally posted on
Matrix. I also had some in-person discussions about this with rgdd, notes on
those below.

## Original proposal

Before the witness network, the process for a log operator for finding witness
was:

 (1) Find out that a witness/witness operator (conflated for now) exists
 (2) Vet the witness
 (3) Ask the witness to configure the log and wait for that to happen

(or more likely: Ask somebody to set up a witness for them)

The witness operator perspective mostly comes in at step 3 where they need to
decide if they can (and want, but let's suppose this is always the case for
now) support the log. This is not a trivial "yes" because there always will be
capacity limits to consider.

The witness network helps with 1 and 3 by providing a list of witnesses and
managing their configuration. This means that for the log operator, they only
need to talk to the witness network maintainers and once their log gets
accepted, they can have their pick from any witness that configures the
respective log lists (after a short while once they have updated them).

This currently works by the witness network maintainers packing logs into
specific "performance tiers" defined by the maximum number of logs and
witnessing requests per second (qps) a given witness can support. Witnesses are
supposed to configure log lists starting from the lowest performance profile up
to the highest one that they can still support in aggregate (i.e. including the
lower tiers).

In the following, we'll mostly consider the qps dimension of resource usage,
since this is the much tighter bottleneck in practice.

The way logs get assigned to performance profiles is currently not documented,
but seems to follow a strategy that tries to minimize qps utilization per
performance profile (i.e. even though the 10qps list would have capacity to
accommodate a 1qps log, 1qps logs get allocated to the 100qps list first since
adding a 1qps log to the former would use up 10% of its capacity).

There are a number of problems with the current approach:

 - Logs will only make use of a subset of witnesses available to them, but the
   log list capacity planning can't know which, so it must happen based on
   worst case (all logs in a list fully utilize all witnesses available to
   them), which leads to underutilization within a given performance profile.

 - Similarly, a witness must assume that all logs from the lists it configures
   will make use of it, so it configures less log lists than it could
   actually handle in practice, also leading to underutilization and less
   witnesses being available to logs in higher profiles.

 - On top, this scheme requires picking somewhat arbitrary bucket sizes for the
   performance profiles.

The core issue here is that the witness network itself tries to do capacity
planning for the witness operators and tries to do so for all of them
simultaneously. Furthermore, it does this in advance, without knowing how log
operators will make use of witnesses.

An alternative approach could be to remove this capacity planning component
from the witness network and only have it be a place where witnesses and logs
can advertise their existence. To that end, consider the following architecture:

 - There is only one list/pool of logs (containing the same data as
   today, including estimated qps). When a log is retired it is marked as
   inactive. (Could also be removed from the list, but the following description
   is clearer that way)

 - As long as they are below capacity, witnesses keep importing *all* logs in
   this list. They also provide an interface where log operators can query if
   their log is configured (i.e. the witness would accept add-checkpoints
   requests for it). A log will be advertised as supported as long as
   activating that individual log would not exceed that witness's capacity.

 - Witnesses keep track of utilized qps. They do this by aggregating the
   advertised qps from the log list over all the logs that have sent them at
   least one checkpoint and are not marked as inactive in the log list.

 - Once a witness has reached its locally-configured qps limit, it stops
   advertising support for/accepting checkpoints from logs from which it
   hasn't received any checkpoints yet.

 - It still keeps updating the log list to see if logs have been marked as
   inactive, which might free up capacity if one such log has previously been
   active on this witness. In that case, it starts advertising logs again.

This resolves the problems described above:

 - If a log does not decide to make use of a witness for one reason or another,
   the capacity for that log is not needlessly reserved on that witness.

 - All witnesses with spare capacity are available to all logs.

 - Witness operators have fine-grained control over the capacity of their
   witness and witnesses can reach closer to 100% utilization.

When a witness is at capacity, a witness operator can easily deploy another
witness which will start picking up new/different logs (since a single log is
unlikely to use multiple witnesses from the same operator unless for
redundancy).

But other than increased complexity, there are also some further downsides:

 - Unlike today, a log can't be certain it will get picked up by the witnesses
   it likes (or possibly any witness) if it has been accepted into the network.
   Thus, a log operator needs to query individual witnesses to see if they have
   picked up the log. But to some extent this already the case today since log
   list downloads might only happen weekly for example.

 - Since logs decide which witnesses they claim, ecosystem diversity can be
   affected by log choices. I.e. a 1qps log takes a hypothetical 1qps witness
   fully out of the ecosystem, but it would likely be better for resilience
   to partition that same witness among 10 0.1qps logs.

 - There is also the potential for race conditions. I.e. a log operator looks
   at all the witnesses with spare capacity, carefully vets them and a few
   seconds before they start making use of it, somebody else claims all spare
   capacity of that witness.

## Discussion notes

Discussion based on this with rgdd during the 2025-05 Sigsum community meeting.
I'm writing this from memory a few days later so it's probably a bit
inaccurate.

 - The main issue with managing the witness network is that we're dealing with
   a scarce resource (witness capacity), if every witness could do 500 qps and
   had unlimited storage, all witnesses could just be free-for-all and we
   wouldn't have to have these discussions. However, this is not the case.

 - The goal of the witness network is not only to help coordinate between logs
   and witnesses, but also manage this scarce resource in a thoughtful way.
   This is the part I wasn't aware of as being a deliberate decision, which
   invalidates the above proposal to an extent.

   At the same time, the witness network aims to serve the long tail of logs.
   The assumption is that heavy hitters such as MTC will curate their own set
   of witnesses for policy reasons anyway.

   So this fits together well - for example an average Sigsum log does 0.1qps.
   Others might do even less or only produce a checkpoint sporadically. Thus
   even 10qps of capacity could serve a lot of logs (possibly *all* of the
   long-tail ones).

 - A part of the awkwardness is that the witness network maintainers do not
   want to be in the position to be able to DoS logs. Thus, they deliberately
   aren't able to cause deconfiguration of logs. However, logs (especially
   things like CT logs) retire frequently. Ideally, logs would be able to
   signal this to the world (and witnesses in particular) cryptographically,
   but the proposed mechanism (tombstones) has not been fleshed out yet.

   Putting this off was "fine" since there was enough spare capacity as not to
   have to worry about this now.

 - However, CT logs getting started to be admitted to the witness network
   compounded this issue and prompted the creation of the 100k log list (which
   in turn prompted the above proposal). Maybe creating such a big list was a
   mistake.

 - Maybe eventually having multiple 10qps lists (maybe grouped somehow so that
   witness operators can choose which parts of the ecosystem to support) would
   be better. This would also help with better bin-packing on the witness side.

   Probably leaving the 100k CT log list as is though?

Further notes added by rgdd:

 - Seems like 100qps might have been an unnecessarily big jump, which, e.g.,
   have made it difficult for some (potential) witness operators to configure
   it.

 - When doing some napkin math, the current 10qps list would likely be able to
   accomodate a lot of the "long tail" / lower-frequency logs; and perhaps one
   or two high-profile ones with higher qps like Go's checksum database.

 - From CT, we're probably expecting something like 10 qps.

 - From MTC, we're probably talking about a qps in the same ballpark (?)

 - We don't have that many other high-qps logs right now, and having something
   like 10qps reserved for that will probably serve us well for some time.

 - So if it increases the number of participating witnesses, then it might be a
   better trade-off to have several 10qps lists (.2, .3) where we basically
   have one which is the "longer tail one" and another which is the "higher-qps
   one". And the "higher qps-one" we expect to fill up a bit quicker, and when
   it's full we will create another one. Or maybe we should even create multipe
   ones right away, and witnesses configure as many as they can even though,
   e.g., .3 is not being populated quite yet?

 - Working on defining tombstone for proper deallocation = worth while to do
   soon since CT is interested in taking part (and sharding is frequent there).
