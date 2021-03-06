package collectors

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/gosuri/uiprogress"

	"github.com/linuxisnotunix/Vulnerobot/modules/collectors/anssi"
	"github.com/linuxisnotunix/Vulnerobot/modules/collectors/nvd"
	"github.com/linuxisnotunix/Vulnerobot/modules/models"
	"github.com/linuxisnotunix/Vulnerobot/modules/settings"
	"github.com/willf/pad"
)

var (
	//listCollector = []func(map[string]interface{}) models.Collector{anssi.New, dummy.New, nvd.New}
	listCollector = []func(map[string]interface{}) models.Collector{anssi.New, nvd.New}
)

//CollectorList list of collector
type CollectorList struct {
	options map[string]interface{}
	list    map[string]models.Collector
}

//Init a collector list and init them
func Init(options map[string]interface{}) *CollectorList {
	return &CollectorList{
		options: options,
		list:    getCollectors(options),
	}
}

//getCollectors Return collector list initalized
func getCollectors(options map[string]interface{}) map[string]models.Collector {
	l := make(map[string]models.Collector, len(listCollector))
	pluginList, ok := options["pluginList"].(*hashset.Set)
	if !ok {
		log.WithFields(log.Fields{
			"options": options,
		}).Debug("No plugins list -> all plugins will be init !")
	}
	for _, builder := range listCollector {
		collector := builder(options)
		if pluginList == nil || pluginList.Contains(strings.ToLower(strings.TrimSpace(collector.ID()))) { //Only init collectors based on options args
			l[collector.ID()] = collector
			if pluginList != nil {
				pluginList.Remove(strings.ToLower(strings.TrimSpace(collector.ID())))
			}
		} else {
			log.Infof("Ignoring plugin %s based on args from cli.", collector.ID())
		}
	}
	if pluginList != nil && !pluginList.Empty() {
		log.WithFields(log.Fields{
			"options":  options,
			"missings": pluginList.Values(),
		}).Warnf("Missing plugins from cli : %v", pluginList.Values())
	}
	return l
}

//Info display information on all available plugins
func (cl *CollectorList) Info(o io.Writer) error {
	keys := []string{}
	for id, collector := range cl.list {
		if collector != nil {
			log.Infof("%s is loaded and %s", collector.ID(), map[bool]string{true: "available", false: "un-available"}[collector.IsAvailable()])
			keys = append(keys, id)
		} else {
			log.Debug("Skipping empty module ", id, " !")
		}
	}
	var cs []map[string]string
	if cl.options["appList"] != nil && cl.options["appList"].(*arraylist.List) != nil {
		cs = make([]map[string]string, cl.options["appList"].(*arraylist.List).Size())
		for i, line := range cl.options["appList"].(*arraylist.List).Values() {
			cs[i] = line.(map[string]string)
		}
	}
	j, _ := json.Marshal(models.ResponseStatus{
		Plugins:    keys,
		Components: cs,
	})
	fmt.Fprint(o, string(j))
	return nil
}

//Collect ask modules to collect and parse data to put in database
func (cl *CollectorList) Collect() error {
	var p *uiprogress.Progress
	if !settings.UIDontDisplayProgress {
		p = uiprogress.New()
		p.RefreshInterval = 50 * time.Millisecond
		p.Start()
		log.SetOutput(p.Bypass())
	}
	var wg sync.WaitGroup
	for id, collector := range cl.list {
		if err := executeCollectorCollect(p, &wg, id, collector); err != nil {
			//return err
			log.Fatalf("A uncatched error occured : %s", err.Error())
		}
	}
	wg.Wait()
	if !settings.UIDontDisplayProgress {
		time.Sleep(p.RefreshInterval) //Wait to UI finish
		log.SetOutput(os.Stdout)
		p.Stop()
	}
	return nil
}

//List ask module to display known CVE stored by them in DB
func (cl *CollectorList) List(o io.Writer) error {
	hl := make(map[string][]interface{}, len(cl.list))
	for id, collector := range cl.list {
		if collector != nil {
			l, err := collector.List()
			if handleCollectorError(err) != nil {
				return err
			}
			if l != nil {
				hl[collector.ID()] = l.Values()
			}
		} else {
			log.Debug("Skipping empty module ", id, " !")
		}
	}

	//Format output
	switch cl.options["outputFormat"].(string) {
	case "csv":
		log.Debug("Exporting to CSV format ...")
		w := csv.NewWriter(o)
		defer w.Flush()
		w.Write([]string{
			"Source",
			"ComponentName", "ComponentVersion", "ComponentFunction", "ComponentCPE",
			"VulnID", "VulnURL",
		})
		for _, collectorResults := range hl {
			for _, e := range collectorResults {
				entry := e.(models.ResponseListEntry)
				for _, vuln := range entry.Vulns {
					w.Write([]string{
						vuln.Source,
						entry.Component["Name"], entry.Component["Version"], entry.Component["Function"], entry.Component["CPE"],
						vuln.Value["ID"], vuln.Value["Summary"], vuln.Value["URL"],
					})
				}
			}
		}
	case "csv-short":
		log.Debug("Exporting to CSV format (shorted format) ...")
		w := csv.NewWriter(o)
		defer w.Flush()
		w.Write([]string{
			"Source", "VulnID", "VulnURL",
		})
		for _, collectorResults := range hl {
			for _, e := range collectorResults {
				entry := e.(models.ResponseListEntry)
				for _, vuln := range entry.Vulns {
					w.Write([]string{
						vuln.Source, vuln.Value["ID"], vuln.Value["URL"],
					})
				}
			}
		}
	case "json":
		log.Debug("Exporting to JSON format ...")
		j, _ := json.Marshal(hl)
		fmt.Fprint(o, string(j))
	default:
		log.Debug("Exporting to JSON format (by default)...")
		j, _ := json.Marshal(hl)
		fmt.Fprint(o, string(j))
	}
	return nil
}

//executeCollector start the collector
func executeCollectorCollect(p *uiprogress.Progress, wg *sync.WaitGroup, id string, collector models.Collector) error {
	wg.Add(1)
	if collector != nil {
		if !collector.IsAvailable() {
			log.Warnf("Skipping plugins %s since it is not available.", collector.ID())
			wg.Done()
			return nil
		}
		log.Info("Starting module ", id, " ...")
		var bar *uiprogress.Bar
		if p != nil {
			bar = p.AddBar(1).AppendCompleted().PrependElapsed()
			bar.PrependFunc(func(b *uiprogress.Bar) string {
				return pad.Right(id, 5, " ")
				//fmt.Sprintf("%s", id)
			})
			bar.AppendFunc(func(b *uiprogress.Bar) string {
				return fmt.Sprintf("(%d/%d)", b.Current(), b.Total)
			})
		}
		go func() {
			defer wg.Done()
			if err := collector.Collect(bar); err != nil {
				handleCollectorError(err)
			}
		}()
	} else {
		log.Debug("Skipping empty module ", id, " !")
	}
	return nil
}

//handleCollectorError handle some common errors
func handleCollectorError(err error) error {
	if err != nil {
		if strings.Contains(err.Error(), "does not implement") {
			log.Warnf("%s", err.Error())
			return nil
		}
		log.Fatalf("Unhandled error : %s", err.Error())
	}
	return err
}
