# Freight Trust & Clearing Network

## Network & Community Forum

<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->

[![All Contributors](https://img.shields.io/badge/all_contributors-2-orange.svg?style=flat-square)](#contributors-)

<!-- ALL-CONTRIBUTORS-BADGE:END -->

## Overview

## Resources

### Identity & Node Management

Here is our guidelines using AWS best-practices

##### ARN Schema

arn:partition:service:region:account-id:resource-id
arn:partition:service:region:account-id:resource-type/resource-id
arn:partition:service:region:account-id:resource-type:resource-id

#### RegEx Formula

(python) --->

```
^arn:(?P<Partition>[^:\n]*):(?P<Service>[^:\n]*):(?P<Region>[^:\n]*):(?P<AccountID>[^:\n]*):(?P<Ignore>(?P<ResourceType>[^:\/\n]*)[:\/])?(?P<Resource>.*)$
```

A pattern to parse Amazon Web Services ARNs into their varying components:

Partition Service Region AccountID ResourceType (optional - empty string is
missing) Resource <----

##### RegEx Simple Formula

`^us-[a-z]*-[0-9]{1}`

[Use this regexer to get an idea of how it works](https://regexr.com/4prv4)

## Guidelines

- Pool Operators
  - Requirements
- Node Operators
  - Requirements

## Documentation Driven Development

There are many ways to drive open source development. Documenting the problem in
the README gives a middle ground between technical and non-technical
specifications. This allows organizing solutions to this challenge around
community and documentation.

> [...] a beautifully crafted library with no documentation is also damn near
> worthless. If your software solves the wrong problem or nobody can figure out
> how to use it, there’s something very bad going on.

- [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html)
  by Tom Preson-Werner

### Conventions and Specifications

Using conventions, documentation and specifications make it easier to:

- communicate the problem you are solving
- ease onboarding
- build and use composable tools
- promote open source contribution and engagement
- promote issue and feature discussion on Github itself

#### Resources

- [opensource.guide](https://opensource.guide/)
- [Github community profiles for public repositories](https://help.github.com/articles/about-community-profiles-for-public-repositories/)
- [Readme Driven Development](http://tom.preston-werner.com/2010/08/23/readme-driven-development.html)
- [pengwynn/flint](https://github.com/pengwynn/flint)
- [Working Backwards](https://www.allthingsdistributed.com/2006/11/working_backwards.html)
- [Literate programming](https://en.wikipedia.org/wiki/Literate_programming)
- [Hammock Driven Development](https://www.youtube.com/watch?v=f84n5oFoZBc)
- [Inversion and The Power of Avoiding Stupidity](https://fs.blog/2013/10/inversion/)
- [choosealicense.com](http://choosealicense.com)
- [The Documentation Compendium](https://github.com/kylelobo/The-Documentation-Compendium)

## Getting Started

To get started, [fork](https://help.github.com/articles/fork-a-repo/) or
[duplicate](https://help.github.com/articles/duplicating-a-repository/) the
repository. Then edit this file and delete everything above this line.

---

### Contributing

How to contribute, build and release are outlined in
[CONTRIBUTING.md](CONTRIBUTING.md), [BUILDING.md](BUILDING.md) and
[RELEASING.md](RELEASING.md) respectively. Commits in this repository follow the
[CONVENTIONAL_COMMITS.md](CONVENTIONAL_COMMITS.md) specification.

## Contributors ✨

Thanks goes to these wonderful people
([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/Wazzag99"><img src="https://avatars2.githubusercontent.com/u/33938431?v=4" width="100px;" alt=""/><br /><sub><b>Warren G</b></sub></a><br /><a href="#userTesting-Wazzag99" title="User Testing">📓</a></td>
    <td align="center"><a href="https://github.com/xcantera"><img src="https://avatars0.githubusercontent.com/u/34890623?v=4" width="100px;" alt=""/><br /><sub><b>Arturo Cantera Carrasco</b></sub></a><br /><a href="#userTesting-xcantera" title="User Testing">📓</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the
[all-contributors](https://github.com/all-contributors/all-contributors)
specification. Contributions of any kind welcome!
