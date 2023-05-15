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
	//本集群信息
	//一，物理信息、组件信息、转发组件信息
	//1, region
	//2, zones
	//3, chassis
	//4, cvn-controller
	//5, region gateways
	//6, zone gateways
	//二，逻辑资源信息
	//1, VPCs
	//2, subnets
	//3, ACLs
	//4, routeTables
	//5, routeEntries
	//
	//

	//cvnController table
	err = session.ExecStmt(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.cvncontroller (
		region text,
		zone text,
		data_center text,
		system_id text,
		nb_restful_ip inet,
		nb_restful_port int,
		nb_rpcx_ip inet,
		nb_rpcx_port int,
		sb_rpcx_ip inet,
		sb_rpcx_port int,
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

type chassisInfo struct {
	chassisName string
	zoneName    string
	regionName  string
	tunnelType  int
	tunnelIp    string
	tunnelPort  int
}
