package trj

import (
	"fmt"
	"time"
)

const (
	HEADER         = "L'oreal check invoice circulation status: "
	START          = HEADER + "JOB STARTED"
	FINISH         = HEADER + "JOB FINISHED"
	TIMEOUT        = 10
	REFRESH_EVERY  = 300
	REFRESH_PERIOD = 1
	MAX_ITERATIONS = 34
)

var (
	startTime        = time.Now()
	endTime          = startTime.Add(TIMEOUT * time.Second)
	heartbeatTime    time.Time
	currentIteration = 0
	heartbeatMessage = "I`m alive..."
)

func init() {
	refreshHeartbeat()

	fmt.Println(START)
	defer fmt.Println(FINISH)
}

func Run() {
	for beat() {
		
	}
}

func beat() bool {
	printUnprocessed()
	iterate()

	return checkLifecycles()
}

func iterate() {
	currentIteration++
}

func checkLifecycles() bool {
	return haveIterations() && canRun()
}

func canRun() bool {
	return haveTime(endTime)
}

func haveIterations() bool {
	return currentIteration < MAX_ITERATIONS
}

func haveTime(finishTime time.Time) bool {
	return time.Now().Before(finishTime)
}

func printUnprocessed() {
	processExternal()
	heartbeat()
}

func processExternal() {}

func heartbeat() {
	if haveTime(heartbeatTime) && noHeartbeatIteration() {
		return
	}

	refreshHeartbeat()

	fmt.Println(heartbeatMessage)
}

func noHeartbeatIteration() bool {
	return currentIteration%REFRESH_EVERY != 0
}

func refreshHeartbeat() {
	heartbeatTime = time.Now().Add(REFRESH_PERIOD * time.Second)
}
