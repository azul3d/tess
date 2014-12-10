====
Tess

A Go port of [libtess.js](https://github.com/brendankenny/libtess.js).

===============
Developer Notes

- GluMesh.checkMesh -> GluMesh.Check
- Do not create types yourself! Use for example:
  - NewGluFace
  - NewGluVertex
  - NewGluHalfEdge
  - NewDictNode
- DictNode
  - getKey is just n.Key
  - getSuccessor is just n.Next
  - getPredecessor is just n.Prev

Left to port:

- `src/`
  - `normal.js`
  - `libtess.js`
  - `sweep.js`
  - `geom.js`
  - `tessmono.js`
  - `render.js`
  - `mesh.js`
  - `priorityq/`
    - `PQHandleElem.js`
    - `PriorityQ.js`
    - `PriorityQHeap.js`
  - `mesh/`
    - `GluMesh.js` (needs work)
  - `libtess/`
    - `GluTesselator.js`
    - `CachedVertex.js`

