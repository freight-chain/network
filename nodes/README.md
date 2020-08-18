# Node Policies.

## Overview.

The following `.json` files are used for enforcing access control and declarative permissions boundries between nodes and pools.
They are generated using AWS IAM policy generator.

## Policies

The policies are as follows.

#### network-diagram

defined network topology and routing _immutable_.

#### cluster-policy

defined restrictions _prohibit_.

#### service-policy.

defined permissions _allowed_.

#### master-node.

potential validator nodes.

#### service-node.

rpc/ws nodelets.

#### worker-node.

redis nodelets.
