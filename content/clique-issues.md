# Clique Issues

> source: https://docs.google.com/document/d/1pIW6Uac5Qanx_L5Y_G4Ucrx_gjITkBgzQKmlBQbW3Cg/edit

## Problem

The Görli network has on a number of occasions forked and become “stuck” on different forks both of which have the same total difficulty, so clients don’t re-org to the other side of the fork.

## Missing In-Turn Blocks

It’s believed that there are cases where an in-turn block is missed even though the validator is online. Theoretically out of turn blocks should be published with a delay after the in-turn block should have been published, however the delay is quite small. Clock drift, malicious actors or buggy code may cause an out-of-turn validator to publish a new block sooner than expected. Alternatively the in-turn block could be slightly delayed due to clock drift on the in-turn validator or some performance problem causing a slight delay.

However it happens, if the in-turn validator receives the out-of-turn block before it publishes, it will import it into it’s chain and consider it the new canonical head (it’s the best block because the in-turn block doesn’t yet exist). Typically getting a new chain head would cause the miner to stop mining it’s current block and start a new one based on the new chain head. This works well for PoW because the new block has higher difficulty and the best chance of being included in the canonical chain, however it’s the wrong behaviour for Clique because the new block will be out of turn whereas the original would have been in-turn and become the highest difficulty.

I’m not sure which clients exhibit this behaviour of cancelling the in-turn block creation but Pantheon definitely does and it looks like Geth does too (based on miner/miner.go line 86).

This behaviour can lead to the network forking. Imagine a network with four validators A, B, C and D. The chain is at block 5 with total difficulty of 10 and block 6 should be produced by A. Ideally we should have block 6, produced by A with total difficulty of 12.

However, if B publishes an out-of turn block to A before it completes producing its block and A cancels creation of its block. We then have:

A: height 6, TD 11
B: height 6, TD 11
C: height 5, TD 10
D: height 5, TD 10

C took longer to receive B’s block so after a delay publishes an out of turn block giving:
A: height 6, TD 11 (from B)
B: height 6, TD 11 (from B)
C: height 6, TD 11 (from C)
D: height 6, TD 11 (from D)

B should publish at height 7 but can’t because on its canonical chain it published block 6 and it doesn’t consider the other chain canonical so both chains will have to use an out of turn block for height 7 and thus have a TD of 12. With the validators evenly split the network is in a very bad state and will find it hard to recover.

Q: Shouldn’t this resolve itself since eventually a block from the other fork should arrive first and given it has the higher TD it would then become canonical head? This fork should be a reasonably unstable state…

If A had published it’s in turn block, it would have had the highest total difficulty and the problem would be avoided.

### Proposal:

We should clarify the Clique spec to indicate that a validator should produce it’s in turn block even if it receives an out of turn block for that height (but not if the chain has progressed more than that one block?) There is work in progress now to apply this to Pantheon.

#### Out of Turn Delay Shorter Than Network Latency

The number of out of turn blocks in the network is significantly higher than necessary because the delay before publishing an out of turn block is randomly selected up to a fairly small maximum delay. Whenever the randomly chosen delay is less than the block propagation time from the in-turn validator an unnecessary out-of-turn block will be published. At minimum this wastes resources. Potentially it could be contributing to the network forking.

If we have four validators, A, B, C and D. If A & B are in Australia and C & D are in Europe then A & B will have low latency to each other but high latency to C & D. Thus, if out of turn blocks are produced by A and C at similar times and the same total difficulty, B will receive A’s block first and treat it as canonical and D will receive C’s block first. As they have the same total difficulty it doesn’t cause a re-org so the network remains forked. This may answer the question posed above about why random timing of received blocks doesn’t eventually cause the network to reach consensus.

#### Proposal:

Add two configurable options for the Clique algorithm:
OUT_OF_TURN_DELAY_MULTIPLIER
Currently explicitly 500ms in the spec
MIN_OUT_OF_TURN_DELAY
New concept so Rinkeby and Görli are effectively using 0ms
Absolute minimum amount of time to wait before publishing an out of turn block
Delay calculation is then:
MIN_OUT_OF_TURN_DELAY +
rand(SIGNER_COUNT \* OUT_OF_TURN_DELAY_MULTIPLER)

On low latency networks the current settings may be enough but a geographically dispersed network like Görli may increase both options to ensure the out of turn delay exceeds the network latency and out of turn blocks aren’t created unnecessarily.
