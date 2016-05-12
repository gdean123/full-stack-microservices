# full-stack-microservices
An experiment: web components rendered live from microservices, integrated together into a shell application

## Technical Vision

* A collection of full-stack microservices that each serve up their own user interface in the form of a collection of web components (using Polymer).
* A thin shell application that glues together the different web components but doesn't have too much logic
* Build pipelines that produce docker images as artifacts that run each microservice for use by other teams in local development and acceptance testing

## Team Organization

* Each team owns one or more microservices and their associated web component frontends
* One team owns the shell application along with its acceptance tests

## Local Development

Every team is responsible for creating a docker image that runs their service (along with its database) when launched. These can be used by other teams that depend on their service during local development. This prevents other teams from depending on a shared, running instance of the service which might go down, might be at the wrong version, or might not be seeded appropriately.

## Testing Approach

Each microservice has the following tests:
1. a small set of full-stack acceptance tests that render the web components in a basic html file and validate that user flows work. These spin up local editions of the dependent services to run against.
2. a set of API-level consumer-driven contract tests, with contracts created by the Javascript frontend and other service consumers
3. a set of Javascript unit tests written using Jasmine that probably don't render the web components but just validate that everything underneath them works.
4. a set of server-side unit tests

The shell application only has a small set of full-stack acceptance tests that render the entire application and validate that user flows work, spinning up local editions of all of the services to run against. Ideally no other tests are required since the only responsibility of the shell application is to glue together full-stack microservices.

## Build Pipelines

**Each microservice**

1. Run javascript unit tests
2. Run server-side unit tests
3. Run server-side consumer-driven contract tests against locally run dependent services
4. Run full-stack acceptance tests against locally run dependent services
5. Run the shell application's acceptance tests to ensure that this commit doesn't break the application as a whole and to avoid blocking other teams by breaking the build at the point of pipeline convergence
6. Create a 'gold-master' tag at this SHA
7. Build and publish a docker container representing the gold-master release
8. Automatically deploy the service to staging (which would ideally be a complete replica of production, including data)

**Shell application**

1. Run full-stack acceptance tests against all locally run services
2. Create a 'gold-master' tag at this SHA
3. Build and publish a docker container representing the gold-master release

## Deployment

I am not yet sure how deployment of such a system to production should work. Simplistically, each microservice could have a manually triggered job at the end of its pipeline that performs the deployment. However, it could be the case that in the time between the last run of the shell application's acceptance tests and the deployment, another service could have changed and an incompatible version could have been deployed to production.

Another approach would be to have a repository that submodules the shell application and all of the microservices. A gold-master tag in this repository would represent a collection of versions of microservices that have been proven to work together by the acceptance tests. A job could deploy all of the changed services to production at once, though order of the deployments could matter and would be difficult to manage. Also, it seems sad to remove the ability for teams to deploy to production independently of one another.

## State of the project

This project is still in its infancy. It is comprised of:
* A tasks service written in Go that serves a hello-world web component
* A shell application written in Ruby with Sinatra that renders the hello-world web component

Right now, the hello-world web component is tested with a Jasmine test, but the test is super slow and that is probably not a sustainable approach.
