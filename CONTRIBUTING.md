# Contributing to Bitcoin Pulse

**First off, thank you for considering contributing to Bitcoin Pulse. It's people like you that make Bitcoin Pulse such a great tool and community.**

## Code of Conduct

This project and everyone participating in it is governed by the [Bitcoin Pulse Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

This section guides you through submitting a bug report for Bitcoin Pulse. Following these guidelines helps maintainers and the community understand your report, reproduce the behavior, and find related reports.

#### How Do I Submit A Good Bug Report?

Bugs are tracked as [GitHub issues](https://guides.github.com/features/issues/). After you've determined [which repository](#before-submitting-a-bug-report) your bug is related to, create an issue on that repository and provide the following information by filling in [the template](.github/ISSUE_TEMPLATE/bug_report.md).

Explain the problem and include additional details to help maintainers reproduce the problem:

* **Use a clear and descriptive title** for the issue to identify the problem.
* **Describe the exact steps which reproduce the problem** in as many details as possible. For example, start by explaining how you started Bitcoin Pulse, e.g. which command exactly you used in the terminal, or how you started Bitcoin Pulse otherwise. When listing steps, **don't just say what you did, but explain how you did it**.
* **Provide specific examples to demonstrate the steps**. Include links to files or GitHub projects, or copy/pasteable snippets, which you use in those examples. If you're providing snippets in the issue, use [Markdown code blocks](https://help.github.com/articles/markdown-basics/#multiple-lines).
* **Describe the behavior you observed after following the steps** and point out what exactly is the problem with that behavior.
* **Explain which behavior you expected to see instead and why.**
* **Include screenshots and animated GIFs** which show you following the described steps and clearly demonstrate the problem. You can use [this tool](https://www.cockos.com/licecap/) to record GIFs on macOS and Windows, and [this tool](https://github.com/colinkeenan/silentcast) or [this tool](https://github.com/GNOME/byzanz) on Linux.
* **If the problem wasn't triggered by a specific action**, describe what you were doing before the problem happened and share more information using the guidelines below.

### Pull Requests

The process described here has several goals:

- Maintain Bitcoin Pulse's quality
- Fix problems that are important to users
- Engage the community in working toward the best possible Bitcoin Pulse
- Enable a sustainable system for Bitcoin Pulse's maintainers to review contributions

Please follow these steps to have your contribution considered by the maintainers:

1. **Fork the Repository**: Before making changes, please fork Bitcoin Pulse to your own GitHub account. This creates a personal copy of the project, allowing you to freely make changes without affecting the main repository. You can do this by clicking on the “Fork” button in the upper-right corner of the repository page.
    - **Clone Your Fork**: After forking, clone your forked repository to your local machine with:
      ```bash
      git clone https://github.com/<your-username>/bitcoin-pulse.git
      ```
    - **Set Up an Upstream Remote**: To stay updated with the main repository, add it as an upstream remote:
      ```bash
      git remote add upstream https://github.com/ZiyadBouazara/bitcoin-pulse.git
      ```
    - **Sync Changes**: Periodically pull changes from the main repository to stay up-to-date:
      ```bash
      git fetch upstream
      git merge upstream/main
      ```

2. **Make Changes**: Create a new branch for your changes. This keeps your work organized and makes it easier for maintainers to review:
   ```bash
   git checkout -b <branch-name>
   ```

3. **Open a Pull Request**: After committing your changes and pushing your branch to your fork, open a pull request (PR) from your branch in your fork to the `main` branch of the main repository. GitHub will guide you through the process after you push your branch.

4. **Ensure Status Checks Pass**: After you submit your pull request, verify that all status checks are passing. If a status check is failing, and you believe that the failure is unrelated to your change, please leave a comment on the pull request explaining why. A maintainer will re-run the status check for you.

5. **Final Review and Merge**: The maintainers will review your PR and may request additional changes before merging.

## Any questions?

Don't hesitate to ask! Open an issue, and we'll do our best to help you out.