[简体中文](https://github.com/khulnasoft/nxgit/blob/master/README_ZH.md)

# Nxgit - Git with a cup of tea

[![Build Status](https://drone.nxgit.io/api/badges/khulnasoft/nxgit/status.svg)](https://drone.nxgit.io/khulnasoft/nxgit)
[![Join the Discord chat at https://discord.gg/NsatcWJ](https://img.shields.io/discord/322538954119184384.svg)](https://discord.gg/NsatcWJ)
[![](https://images.microbadger.com/badges/image/nxgit/nxgit.svg)](https://microbadger.com/images/nxgit/nxgit "Get your own image badge on microbadger.com")
[![codecov](https://codecov.io/gh/khulnasoft/nxgit/branch/master/graph/badge.svg)](https://codecov.io/gh/khulnasoft/nxgit)
[![Go Report Card](https://goreportcard.com/badge/go.khulnasoft.com/nxgit)](https://goreportcard.com/report/go.khulnasoft.com/nxgit)
[![GoDoc](https://godoc.org/go.khulnasoft.com/nxgit?status.svg)](https://godoc.org/go.khulnasoft.com/nxgit)
[![GitHub release](https://img.shields.io/github/release/khulnasoft/nxgit.svg)](https://github.com/khulnasoft/nxgit/releases/latest)
[![Help Contribute to Open Source](https://www.codetriage.com/khulnasoft/nxgit/badges/users.svg)](https://www.codetriage.com/khulnasoft/nxgit)
[![Become a backer/sponsor of nxgit](https://opencollective.com/nxgit/tiers/backer/badge.svg?label=backer&color=brightgreen)](https://opencollective.com/nxgit)

## Purpose

The goal of this project is to make the easiest, fastest, and most
painless way of setting up a self-hosted Git service.
Using Go, this can be done with an independent binary distribution across
**all platforms** which Go supports, including Linux, macOS, and Windows
on x86, amd64, ARM and PowerPC architectures.
Want to try it before doing anything else?
Do it [with the online demo](https://try.nxgit.io/)!
This project has been
[forked](https://blog.nxgit.io/2016/12/welcome-to-nxgit/) from
[Gogs](https://gogs.io) since 2016.11 but changed a lot.

## Building

From the root of the source tree, run:

    TAGS="bindata" make generate all

More info: https://docs.nxgit.io/en-us/install-from-source/

## Using

    ./nxgit web

NOTE: If you're interested in using our APIs, we have experimental
support with [documentation](https://try.nxgit.io/api/swagger).

## Contributing

Expected workflow is: Fork -> Patch -> Push -> Pull Request

NOTES:

1. **YOU MUST READ THE [CONTRIBUTORS GUIDE](CONTRIBUTING.md) BEFORE STARTING TO WORK ON A PULL REQUEST.**
2. If you have found a vulnerability in the project, please write privately to **security@nxgit.io**. Thanks!

## Further information

For more information and instructions about how to install Nxgit, please look
at our [documentation](https://docs.nxgit.io/en-us/). If you have questions
that are not covered by the documentation, you can get in contact with us on
our [Discord server](https://discord.gg/NsatcWJ),
or [forum](https://discourse.nxgit.io/)!

## Authors

* [Maintainers](https://github.com/orgs/khulnasoft/people)
* [Contributors](https://github.com/khulnasoft/nxgit/graphs/contributors)
* [Translators](options/locale/TRANSLATORS)

## Backers

Thank you to all our backers! 🙏 [[Become a backer](https://opencollective.com/nxgit#backer)]

<a href="https://opencollective.com/nxgit#backers" target="_blank"><img src="https://opencollective.com/nxgit/backers.svg?width=890"></a>

## Sponsors

Support this project by becoming a sponsor. Your logo will show up here with a link to your website. [[Become a sponsor](https://opencollective.com/nxgit#sponsor)]

<a href="https://opencollective.com/nxgit/sponsor/0/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/0/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/1/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/1/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/2/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/2/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/3/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/3/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/4/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/4/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/5/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/5/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/6/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/6/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/7/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/7/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/8/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/8/avatar.svg"></a>
<a href="https://opencollective.com/nxgit/sponsor/9/website" target="_blank"><img src="https://opencollective.com/nxgit/sponsor/9/avatar.svg"></a>

## FAQ

**How do you pronounce Nxgit?**

Nxgit is pronounced [/ɡɪ’ti:/](https://youtu.be/EM71-2uDAoY) as in "gi-tea" with a hard g.

**Why is this not hosted on a Nxgit instance?**

We're [working on it](https://github.com/khulnasoft/nxgit/issues/1029).

## License

This project is licensed under the MIT License.
See the [LICENSE](https://github.com/khulnasoft/nxgit/blob/master/LICENSE) file
for the full license text.

## Screenshots
Looking for an overview of the interface? Check it out!

| | | |
|:---:|:---:|:---:|
|![Dashboard](https://image.ibb.co/dms6DG/1.png)|![Repository](https://image.ibb.co/m6MSLw/2.png)|![Commits History](https://image.ibb.co/cjrSLw/3.png)|
|![Branches](https://image.ibb.co/e6vbDG/4.png)|![Issues](https://image.ibb.co/bJTJSb/5.png)|![Pull Request View](https://image.ibb.co/e02dSb/6.png)|
|![Releases](https://image.ibb.co/cUzgfw/7.png)|![Activity](https://image.ibb.co/eZgGDG/8.png)|![Wiki](https://image.ibb.co/dYV9YG/9.png)|
|![Diff](https://image.ibb.co/ewA9YG/10.png)|![Organization](https://image.ibb.co/ceOwDG/11.png)|![Profile](https://image.ibb.co/c44Q7b/12.png)|
