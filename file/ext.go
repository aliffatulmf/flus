package file

import (
	"fmt"
	"path/filepath"
)

// Supported file extensions
var (
	DocumentExtensions = []string{
		".pdf", ".xlsx", ".doc", ".docx", ".xls", ".csv", ".txt", ".ppt", ".pptx",
		".odt", ".ods", ".odp", ".odg", ".odf", ".rtf", ".tex", ".wks", ".wps",
		".wpd", ".yaml", ".yml", ".htm", ".html", ".xml", ".json", ".log", ".md",
		".markdown", ".msg", ".eml", ".emlx", ".mbox", ".pages", ".numbers", ".key",
		".epub", ".azw", ".azw3", ".fb2", ".djvu", ".djv", ".pdb", ".pml", ".pmlz",
		".rb", ".c", ".cpp", ".h", ".hpp", ".java", ".class", ".cs", ".vb", ".js",
		".php", ".swf", ".fla", ".pl", ".pm", ".sh", ".bat", ".cmd", ".ps1", ".vbs",
		".go", ".py", ".pyc", ".pyo", ".pyd", ".pyw", ".pyz", ".pyzw", ".swift",
		".kt", ".kts", ".rs", ".r", ".lua", ".coffee", ".sass", ".scss", ".less",
		".css", ".cshtml", ".vbhtml", ".jsp", ".asp", ".aspx", ".cer", ".cfm",
		".yaws", ".ejs", ".phtml", ".hbs", ".mustache", ".twig", ".jade", ".haml",
		".d", ".erl", ".hrl", ".ex", ".exs", ".eex", ".elm", ".ls", ".factor", ".fth",
		".fs", ".fsi", ".fsx", ".ml", ".mli", ".mll", ".mly", ".re", ".rei", ".res",
		".resi", ".resx", ".bas", ".frm", ".cls", ".ctl", ".vbp", ".vbg", ".vbscript",
	}

	ImageExtensions = []string{
		".jpeg", ".jpg", ".png", ".gif", ".tiff", ".tif", ".bmp", ".svg", ".eps",
		".raw", ".cr2", ".nef", ".orf", ".sr2", ".psd", ".ai", ".ico", ".webp",
	}

	AudioExtensions = []string{
		".mp3", ".wav", ".wma", ".mpa", ".aif", ".iff", ".m3u", ".m4a", ".mid", ".mpa",
		".ra", ".ogg", ".oga", ".opus", ".flac", ".aac", ".ac3", ".mka", ".wv", ".ape",
	}

	VideoExtensions = []string{
		".avi", ".flv", ".wmv", ".mov", ".mp4", ".webm", ".vob", ".mng", ".qt", ".mpg",
		".mpeg", ".3gp", ".mkv", ".m4v", ".h264", ".rm", ".swf", ".asf", ".asx", ".rmvb",
		".srt", ".mpv", ".ogv", ".dv", ".divx", ".xvid", ".m2ts", ".mts",
	}

	ArchiveExtensions = []string{
		".a", ".ar", ".cpio", ".iso", ".tar", ".gz", ".rz", ".7z", ".dmg", ".rar",
		".xar", ".zip", ".jar", ".bz2", ".z", ".lz", ".lzma", ".lzo", ".xz", ".tz",
		".deb", ".rpm", ".zipx", ".sit", ".sitx", ".pkg", ".bz2", ".tbz2", ".tgz",
		".tlz", ".txz", ".war", ".ear", ".sar", ".rar", ".alz", ".ace", ".zoo",
		".cpz", ".pak", ".arc", ".cab", ".msu", ".msix", ".appx", ".gz", ".xz", ".lz4",
	}

	BinaryExtensions = []string{
		".exe", ".msi", ".bin", ".com", ".apk", ".app", ".bat", ".cgi", ".pl", ".gadget",
		".jar", ".py", ".wsf", ".dmg", ".iso", ".elf", ".o", ".so", ".dll", ".dylib",
	}
)

// Directory names for file types
const (
	DocumentsDir = "Documents"
	ImagesDir    = "Images"
	AudioDir     = "Audio"
	VideoDir     = "Video"
	ArchivesDir  = "Archives"
	BinaryDir    = "Binary"
)

var ErrUnsupported = fmt.Errorf("unsupported file extension")

// IsIn checks if the provided extension is in the list of extensions.
func IsIn(ext string, exts []string) bool {
	for _, e := range exts {
		if e == ext {
			return true
		}
	}
	return false
}

// IsSupported checks if the provided file is supported.
func IsSupported(file string) bool {
	ext := filepath.Ext(file)

	return IsIn(ext, DocumentExtensions) ||
		IsIn(ext, ImageExtensions) ||
		IsIn(ext, AudioExtensions) ||
		IsIn(ext, VideoExtensions) ||
		IsIn(ext, ArchiveExtensions) ||
		IsIn(ext, BinaryExtensions)
}

// FileToDir returns the directory name for the provided file.
func FileToDir(file string) (string, error) {
	ext := filepath.Ext(file)

	switch {
	case IsIn(ext, DocumentExtensions):
		return DocumentsDir, nil
	case IsIn(ext, ImageExtensions):
		return ImagesDir, nil
	case IsIn(ext, AudioExtensions):
		return AudioDir, nil
	case IsIn(ext, VideoExtensions):
		return VideoDir, nil
	case IsIn(ext, ArchiveExtensions):
		return ArchivesDir, nil
	case IsIn(ext, BinaryExtensions):
		return BinaryDir, nil
	}

	return "", ErrUnsupported
}
