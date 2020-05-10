<p style="align-content: center">
    <img src="https://user-images.githubusercontent.com/19979068/77257398-9cedfe80-6c39-11ea-890a-9167ffd1b374.png" alt="">
    <br /><i>An all-encompassing Bioinformatics tool for genome assembly and annotation projects</i><br>
</p>

---

# Table of contents

1. [About](#1-about) </br>
1. [Architecture](#2-architecture)
1. [Installation](#3-installation)
1. [Running](#4-running-genomagic)
1. [Feedback and bug reports](#5-feedback-and-bug-reports)

## 1. About

An all-encompassing Bioinformatics tool for genome assembly and annotation projects. 
This project is still under development. 


## 2. Architecture

The overall model follows the master/slave architecture. The master is what users interact with. 
The users specify the files containing the contigs and what type of read they have e.g Illumina. 
The master takes the user's input and schedules assembly, parsing of results, and reporting, in that order. 

![](./architecture.png)
 
## 3. Installation

1. Clone the repository  
`$ git clone https://github.com/genomagic/genomagic.git`  

2. Build the `main.go` file  
`$ go build main.go`

## 4. Running GenoMagic

After building the binary, the program can be run as follows:
```
$ ./main -fastq $(pwd)/reads.fastq
```

The available flags are:

<table>
    <tr>
        <th>Flag</th>
        <th>Value</th>
        <th>Required</th>
        <th>Description</th>
    </tr>
    <tr>
        <td><code>fastq</code></td>
        <td>./path/to/sequence.fastq</td>
        <td>Yes</td>
        <td>The path to the FASTQ file containing raw sequence data for assembly</td>
    </tr>
    <tr>
        <td><code>prep</code></td>
        <td><code>true</code> / <code>false</code></td>
        <td>No</td>
        <td>whether to install all the necessary Docker containers for assembly as a preparatory step. 
            Should be done at least once</td>
    </tr>
    <tr>
        <td><code>out</code></td>
        <td>./path/to/output/directory</td>
        <td>No</td>
        <td>The path to the directory where results will be stored, defaults to current working directory</td>
    </tr>
    <tr>
        <td><code>threads</code></td>
        <td><code>integer</code></td>
        <td>No</td>
        <td>The number of threads to use for assembly and output parsing processes</td>
    </tr>
</table>

## 5. Feedback and bug reports
Submit feedback and bug reports by using the Issues section of the repository.

