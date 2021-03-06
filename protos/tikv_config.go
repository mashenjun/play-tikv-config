package protos

type ReadPoolConfig struct {
	Unified struct {
		MinThreadCount    int    `toml:"min-thread-count"`
		MaxThreadCount    int    `toml:"max-thread-count"`
		StackSize         string `toml:"stack-size"`
		MaxTasksPerWorker int    `toml:"max-tasks-per-worker"`
	} `toml:"unified"`
	Storage struct {
		HighConcurrency         int    `toml:"high-concurrency"`
		NormalConcurrency       int    `toml:"normal-concurrency"`
		LowConcurrency          int    `toml:"low-concurrency"`
		MaxTasksPerWorkerHigh   int    `toml:"max-tasks-per-worker-high"`
		MaxTasksPerWorkerNormal int    `toml:"max-tasks-per-worker-normal"`
		MaxTasksPerWorkerLow    int    `toml:"max-tasks-per-worker-low"`
		StackSize               string `toml:"stack-size"`
	} `toml:"storage"`
	Coprocessor struct {
		HighConcurrency         int    `toml:"high-concurrency"`
		NormalConcurrency       int    `toml:"normal-concurrency"`
		LowConcurrency          int    `toml:"low-concurrency"`
		MaxTasksPerWorkerHigh   int    `toml:"max-tasks-per-worker-high"`
		MaxTasksPerWorkerNormal int    `toml:"max-tasks-per-worker-normal"`
		MaxTasksPerWorkerLow    int    `toml:"max-tasks-per-worker-low"`
		StackSize               string `toml:"stack-size"`
	} `toml:"coprocessor"`
}

type ServerConfig struct {
	Addr                             string  `toml:"addr"`
	AdvertiseAddr                    string  `toml:"advertise-addr"`
	StatusAddr                       string  `toml:"status-addr"`
	AdvertiseStatusAddr              string  `toml:"advertise-status-addr"`
	StatusThreadPoolSize             int     `toml:"status-thread-pool-size"`
	MaxGrpcSendMsgLen                int     `toml:"max-grpc-send-msg-len"`
	RaftClientGrpcSendMsgBuffer      int     `toml:"raft-client-grpc-send-msg-buffer"`
	RaftClientQueueSize              int     `toml:"raft-client-queue-size"`
	RaftMsgMaxBatchSize              int     `toml:"raft-msg-max-batch-size"`
	GrpcCompressionType              string  `toml:"grpc-compression-type"`
	GrpcConcurrency                  int     `toml:"grpc-concurrency"`
	GrpcConcurrentStream             int     `toml:"grpc-concurrent-stream"`
	GrpcRaftConnNum                  int     `toml:"grpc-raft-conn-num"`
	GrpcMemoryPoolQuota              int64   `toml:"grpc-memory-pool-quota"`
	GrpcStreamInitialWindowSize      string  `toml:"grpc-stream-initial-window-size"`
	GrpcKeepaliveTime                string  `toml:"grpc-keepalive-time"`
	GrpcKeepaliveTimeout             string  `toml:"grpc-keepalive-timeout"`
	ConcurrentSendSnapLimit          int     `toml:"concurrent-send-snap-limit"`
	ConcurrentRecvSnapLimit          int     `toml:"concurrent-recv-snap-limit"`
	EndPointRecursionLimit           int     `toml:"end-point-recursion-limit"`
	EndPointStreamChannelSize        int     `toml:"end-point-stream-channel-size"`
	EndPointBatchRowLimit            int     `toml:"end-point-batch-row-limit"`
	EndPointStreamBatchRowLimit      int     `toml:"end-point-stream-batch-row-limit"`
	EndPointEnableBatchIfPossible    bool    `toml:"end-point-enable-batch-if-possible"`
	EndPointRequestMaxHandleDuration string  `toml:"end-point-request-max-handle-duration"`
	EndPointMaxConcurrency           int     `toml:"end-point-max-concurrency"`
	SnapMaxWriteBytesPerSec          string  `toml:"snap-max-write-bytes-per-sec"`
	SnapMaxTotalSize                 string  `toml:"snap-max-total-size"`
	StatsConcurrency                 int     `toml:"stats-concurrency"`
	HeavyLoadThreshold               int     `toml:"heavy-load-threshold"`
	HeavyLoadWaitDuration            string  `toml:"heavy-load-wait-duration"`
	EnableRequestBatch               bool    `toml:"enable-request-batch"`
	BackgroundThreadCount            int     `toml:"background-thread-count"`
	EndPointSlowLogThreshold         string  `toml:"end-point-slow-log-threshold"`
	ForwardMaxConnectionsPerAddress  int     `toml:"forward-max-connections-per-address"`
	RejectMessagesOnMemoryRatio      float64 `toml:"reject-messages-on-memory-ratio"`
	Labels                           struct{} `toml:"labels"`
}

type StorageConfig struct {
	DataDir                        string  `toml:"data-dir"`
	GcRatioThreshold               float64 `toml:"gc-ratio-threshold"`
	MaxKeySize                     int     `toml:"max-key-size"`
	SchedulerConcurrency           int     `toml:"scheduler-concurrency"`
	SchedulerWorkerPoolSize        int     `toml:"scheduler-worker-pool-size"`
	SchedulerPendingWriteThreshold string  `toml:"scheduler-pending-write-threshold"`
	ReserveSpace                   string  `toml:"reserve-space"`
	EnableAsyncApplyPrewrite       bool    `toml:"enable-async-apply-prewrite"`
	EnableTTL                      bool    `toml:"enable-ttl"`
	TTLCheckPollInterval           string  `toml:"ttl-check-poll-interval"`
	FlowControl                    struct {
		Enable                          bool   `toml:"enable"`
		SoftPendingCompactionBytesLimit string `toml:"soft-pending-compaction-bytes-limit"`
		HardPendingCompactionBytesLimit string `toml:"hard-pending-compaction-bytes-limit"`
		MemtablesThreshold              int    `toml:"memtables-threshold"`
		L0FilesThreshold                int    `toml:"l0-files-threshold"`
	} `toml:"flow-control"`
	BlockCache struct {
		Shared              bool    `toml:"shared"`
		NumShardBits        int     `toml:"num-shard-bits"`
		StrictCapacityLimit bool    `toml:"strict-capacity-limit"`
		HighPriPoolRatio    float64 `toml:"high-pri-pool-ratio"`
		MemoryAllocator     string  `toml:"memory-allocator"`
	} `toml:"block-cache"`
	IoRateLimit struct {
		MaxBytesPerSec              string `toml:"max-bytes-per-sec"`
		Mode                        string `toml:"mode"`
		Strict                      bool   `toml:"strict"`
		ForegroundReadPriority      string `toml:"foreground-read-priority"`
		ForegroundWritePriority     string `toml:"foreground-write-priority"`
		FlushPriority               string `toml:"flush-priority"`
		LevelZeroCompactionPriority string `toml:"level-zero-compaction-priority"`
		CompactionPriority          string `toml:"compaction-priority"`
		ReplicationPriority         string `toml:"replication-priority"`
		LoadBalancePriority         string `toml:"load-balance-priority"`
		GcPriority                  string `toml:"gc-priority"`
		ImportPriority              string `toml:"import-priority"`
		ExportPriority              string `toml:"export-priority"`
		OtherPriority               string `toml:"other-priority"`
	} `toml:"io-rate-limit"`
}

type PDConfig struct {
	Endpoints        []string `toml:"endpoints"`
	RetryInterval    string   `toml:"retry-interval"`
	RetryMaxCount    int64    `toml:"retry-max-count"`
	RetryLogEvery    int      `toml:"retry-log-every"`
	UpdateInterval   string   `toml:"update-interval"`
	EnableForwarding bool     `toml:"enable-forwarding"`
}

type RaftStoreConfig struct {
	Prevote                          bool    `toml:"prevote"`
	RaftdbPath                       string  `toml:"raftdb-path"`
	Capacity                         string  `toml:"capacity"`
	RaftBaseTickInterval             string  `toml:"raft-base-tick-interval"`
	RaftHeartbeatTicks               int     `toml:"raft-heartbeat-ticks"`
	RaftElectionTimeoutTicks         int     `toml:"raft-election-timeout-ticks"`
	RaftMinElectionTimeoutTicks      int     `toml:"raft-min-election-timeout-ticks"`
	RaftMaxElectionTimeoutTicks      int     `toml:"raft-max-election-timeout-ticks"`
	RaftMaxSizePerMsg                string  `toml:"raft-max-size-per-msg"`
	RaftMaxInflightMsgs              int     `toml:"raft-max-inflight-msgs"`
	RaftEntryMaxSize                 string  `toml:"raft-entry-max-size"`
	RaftLogGcTickInterval            string  `toml:"raft-log-gc-tick-interval"`
	RaftLogGcThreshold               int     `toml:"raft-log-gc-threshold"`
	RaftLogGcCountLimit              int     `toml:"raft-log-gc-count-limit"`
	RaftLogGcSizeLimit               string  `toml:"raft-log-gc-size-limit"`
	RaftLogReserveMaxTicks           int     `toml:"raft-log-reserve-max-ticks"`
	RaftEnginePurgeInterval          string  `toml:"raft-engine-purge-interval"`
	RaftEntryCacheLifeTime           string  `toml:"raft-entry-cache-life-time"`
	RaftRejectTransferLeaderDuration string  `toml:"raft-reject-transfer-leader-duration"`
	SplitRegionCheckTickInterval     string  `toml:"split-region-check-tick-interval"`
	RegionSplitCheckDiff             string  `toml:"region-split-check-diff"`
	RegionCompactCheckInterval       string  `toml:"region-compact-check-interval"`
	RegionCompactCheckStep           int     `toml:"region-compact-check-step"`
	RegionCompactMinTombstones       int     `toml:"region-compact-min-tombstones"`
	RegionCompactTombstonesPercent   int     `toml:"region-compact-tombstones-percent"`
	PdHeartbeatTickInterval          string  `toml:"pd-heartbeat-tick-interval"`
	PdStoreHeartbeatTickInterval     string  `toml:"pd-store-heartbeat-tick-interval"`
	SnapMgrGcTickInterval            string  `toml:"snap-mgr-gc-tick-interval"`
	SnapGcTimeout                    string  `toml:"snap-gc-timeout"`
	LockCfCompactInterval            string  `toml:"lock-cf-compact-interval"`
	LockCfCompactBytesThreshold      string  `toml:"lock-cf-compact-bytes-threshold"`
	NotifyCapacity                   int     `toml:"notify-capacity"`
	MessagesPerTick                  int     `toml:"messages-per-tick"`
	MaxPeerDownDuration              string  `toml:"max-peer-down-duration"`
	MaxLeaderMissingDuration         string  `toml:"max-leader-missing-duration"`
	AbnormalLeaderMissingDuration    string  `toml:"abnormal-leader-missing-duration"`
	PeerStaleStateCheckInterval      string  `toml:"peer-stale-state-check-interval"`
	LeaderTransferMaxLogLag          int     `toml:"leader-transfer-max-log-lag"`
	SnapApplyBatchSize               string  `toml:"snap-apply-batch-size"`
	ConsistencyCheckInterval         string  `toml:"consistency-check-interval"`
	ReportRegionFlowInterval         string  `toml:"report-region-flow-interval"`
	RaftStoreMaxLeaderLease          string  `toml:"raft-store-max-leader-lease"`
	RightDeriveWhenSplit             bool    `toml:"right-derive-when-split"`
	AllowRemoveLeader                bool    `toml:"allow-remove-leader"`
	MergeMaxLogGap                   int     `toml:"merge-max-log-gap"`
	MergeCheckTickInterval           string  `toml:"merge-check-tick-interval"`
	UseDeleteRange                   bool    `toml:"use-delete-range"`
	CleanupImportSstInterval         string  `toml:"cleanup-import-sst-interval"`
	LocalReadBatchSize               int     `toml:"local-read-batch-size"`
	ApplyMaxBatchSize                int     `toml:"apply-max-batch-size"`
	ApplyPoolSize                    int     `toml:"apply-pool-size"`
	ApplyRescheduleDuration          string  `toml:"apply-reschedule-duration"`
	ApplyLowPriorityPoolSize         int     `toml:"apply-low-priority-pool-size"`
	StoreMaxBatchSize                int     `toml:"store-max-batch-size"`
	StorePoolSize                    int     `toml:"store-pool-size"`
	StoreRescheduleDuration          string  `toml:"store-reschedule-duration"`
	StoreLowPriorityPoolSize         int     `toml:"store-low-priority-pool-size"`
	StoreIoPoolSize                  int     `toml:"store-io-pool-size"`
	StoreIoNotifyCapacity            int     `toml:"store-io-notify-capacity"`
	FuturePollSize                   int     `toml:"future-poll-size"`
	HibernateRegions                 bool    `toml:"hibernate-regions"`
	DevAssert                        bool    `toml:"dev-assert"`
	ApplyYieldDuration               string  `toml:"apply-yield-duration"`
	PerfLevel                        int     `toml:"perf-level"`
	EvictCacheOnMemoryRatio          float64 `toml:"evict-cache-on-memory-ratio"`
	CmdBatch                         bool    `toml:"cmd-batch"`
	RaftWriteSizeLimit               string  `toml:"raft-write-size-limit"`
	WaterfallMetrics                 bool    `toml:"waterfall-metrics"`
	IoRescheduleConcurrentMaxCount   int     `toml:"io-reschedule-concurrent-max-count"`
	IoRescheduleHotpotDuration       string  `toml:"io-reschedule-hotpot-duration"`
	InspectInterval                  string  `toml:"inspect-interval"`
}

type CoprocessorConfig struct {
	SplitRegionOnTable     bool   `toml:"split-region-on-table"`
	BatchSplitLimit        int    `toml:"batch-split-limit"`
	RegionMaxSize          string `toml:"region-max-size"`
	RegionSplitSize        string `toml:"region-split-size"`
	RegionMaxKeys          int    `toml:"region-max-keys"`
	RegionSplitKeys        int    `toml:"region-split-keys"`
	ConsistencyCheckMethod string `toml:"consistency-check-method"`
	PerfLevel              int    `toml:"perf-level"`
}

type CFConfig struct {
	BlockSize                           string   `toml:"block-size"`
	BlockCacheSize                      string   `toml:"block-cache-size"`
	DisableBlockCache                   bool     `toml:"disable-block-cache"`
	CacheIndexAndFilterBlocks           bool     `toml:"cache-index-and-filter-blocks"`
	PinL0FilterAndIndexBlocks           bool     `toml:"pin-l0-filter-and-index-blocks"`
	UseBloomFilter                      bool     `toml:"use-bloom-filter"`
	OptimizeFiltersForHits              bool     `toml:"optimize-filters-for-hits"`
	WholeKeyFiltering                   bool     `toml:"whole-key-filtering"`
	BloomFilterBitsPerKey               int      `toml:"bloom-filter-bits-per-key"`
	BlockBasedBloomFilter               bool     `toml:"block-based-bloom-filter"`
	ReadAmpBytesPerBit                  int      `toml:"read-amp-bytes-per-bit"`
	CompressionPerLevel                 []string `toml:"compression-per-level"`
	WriteBufferSize                     string   `toml:"write-buffer-size"`
	MaxWriteBufferNumber                int      `toml:"max-write-buffer-number"`
	MinWriteBufferNumberToMerge         int      `toml:"min-write-buffer-number-to-merge"`
	MaxBytesForLevelBase                string   `toml:"max-bytes-for-level-base"`
	TargetFileSizeBase                  string   `toml:"target-file-size-base"`
	Level0FileNumCompactionTrigger      int      `toml:"level0-file-num-compaction-trigger"`
	Level0SlowdownWritesTrigger         int      `toml:"level0-slowdown-writes-trigger"`
	Level0StopWritesTrigger             int      `toml:"level0-stop-writes-trigger"`
	MaxCompactionBytes                  string   `toml:"max-compaction-bytes"`
	CompactionPri                       int      `toml:"compaction-pri"`
	DynamicLevelBytes                   bool     `toml:"dynamic-level-bytes"`
	NumLevels                           int      `toml:"num-levels"`
	MaxBytesForLevelMultiplier          int      `toml:"max-bytes-for-level-multiplier"`
	CompactionStyle                     int      `toml:"compaction-style"`
	DisableAutoCompactions              bool     `toml:"disable-auto-compactions"`
	DisableWriteStall                   bool     `toml:"disable-write-stall"`
	SoftPendingCompactionBytesLimit     string   `toml:"soft-pending-compaction-bytes-limit"`
	HardPendingCompactionBytesLimit     string   `toml:"hard-pending-compaction-bytes-limit"`
	ForceConsistencyChecks              bool     `toml:"force-consistency-checks"`
	PropSizeIndexDistance               int      `toml:"prop-size-index-distance"`
	PropKeysIndexDistance               int      `toml:"prop-keys-index-distance"`
	EnableDoublySkiplist                bool     `toml:"enable-doubly-skiplist"`
	EnableCompactionGuard               bool     `toml:"enable-compaction-guard"`
	CompactionGuardMinOutputFileSize    string   `toml:"compaction-guard-min-output-file-size"`
	CompactionGuardMaxOutputFileSize    string   `toml:"compaction-guard-max-output-file-size"`
	BottommostLevelCompression          string   `toml:"bottommost-level-compression"`
	BottommostZstdCompressionDictSize   int      `toml:"bottommost-zstd-compression-dict-size"`
	BottommostZstdCompressionSampleSize int      `toml:"bottommost-zstd-compression-sample-size"`
	Titan                               struct {
		MinBlobSize             string  `toml:"min-blob-size"`
		BlobFileCompression     string  `toml:"blob-file-compression"`
		BlobCacheSize           string  `toml:"blob-cache-size"`
		MinGcBatchSize          string  `toml:"min-gc-batch-size"`
		MaxGcBatchSize          string  `toml:"max-gc-batch-size"`
		DiscardableRatio        float64 `toml:"discardable-ratio"`
		SampleRatio             float64 `toml:"sample-ratio"`
		MergeSmallFileThreshold string  `toml:"merge-small-file-threshold"`
		BlobRunMode             string  `toml:"blob-run-mode"`
		LevelMerge              bool    `toml:"level-merge"`
		RangeMerge              bool    `toml:"range-merge"`
		MaxSortedRuns           int     `toml:"max-sorted-runs"`
		GcMergeRewrite          bool    `toml:"gc-merge-rewrite"`
	} `toml:"titan"`
}


// TiKV config
type TiKVConfig struct {
	LogLevel                     string  `toml:"log-level"`
	LogFile                      string  `toml:"log-file"`
	LogFormat                    string  `toml:"log-format"`
	SlowLogFile                  string  `toml:"slow-log-file"`
	SlowLogThreshold             string  `toml:"slow-log-threshold"`
	LogRotationTimespan          string  `toml:"log-rotation-timespan"`
	LogRotationSize              string  `toml:"log-rotation-size"`
	PanicWhenUnexpectedKeyOrData bool    `toml:"panic-when-unexpected-key-or-data"`
	EnableIoSnoop                bool    `toml:"enable-io-snoop"`
	AbortOnPanic                 bool    `toml:"abort-on-panic"`
	MemoryUsageHighWater         float64 `toml:"memory-usage-high-water"`
	ReadPool                     ReadPoolConfig `toml:"readpool"`
	Server ServerConfig `toml:"server"`
	Storage StorageConfig `toml:"storage"`
	Pd PDConfig `toml:"pd"`
	Metric struct {
		Job string `toml:"job"`
	} `toml:"metric"`
	RaftStore RaftStoreConfig `toml:"raftstore"`
	Coprocessor CoprocessorConfig `toml:"coprocessor"`
	CoprocessorV2 struct { // todo maybe option in rust
	} `toml:"coprocessor-v2"`
	Rocksdb struct {
		InfoLogLevel                     string `toml:"info-log-level"`
		WalRecoveryMode                  int    `toml:"wal-recovery-mode"`
		WalDir                           string `toml:"wal-dir"`
		WalTTLSeconds                    int    `toml:"wal-ttl-seconds"`
		WalSizeLimit                     string `toml:"wal-size-limit"`
		MaxTotalWalSize                  string `toml:"max-total-wal-size"`
		MaxBackgroundJobs                int    `toml:"max-background-jobs"`
		MaxBackgroundFlushes             int    `toml:"max-background-flushes"`
		MaxManifestFileSize              string `toml:"max-manifest-file-size"`
		CreateIfMissing                  bool   `toml:"create-if-missing"`
		MaxOpenFiles                     int    `toml:"max-open-files"`
		EnableStatistics                 bool   `toml:"enable-statistics"`
		StatsDumpPeriod                  string `toml:"stats-dump-period"`
		CompactionReadaheadSize          string `toml:"compaction-readahead-size"`
		InfoLogMaxSize                   string `toml:"info-log-max-size"`
		InfoLogRollTime                  string `toml:"info-log-roll-time"`
		InfoLogKeepLogFileNum            int    `toml:"info-log-keep-log-file-num"`
		InfoLogDir                       string `toml:"info-log-dir"`
		RateBytesPerSec                  string `toml:"rate-bytes-per-sec"`
		RateLimiterRefillPeriod          string `toml:"rate-limiter-refill-period"`
		RateLimiterMode                  int    `toml:"rate-limiter-mode"`
		RateLimiterAutoTuned             bool   `toml:"rate-limiter-auto-tuned"`
		BytesPerSync                     string `toml:"bytes-per-sync"`
		WalBytesPerSync                  string `toml:"wal-bytes-per-sync"`
		MaxSubCompactions                int    `toml:"max-sub-compactions"`
		WritableFileMaxBufferSize        string `toml:"writable-file-max-buffer-size"`
		UseDirectIoForFlushAndCompaction bool   `toml:"use-direct-io-for-flush-and-compaction"`
		EnablePipelinedWrite             bool   `toml:"enable-pipelined-write"`
		EnableMultiBatchWrite            bool   `toml:"enable-multi-batch-write"`
		EnableUnorderedWrite             bool   `toml:"enable-unordered-write"`
		DefaultCF                        CFConfig `toml:"defaultcf"`
		WriteCF CFConfig `toml:"writecf"`
		LockCF CFConfig `toml:"lockcf"`
		RaftCF CFConfig `toml:"raftcf"`
		Titan struct {
			Enabled                  bool   `toml:"enabled"`
			Dirname                  string `toml:"dirname"`
			DisableGc                bool   `toml:"disable-gc"`
			MaxBackgroundGc          int    `toml:"max-background-gc"`
			PurgeObsoleteFilesPeriod string `toml:"purge-obsolete-files-period"`
		} `toml:"titan"`
	} `toml:"rocksdb"`
	RaftDB struct {
		WalRecoveryMode                  int    `toml:"wal-recovery-mode"`
		WalDir                           string `toml:"wal-dir"`
		WalTTLSeconds                    int    `toml:"wal-ttl-seconds"`
		WalSizeLimit                     string `toml:"wal-size-limit"`
		MaxTotalWalSize                  string `toml:"max-total-wal-size"`
		MaxBackgroundJobs                int    `toml:"max-background-jobs"`
		MaxBackgroundFlushes             int    `toml:"max-background-flushes"`
		MaxManifestFileSize              string `toml:"max-manifest-file-size"`
		CreateIfMissing                  bool   `toml:"create-if-missing"`
		MaxOpenFiles                     int    `toml:"max-open-files"`
		EnableStatistics                 bool   `toml:"enable-statistics"`
		StatsDumpPeriod                  string `toml:"stats-dump-period"`
		CompactionReadaheadSize          string `toml:"compaction-readahead-size"`
		InfoLogMaxSize                   string `toml:"info-log-max-size"`
		InfoLogRollTime                  string `toml:"info-log-roll-time"`
		InfoLogKeepLogFileNum            int    `toml:"info-log-keep-log-file-num"`
		InfoLogDir                       string `toml:"info-log-dir"`
		InfoLogLevel                     string `toml:"info-log-level"`
		MaxSubCompactions                int    `toml:"max-sub-compactions"`
		WritableFileMaxBufferSize        string `toml:"writable-file-max-buffer-size"`
		UseDirectIoForFlushAndCompaction bool   `toml:"use-direct-io-for-flush-and-compaction"`
		EnablePipelinedWrite             bool   `toml:"enable-pipelined-write"`
		EnableUnorderedWrite             bool   `toml:"enable-unordered-write"`
		AllowConcurrentMemtableWrite     bool   `toml:"allow-concurrent-memtable-write"`
		BytesPerSync                     string `toml:"bytes-per-sync"`
		WalBytesPerSync                  string `toml:"wal-bytes-per-sync"`
		DefaultCF                        CFConfig `toml:"defaultcf"`
		Titan struct {
			Enabled                  bool   `toml:"enabled"`
			Dirname                  string `toml:"dirname"`
			DisableGc                bool   `toml:"disable-gc"`
			MaxBackgroundGc          int    `toml:"max-background-gc"`
			PurgeObsoleteFilesPeriod string `toml:"purge-obsolete-files-period"`
		} `toml:"titan"`
	} `toml:"raftdb"`
	RaftEngine struct {
		Enable                    bool   `toml:"enable"`
		Dir                       string `toml:"dir"`
		RecoveryMode              string `toml:"recovery-mode"`
		BytesPerSync              string `toml:"bytes-per-sync"`
		TargetFileSize            string `toml:"target-file-size"`
		PurgeThreshold            string `toml:"purge-threshold"`
		BatchCompressionThreshold string `toml:"batch-compression-threshold"`
	} `toml:"raft-engine"`
	Security struct {
		CaPath        string        `toml:"ca-path"`
		CertPath      string        `toml:"cert-path"`
		KeyPath       string        `toml:"key-path"`
		CertAllowedCn []interface{} `toml:"cert-allowed-cn"`
		Encryption    struct {
			DataEncryptionMethod           string `toml:"data-encryption-method"`
			DataKeyRotationPeriod          string `toml:"data-key-rotation-period"`
			EnableFileDictionaryLog        bool   `toml:"enable-file-dictionary-log"`
			FileDictionaryRewriteThreshold int    `toml:"file-dictionary-rewrite-threshold"`
			MasterKey                      struct {
				Type string `toml:"type"`
			} `toml:"master-key"`
			PreviousMasterKey struct {
				Type string `toml:"type"`
			} `toml:"previous-master-key"`
		} `toml:"encryption"`
	} `toml:"security"`
	Import struct {
		NumThreads          int    `toml:"num-threads"`
		StreamChannelWindow int    `toml:"stream-channel-window"`
		ImportModeTimeout   string `toml:"import-mode-timeout"`
	} `toml:"import"`
	Backup struct {
		NumThreads int    `toml:"num-threads"`
		BatchSize  int    `toml:"batch-size"`
		SstMaxSize string `toml:"sst-max-size"`
	} `toml:"backup"`
	PessimisticTxn struct {
		WaitForLockTimeout  string `toml:"wait-for-lock-timeout"`
		WakeUpDelayDuration string `toml:"wake-up-delay-duration"`
		Pipelined           bool   `toml:"pipelined"`
	} `toml:"pessimistic-txn"`
	Gc struct {
		RatioThreshold                   float64 `toml:"ratio-threshold"`
		BatchKeys                        int     `toml:"batch-keys"`
		MaxWriteBytesPerSec              string  `toml:"max-write-bytes-per-sec"`
		EnableCompactionFilter           bool    `toml:"enable-compaction-filter"`
		CompactionFilterSkipVersionCheck bool    `toml:"compaction-filter-skip-version-check"`
	} `toml:"gc"`
	Split struct {
		QPSThreshold        int     `toml:"qps-threshold"`
		SplitBalanceScore   float64 `toml:"split-balance-score"`
		SplitContainedScore float64 `toml:"split-contained-score"`
		DetectTimes         int     `toml:"detect-times"`
		SampleNum           int     `toml:"sample-num"`
		SampleThreshold     int     `toml:"sample-threshold"`
		ByteThreshold       int     `toml:"byte-threshold"`
	} `toml:"split"`
	Cdc struct {
		MinTsInterval              string `toml:"min-ts-interval"`
		HibernateRegionsCompatible bool   `toml:"hibernate-regions-compatible"`
		IncrementalScanThreads     int    `toml:"incremental-scan-threads"`
		IncrementalScanConcurrency int    `toml:"incremental-scan-concurrency"`
		IncrementalScanSpeedLimit  string `toml:"incremental-scan-speed-limit"`
		SinkMemoryQuota            string `toml:"sink-memory-quota"`
		OldValueCacheMemoryQuota   string `toml:"old-value-cache-memory-quota"`
	} `toml:"cdc"`
	ResolvedTs struct {
		Enable            bool   `toml:"enable"`
		AdvanceTsInterval string `toml:"advance-ts-interval"`
		ScanLockPoolSize  int    `toml:"scan-lock-pool-size"`
	} `toml:"resolved-ts"`
	ResourceMetering struct {
		Enabled             bool   `toml:"enabled"`
		AgentAddress        string `toml:"agent-address"`
		ReportAgentInterval string `toml:"report-agent-interval"`
		MaxResourceGroups   int    `toml:"max-resource-groups"`
		Precision           string `toml:"precision"`
	} `toml:"resource-metering"`
}

var DefaultTiKVCfg = &TiKVConfig{RaftStore: RaftStoreConfig{
	ApplyMaxBatchSize: 256,
	StoreMaxBatchSize: 256,
}}