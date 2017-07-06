package main

import (
	"bufio"
	"bytes"
	"context"
	"log"
	osExec "os/exec"
	"strconv"
	"strings"
	"time"

	"collectd.org/api"
	"collectd.org/exec"
)

func main() {
	e := exec.NewExecutor()
	e.VoidCallback(nsdStats, exec.Interval())
	e.Run(context.Background())
}

func nsdStats(ctx context.Context, interval time.Duration) {
	var vl *api.ValueList

	buf := &bytes.Buffer{}
	cmd := osExec.Command("/bin/sh", "-c", "nsd-control stats_noreset")
	cmd.Stdout = buf

	if err := cmd.Run(); err != nil {
		log.Fatalf("unable to execute nsd-control: %v", err)
	}

	now := time.Now()

	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.HasPrefix(line, "num.") {
			continue
		}

		fields := strings.SplitN(line, "=", 2)
		if len(fields) != 2 {
			continue
		}

		metric := fields[0]
		metric = strings.TrimPrefix(metric, "num.")
		metric = strings.Replace(metric, ".", "_", -1)

		value, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Printf("error: unable to parse metric value: %v", err)
			continue
		}

		vl = &api.ValueList{
			Identifier: api.Identifier{
				Host:         exec.Hostname(),
				Plugin:       "nsd",
				Type:         "derive",
				TypeInstance: metric,
			},
			Time:     now,
			Interval: interval,
			Values:   []api.Value{api.Derive(value)},
		}

		exec.Putval.Write(ctx, vl)
	}
}
