---
order: 7
---

# Plug Chain Hub Upgrade, Instructions

This document describes the steps for validator and full node operators for the successful execution of the *upgrade plan*. 

TOC:
- [Preparing for the upgrade](#preparing-for-the-upgrade)
  - [Backups](#backups)
  - [Testing](#testing)
- [Upgrade process](./upgrade-process.md)
- [Rollback plan](#rollback-plan)
- [Risks](#risks)


## Preparing for the upgrade

### Backups

Prior to the upgrade, validators are encouraged to take a full data snapshot. Snapshotting depends heavily on infrastructure, but generally this can be done by backing up the `.plugchain` directory.

It is critically important for validator operators to back-up the `.plugchain/data/priv_validator_state.json` file after stopping the plugchaind process. This file is updated every block as your validator participates in consensus rounds. It is a critical file needed to prevent double-signing, in case the upgrade fails and the previous chain needs to be restarted.

### Testing

For those validator and full node operators that are interested in ensuring preparedness for the impending upgrade , use plugchain-tesetnet-1 network to test upgrade.

## Upgrade duration

The upgrade may take several hours to complete because Plug Chain Hub participants operate globally with differing operating hours and it may take some time for operators to upgrade their binaries and connect to the network.

## Rollback plan

During the network upgrade, core Plug Chain teams will be keeping an ever vigilant eye and communicating with operators on the status of their upgrades. During this time, the core teams will listen to operator needs to determine if the upgrade is experiencing unintended challenges. In the event of unexpected challenges, the core teams, after conferring with operators and attaining social consensus, may choose to declare that the upgrade will be skipped. 

Steps to skip this upgrade proposal are simply to resume the Plug Chain Hub network with the (downgraded) v0.5.0 binary using the following command:

> plugchaind start --unsafe-skip-upgrade 762880

Note: There is no particular need to restore a state snapshot prior to the upgrade height, unless specifically directed by core Plug Chain teams.

Important: A social consensus decision to skip the upgrade will be based solely on technical merits, thereby respecting and maintaining the decentralized governance process of the upgrade proposal's successful YES vote.

## Communications

Operators are encouraged to join the `#validators-verified` channel of the Plug Chain Community Discord. This channel is the primary communication tool for operators to ask questions, report upgrade status, report technical issues, and to build social consensus should the need arise. This channel is restricted to known operators and requires verification beforehand - requests to join the `#validators-verified` channel can be sent to the `#validators-public` channel.  

## Risks

As a validator performing the upgrade procedure on your consensus nodes carries a heightened risk of double-signing and being slashed. The most important piece of this procedure is verifying your software version and genesis file hash before starting your validator and signing.

The riskiest thing a validator can do is discover that they made a mistake and repeat the upgrade procedure again during the network startup. If you discover a mistake in the process, the best thing to do is wait for the network to start before correcting it. 