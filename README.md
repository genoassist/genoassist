<p align="center">
    <img src="https://user-images.githubusercontent.com/19979068/77257398-9cedfe80-6c39-11ea-890a-9167ffd1b374.png">
</p>

# GenoMagic
An all-encompasing bioinformatics tool for genome assembly and annotation projects. 

## Architecture
The overall model follows the master/slave architecture. The master is what users interact with. The users specify the files containing the contigs and what type of read they have e.g Illumina. The master takes the user's input and schedules assembly, parsing of results, and reporting, in that order. 


![](./architecture.png)

