package trj

import (
	"fmt"
	"strconv"
	"time"
)

const (
	HEADER         = "L'oreal check invoice circulation status: "
	START          = HEADER + "JOB STARTED"
	FINISH         = HEADER + "JOB FINISHED"
	TIMEOUT        = 10
	REFRESH_EVERY  = 300
	REFRESH_PERIOD = 1
	MAX_ITERATIONS = 3400
)

var (
	startTime        = time.Now()
	endTime          = startTime.Add(TIMEOUT * time.Second)
	heartbeatTime    time.Time
	currentIteration = 0
	heartbeatMessage = "I`m alive..."
)

func init() {
	fmt.Println(START)

	refreshHeartbeatTime()
}

func Run() {
	for beat() {

	}

	fmt.Println(FINISH)
}

func beat() bool {
	printUnprocessed()
	iterate()

	return canRun()
}

func iterate() {
	currentIteration++
}

func canRun() bool {
	canRun := haveIterations() && haveTime(endTime)

	if !canRun {
		message := "Stopped because of: "
		if !haveIterations() {
			message += "noIterations "
		}

		if !haveTime(endTime) {
			message += "noTime"
		}

		fmt.Println(message)
	}

	return canRun
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

func processExternal() { fmt.Println("wow itr " + strconv.Itoa(currentIteration)) }

func heartbeat() {
	if haveTime(heartbeatTime) && noHeartbeatIteration() {
		return
	}

	if !haveTime(heartbeatTime) {
		refreshHeartbeatTime()
		sendHeartbeat("BY TIME")
	}

	if !haveIterations() {
		sendHeartbeat("BY ITERATIONS")
	}
}

func sendHeartbeat(postfix string) {
	fmt.Println(heartbeatMessage + " Heartbeated " + postfix)
}

func noHeartbeatIteration() bool {
	return currentIteration%REFRESH_EVERY != 0
}

func refreshHeartbeatTime() {
	heartbeatTime = time.Now().Add(REFRESH_PERIOD * time.Second)
}
