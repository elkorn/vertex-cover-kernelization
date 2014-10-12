# Problems
## Blockers
- `findCrown` treats independent sets of single vertices as crowns with |I|=1 and |H|=0. This leads to trimming the whole graph in `Reducing` before the Chen,Kanj,Xia algorithm can be applied.
- The primal LP formulation cause all vertices to have assigned Xu=0.5, meaning that they may or may not be in the minimum vertex cover, which is inconclusive.

## Non-blockers
- network flow kernelization is much slower than it should be.