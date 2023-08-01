package ext

import (
	"fmt"
	"path/filepath"
	"strings"
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
var (
	DocumentsDir = "Documents"
	ImagesDir    = "Images"
	AudioDir     = "Audio"
	VideoDir     = "Video"
	ArchivesDir  = "Archives"
	BinaryDir    = "Binary"
)

var ErrUnsupported = fmt.Errorf("unsupported file extension")

func isDocument(ext string) bool {
	for _, e := range DocumentExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func isAudio(ext string) bool {
	for _, e := range AudioExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func isImage(ext string) bool {
	for _, e := range ImageExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func isVideo(ext string) bool {
	for _, e := range VideoExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func isArchive(ext string) bool {
	for _, e := range ArchiveExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

func isBinary(ext string) bool {
	for _, e := range BinaryExtensions {
		if e == ext {
			return true
		}
	}
	return false
}

// FileClassifier classifies a file based on its extension.
// It returns the corresponding directory for the file type (DocumentsDir, ImagesDir, AudioDir, VideoDir).
// If the file extension is not supported, it returns an error.
func FileClassifier(file string) (string, error) {
	ext := strings.ToLower(filepath.Ext(file))

	switch {
	case isDocument(ext):
		return DocumentsDir, nil
	case isImage(ext):
		return ImagesDir, nil
	case isAudio(ext):
		return AudioDir, nil
	case isVideo(ext):
		return VideoDir, nil
	}

	return "", fmt.Errorf("unsupported file extension: %s", ext)
}

type File struct {
	Path string
	Name string
	Ext  string
}

func NewFile(file string) *File {
	return &File{
		Path: filepath.Dir(file),
		Name: filepath.Base(file),
		Ext:  strings.ToLower(filepath.Ext(file)),
	}
}

func (f *File) IsSupported() bool {
	return isDocument(f.Ext) || isImage(f.Ext) || isAudio(f.Ext) || isVideo(f.Ext) || isArchive(f.Ext) || isBinary(f.Ext)
}

func (f *File) GetDirName() (string, error) {
	switch {
	case isDocument(f.Ext):
		return DocumentsDir, nil
	case isImage(f.Ext):
		return ImagesDir, nil
	case isAudio(f.Ext):
		return AudioDir, nil
	case isVideo(f.Ext):
		return VideoDir, nil
	case isArchive(f.Ext):
		return ArchivesDir, nil
	case isBinary(f.Ext):
		return BinaryDir, nil
	}
	return "", ErrUnsupported
}
