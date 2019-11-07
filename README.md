## CPU-BURN

A simple app that generates cpu load for testing purposes


### Environment variables

**NUM_CORE**: Specificy the number of cores the progam is aloud to use.  Defaults to 0 which indicates no limit and will use all cores.  
**BURN_SPEED**: Specify how agressivly we should generate cpu load

* light - A simple app that generates a moderate level of cpu usage.  Runs a signel goroutine per core and sleeps for .5 millisecond in between iterations
* ridiculous - Slightly more aggressive.  Runs 5 goroutines per core and sleeps for about .5 millisecond in between iterations
* ludicrous - Very aggressive.  Runs 10 goroutrines per core and sleeps for about .5 millisecond in between iterations


### Push the app

```
cf push
```