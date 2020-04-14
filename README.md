<p align="center">
    <img src="https://user-images.githubusercontent.com/19979068/77257398-9cedfe80-6c39-11ea-890a-9167ffd1b374.png">
    <br /><i>An all-encompasing bioinformatics tool for genome assembly and annotation projects</i><br>
</p>

---


### TABLE OF CONTENTS

1. [About GenoMagic](#sec1) </br>
2. [Software Architecture](#sec2)</br>
3. [Installation](#sec3)</br>
4. [Feedback and bug reports](#sec4)<br />

<a name="sec1"></a>
## 1. About GenoMagic

An all-encompasing bioinformatics tool for genome assembly and annotation projects. This project is still under development. Therefore, it is missing a *Usage* section; this will be added in the future.


<a name="sec2"></a>
## 2. Software Architecture

The overall model follows the master/slave architecture. The master is what users interact with. The users specify the files containing the contigs and what type of read they have e.g Illumina. The master takes the user's input and schedules assembly, parsing of results, and reporting, in that order. 

![](./architecture.png)
 
<a name="sec3"></a>
## 3. Installation

1. Clone the repository  
`$ git clone https://github.com/genomagic/genomagic.git`  

2. Build the `main.go` file  
`$ go build main.go`

3. Run the binary  
`$ ./main`  
<a name="sec4"></a>
## 4. Feedback and bug reports
Please submit any and all feedbacks as issues to this repository.

