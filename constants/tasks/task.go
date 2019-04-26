package tasks

const (
	//FileMimeDetect detect file mime
	FileMimeDetect int32 = 1001
	// OfflineDownloadDetect detect
	// Detect Unkown Download Task
	OfflineDownloadDetect int32 = 5001

	// OfflineDownloadURL URL Only
	// URL type, http/https prifix
	OfflineDownloadURL int32 = 5020

	// OfflineDownloadTorrent Torrent Only
	// URL type, http/https prifix
	OfflineDownloadTorrent int32 = 5030

	// OfflineDownloadThunder thunder
	OfflineDownloadThunder int32 = 5040

	// OfflineDownloadMagnet magnet
	OfflineDownloadMagnet int32 = 5050
	// OfflineDownloadEd2k ED2K
	OfflineDownloadEd2k int32 = 5060
)
