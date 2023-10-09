
# Git Branching Strategy and CI/CD

This project adopts trunk based development approach, complemented with monorepo strategy. The main source of truth of this project is the `main` branch.

## Development Workflow

To propose changes to this project, firstly make sure you have the latest version of `main` bracnh in your local workstation, then checkout from it to create new branch. Make sure the scope of changes is as small and specific as possible, so it would easier to test. Please do pull and rebase from the `main` branch regularly.

If the changes are already completed, create pull request from your feature branch to `main` branch. If your changes including those files in `services/<name of the service>` folder, a CI pipeline will be triggered to execute test on your changes. Make sure that your PR has been approved by the service owners or other related stakeholders and all test has been executed successfully, then you can merge the PR to `main` branch.

Currently, the CI pipeline on pull request supports these execution:
- Golang service
    - Unit test
    - Lint using `golangci-lint`
    - Build test using Docker
- NodeJS service
    - Unit test
    - Build test using Docker

#### Development environment

Upon merging a PR to `main` branch, if the changes including the files beloning to `services/<name of the service>` folder, a CI pipeline will be triggered to execute the build of the specific service.

If the build success, the workflow will be continued to deployment of the service to development environment's Kubernetes cluster using container image built during the previous job.

#### Staging environment

To promote your service to staging environment, firstly create a release branch from `main` branch with format `release/<service name>/<semantic version>`. Then, create a git tag from that release branch, following this format `<service name>-v<semantic version>-rc<rc number>`. Upon the tag creation, a CI/CD workflow will be triggered to build the container image of the service based on the commit referenced by the tag, then deploy the service to staging environment's Kubernetes cluster.

For example, we want to promote `user` service to staging environment with semantic version of `1.0.0`. Firstly, we need to create a release branch from `main` branch with branch name `release/user/1.0.0`. Then, we create git tag from that branch with format `user-v1.0.0-rc1`. This tag creation will trigger build and deployment to staging environment.

Assuming that the tag `user-v1.0.0-rc1` still contains some bugs and the fixes have been pushed in the later commit, we could tag this later commit with `user-v1.0.0-rc2` and it will trigger the deployment to staging environment with the bugfixes.

#### Production environment

The trigger to production deployment is also tag creation with rule that the tag should reference the same commit as the tag used to trigger deployment to staging environment. The tag format for production deployment is `<service name>-v<semantic version>`, note that there is no `rc` in production tag. The tag creation will trigger re-tagging of the container image used in staging environment, we call this step as `publish` step. Then, the workflow will deploy the re-tagged container image to production environment's Kubernetes cluster.

For example, we want to promote user service tagged with `user-v1.0.0-rc2` to production environment. We should create a git tag `user-v1.0.0` referencing the same commit as the previous tag, then the workflow will be triggered to deploy to production.

Overall, this is the development flow starting from when to code is still developed in engineers' local workstation, until it is deployed to production environment.

