# Contributing to Boostly

Thank you for your interest in contributing to Boostly! Your efforts will help make this library even better for everyone.

## Getting Started

1. **Fork the Repository**: Click on the 'Fork' button at the top right of this page. This will create a copy of the Boostly repository in your GitHub account.

2. **Clone Your Fork**: Open a terminal and run:
   ```
   git clone https://github.com/filecoin-shipyard/boostly.git
   ```

3. **Add the Upstream Remote**: This will be useful to sync your fork with the latest changes:
   ```
   git remote add upstream https://github.com/filecoin-shipyard/boostly.git
   ```

4. **Create a New Branch**: Always base your work on a new branch:
   ```
   git checkout -b feature/my-new-feature
   ```

## Guidelines

- **Code Style**: Make sure to follow Go's official [coding standards](https://golang.org/doc/effective_go.html).

- **Commit Messages**: Keep your commit messages clear and descriptive.

- **Update Documentation**: If your changes add or modify functionality, ensure that corresponding documentation is updated as well.

- **Write Tests**: Ensure your code has adequate test coverage.

## Submitting a Pull Request

1. **Commit Your Changes**: Once you are happy with your changes, add them to the staging area and commit them:
   ```
   git add .
   git commit -m "Add a descriptive commit message"
   ```

2. **Sync with Upstream**: Fetch the latest changes from the original repository:
   ```
   git fetch upstream
   git rebase upstream/main
   ```

3. **Push Your Branch**: Push your branch to your fork:
   ```
   git push origin feature/my-new-feature
   ```

4. **Open a Pull Request**: Go to your fork on GitHub and click on the 'Compare & pull request' button. Fill in a clear title and description and submit the pull request.

## Feedback

If you're unsure about any aspect of the contribution process, please open an [issue](https://github.com/filecoin-shipyard/boostly/issues). We're here to help!

---

Thank you for making Boostly better! :rocket: