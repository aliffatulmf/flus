package ext

import (
	"fmt"
	"path/filepath"
)

// Supported file extensions
var (
	DocumentExtensions = []string{".pdf", ".xlsx", ".doc", ".docx", ".xls", ".csv", ".txt", ".ppt", ".pptx", ".odt", ".ods", ".odp", ".odg", ".odf", ".rtf", ".tex", ".wks", ".wps", ".wpd", ".yaml", ".yml", ".htm", ".html"}
	ImageExtensions    = []string{".jpeg", ".jpg", ".png", ".gif", ".tiff", ".tif", ".bmp", ".svg", ".eps", ".raw", ".cr2", ".nef", ".orf", ".sr2", ".psd", ".ai"}
	AudioExtensions    = []string{".mp3", ".wav", ".wma", ".mpa", ".aif", ".iff", ".m3u", ".m4a", ".mid", ".mpa", ".ra", ".ogg", ".oga", ".opus", ".flac", ".aac"}
	VideoExtensions    = []string{".avi", ".flv", ".wmv", ".mov", ".mp4", ".webm", ".vob", ".mng", ".qt", ".mpg", ".mpeg", ".3gp", ".mkv", ".m4v", ".h264", ".rm", ".swf", ".asf", ".asx", ".rmvb", ".srt", ".mpv"}
	ArchiveExtensions  = []string{".a", ".ar", ".cpio", ".iso", ".tar", ".gz", ".rz", ".7z", ".dmg", ".rar", ".xar", ".zip", ".jar", ".bz2", ".z", ".lz", ".lzma", ".lzo", ".xz", ".tz", ".deb", ".rpm", ".zipx", ".sit", ".sitx", ".pkg", ".bz2", ".tbz2", ".tgz", ".tlz", ".txz", ".war", ".ear", ".sar", ".rar", ".alz", ".ace", ".zoo", ".cpz", ".pak", ".arc"}
	BinaryExtensions   = []string{".exe", ".msi", ".bin", ".com", ".apk", ".app", ".bat", ".cgi", ".pl", ".gadget", ".jar", ".py", ".wsf", ".dmg", ".iso"}
)

// Directory names for file types
var (
	DocumentsDir = "Documents"
	ImagesDir    = "Images"
	AudioDir     = "Audio"
	VideoDir     = "Video"
	ArchivesDir  = "Archives"
	BinaryDir    = "Binary"
)

// Check if the file extension is in the list of supported extensions.
func isSupported(ext string, extensions []string) bool {
	for _, e := range extensions {
		if e == ext {
			return true
		}
	}
	return false
}

func IsSupported(name string) bool {
	ext := filepath.Ext(name)

	switch {
	case isSupported(ext, DocumentExtensions):
		return true
	case isSupported(ext, ImageExtensions):
		return true
	case isSupported(ext, AudioExtensions):
		return true
	case isSupported(ext, VideoExtensions):
		return true
		// case isSupported(ext, ArchiveExtensions):
		// 	return true
		// case isSupported(ext, BinaryExtensions):
		// 	return true
	}

	return false
}

// Get the directory name for the file extension. Returns an error if the extension is not supported.
func GetFileDirectory(name string) (string, error) {
	ext := filepath.Ext(name)

	switch {
	case isSupported(ext, DocumentExtensions):
		return DocumentsDir, nil
	case isSupported(ext, ImageExtensions):
		return ImagesDir, nil
	case isSupported(ext, AudioExtensions):
		return AudioDir, nil
	case isSupported(ext, VideoExtensions):
		return VideoDir, nil
	// case isSupported(ext, ArchiveExtensions):
	// 	return ArchivesDir, nil
	// case isSupported(ext, BinaryExtensions):
	// 	return BinaryDir, nil
	default:
		return "", fmt.Errorf("unsupported file extension: %s", ext)
	}
}

// IsDirSupported checks if the directory name is supported.
// Returns true if the directory name is supported, false otherwise.
// Supported directory names are: Documents, Images, Audio, Video, Archives, Binary.
func IsDirSupported(name string) bool {
	dtype := []string{DocumentsDir, ImagesDir, AudioDir, VideoDir, ArchivesDir, BinaryDir}

	for _, dir := range dtype {
		if dir == name {
			return true
		}
	}
	return false
}
