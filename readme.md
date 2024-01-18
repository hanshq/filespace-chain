
# Project Name

Brief description of your project.

## Getting Started

These instructions will get your copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

What things you need to install the software and how to install them.

```
Give examples
```

### Installation

A step-by-step series of examples that tell you how to get a development environment running.

```
Give examples
```

## Usage

### Pushing Docker Image

To push the Docker image to a registry, use the `docker_push.sh` script with the version number as an argument.

```bash
./scripts/src/docker_push.sh 22
```

### Deploying on Akash

To deploy on Akash, use the SDL file located at `./data/deploy/akash_deploy.sdl`.

#### Steps for Akash Console

1. **Add Owner Key**: Recover the owner's key in the Akash console.

    ```bash
    filespace-chaind keys add owner --recover
    ```

2. **Enable Validator**: Run the script to enable the validator.

    ```bash
    ./scripts/src/enable_validator.sh
    ```

## Contributing

Please read [CONTRIBUTING.md](link-to-contributing-file) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](link-to-tags).

## Authors

* **Your Name** - *Initial work* - [YourUsername](link-to-your-profile)

See also the list of [contributors](link-to-contributors) who participated in this project.

## License

This project is licensed under the XYZ License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc