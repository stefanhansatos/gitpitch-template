### A Group Membership Solution

- Scalable |
- Weakly-consistent |
- Infection-style |
- Process Group Membership Protocol |
- SWIM++ implemented in Go by HashiCorp |
- \+ Piggybacking |
- \+ Suspicion Mechanism, ... |

---
### Addressing Categories
<br>
- Location Addressing |
- Content Addressing |
- Context Addressing |

+++
### Network of Isomorphic Applications <br>
### Operating on Peer Devices

---
### Distributed Hash Table

Core component in distributed systems like

- Distributed VCS (Git)
- Blockchain |
- Peer-Too-Peer Infrastructure |

---
### Content Addressing Solutions
<br>
- binary tree of common size |
- XOR-Metric |
- Kademlia (used by Ethereum, IPFS, ...) |

---

![Kademlia](assets/image/kademlia.png)

--- 
### Context Addressing 

- Context Functions instead of hash functions |
- Distributed Context Tables instead of DHT |
- Context Engineering <br>    i.e. creating individual context functions |


--- 
### Context Functions 

- Functions which can associate potentially anything <br>    to a certain kind of binary structure |
- The binary structure consists of two parts |
    - one represents information
    - one represents relevance or interest
- As a binary tree
    - one represents a subtree (relevance)
    - one represents a leaf in it (information)
- Another representation is <br>    a double stranded line of bits  |
  


