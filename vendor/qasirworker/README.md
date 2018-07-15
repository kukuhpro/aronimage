#### Qasir Worker 
Simple Worker that been build for handling any processing using queue and worker.

```
    type Payload struct {
	    Helo string
    }

    // THIS IS WHERE PROCESSING WILL RUNNING, EVERY STRUCT MUST HAVE FUNCTION HANDLE
    func (p *Payload) Handle() error {
        log.Println(p.Helo)
        return nil
    }

	qasirworker.NewDispatcher(MaxWorker, MaxQueue).Run()
    
    func main() {
        payload := &Payload{Helo: "kukuh"}
		work := qasirworker.Job{Executor: payload}
		qasirworker.JobQueue <- work
    }
```