package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

var (
	cqlClusterConfig *gocql.ClusterConfig
	cqlSession       *gocqlx.Session
)

func initCql() error {
	clusterHosts := strings.Split(cqlCluster, ",")

	cqlClusterConfig = gocql.NewCluster(clusterHosts...)
	cqlClusterConfig.CQLVersion = cqlVersion
	cqlClusterConfig.Timeout = time.Duration(cqlTimeout) * time.Millisecond
	cqlClusterConfig.Consistency = gocql.Quorum
	cqlClusterConfig.MaxWaitSchemaAgreement = 2 * time.Minute
	if cqlRetry > 0 {
		cqlClusterConfig.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: cqlRetry}
	}

	switch cqlCompress {
	case "snappy":
		cqlClusterConfig.Compressor = &gocql.SnappyCompressor{}
	case "":
	default:
		vlog.Errorf("invalid compressor: %s ignore it", cqlCompress)
	}

	session, err := gocqlx.WrapSession(cqlClusterConfig.CreateSession())
	if err != nil {
		vlog.Errorf("CreateSession:", err)
		return err
	}

	//创建keyspace
	err = session.ExecStmt(fmt.Sprintf(
		`CREATE KEYSPACE %s WITH replication = {'class' : 'SimpleStrategy', 'replication_factor' : %d}`,
		cqlKeyspace,
		cqlRF))
	if err != nil {
		vlog.Errorf("create keyspace: %w", err)
		return err
	}

	//创建table
	//cvnController table
	err = session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.cvncontroller (
		region text,
		zone text,
		data_center text,
		system_id text,
		north_restful_ip inet,
		north_restful_port int,
		north_rpcx_ip inet,
		north_rpcx_port int,
		south_rpcx_ip inet,
		south_rpcx_port int,
		create_time timestamp,
		last_heartbeat timestamp,
		PRIMARY KEY (region, systemid))
		WITH default_time_to_live = 60`, cqlKeyspace))
	if err != nil {
		vlog.Errorf("create table datacenter fail: %w", err)
		return err
	}

	//datacenter table
	err = session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.datacenter (
		region text,
		zone text,
		data_center text,
		create_time timestamp,
		PRIMARY KEY (region, zone, data_center))`, cqlKeyspace))
	if err != nil {
		vlog.Errorf("create table datacenter fail: %w", err)
		return err
	}

	//chassis table

	//vpc table

	//subnet table

	//port table

	//routeTable table

	//routeEntry table

	//acl table

	//securityGroup table

	//CEN table

	//peering table

	return nil
}

func releaseCql() {
	cqlSession.Close()
}

type cvnController struct {
	regionName       string
	zoneName         string
	dcName           string
	hostname         string
	northRestfulIp   string
	northRestfulPort int
	northRpcxIp      string
	northRpcxPort    int
	southRpcxIp      string
	southRpcxPort    int
	create_time      time.Time
	last_heartbeat   time.Time
}

type dataCenter struct {
	regionName  string
	zoneName    string
	dcName      string
	create_time time.Time
}

type chassis struct {
}
