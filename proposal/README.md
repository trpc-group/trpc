English | [中文](README_zh_CN.md)

# tRPC RFCs

## Introduction

Before getting started, please familiarize yourself with the tRPC project through the [official website](https://trpc.group/).

This repository is mainly used to record proposals for major feature evolutions in tRPC.
It can:

- Provide upcoming major feature updates and future plans for tRPC.
- Record oversize features that are not suitable for issues.
- Standardize the development process for large features.
- Encourage open discussions based on design documents.

## Proposal Types

- `An` General features that affect all languages.
- `Ln` Language-specific changes.
- `Gn` Protocol layer changes.
- `Pn` Changes to this process itself.

Use 4 labels to record the maximum tRPC number.
When assigning a number to a new proposal, take the corresponding label number, and then increment the label number by one.

## Process

1. Create an issue to determine the tRPC number for the feature.
2. Clone the repo, copy the TEMPLATE.md, and rename it to $Type-$Summary.
For proposals specific to a programming language, insert the language name in the name as $Type-$Language-$Summary.
3. Submit an MR. The title of the MR should include the tRPC number.
4. After the OWNER completes the first draft of the proposal, the review process can be initiated.
5. Invite as many people as possible for the review, collect feedback, and update the proposal. 
Choose the appropriate discussion method:
If the proposal has a wide impact (e.g., An type proposals), initiate a discussion in the corresponding issue and update the discussion link in the proposal.
If the proposal is related to implementation details (e.g., `Ln` type proposals), discussion can be limited to the MR.
6. Once the proposal is approved, merge the MR. 
To ensure a clean history, merging must be done using the rebase + squash method.
Commit messages must start with the tRPC number.
7. Updates to accepted proposals should indicate the tRPC number in the MR.
Depending on the extent of the changes, decide whether a new discussion is needed.