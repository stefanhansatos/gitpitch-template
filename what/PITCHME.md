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

- Associate anything to a binary structure |
- The binary structure consists of two parts |
    - one represents information (content) |
    - one represents relevance / interest (mask) |
- A double stranded line of bits (context)  |

  


---?image=assets/image/contextingTree1.png&size=contain&transition=none

---?image=assets/image/contextingTree2.png&size=contain&transition=none

---?image=assets/image/contextingTree3.png&size=contain&transition=none

---?image=assets/image/contextingTree4.png&size=contain&transition=none

---?image=assets/image/contextingTree5.png&size=contain&transition=none

---?image=assets/image/contextingTree6.png&size=contain&transition=none

---?image=assets/image/contextingTree7.png&size=contain&transition=none

---?image=assets/image/contextingTree8.png&size=contain&transition=none

---?image=assets/image/contextingTree9.png&size=contain&transition=none

---

### Kinds of context

- standards, e.g. GPS |
- taxonomies, e.g. languages |
- dictionaries, e.g. tags, categories |
- all, binary representable, i.e. everything |
- hashes, ai, blockchain |


