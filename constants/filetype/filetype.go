package filetype

// github.com/zzzhr1990/go-server-common/constants/filetype

const (
	//Unkown cannot detect file
	Unkown int32 = 0

	// Directory directory
	Directory int32 = 10

	// Image image, NOT include GIF (GIF is video)
	Image int32 = 200

	// Video type, may be Gay Video
	Video int32 = 300

	// Audio type, GV -LOL-
	Audio int32 = 400

	// Document Torrent Only
	Document int32 = 500

	// Archive file
	Archive int32 = 600

	// Torrent BT
	Torrent int32 = 700
)
