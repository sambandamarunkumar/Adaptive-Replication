
# adaptive-replication
**Adaptive Replication for Low Latency Distributed Clusters**

### Paper Information
- **Author(s):** Arunkumar Sambandam
- **Published In:** *********************************************International Journal of Leading Research Publication (IJLRP)
- **Publication Date:** ******Aug 2021
- **ISSN:** E-ISSN: **********2582-8010
- **DOI:**
- **Impact Factor:** *******9.56

### Abstract
This paper addresses the limitations of static data replication policies in distributed systems operating under heterogeneous networks and dynamic workloads. It proposes an adaptive replication strategy that leverages real-time telemetry, including network latency, access patterns, and I/O behavior, to optimize replica placement. By dynamically relocating replicas closer to demand, the approach reduces cross-node communication, read latency, and tail delays. The framework provides a practical, telemetry-driven architecture for improving performance and locality in modern distributed clusters.

### Key Contributions
- **Adaptive Replication Framework:**
  Proposed a telemetry-driven replication architecture that dynamically adjusts replica placement to reduce latency under changing network and workload conditions.
  
- **Latency-Aware Placement:**
  Introduced a decision model that uses network latency, access patterns, and I/O signals to select optimal replica locations across distributed clusters.
    
- **Dynamic Replica Migration:**
  Developed a systematic method to relocate replicas closer to demand while maintaining consistency and fault tolerance.
   
- **End-to-End Validation:**
  Designed, implemented, and evaluated the approach using multi-node clusters, demonstrating consistent latency reductions over baseline replication.
  
### Relevance & Real-World Impact
- **Lower Replication Latency:**
  Achieved substantial and consistent reductions in replication and data access latency across varying cluster sizes.
   
- **Improved Data Locality:**
Enhanced read performance by dynamically positioning replicas closer to frequently accessing nodes.

- **Scalable Performance:**
    Maintained low latency as cluster size increased despite growing network and coordination complexity.
  
  **Adaptive Stability:**
  Improved system responsiveness by continuously adapting to network and I/O dynamics.
   
- **Production Ready:**
    Delivered a lightweight, practical framework suitable for cloud deployments, research, and advanced systems education.

### Experimental Results (Summary)

  | Nodes | Baseline Replication latency (ms) | Adaptive Replication latency (ms) | Improvment (%)  |
  |-------|-----------------------------------| ----------------------------------| ----------------|
  | 3     |  1400                             | 620                               | 55.71           |
  | 5     |  1330                             | 560                               | 57.89           |
  | 7     |  1250                             | 515                               | 58.80           |
  | 9     |  1200                             | 490                               | 59.17           |
  | 11    |  1160                             | 470                               | 59.48           |

### Citation
Adaptive Replication for Low Latency Distributed Clusters
* Arunkumar Sambandam
* ***********************************International Journal of Leading Research Publication 
* ISSN E-ISSN: *****************************2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijlrp.com*****************/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/**** | **Email**: arunkumar.sambandam@yahoo.com






