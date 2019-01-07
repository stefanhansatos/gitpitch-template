package funding

func NewFundServer(initialBalance int) *FundServer {
    server := &FundServer{
        // make() creates builtins like channels
        Commands: make(chan interface{}),
        fund: NewFund(initialBalance),
    }

    // Spawn off the server's main loop immediately
    go server.loop()
    return server
}


