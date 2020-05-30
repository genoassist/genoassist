<p align="center">
    <img src="https://user-images.githubusercontent.com/19979068/77257398-9cedfe80-6c39-11ea-890a-9167ffd1b374.png" alt="">
    <br /><i>An all-encompassing Bioinformatics tool for genome assembly and annotation projects</i><br>
</p>

---

# Table of contents

1. [About](#1-about) </br>
1. [Installation](#2-installation)
1. [GenoMagic Usage](#3-genomagic-usage)
1. [Architecture](#4-architecture)
2. [Maintainers](#5-maintainers)
1. [Feedback and bug reports](#5-feedback-and-bug-reports)

## 1. About

One of the challenges that biologists and bioinformaticians face during genome assembly projects is choosing from the plethora of assembly softwares. This is highly time consuming as there are various parameters for each of the assemblers that the user needs to learn about. And even if the user learns about these various parameters for each assembler; there is still the tedious job of running various assemblers, and comparing the statistics to identify the best assembly.

My aim with GenoMagic is to give the biologist an application which performs genome assemblies with minimal configuration required, which visualizes the comparative statistics for the chosen assemblers. This app enables the bioinformaticians to alter the default parameters of their choosing, making it applicable for users with diverse experiences and skills.

 
## 2. Installation

1. You can either use go (will be added to `$GOPATH/`):
    ```sh
    $ go get -u github.com/genomagic
    ```
    
    Or clone the repository:  
    ```sh
    $ git clone https://github.com/genomagic/genomagic
    ```
2. Build the `main.go` file
    ```sh
    $ go build main.go
    ```

If you are missing packages, run `go mod vendor` to collect the necessary packages

## 3. GenoMagic usage

GenoMagic only requires a YAML file that contains the configuration it should use to run its processes. A template can 
be found in this repository. For convenience, here's an example specification:

```yaml
assemblers:
  megahit:
    kmers: "27"
  abyss:
    kmers: "27"
genomagic:
  inputFilePath: "/test/raw_sequences.fastq"
  outputPath: "/test/output"
  threads: 2
  prep: true
```

Note: all paths used with GenoMagic have to be absolute paths (a Docker requirement).

### Installing Docker images through GenoMagic

If you are encountering problems with Docker, make sure that:
1. The Docker daemon is running in the background
1. You have the necessary Docker images, which can be installed via GenoMagic specifying `prep: true` under `genomagic`
in the YAML configuration. This will install the necessary Docker images for the containers that GenoMagic 
runs.

## 4. Architecture

The overall model follows the master/slave architecture. The master is what users interact with. 
The users specify the files containing the contigs and what type of read they have e.g Illumina. 
The master takes the user's input and schedules assembly, parsing of results, and reporting, in that order. 

![](./architecture.png)

## 5. Maintainers

[Tayab Soomro](https://github.com/tayabsoomro)  
[Flaviu Vadan](https://github.com/flaviuvadan)

Feel free to contact any of the maintainers if you would like to be an active 
maintainer and contributor to GenoMagic! If you would like to contribute only,
you are encouraged to grab an issue and submit a pull request with proposed
changes for review! 

## 6. Feedback and bug reports

Submit feedback and bug reports by using the Issues section of the repository.

